default: testacc

.PHONY: deps
deps:
	go mod download

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

version = 0.1.0
provider_path = registry.terraform.io/simiancreative/sentry/$(version)/darwin_arm64/
root_path = /Users/ross/code/reguard/endo/subscriptions/envs/global/.terraform/providers/

install_macos:
	go build -o terraform-provider-sentry_$(version)

	mkdir -p $(root_path)$(provider_path)
	cp terraform-provider-sentry_$(version) $(root_path)$(provider_path)

