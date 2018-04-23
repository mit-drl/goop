package goop_test

import (
	"testing"

	"github.com/mit-drl/goop/solvers"
)

func TestLPSolve(t *testing.T) {
	t.Run("SimpleMIP", func(t *testing.T) {
		solveSimpleMIPModel(t, solvers.NewLPSolveSolver())
	})

	t.Run("SumRowsCols", func(t *testing.T) {
		solveSumRowsColsModel(t, solvers.NewLPSolveSolver())
	})
}
