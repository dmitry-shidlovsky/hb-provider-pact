SHELL = "/bin/bash"

export PATH := $(PWD)/pact/bin:$(PATH)
export PATH
export PROVIDER_NAME = TestProvider
export CONSUMER_NAME = TestConsumer
export PACT_DIR = $(PWD)/pacts
export LOG_DIR = $(PWD)/log
