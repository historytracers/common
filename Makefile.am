# Automake options
AUTOMAKE_OPTIONS = foreign

# Build configuration
PUBLISHERBINARY_NAME = historytracers-publisher$(GO_EXEEXT)
BUILD_DIR = build
PUBLISHERSRC_DIR = src/publisher

# Go compiler wrapper that ensures its bin dir is in PATH and GOROOT is set
GO_CMD = PATH="$(GO_BIN_DIR):$(PATH)" GOROOT="$(GO_GOROOT)" "$(GO_COMPILER)"
GO_RUN = cd "$(abs_top_srcdir)/src/go" && $(GO_CMD)

# What to distribute
EXTRA_DIST = src/go/config.go src/go/creator.go src/go/data-type.go src/go/timestamp.go src/go/go.mod \
	$(PUBLISHERSRC_DIR)/config.go $(PUBLISHERSRC_DIR)/main.go $(PUBLISHERSRC_DIR)/go.mod \
	$(PUBLISHERSRC_DIR)/smartphone/audio.go $(PUBLISHERSRC_DIR)/smartphone/common.go \
	$(PUBLISHERSRC_DIR)/smartphone/config.go $(PUBLISHERSRC_DIR)/smartphone/doc.go \
	$(PUBLISHERSRC_DIR)/smartphone/git.go $(PUBLISHERSRC_DIR)/smartphone/hash.go \
	$(PUBLISHERSRC_DIR)/smartphone/minify.go $(PUBLISHERSRC_DIR)/smartphone/runner.go \
	$(PUBLISHERSRC_DIR)/smartphone/smgame.go $(PUBLISHERSRC_DIR)/smartphone/sources.go \
	$(PUBLISHERSRC_DIR)/smartphone/smartphone_test.go \
	$(PUBLISHERSRC_DIR)/README.md $(PUBLISHERSRC_DIR)/README.es.md $(PUBLISHERSRC_DIR)/README.pt-BR.md \
	build-publisher.sh

# Verify the library compiles and build the local smartphone publisher
all-local: publisher
	@echo "Building $(PACKAGE_NAME) $(PACKAGE_VERSION)..."
	$(GO_RUN) mod download
	$(GO_RUN) mod tidy
	$(GO_RUN) fmt ./...
	$(GO_RUN) build ./...
	$(GO_RUN) vet ./...

# Build the local smartphone publisher, which rewrites and minifies the
# content under src/smartphone/ into the content directory (build/www/).
publisher:
	@echo "Building $(PUBLISHERBINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	cd $(PUBLISHERSRC_DIR) && $(GO_CMD) mod download && $(GO_CMD) mod tidy && $(GO_CMD) fmt ./...
	cd $(PUBLISHERSRC_DIR) && $(GO_CMD) build $(GO_BUILD_FLAGS) -o ../../$(BUILD_DIR)/$(PUBLISHERBINARY_NAME) .

# Run tests
test:
	$(GO_RUN) test ./... -v
	cd $(PUBLISHERSRC_DIR) && $(GO_CMD) test ./... -v

# Run benchmarks
bench:
	$(GO_RUN) test ./... -bench=. -run=^$

# Install the Go package and the publisher binary
install-exec-local:
	@echo "Installing $(PACKAGE_NAME)..."
	$(GO_RUN) install ./...
	$(MKDIR_P) $(DESTDIR)$(bindir)
	$(INSTALL_PROGRAM) $(BUILD_DIR)/$(PUBLISHERBINARY_NAME) $(DESTDIR)$(bindir)/$(PUBLISHERBINARY_NAME)

# Clean
clean-local:
	@echo "Cleaning..."
	$(GO_RUN) clean ./...
	-cd $(PUBLISHERSRC_DIR) && $(GO_CMD) clean ./...
	rm -rf $(BUILD_DIR)

# Phony targets
.PHONY: test bench publisher
