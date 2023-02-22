package boolal

const (
	truth   = "■"
	falsity = "□"
)

func Truth() UnaryOp {
	return Var(truth)
}

func Falsity() UnaryOp {
	return Var(falsity)
}

func asUnaryOp(term any) UnaryOp {
	switch v := term.(type) {
	case string:
		return UnaryOp{Var: v}

	case UnaryOp:
		return v

	case *Expr:
		if v.Rhs == nil {
			return v.Lhs
		} else {
			return UnaryOp{Al: v}
		}
	default:
		panic("invalid value")
	}
}

func asUnaryOps(term any, terms []any) []UnaryOp {
	vs := make([]UnaryOp, len(terms)+1)
	vs[0] = asUnaryOp(term)
	for i, t := range terms {
		vs[i+1] = asUnaryOp(t)
	}

	return vs
}

func Var(v string) UnaryOp {
	return UnaryOp{Var: v}
}

func Not(term any) UnaryOp {
	op := asUnaryOp(term)
	if op.Name == "" {
		op.Name = "!"
	} else {
		op.Name = ""
	}

	return op
}

func And(lhs any, rhs any, terms ...any) *Expr {
	return &Expr{
		Lhs: asUnaryOp(lhs),
		Rhs: []BinaryOp{{And: asUnaryOps(rhs, terms)}},
	}
}

func Or(lhs any, rhs any, terms ...any) *Expr {
	return &Expr{
		Lhs: asUnaryOp(lhs),
		Rhs: []BinaryOp{{Or: asUnaryOps(rhs, terms)}},
	}
}

func (a *Expr) concat(b BinaryOp) {
	if a.Rhs == nil {
		a.Rhs = make([]BinaryOp, 0, 1)
	}
	a.Rhs = append(a.Rhs, b)
}

func (a *Expr) And(term any, terms ...any) *Expr {
	a.concat(BinaryOp{And: asUnaryOps(term, terms)})

	return a
}

func (a *Expr) Or(term any, terms ...any) *Expr {
	a.concat(BinaryOp{Or: asUnaryOps(term, terms)})

	return a
}
