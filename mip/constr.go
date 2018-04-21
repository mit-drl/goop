package mip

// Constr represnts a linear constraint of the form x <= y, x >= y, or
// x == y. Constr uses a left and right hand side expressions along with a
// constraint sense (<=, >=, ==) to represent a generalized linear constraint
type Constr struct {
	lhs   Expr
	rhs   Expr
	sense ConstrSense
}

// LessEq returns a constraint representing lhs <= rhs
func LessEq(lhs, rhs Expr) *Constr {
	return &Constr{lhs, rhs, SenseLessThanEqual}
}

// Eq returns a constraint representing lhs == rhs
func Eq(lhs, rhs Expr) *Constr {
	return &Constr{lhs, rhs, SenseEqual}
}

// GreaterEq returns a constraint representing lhs >= rhs
func GreaterEq(lhs, rhs Expr) *Constr {
	return &Constr{lhs, rhs, SenseGreaterThanEqual}
}

// ConstrSense represents if the constraint x <= y, x >= y, or x == y. For easy
// integration with Gurobi, the senses have been encoding using a byte in
// the same way Gurobi encodes the constraint senses.
type ConstrSense byte

// Different constraint senses conforming to Gurobi's encoding.
const (
	SenseEqual            ConstrSense = '='
	SenseLessThanEqual                = '<'
	SenseGreaterThanEqual             = '>'
)
