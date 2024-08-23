# API
.PHONY: bundle
bundle:
	docker run --rm -v ${PWD}:/spec redocly/cli bundle docs/api/api.yaml -o docs/api/bundled.yaml

.PHONY: build-docs
build-docs:
	docker run --rm -v ${PWD}:/spec redocly/cli build-docs docs/api/api.yaml --o docs/api/redoc.html

# ogen
.PHONY: install
install:
	go install -v github.com/ogen-go/ogen/cmd/ogen@latest

.PHONY: gen
gen:
	go generate ./...
