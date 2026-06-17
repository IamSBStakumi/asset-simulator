REVIVE_VERSION := v1.5.1
MOCKGEN_VERSION := v0.6.0

.PHONY: setup
setup:
	go install github.com/mgechev/revive@$(REVIVE_VERSION)
	go install go.uber.org/mock/mockgen@$(MOCKGEN_VERSION)
	git config --local core.hooksPath .githooks
	chmod +x .githooks/pre-commit

run-cli:
	go run ./cmd/asset-simulator

build-cli:
	go build -o assetsim ./cmd/asset-simulator

lintAll:
	go vet ./...
	go fmt ./...
	revive -config reviveConfig.toml -formatter friendly ./...

lint:
	go vet ${FILENAME}
	go fmt ${FILENAME}
	revive -config reviveConfig.toml -formatter friendly ${FILENAME}

lint-dir:
	go vet ${DIR}/*
	go fmt ${DIR}/*
	revive -config reviveConfig.toml -formatter friendly ${DIR}/...
