include ./config.mk

install:
	@if [ ! -d pact/bin ]; then\
		curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash;\
    fi

provider-unit:
	go test ./provider -run unit -count=1

provider-pact: install
	go test ./provider -run pact -count=1

.PHONY: install provider-pact provider-unit
