.PHONY: build deploy

deploy: build
	AWS_PROFILE=mine sam deploy

build:
	CGO_ENABLED=0 sam build