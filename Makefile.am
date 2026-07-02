# Automake options
AUTOMAKE_OPTIONS = foreign

# Go compiler wrapper that ensures its bin dir is in PATH and GOROOT is set
GO_RUN = cd "$(abs_top_srcdir)/src/go" && PATH="$(GO_BIN_DIR):$(PATH)" GOROOT="$(GO_GOROOT)" "$(GO_COMPILER)"

# What to distribute
EXTRA_DIST = src/go/config.go src/go/creator.go src/go/data-type.go src/go/timestamp.go src/go/go.mod

# Verify the library compiles
all-local:
	@echo "Building $(PACKAGE_NAME) $(PACKAGE_VERSION)..."
	$(GO_RUN) mod download
	$(GO_RUN) mod tidy
	$(GO_RUN) fmt ./...
	$(GO_RUN) build ./...
	$(GO_RUN) vet ./...

# Run tests
test:
	$(GO_RUN) test ./... -v

# Run benchmarks
bench:
	$(GO_RUN) test ./... -bench=. -run=^$

# Install the Go package
install-exec-local:
	@echo "Installing $(PACKAGE_NAME)..."
	$(GO_RUN) install ./...

# Clean
clean-local:
	@echo "Cleaning..."
	$(GO_RUN) clean ./...

# Phony targets
.PHONY: test bench
