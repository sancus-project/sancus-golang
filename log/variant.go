package log

// variant identifier
type Variant uint

type loggerVariant struct {
	prefix string
	flags  uint
}

// SetDefaultVariant
func (logger *Logger) SetDefaultVariant(v Variant) *Logger {
	logger.ctx.SetDefaultVariant(v)
	return logger
}

func (ctx *LoggerContext) SetDefaultVariant(v Variant) *LoggerContext {
	ctx.defaultVariant = v
	return ctx
}

// DefaultVariant
func (logger *Logger) DefaultVariant() Variant {
	return logger.ctx.DefaultVariant()
}

func (ctx *LoggerContext) DefaultVariant() Variant {
	return ctx.defaultVariant
}

// SetVariant
func (logger *Logger) SetVariant(k Variant, prefix string, flags uint) *Logger {
	logger.ctx.SetVariant(k, prefix, flags)
	return logger
}

func (ctx *LoggerContext) SetVariant(k Variant, prefix string, flags uint) *LoggerContext {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	ctx.variants[k] = loggerVariant{
		prefix: prefix,
		flags:  flags,
	}

	return ctx
}

// RemoveVariant
func (logger *Logger) RemoveVariant(k Variant) *Logger {
	logger.ctx.RemoveVariant(k)
	return logger
}

func (ctx *LoggerContext) RemoveVariant(k Variant) *LoggerContext {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	delete(ctx.variants, k)
	return ctx
}

// GetVariantPrefixFlags
func (ctx *LoggerContext) GetVariantPrefixFlags(k Variant, flags uint) (string, uint) {
	if v, ok := ctx.variants[k]; ok {
		return v.prefix, apply_flags(flags, v.flags)
	} else {
		return "", flags
	}
}
