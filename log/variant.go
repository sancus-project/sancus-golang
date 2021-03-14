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

// SetErrorVariant
func (logger *Logger) SetErrorVariant(v Variant) *Logger {
	logger.ctx.SetErrorVariant(v)
	return logger
}

func (ctx *LoggerContext) SetErrorVariant(v Variant) *LoggerContext {
	ctx.errorVariant = v
	return ctx
}

// ErrorVariant
func (logger *Logger) ErrorVariant() Variant {
	return logger.ctx.ErrorVariant()
}

func (ctx *LoggerContext) ErrorVariant() Variant {
	return ctx.errorVariant
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

// VariantEnabled
func (logger *Logger) VariantEnabled(k Variant) bool {
	return logger.ctx.VariantEnabled(k, logger.Flags())
}

func (ctx *LoggerContext) VariantEnabled(k Variant, mask uint) bool {
	if mask == 0 {
		mask = ctx.Flags()
	}

	n := uint(k)
	if n == 0 || k == ctx.errorVariant || (n >= Lvariants && (n&mask) == n) {
		return true
	}
	return false
}
