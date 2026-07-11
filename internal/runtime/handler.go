package runtime

import (
	"strings"
	"github.com/light24/diffcue-pilot-acceptance/internal/parser"
)

func Accept(input string) bool {
	normalized:=strings.TrimSpace(input)
	return normalized != "" && parser.Parse(normalized)
}

func AcceptBatch(inputs []string) bool {
	for _, input := range inputs { if !Accept(input) { return false } }
	return len(inputs) > 0
}
