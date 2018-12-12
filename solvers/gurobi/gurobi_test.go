package gurobi

import (
	"testing"
)

func TestSwig(t *testing.T) {
	s := NewWrapped_GurobiSolver()
	DeleteWrapped_GurobiSolver(s)
}
