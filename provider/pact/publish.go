package main

import (
	"log"
	"os"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func main() {
	p := dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{os.Getenv("PACT_BROKER_URL") + "/contracts/provider/pactflow-example-bi-directional-provider-dredd/version/" + os.Getenv("PACT_SERVICE_VERSION")},
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
