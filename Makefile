.PHONY: dev
dev:
	@bun run turbowatch

.PHONY: install
install: bun_install install_sqlc

.PHONY: bun_install
bun_install:
	@bun install

.PHONY: install_sqlc
install_sqlc:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
