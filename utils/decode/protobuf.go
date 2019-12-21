package decode

import (
	"bufio"
	"errors"
	"io"

	"github.com/gogo/protobuf/proto"
)

var (
	ErrorInvalidProtobuf = errors.New("Invalid Protobuf data")
)

func NewProtobufScanner(in io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(in)
	s.Split(ScanProtobuf)
	return s
}

func ScanProtobuf(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if x, n := proto.DecodeVarint(data); n > 0 {
		wire := x & 0b111 // wire type

		switch wire {
		case 0: // Varint
			advance = n
		case 5: // fixed32
			advance = n + 4
		case 1: // fixed64
			advance = n + 8
		case 2: // length delimited
			if len(data) > n {
				// payload length
				plen := int(data[n : n+1][0])

				// message length
				advance = n + 1 + plen
			}
		default:
			err = ErrorInvalidProtobuf
		}

	} else if atEOF {
		// incomplete and closed.
		err = io.EOF
	} else if len(data) > 4 {
		// message too long to be credible
		err = ErrorInvalidProtobuf
	}

	if err != nil {
		// failed. discard all.
		advance = len(data)
	} else if advance > len(data) {
		// incomplete. wait.
		advance = 0
	} else if advance > 0 {
		// all bytes accounted for
		token = data[:advance]
	}

	return
}
