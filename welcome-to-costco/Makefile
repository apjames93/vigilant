.PHONY: build start-dev

start-dev:
	sam build
	sam local start-api --debug
build-MyLayer:
	go mod tidy

	cp go.mod "$(ARTIFACTS_DIR)"
	cp go.sum "$(ARTIFACTS_DIR)"
build:
	sam build
