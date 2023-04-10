package transactiontracker

import (
	_ "github.com/ogen-go/ogen/gen"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target internal/tracker/models --clean openapi.yaml
