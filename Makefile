# SPDX-FileCopyrightText: 2022 Kalle Fagerberg
# 
# SPDX-License-Identifier: CC0-1.0

.PHONY: build
build: hns-tree

hns-tree: *.go internal/*/*.go
	go build

.PHONY: install
install:
	go install .

.PHONY: run
run:
	go run .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -f hns-tree

.PHONY: check
check:
	go test ./...

.PHONY: deps
deps: deps-go deps-pip deps-npm

.PHONY: deps-go
deps-go:
	go install github.com/mgechev/revive@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go get

.PHONY: deps-pip
deps-pip:
	python3 -m pip install --upgrade --user reuse

.PHONY: deps-npm
deps-npm: node_modules

node_modules:
	npm install

.PHONY: lint
lint: lint-md lint-go lint-license

.PHONY: lint-fix
lint-fix: lint-md-fix lint-go-fix

.PHONY: lint-md
lint-md: node_modules
	npx remark .

.PHONY: lint-md-fix
lint-md-fix: node_modules
	npx remark . -o

.PHONY: lint-go
lint-go:
	@echo goimports -d '**/*.go'
	@goimports -d $(shell git ls-files "*.go")
	revive -formatter stylish -config revive.toml ./...

.PHONY: lint-go-fix
lint-go-fix:
	@echo goimports -d -w '**/*.go'
	@goimports -d -w $(shell git ls-files "*.go")

.PHONY: lint-license
lint-license:
	reuse lint
