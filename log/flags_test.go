package log

import (
	"testing"
)

func inteq(t *testing.T, n uint, wants, got uint) {
	if wants != got {
		t.Errorf("%v: Expected %v but got %v", n, wants, got)
	}
}

func TestSetFlags(t *testing.T) {

	ctx := NewLoggerContext(0)
	inteq(t, 1, ctx.Flags(), Lstdflags)

	ctx.SetFlags(Lpackage)
	inteq(t, 2, ctx.Flags(), Lpackage)

	ctx.SetFlags(0)
	inteq(t, 3, ctx.Flags(), Lpackage)

	ctx.SetFlags(Lor | Lfileline)
	inteq(t, 4, ctx.Flags(), Lpackage|Lfileline)

	l1 := ctx.NewLogger("test: ").SetFlags(Lor | Lshortfile)
	inteq(t, 5, ctx.Flags(), Lpackage|Lfileline)
	inteq(t, 6, l1.Flags(), Lpackage|Lfileline|Lshortfile)

	l2 := ctx.NewLogger("test: ").SetFlags(Lnot | Lpackage).SetFlags(Lor | Lfunc)
	inteq(t, 7, ctx.Flags(), Lpackage|Lfileline)
	inteq(t, 8, l1.Flags(), Lpackage|Lfileline|Lshortfile)
	inteq(t, 9, l2.Flags(), Lfunc|Lfileline)

	l2.SetFlags(0)
	inteq(t, 10, l2.Flags(), ctx.Flags())
}
