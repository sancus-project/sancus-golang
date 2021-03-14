package log

const (
	Lnoprefix = 1 << iota

	Ldate
	Ltime
	Lseconds
	Lelapsed
	Lrelative
	Lmicroseconds
	LUTC

	Lshortfile
	Llongfile
	Lfileline
	Lpackage
	Lfunc
	LPID

	Lor
	Land
	Lnot

	Lvariants

	Lstdflags = Ldate | Ltime
)

func apply_flags(old, flags uint) uint {
	mod := uint(Lor | Land | Lnot)

	if old == 0 {
		old = Lstdflags
	}

	if flags == 0 {
		// NOP
		flags = old
	} else if flags&mod == 0 {
		// direct set
	} else {
		// apply mask

		if flags&Lor != 0 {
			flags |= old
		} else if flags&Land != 0 {
			flags &= old
		} else if flags&Lnot != 0 {
			flags = old & ^flags
		}

		flags &= ^mod
	}

	return flags
}

//
//
func (ctx *LoggerContext) SetFlags(flags uint) *LoggerContext {
	ctx.flags = apply_flags(ctx.flags, flags)

	return ctx
}

func (ctx *LoggerContext) Flags() uint {
	return ctx.flags
}

// SetFlags of Logger
// zero implies it has to be reset back to what the context wants
// if not, the mask gets applied over the current
func (p *Logger) SetFlags(flags uint) *Logger {
	if flags == 0 {
		// reset
	} else {
		var old uint

		if p.flags == 0 {
			old = p.ctx.flags
		} else {
			old = p.flags
		}

		flags = apply_flags(old, flags)
	}

	p.flags = flags
	return p
}

func (p *Logger) Flags() uint {
	if p.flags == 0 {
		return p.ctx.flags
	} else {
		return p.flags
	}
}
