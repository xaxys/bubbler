# Copyright (C) 2023 xaxys. All rights reserved.
PACKAGE_NAME   := bubbler

# GEN_FILES := $(wildcard modules/*/cmd/generate.go)

ifeq ($(OS),Windows_NT)  # is Windows_NT on XP, 2000, 7, Vista, 10...
    GO         ?= go.exe
    PWD        := ${CURDIR}
    TARGET     := $(PACKAGE_NAME).exe
    BUILD_TAGS := $(shell git describe --tags --always --dirty="-dev")
    BUILD_TIME := $(shell echo %date% %time%)
    GIT_COMMIT := $(shell git rev-parse --short HEAD)
    GO_VERSION := $(subst go version ,,$(shell go version))
    GOPATH     := $(subst ;,,$(shell go env GOPATH))
    RM_CMD_1   := del /s /q
    RM_CMD_2   := 
    EXPORT     := set
    SCRIPT_EXT := .bat
else ifeq ($(shell uname),Darwin)
    GO         ?= go
    PWD        := ${CURDIR}
    TARGET     := $(PACKAGE_NAME)
    BUILD_TAGS := $(shell git describe --tags --always --dirty="-dev")
    BUILD_TIME := $(shell date -z UTC)
    GIT_COMMIT := $(shell git rev-parse --short HEAD)
    GO_VERSION := $(subst go version ,,$(shell go version))
    GOPATH     := $(shell go env GOPATH)
    RM_CMD_1   := find . -type f -name
    RM_CMD_2   := -delete
    EXPORT     := export
    SCRIPT_EXT := .sh
else
    GO         ?= go
    PWD        := ${CURDIR}
    TARGET     := $(PACKAGE_NAME)
    BUILD_TAGS := $(shell git describe --tags --always --dirty="-dev")
    BUILD_TIME := $(shell date --utc)
    GIT_COMMIT := $(shell git rev-parse --short HEAD)
    GO_VERSION := $(subst go version ,,$(shell go version))
    GOPATH     := $(shell go env GOPATH)
    RM_CMD_1   := find . -type f -name
    RM_CMD_2   := -delete
    EXPORT     := export
    SCRIPT_EXT := .sh
endif

define exec-cmd
$(1)

endef

all: gen build

gen:
	@echo Generating $(PACKAGE_NAME) ...
#	@$(foreach file, $(GEN_FILES), $(call exec-cmd, $(GO) run $(file)))
	@$(call exec-cmd, $(PWD)/tools/gen$(SCRIPT_EXT))

build:
	@echo Building $(PACKAGE_NAME) ...
	@$(GO) env -w CGO_ENABLED="1"
	@$(GO) build \
		-ldflags="-X 'main.BuildTags=$(BUILD_TAGS)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.GoVersion=$(GO_VERSION)'" \
		-o $(TARGET) $(PWD)/main.go

run: build
	@echo Running $(PACKAGE_NAME) ...
	@$(PWD)/$(TARGET)

test: test-short

test-full: clean
	@echo Testing $(PACKAGE_NAME) ...
	@$(GO) env -w CGO_ENABLED="1"
	@$(GO) test \
		-ldflags="-X 'main.BuildTags=$(BUILD_TAGS)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.GoVersion=$(GO_VERSION)'" \
		-timeout=30m -race -coverprofile=coverage.out ./...

test-short: clean
	@echo Testing $(PACKAGE_NAME) ...
	@$(GO) env -w CGO_ENABLED="1"
	@$(GO) test \
		-ldflags="-X 'main.BuildTags=$(BUILD_TAGS)' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.GoVersion=$(GO_VERSION)'" \
		-timeout=30m -race -short -coverprofile=coverage.out ./...

clean:
	@echo Cleaning $(PACKAGE_NAME) ...
	-@$(RM_CMD_1) $(TARGET)    $(RM_CMD_2)
	-@$(RM_CMD_1) coverage.out $(RM_CMD_2)
	-@$(RM_CMD_1) "*.db"       $(RM_CMD_2)
	-@$(RM_CMD_1) "*.exe"      $(RM_CMD_2)
	-@$(RM_CMD_1) "*.out"      $(RM_CMD_2)
	-@$(RM_CMD_1) "*.yaml"     $(RM_CMD_2)
ifeq ($(OS),Windows_NT)
	-@for /r e2e\tests %%f in (*.bb.go) do @del /q "%%f" >nul 2>&1 || ver >nul
	-@for /f "delims=" %%f in ('dir /s /b /a-d e2e\tests\* 2^>nul') do @echo %%~dpf | findstr /i "\\gen\\" >nul && del /q "%%f" >nul 2>&1 || ver >nul
	-@for /f "delims=" %%d in ('dir /s /b /ad e2e\tests 2^>nul ^| sort /R') do @rd "%%d" 2>nul || ver >nul
else
	@find e2e/tests -type f \( -path "*/gen/*" -o -name "*.bb.go" \) -delete
	@find e2e/tests -type d -empty -delete
endif

e2e: $(TARGET)
ifeq ($(OS),Windows_NT)
	@echo Running e2e tests via Docker on Windows ...
	@docker-compose -f e2e/docker-compose.yml run --build --rm e2e
else
	@echo Running e2e tests ...
	@BUBBLER=$(PWD)/$(TARGET) bash e2e/run_tests.sh
endif

e2e-docker:
	@echo Running e2e tests in Docker ...
	@docker-compose -f e2e/docker-compose.yml run --build --rm e2e

.PHONY: all gen build run test test-full test-short clean e2e e2e-docker
