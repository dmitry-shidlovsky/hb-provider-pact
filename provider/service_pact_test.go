package provider

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

func TestProvider_pact(t *testing.T) {
	go startInstrumentedProvider()

	pact := createPact()

	publishVerificationResults := false
	env, found := os.LookupEnv("PUBLISH_VERIFICATION_RESULTS")
	if found {
		result, err := strconv.ParseBool(env)
		publishVerificationResults = result
		if err != nil {
			t.Fatal(err)
		}
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", port),
		PactURLs:                   []string{filepath.FromSlash(fmt.Sprintf("%s/testconsumer-testprovider.json", os.Getenv("PACT_DIR")))},
		ProviderVersion:            os.Getenv("PACT_SERVICE_VERSION"),
		Tags:                       []string{"master"},
		BrokerURL:                  os.Getenv("PACT_BROKER_URL"),
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		PublishVerificationResults: publishVerificationResults,
	})

	if err != nil {
		t.Log(err)
	}
}

// Starts the provider API with hooks for provider states.
// This essentially mirrors the main.go file, with extra routes added.
func startInstrumentedProvider() {
	userRepository.TestInit()
	mux := GetHTTPHandler()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %d (%s)", port, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))

}

// Configuration / Test Data
var dir, _ = os.Getwd()
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort()

// Setup the Pact client.
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "TestProvider",
		LogDir:   logDir,
		LogLevel: "INFO",
	}
}
