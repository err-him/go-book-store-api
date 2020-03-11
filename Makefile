.DEFAULT_GOAL=build

PROJECT?=$(shell basename $(shell pwd))

GO_BUILD_DIR=build
.PHONY: build

build:
				make clean
				mkdir -p $(GO_BUILD_DIR)
				go build -o $(GO_BUILD_DIR) -v

run:
				make clean
				make build
				cd $(GO_BUILD_DIR)
				./$(GO_BUILD_DIR)/$(PROJECT)



clean:
				rm -rf build

check-swagger:
				which swagger || (GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger:
				GO111MODULE=auto swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger:
				swagger serve -F=swagger ./swagger.yml
