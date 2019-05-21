package log // import "github.com/amery/go-misc/log"

import (
	"sync"
)

type Logger struct {
	mu     sync.Mutex
	prefix string
}
