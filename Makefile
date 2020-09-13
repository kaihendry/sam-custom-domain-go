.PHONY: build deploy validate destroy

DOMAINNAME = domaingo.natalian.org
ACMCERTIFICATEARN = arn:aws:acm:us-east-1:407461997746:certificate/8e3ee384-d4f3-47b7-b0a6-c28fa0b4a26b
HOSTEDZONEID = Z2OT4MN00JO5F8

deploy: build
	AWS_PROFILE=mine sam deploy --parameter-overrides DomainName=$(DOMAINNAME) ACMCertificateArn=$(ACMCERTIFICATEARN) HostedZoneId=$(HOSTEDZONEID)

build:
	CGO_ENABLED=0 sam build

validate:
	aws cloudformation validate-template --template-body file://template.yaml

destroy:
	# TODO match with samconfig.toml
	AWS_PROFILE=mine aws cloudformation delete-stack --stack-name go-sam-sanity
