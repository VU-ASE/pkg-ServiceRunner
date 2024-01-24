# Makefile in accordance with the docs on git management (to use in combination with meta)
.PHONY: build start clean test

BUILD_DIR=bin/
BINARY_NAME=ServiceRunner

build:
	@echo "building ${BINARY_NAME}"
	@cd src/ && go build -o "../$(BUILD_DIR)/${BINARY_NAME}"

start: build
	@echo "starting ${BINARY_NAME}"
	@cd "./${BUILD_DIR}/${BINARY_NAME}"

clean:
	@echo "Cleaning all targets for mod/Actuator"
	rm -rf $(BUILD_DIR)

test:
	go test ./src -v -count=1 -timeout 0