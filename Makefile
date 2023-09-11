MAIN_PACKAGE_PATH := ./app/...
BINARY_NAME := test_quadro.exe

## build: build the application
.PHONY: build
build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	${BINARY_NAME}
