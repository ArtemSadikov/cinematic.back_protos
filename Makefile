GIT_CMD = git
BUMP2VERSION_CMD = bump2version

# Default target
.PHONY: all
all: gen_proto bump-version


PROTO_FILES := $(shell find proto -name "*.proto")
# PROTO_FODLERS := $(shell find ./proto -type d ! -name "proto")
gen_proto:
	protoc -I=proto $(PROTO_FILES) \
	--go_out=./generated/go				--go_opt=paths=source_relative \
	--go-grpc_out=./generated/go	--go-grpc_opt=paths=source_relative

# Bump version
.PHONY: bump-version
bump-version:
	@echo "Bumping version..."
	$(BUMP2VERSION_CMD) patch

# Display help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make gen_proto  		- Build proto files"
	@echo "  make bump-version  - Bump the git tag version"
	@echo "  make all           - Build proto files, bump version"
