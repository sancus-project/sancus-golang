package uritemplate

import (
	"fmt"
)

type exprLiteral struct {
	literal string
}

func (e *exprLiteral) Type() exprType {
	return exprLITERAL
}

func (e *exprLiteral) String() string {
	return fmt.Sprintf("%q", e.literal)
}
