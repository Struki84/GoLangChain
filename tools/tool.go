package tools

import (
	"context"

	"github.com/tmc/langchaingo/llms"
)

// Tool is a tool for the llm agent to interact with different applications.
type Tool interface {
	Name() string
	Description() string
	Definition() *llms.FunctionDefinition
	Call(ctx context.Context, input string) (string, error)
}
