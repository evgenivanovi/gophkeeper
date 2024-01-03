# ____________________________________________________________________________________________________ #

# __________________________________________________ #
# Functions
# __________________________________________________ #

### GIT functions
git_hash=$(shell git log --format="%h" -n 1 2> /dev/null)
git_tag=$(shell git describe --exact-match --abbrev=0 --tags 2> /dev/null)

# __________________________________________________ #
# Variables
# __________________________________________________ #

### GO properties
GOBIN?=$(GOPATH)/bin
LOCAL_BIN:=$(CURDIR)/bin
export PATH:=$(GOBIN):$(PATH)

### VCS properties
GIT_HASH:=$(call git_hash)
GIT_TAG:=$(call git_tag)

### Project properties
MODULE_NAME=github.com/evgenivanovi/gophkeeper
MODULE_PATH=$(CURDIR)

SPEC_PB_MODULE_NAME=$(MODULE_NAME)/spec/pb
SPEC_PB_MODULE_PATH=$(CURDIR)/spec/pb

API_PB_MODULE_NAME=$(MODULE_NAME)/api/pb
API_PB_MODULE_PATH=$(CURDIR)/api/pb

### APP properties
APP_VERSION:=$(if $(GIT_TAG),$(GIT_TAG)-$(GIT_HASH),$(shell git describe --all --long HEAD))

### Meta properties
BIN_DIR=$(CURDIR)/bin
BUILD_DIR=$(CURDIR)/build

APP_SRC_DIR=$(CURDIR)/cmd/gophkeeper
APP_BIN_NAME=gophkeeper

CLI_SRC_DIR=$(CURDIR)/cmd/gophkeepctl
CLI_BIN_NAME=gophkeepctl

MIG_SRC_DIR=$(CURDIR)/cmd/migrator
MIG_BIN_DIR=migrator

### App properties
MIGRATION_DIR="./migrations"
DATABASE_DSN="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

# __________________________________________________ #
# Scripts
# __________________________________________________ #

# ____________________________________________________________________________________________________ #

.PHONY: init
init:
	@echo 'Project initialization.'

	@echo 'Evaluating GIT properties:'
	@echo GIT_HASH=$(GIT_HASH)
	@echo GIT_TAG=$(GIT_TAG)

	@echo 'Evaluating Application properties'
	@echo APP_VERSION=$(APP_VERSION)

	@echo 'Installing dependencies.'

	@mkdir -p $(LOCAL_BIN)

	@ls $(LOCAL_BIN)/golangci-lint &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

	@ls $(LOCAL_BIN)/goimports &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@latest

	@ls $(LOCAL_BIN)/protoc-gen-go &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	@ls $(LOCAL_BIN)/jet &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install github.com/go-jet/jet/v2/cmd/jet@latest

	@ls $(LOCAL_BIN)/protoc-gen-go-grpc &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# __________________________________________________ #

.PHONY: tidy
tidy:
	@find $(CURDIR) \
		-name 'go.mod' \
		-exec bash -c 'pushd "$${1%go.mod}" && go mod tidy && popd' _ {} \; \
		> /dev/null

# __________________________________________________ #

.PHONY: app/build/all
app/build/all: init
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(APP_BIN_NAME)/$(APP_BIN_NAME)-linux-amd64 $(APP_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=linux \
	GOARCH=arm64 \
	go build -o $(BUILD_DIR)/$(APP_BIN_NAME)/$(APP_BIN_NAME)-linux-arm64 $(APP_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(APP_BIN_NAME)/$(APP_BIN_NAME)-darwin-amd64 $(APP_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=darwin \
	GOARCH=arm64 \
	go build -o $(BUILD_DIR)/$(APP_BIN_NAME)/$(APP_BIN_NAME)-darwin-arm64 $(APP_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=windows \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(APP_BIN_NAME)/$(APP_BIN_NAME)-windows-amd64 $(APP_SRC_DIR);

# __________________________________________________ #

.PHONY: app/build
app/build: init
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	go build -o $(BUILD_DIR)/$(APP_BIN_NAME)/$(APP_BIN_NAME) $(APP_SRC_DIR)

# __________________________________________________ #

.PHONY: app/run
app/run: init app/build
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	$(BUILD_DIR)/$(APP_BIN_NAME)

# __________________________________________________ #

.PHONY: cli/build/all
cli/build/all: init
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(CLI_BIN_NAME)/$(CLI_BIN_NAME)-linux-amd64 $(CLI_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=linux \
	GOARCH=arm64 \
	go build -o $(BUILD_DIR)/$(CLI_BIN_NAME)/$(CLI_BIN_NAME)-linux-arm64 $(CLI_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(CLI_BIN_NAME)/$(CLI_BIN_NAME)-darwin-amd64 $(CLI_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=darwin \
	GOARCH=arm64 \
	go build -o $(BUILD_DIR)/$(CLI_BIN_NAME)/$(CLI_BIN_NAME)-darwin-arm64 $(CLI_SRC_DIR);

	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	GOOS=windows \
	GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(CLI_BIN_NAME)/$(CLI_BIN_NAME)-windows-amd64 $(CLI_SRC_DIR);

# __________________________________________________ #

.PHONY: cli/build
cli/build: init
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	go build -o $(BUILD_DIR)/$(CLI_BIN_NAME)/$(CLI_BIN_NAME) $(CLI_SRC_DIR)

# __________________________________________________ #

.PHONY: cli/run
cli/run: init cli/build
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	GIT_HASH=$(GIT_HASH) \
	GIT_TAG=$(GIT_TAG) \
	APP_VERSION=$(APP_VERSION) \
	$(BUILD_DIR)/$(CLI_BIN_NAME)

# __________________________________________________ #

.PHONY: db/jet
db/jet: init
	@echo 'Running schema initialization...'
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	$(LOCAL_BIN)/jet -dsn=$(DATABASE_DSN) -path=./internal/server

# __________________________________________________ #

.PHONY: db/migrations/up
db/migrations/up: init
	@echo 'Running migration...'
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	go run $(MIG_SRC_DIR) \
		-dir $(MIGRATION_DIR) \
		-dsn $(DATABASE_DSN) \
		-command up

# __________________________________________________ #

.PHONY: db/migrations/down
db/migrations/down: init
	@echo 'Running rollback migration...'
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	go run $(MIG_SRC_DIR) \
		-dir $(MIGRATION_DIR) \
		-dsn $(DATABASE_DSN) \
		-command down

# __________________________________________________ #

.PHONY: proto/init
proto/init:
	@echo 'Proto initialization.'

	@echo 'Installing dependencies.'

	@mkdir -p $(LOCAL_BIN)

	@ls $(LOCAL_BIN)/protoc-gen-go &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	@ls $(LOCAL_BIN)/protoc-gen-go-grpc &> /dev/null || \
		GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

	@echo 'Create directories for proto files.'
	@mkdir -p $(SPEC_PB_MODULE_PATH)
	@echo 'Create directories for generated proto files.'
	@mkdir -p $(API_PB_MODULE_PATH)

	@echo 'Generating SPEC module.'
	@pushd $(SPEC_PB_MODULE_PATH) > /dev/null && go mod init $(SPEC_PB_MODULE_NAME) || true && popd > /dev/null
	@pushd $(SPEC_PB_MODULE_PATH) > /dev/null && go mod tidy || true && popd > /dev/null

# __________________________________________________ #

.PHONY: proto/clean
proto/clean: init
	@echo 'Clean generated proto files.'
	@$(shell find $(API_PB_MODULE_PATH) -name '*.go' -exec rm -rf {} \;)

# __________________________________________________ #

.PHONY: proto/compile
proto/compile: init proto/init
	@echo 'Install dependencies for API module'
	@GOPATH=$(GOPATH) \
	GOBIN=$(GOBIN) \
	pushd $(MODULE_PATH) > /dev/null && \
		go get google.golang.org/protobuf@latest && \
		go get google.golang.org/grpc@latest && \
		popd > /dev/null

	@echo 'Compile proto files.'
	@protoc \
        --proto_path $(SPEC_PB_MODULE_PATH) \
        --go_out=$(API_PB_MODULE_PATH) \
        --go_opt=paths=source_relative \
        --go-grpc_out=$(API_PB_MODULE_PATH) \
        --go-grpc_opt=paths=source_relative \
        $(shell find $(SPEC_PB_MODULE_PATH) -name '*.proto')

# ____________________________________________________________________________________________________ #
