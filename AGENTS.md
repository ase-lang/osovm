# AGENTS.md

## Commands
- **Build**: `go build ./...`
- **Test**: `go test ./...`
- **Test Single Package**: `go test ./pkg/parser` (replace with specific package)
- **Test Single Test**: `go test -run TestName ./pkg/parser`
- **Format**: `go fmt ./...`
- **Lint**: `golangci-lint run` (if available)

## Architecture
- **osovm**: Virtual machine implementation
- **pkg/parser**: Parser for the VM language
- **pkg/runtime**: Runtime execution engine
- **pkg/telemetry**: Telemetry and monitoring
- **pkg/witness**: Witness/verification system
- **grammar/**: Language grammar definitions
- **examples/**: Example programs
- **vm/**: Core VM implementation

## Code Style
- Follow standard Go conventions and `gofmt` formatting
- Use descriptive names; prefer clarity over brevity
- Handle errors explicitly; never ignore returned errors
- Use package comments for public APIs
