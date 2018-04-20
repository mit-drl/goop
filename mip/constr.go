package mip

type Constr struct {
	lhs   Expr
	rhs   Expr
	sense ConstrSense
}

type ConstrSense byte

const (
	SenseEqual            ConstrSense = '='
	SenseLessThanEqual                = '<'
	SenseGreaterThanEqual             = '>'
)

func LessThanEqual(lhs, rhs Expr) *Constr {
	return &Constr{lhs, rhs, SenseLessThanEqual}
}

func Equal(lhs, rhs Expr) *Constr {
	return &Constr{lhs, rhs, SenseEqual}
}

func GreaterThanEqual(lhs, rhs Expr) *Constr {
	return &Constr{lhs, rhs, SenseGreaterThanEqual}
}
