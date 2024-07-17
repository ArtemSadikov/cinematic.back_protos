GIT_CMD = git
BUMP2VERSION_CMD = bump2version

# Default target
.PHONY: all
all: gen_proto bump-version push

gen_proto:
	protoc -I proto ./proto/**/*.proto \
	--go_out=./generated/go				--go_opt=paths=source_relative \
	--go-grpc_out=./generated/go	--go-grpc_opt=paths=source_relative

# Bump version
.PHONY: bump-version
bump-version:
	@echo "Bumping version..."
	$(GIT_CMD) add .
	$(GIT_CMD) commit -m "Bump version"
	$(BUMP2VERSION_CMD) patch
	$(GIT_CMD) tag `$(BUMP2VERSION_CMD) --list | grep new_version= | sed -r s,"^.*=",,`

# Display help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make gen_proto   - Build proto files"
	@echo "  make bump-version  - Bump the git tag version"
	@echo "  make all           - Build proto files, bump version, and push to git"
