package uritemplate

// Special logic elements
//
type exprSpecial struct {
	typ exprType
}

func (e *exprSpecial) Type() exprType {
	return e.typ
}

func (e *exprSpecial) String() string {
	switch e.typ {
	case exprEOL:
		return "EOL"
	default:
		return "unknown"
	}
}
