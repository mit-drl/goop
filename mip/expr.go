package mip

// Expr represents a linear general expression of the form
// c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
// variables and k is a constant
type Expr interface {
	// NumVars returns the number of variables in the expression
	NumVars() int

	// Vars returns a slice of the Var ids in the expression
	Vars() []uint64

	// Coeffs returns a slice of the coefficients in the expression
	Coeffs() []float64

	// Constant returns the constant additive value in the expression
	Constant() float64

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e Expr) Expr

	// Mult multiplies the current expression to another and returns the
	// resulting expression
	Mult(c float64) Expr

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(e Expr) *Constr

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(e Expr) *Constr

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(e Expr) *Constr
}

type LinearExpr struct {
	vars     []uint64
	coeffs   []float64
	constant float64
}

// NewExpr returns a new expression with a single additive constant value, c,
// and no variables. Creating an expression like sum := NewExpr(0) is useful
// for creating new empty expressions that you can perform operatotions on
// later
func NewExpr(c float64) Expr {
	return &LinearExpr{constant: c}
}

func (e *LinearExpr) NumVars() int {
	return len(e.vars)
}

func (e *LinearExpr) Vars() []uint64 {
	return e.vars
}

func (e *LinearExpr) Coeffs() []float64 {
	return e.coeffs
}

func (e *LinearExpr) Constant() float64 {
	return e.constant
}

func (e *LinearExpr) Plus(other Expr) Expr {
	e.vars = append(e.vars, other.Vars()...)
	e.coeffs = append(e.coeffs, other.Coeffs()...)
	e.constant += other.Constant()
	return e
}

func (e *LinearExpr) Mult(c float64) Expr {
	for i, coeff := range e.coeffs {
		e.coeffs[i] = coeff * c
	}
	e.constant *= c

	return e
}

func (e *LinearExpr) LessEq(other Expr) *Constr {
	return LessThanEqual(e, other)
}

func (e *LinearExpr) GreaterEq(other Expr) *Constr {
	return GreaterThanEqual(e, other)
}

func (e *LinearExpr) Eq(other Expr) *Constr {
	return Equal(e, other)
}

func getVarsPtr(e Expr) *uint64 {
	if e.NumVars() > 0 {
		return &e.Vars()[0]
	}

	return nil
}

func getCoeffsPtr(e Expr) *float64 {
	if e.NumVars() > 0 {
		return &e.Coeffs()[0]
	}

	return nil
}
