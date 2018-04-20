package mip

type Var struct {
	id    uint64
	lower float64
	upper float64
	vtype VarType
}

func (v *Var) NumVars() int {
	return 1
}

func (v *Var) Vars() []uint64 {
	return []uint64{v.id}
}

func (v *Var) Coeffs() []float64 {
	return []float64{1}
}

func (v *Var) Constant() float64 {
	return 0
}

func (v *Var) Plus(e Expr) Expr {
	vars := append([]uint64{v.id}, e.Vars()...)
	coeffs := append([]float64{1}, e.Coeffs()...)
	newExpr := &LinearExpr{
		vars:     vars,
		coeffs:   coeffs,
		constant: e.Constant(),
	}
	return newExpr
}

func (v *Var) Mult(c float64) Expr {
	vars := []uint64{v.id}
	coeffs := []float64{c}
	newExpr := &LinearExpr{
		vars:     vars,
		coeffs:   coeffs,
		constant: 0,
	}
	return newExpr
}

func (v *Var) LessEq(other Expr) *Constr {
	return LessThanEqual(v, other)
}

func (v *Var) GreaterEq(other Expr) *Constr {
	return GreaterThanEqual(v, other)
}

func (v *Var) Eq(other Expr) *Constr {
	return Equal(v, other)
}

func (v *Var) ID() uint64 {
	return v.id
}

func (v *Var) Lower() float64 {
	return v.lower
}

func (v *Var) Upper() float64 {
	return v.upper
}

func (v *Var) Type() VarType {
	return v.vtype
}

type VarType byte

const (
	Continuous VarType = 'C'
	Binary             = 'B'
	Integer            = 'I'
)
