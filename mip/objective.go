package mip

// Objective represents an optimization objective given an expression and
// objective sense (maximize or minimize).
type Objective struct {
	Expr
	sense ObjSense
}

// NewObjective returns a new optimization objective given an expression and
// objective sense
func NewObjective(e Expr, sense ObjSense) *Objective {
	return &Objective{e, sense}
}

// ObjSense represents whether an optimization objective is to be maximized or
// minimized. This implementation conforms to the Gurobi encoding
type ObjSense int

// Objective senses (minimize and maximize) encoding using Gurobi's standard
const (
	SenseMinimize ObjSense = 1
	SenseMaximize          = -1
)
