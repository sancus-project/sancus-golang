package conf

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

type cobraFlag struct {
	f *pflag.Flag
	v interface{}
}

func (f cobraFlag) Changed() bool {
	return f.f.Changed
}

func (f cobraFlag) Flag() *pflag.Flag {
	return f.f
}

func (f cobraFlag) Raw() interface{} {
	return f.v
}

type CobraMapper struct {
	mapper

	values map[string]Flag
	set    *pflag.FlagSet
}

func NewCobraMapper(set *pflag.FlagSet) *CobraMapper {
	if set != nil {
		m := &CobraMapper{
			set:    set,
			values: make(map[string]Flag),
		}
		m.mapper = m.Lookup
		registerMapper(set, m)
		return m
	}
	return nil
}

func (m *CobraMapper) Lookup(name string) Flag {
	if v, ok := m.values[name]; ok {
		return v
	}
	return nil
}

func (m *CobraMapper) addFlag(name string, v interface{}) *CobraMapper {
	p := &cobraFlag{
		f: m.set.Lookup(name),
		v: v,
	}
	m.values[name] = p
	return m
}

// Uint16
func (f cobraFlag) GetUint16() (uint16, bool) {
	if p, ok := f.Raw().(*uint16); ok {
		ok = f.Changed()
		return *p, ok
	} else {
		return 0, false
	}
}

func (m *CobraMapper) Uint16P(name, shorthand string, value uint16, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = value
	m.set.Uint16VarP(v, name, shorthand, value, usage)
	return m.addFlag(name, v)
}

// Duration
func (f cobraFlag) Duration() (time.Duration, bool) {
	if p, ok := f.Raw().(*time.Duration); ok {
		ok = f.Changed()
		return *p, ok
	} else {
		return 0, false
	}
}

func (m *CobraMapper) DurationP(name, shorthand string, value time.Duration, usage string, args ...interface{}) *CobraMapper {
	v := new(time.Duration)
	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = value
	m.set.DurationVarP(v, name, shorthand, value, usage)
	return m.addFlag(name, v)
}
