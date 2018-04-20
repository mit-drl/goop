package mip

import (
	"github.com/mit-drl/goop/mip/solvers"
)

const (
	tinyNum float64 = 0.01
)

type Solution struct {
	vals      solvers.DoubleVector
	Objective float64
	Optimal   bool
	Gap       float64
}

func NewSolution(mipSol solvers.MIPSolution) *Solution {
	return &Solution{
		vals:      mipSol.GetValues(),
		Objective: mipSol.GetObj(),
		Optimal:   mipSol.GetOptimal(),
		Gap:       mipSol.GetGap(),
	}
}

func (s *Solution) Value(v *Var) float64 {
	return s.vals.Get(int(v.ID()))
}

func (s *Solution) IsOne(v *Var) bool {
	return s.Value(v) > tinyNum
}
