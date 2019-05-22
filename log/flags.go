package log

const (
	Lnoprefix = 1 << iota

	Lstdflags = Lnoprefix
)

func apply_flags(old, mask uint) uint {
	if mask == 0 {
		mask = Lstdflags
	}

	return mask
}
