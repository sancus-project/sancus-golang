package uritemplate

// expressions
//
type exprType int

const (
	exprLITERAL exprType = iota + 1
	exprEOL
	exprSEQUENCE
	exprCAPTURE
	exprOPTIONAL
)

type expression interface {
	Type() exprType
	String() string
}

type container interface {
	addToken(t *token, p *parser) bool
	addExpression(e expression, p *parser) bool
}
