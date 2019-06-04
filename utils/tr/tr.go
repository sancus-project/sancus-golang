package tr

import (
	"fmt"
)

func translateRune(in rune, s0, s1 []rune) rune {
	for i, c := range s0 {
		if in == c {
			// match
			return s1[i]
		}
	}
	return in
}

func translateRunes(in, s0, s1 []rune) (out []rune, changed bool) {
	if len(in) == 0 {
		return
	} else if len(s0) != len(s1) {
		panic(fmt.Sprintf("Set lengths differ"))
	}

	out = make([]rune, len(in))
	for i, c := range in {
		if c1 := translateRune(c, s0, s1); c1 != c {
			changed = true
			c = c1
		}

		out[i] = c
	}
	if !changed {
		out = in
	}

	return
}

func TranslateRunes(in, s0, s1 string) string {
	if out, changed := translateRunes([]rune(in), []rune(s0), []rune(s1)); changed {
		return string(out)
	} else {
		return in
	}
}
