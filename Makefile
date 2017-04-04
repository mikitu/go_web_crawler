#get current dir
DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
IMG_NAME = "go-web-crawler-example"

GO := $(shell command -v go 2> /dev/null)
GLIDE := $(shell command -v glide 2> /dev/null)

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:  ## Build binaries and docker image
	glide install
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "$(DIR)/deploy/example1" "$(DIR)/src/main.go"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "$(DIR)/deploy/example2" "$(DIR)/src/examples/simple_regexp/main.go"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "$(DIR)/deploy/example3" "$(DIR)/src/examples/urls_and_images/main.go"

	docker build --no-cache=true -t $(IMG_NAME) .

dev:  ## Generates the binaries for development environment
	glide install
	go build -o ./bin/example1 "$(DIR)/src/main.go"
	go build -o ./bin/example2 "$(DIR)/src/examples/simple_regexp/main.go"
	go build -o ./bin/example3 "$(DIR)/src/examples/urls_and_images/main.go"

all:
    ifndef GO
        $(error "go is not available. please install go before continue")
    endif
    ifndef GLIDE
        $(echo "glide is not available. installing glide")
        $(shell command go get -v github.com/Masterminds/glide)
    endif