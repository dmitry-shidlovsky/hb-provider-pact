SHELL = "/bin/bash"

export PATH := $(PWD)/pact/bin:$(PATH)
export PATH
export PROVIDER_NAME = SoverenProvider
export CONSUMER_NAME = SoverenConsumer
export PACT_DIR = $(PWD)/pacts
export LOG_DIR = $(PWD)/log
