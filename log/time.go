package log

import (
	"bufio"
	"fmt"
	"time"
)

type TimeContext interface {
	Now() time.Time
}

// writeTimestamp writes a timestamp according to flags and TimeContext
func writeTimestamp(w *bufio.Writer, flags uint, tctx TimeContext) {
	if flags&(Ldate|Ltime|Lseconds) != 0 {
		var use_seconds bool

		var accuracy int
		var dmin time.Duration

		if flags&Lmicroseconds != 0 {
			accuracy = 6
			dmin = time.Microsecond
		} else {
			accuracy = 3
			dmin = time.Millisecond
		}

		if flags&Lseconds != 0 {
			use_seconds = true
		}

		now := tctx.Now()
		if flags&LUTC != 0 {
			now = now.UTC()
		}

		w.WriteRune('[')
		if use_seconds {
			// Seconds
			var sec = now.Unix()
			var subsec = now.Nanosecond() / int(dmin)

			w.WriteString(fmt.Sprintf("%v.%0*v", sec, accuracy, subsec))
		} else {
			// Walltime
			var show_date = (flags&Ldate != 0)
			var show_time = (flags&Ltime != 0)

			if show_date {
				var year, month, day = now.Date()

				// Date
				w.WriteString(fmt.Sprintf("%04v-%02v-%02v", year, int(month), day))

				if show_time {
					w.WriteRune(' ')
				}
			}

			if show_time {
				var hour, min, sec = now.Clock()
				var subsec = now.Nanosecond() / int(dmin)

				w.WriteString(fmt.Sprintf("%02v:%02v:%02v.%0*v", hour, min, sec, accuracy, subsec))
			}
		}
		w.WriteString("] ")
	}
}

//
//
type wallTimeContext struct{}

func (wallTimeContext) Now() time.Time {
	return time.Now()
}

var StdTimeContext TimeContext = &wallTimeContext{}
