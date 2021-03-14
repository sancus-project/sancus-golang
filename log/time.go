package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type TimeContext interface {
	Now() time.Time
	StartTime() time.Time
	LastWrite() time.Time
	SetLastWrite(time.Time)
}

// writeTimestamp writes a timestamp according to flags and TimeContext
func writeTimestamp(w *bufio.Writer, flags uint, tctx TimeContext) {
	if flags&LPID != 0 {
		w.WriteString(fmt.Sprintf("%v: ", os.Getpid()))
	}

	if flags&(Ldate|Ltime|Lseconds|Lrelative) != 0 {
		var use_seconds bool
		var use_elapsed bool
		var use_relative bool

		var accuracy int
		var dmin, elapsed, relative time.Duration

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

		if flags&Lrelative != 0 {
			use_relative = true
		}

		if flags&Lseconds != 0 {
			use_seconds = true
		}

		tctx.SetLastWrite(now)
		if use_relative {
			relative = now.Sub(tctx.StartTime()).Round(dmin)
		} else if flags&LUTC != 0 {
			now = now.UTC()
		}

		w.WriteRune('[')
		if use_relative {
			// Relative to application start
			var sec = int(relative / time.Second)
			var subsec = int(relative%time.Second) / int(dmin)

			if use_seconds {
				w.WriteString(fmt.Sprintf("%v.%0*v", sec, accuracy, subsec))
			} else {
				var hour, min int

				min = sec / 60
				hour = min / 60
				min = min % 60
				sec = sec % 60

				w.WriteString(fmt.Sprintf("%02v:%02v:%02v.%0*v", hour, min, sec, accuracy, subsec))
			}
		} else if use_seconds {
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
	startTime time.Time // relative
	lastWrite time.Time // elapsed
}

func (wallTimeContext) Now() time.Time {
	return time.Now()
}

func (tctx *wallTimeContext) StartTime() time.Time {
	return tctx.startTime
}

func (tctx *wallTimeContext) LastWrite() time.Time {
	return tctx.lastWrite
}

func (tctx *wallTimeContext) SetLastWrite(t time.Time) {
	tctx.lastWrite = t
}

var StdTimeContext TimeContext = &wallTimeContext{
	startTime: time.Now(),
}
