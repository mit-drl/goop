package goop

import (
	"github.com/mit-drl/goop/solvers"
)

const (
	tinyNum float64 = 0.01
)

// Solution stores the solution of an optimization problem and associated
// metatdata
type Solution struct {
	vals solvers.DoubleVector

	// The objective for the solution
	Objective float64

	// Whether or not the solution is within the optimality threshold
	Optimal bool

	// The optimality gap returned from the solver. For many solvers, this is
	// the gap between the best possible solution with integer relaxation and
	// the best integer solution found so far.
	Gap float64
}

func newSolution(mipSol solvers.MIPSolution) *Solution {
	return &Solution{
		vals:      mipSol.GetValues(),
		Objective: mipSol.GetObj(),
		Optimal:   mipSol.GetOptimal(),
		Gap:       mipSol.GetGap(),
	}
}

// Value returns the value assigned to the variable in the solution
func (s *Solution) Value(v *Var) float64 {
	return s.vals.Get(int(v.ID()))
}

// IsOne returns true if the value assigned to the variable is an integer,
// and assigned to one. This is a convenience method which should not be
// super trusted...
func (s *Solution) IsOne(v *Var) bool {
	return v.Type() == Integer && s.Value(v) > tinyNum
}
