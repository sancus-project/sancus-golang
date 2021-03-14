package log

import (
	"io"
	"log"
	"strings"
	"testing"
)

func streq(t *testing.T, wants, got string) {
	if wants != got {
		t.Errorf("Expected %q but got %q", wants, got)
	}
}

func newTestLogger(prefix string, flags uint, w io.Writer) *Logger {
	ctx := NewLoggerContext(flags)
	ctx.SetBackend(w)
	return ctx.NewLogger(prefix)
}

func TestStdLogger(t *testing.T) {
	var buf strings.Builder

	logger := newTestLogger("test_prefix: ", Lnoprefix, &buf)

	logger.SetStandard()

	// single
	log.Printf("test:%v", 1)
	streq(t, "test_prefix: test:1\n", buf.String())
	buf.Reset()

	// multicall
	log.Printf("test:%v", 1)
	log.Printf("test:%v", 2)
	streq(t, "test_prefix: test:1\ntest_prefix: test:2\n", buf.String())
	buf.Reset()

	// multiline
	log.Printf("%v\n%v\n%v\n%v", 1, 2, 3, 4)
	streq(t, "test_prefix: 1\ntest_prefix: 2\ntest_prefix: 3\ntest_prefix: 4\n", buf.String())
	buf.Reset()

	// empty
	log.Print("")
	streq(t, "test_prefix:\n", buf.String())
	buf.Reset()

	// only whitespace
	log.Print("  \n    \n\t ")
	streq(t, "test_prefix:\n", buf.String())
	buf.Reset()

	// trailing whitespace
	log.Print("hello    \nworld\t\n\n ")
	streq(t, "test_prefix: hello\ntest_prefix: world\n", buf.String())
	buf.Reset()
}
