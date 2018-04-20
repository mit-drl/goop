package mip

type Objective struct {
	Expr
	sense ObjSense
}

type ObjSense int

const (
	SenseMinimize ObjSense = 1
	SenseMaximize          = -1
)

func NewObjective(e Expr, sense ObjSense) *Objective {
	return &Objective{e, sense}
}
