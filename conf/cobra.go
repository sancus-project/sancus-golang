package conf

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

type cobraFlag struct {
	f   *pflag.Flag
	v   interface{}
	out interface{}
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

	values map[string]*cobraFlag
	set    *pflag.FlagSet
}

func NewCobraMapper(set *pflag.FlagSet) *CobraMapper {
	if set != nil {
		m := &CobraMapper{
			set:    set,
			values: make(map[string]*cobraFlag),
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

func (m *CobraMapper) addFlag(name string, v interface{}, out interface{}) *CobraMapper {
	p := &cobraFlag{
		f:   m.set.Lookup(name),
		v:   v,
		out: out,
	}
	m.values[name] = p
	return m
}

func (m *CobraMapper) Parse() {
	for _, p := range m.values {
		if p.out == nil {
			// skip
		} else if v, ok := p.GetUint16(); ok {
			// Uint16
			out := p.out.(*uint16)
			*out = v
		} else if v, ok := p.GetDuration(); ok {
			// Duration
			out := p.out.(*time.Duration)
			*out = v
		}
	}
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

func (m *CobraMapper) uint16VarP(out *uint16, v *uint16, name, shorthand string, value uint16, usage string, args ...interface{}) *CobraMapper {
	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = value
	m.set.Uint16VarP(v, name, shorthand, value, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) UintVar16P(out *uint16, name, shorthand string, value uint16, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	return m.uint16VarP(out, v, name, shorthand, value, usage, args...)
}

func (m *CobraMapper) Uint16P(name, shorthand string, value uint16, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	return m.uint16VarP(nil, v, name, shorthand, value, usage, args...)
}

// Duration
func (f cobraFlag) GetDuration() (time.Duration, bool) {
	if p, ok := f.Raw().(*time.Duration); ok {
		ok = f.Changed()
		return *p, ok
	} else {
		return 0, false
	}
}

func (m *CobraMapper) durationVarP(out *time.Duration, v *time.Duration, name, shorthand string, value time.Duration, usage string, args ...interface{}) *CobraMapper {
	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = value
	m.set.DurationVarP(v, name, shorthand, value, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) DurationVarP(out *time.Duration, name, shorthand string, value time.Duration, usage string, args ...interface{}) *CobraMapper {
	v := new(time.Duration)
	return m.durationVarP(out, v, name, shorthand, value, usage, args...)
}
func (m *CobraMapper) DurationP(name, shorthand string, value time.Duration, usage string, args ...interface{}) *CobraMapper {
	v := new(time.Duration)
	return m.durationVarP(nil, v, name, shorthand, value, usage, args...)
}
