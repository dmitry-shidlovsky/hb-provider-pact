package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func main() {
	p := dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/testconsumer-testprovider.json", os.Getenv("PACT_DIR")))},
		PactBroker:      os.Getenv("PACT_BROKER_URL"),
		ConsumerVersion: os.Getenv("PACT_SERVICE_VERSION"),
		Tags:            []string{os.Getenv("PACT_SERVICE_TAG")},
		BrokerToken:     os.Getenv("PACT_BROKER_TOKEN"),
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
