package boolal

type EvalOption struct {
	FallbackValue bool
}

type EvalOptionModifier func(o *EvalOption)

func WithOption(opt EvalOption) EvalOptionModifier {
	return func(o *EvalOption) { *o = opt }
}

func WithFallbackAs(v bool) EvalOptionModifier {
	return func(o *EvalOption) { o.FallbackValue = v }
}

func eval(data map[string]bool, op UnaryOp, opt EvalOption) bool {
	neg := op.Name != ""

	if op.Al != nil {
		return neg != op.Al.Eval(data, WithOption(opt))
	}

	v, ok := data[op.Var]
	if !ok {
		v = opt.FallbackValue
	}

	return neg != v
}

func (e *Expr) Eval(data map[string]bool, modifiers ...EvalOptionModifier) bool {
	opt := EvalOption{
		FallbackValue: false,
	}

	for _, m := range modifiers {
		m(&opt)
	}

	rst := eval(data, e.Lhs, opt)
	for _, next := range e.Rhs {
		var (
			ops []UnaryOp = nil
			neg bool
		)

		if len(next.And) > 0 {
			neg = true
			ops = next.And
		} else {
			neg = false
			ops = next.Or
		}

		if len(ops) == 0 {
			continue
		}

		if neg != rst {
			rst = !neg
			continue
		}

		for _, v := range ops {
			if neg != eval(data, v, opt) {
				rst = !neg
				break
			}
		}
	}

	return rst
}
