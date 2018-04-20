package solvers

import "github.com/sirupsen/logrus"

// SolverType represnts the type of solver such as Gurobi or LPSolve
type SolverType int

// Supported solvers for quick instantiation
const (
	Gurobi SolverType = iota
)

// NewSolver returns an instantiated solver given the solver type
func NewSolver(solverType SolverType) Solver {
	switch solverType {
	case Gurobi:
		return NewGurobiSolver()
	default:
		logrus.WithField("solverType", solverType).Panic("Solver type unknown")
		return nil
	}
}
