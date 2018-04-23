package goop

// Integer constants represnting commonly used numbers. Makes for better
// readability
const (
	Zero = K(0)
	One  = K(1)
)

// K is a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
type K float64

// NumVars returns the number of variables in the expression. For constants,
// this is always 0
func (c K) NumVars() int {
	return 0
}

// Vars returns a slice of the Var ids in the expression. For constants,
// this is always nil
func (c K) Vars() []uint64 {
	return nil
}

// Coeffs returns a slice of the coefficients in the expression. For constants,
// this is always nil
func (c K) Coeffs() []float64 {
	return nil
}

// Constant returns the constant additive value in the expression. For
// constants, this is just the constants value
func (c K) Constant() float64 {
	return float64(c)
}

// Plus adds the current expression to another and returns the resulting
// expression
func (c K) Plus(e Expr) Expr {
	newExpr := new(LinearExpr)
	newExpr.vars = append([]uint64{}, e.Vars()...)
	newExpr.coeffs = append([]float64{}, e.Coeffs()...)
	newExpr.constant = e.Constant() + c.Constant()
	return newExpr
}

// Mult multiplies the current expression to another and returns the
// resulting expression
func (c K) Mult(val float64) Expr {
	return K(float64(c) * val)
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (c K) LessEq(other Expr) *Constr {
	return LessEq(c, other)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (c K) GreaterEq(other Expr) *Constr {
	return GreaterEq(c, other)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (c K) Eq(other Expr) *Constr {
	return Eq(c, other)
}
