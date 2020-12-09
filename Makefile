STACK = gosamsanity2

.PHONY: build deploy validate destroy

DOMAINNAME = domaingo.webconverger.com
ACMCERTIFICATEARN = arn:aws:acm:us-east-1:407461997746:certificate/64cb77f4-763f-45fc-8536-e3a8a848fe3b

deploy: build
	AWS_PROFILE=mine sam deploy --stack-name $(STACK) --parameter-overrides DomainName=$(DOMAINNAME) ACMCertificateArn=$(ACMCERTIFICATEARN)

build:
	CGO_ENABLED=0 sam build

validate:
	aws cloudformation validate-template --template-body file://template.yaml

destroy:
	AWS_PROFILE=mine aws cloudformation delete-stack --stack-name $(STACK)
