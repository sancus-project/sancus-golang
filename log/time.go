package log

import (
	"bufio"
	"fmt"
	"time"
)

type TimeContext interface {
	Now() time.Time
	LastWrite() time.Time
	SetLastWrite(time.Time)
}

// writeTimestamp writes a timestamp according to flags and TimeContext
func writeTimestamp(w *bufio.Writer, flags uint, tctx TimeContext) {
	if flags&(Ldate|Ltime|Lseconds) != 0 {
		var use_seconds bool
		var use_elapsed bool

		var accuracy int
		var dmin, elapsed time.Duration

		if flags&Lmicroseconds != 0 {
			accuracy = 6
			dmin = time.Microsecond
		} else {
			accuracy = 3
			dmin = time.Millisecond
		}

		now := tctx.Now()
		if flags&Lelapsed != 0 {
			use_elapsed = true

			if t0 := tctx.LastWrite(); !t0.IsZero() {
				elapsed = now.Sub(t0).Round(dmin)
			}
		}

		if flags&Lseconds != 0 {
			use_seconds = true
		}

		tctx.SetLastWrite(now)
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

		if use_elapsed {
			w.WriteString(fmt.Sprintf(" +%.*f", accuracy, elapsed.Seconds()))
		}
		w.WriteString("] ")
	}
}

//
//
type wallTimeContext struct {
	lastWrite time.Time // elapsed
}

func (wallTimeContext) Now() time.Time {
	return time.Now()
}

func (tctx *wallTimeContext) LastWrite() time.Time {
	return tctx.lastWrite
}

func (tctx *wallTimeContext) SetLastWrite(t time.Time) {
	tctx.lastWrite = t
}

var StdTimeContext TimeContext = &wallTimeContext{}
