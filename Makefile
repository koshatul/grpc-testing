MATRIX_OS ?= darwin linux windows
MATRIX_ARCH ?= amd64

GIT_HASH ?= $(shell git show -s --format=%h)
GIT_TAG ?= $(shell git tag -l --merged $(GIT_HASH) | tail -n1)
APP_VERSION ?= $(if $(TRAVIS_TAG),$(TRAVIS_TAG),$(if $(GIT_TAG),$(GIT_TAG),$(GIT_HASH)))
APP_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

CLI_ARGS ?= world

-include artifacts/make/go/Makefile

artifacts/make/%/Makefile:
	curl -sf https://jmalloc.github.io/makefiles/fetch | bash /dev/stdin $*

# .PHONY: install
# install: vendor $(REQ) $(_SRC) | $(USE)
# 	$(eval PARTS := $(subst /, ,$*))
# 	$(eval BUILD := $(word 1,$(PARTS)))
# 	$(eval OS    := $(word 2,$(PARTS)))
# 	$(eval ARCH  := $(word 3,$(PARTS)))
# 	$(eval BIN   := $(word 4,$(PARTS)))
# 	$(eval ARGS  := $(if $(findstring debug,$(BUILD)),$(DEBUG_ARGS),$(RELEASE_ARGS)))

# 	CGO_ENABLED=$(CGO_ENABLED) GOOS="$(OS)" GOARCH="$(ARCH)" go install $(ARGS) "./src/cmd/..."

.PHONY: run
run: artifacts/build/debug/$(GOOS)/$(GOARCH)/grpc-testing cfssl
	$< $(RUN_ARGS)

.PHONY: server
server: artifacts/build/debug/$(GOOS)/$(GOARCH)/grpc-testing cfssl
	$< server

.PHONY: client
client: artifacts/build/debug/$(GOOS)/$(GOARCH)/grpc-testing cfssl
	$< client $(CLI_ARGS)

.PHONY: cfssl
cfssl: artifacts/cfssl/localhost.pem artifacts/cfssl/ca.pem artifacts/cfssl/ca-config.json

.PRECIOUS: artifacts/cfssl/ca-config.json
artifacts/cfssl/ca-config.json:
	-@mkdir -p artifacts/cfssl
	cp test/ca-config.json artifacts/cfssl/ca-config.json

.PRECIOUS: artifacts/cfssl/ca.pem
artifacts/cfssl/ca.pem: artifacts/cfssl/ca-config.json
	-@mkdir -p artifacts/cfssl
	cfssl gencert -initca -config="artifacts/cfssl/ca-config.json" -profile="ca" test/ca-csr.json | cfssljson -bare artifacts/cfssl/ca -
	cfssl sign -ca="artifacts/cfssl/ca.pem" -ca-key="artifacts/cfssl/ca-key.pem" -config="artifacts/cfssl/ca-config.json" -profile="ca" -csr=artifacts/cfssl/ca.csr test/ca-csr.json | cfssljson -bare artifacts/cfssl/ca

.PRECIOUS: artifacts/cfssl/localhost.pem
artifacts/cfssl/localhost.pem: artifacts/cfssl/ca-config.json artifacts/cfssl/ca.pem
	-@mkdir -p artifacts/cfssl
	cfssl gencert -ca="artifacts/cfssl/ca.pem" -ca-key="artifacts/cfssl/ca-key.pem" -config="artifacts/cfssl/ca-config.json" -profile="server" test/localhost.json | cfssljson -bare artifacts/cfssl/localhost