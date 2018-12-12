package gurobi

// #cgo CXXFLAGS: --std=c++11 -I${SRCDIR}/../../.third_party/gurobi/include
// #cgo CXXFLAGS: -I${SRCDIR}/..
// #cgo LDFLAGS: -L${SRCDIR}/../../.third_party/gurobi/lib -lgurobi_g++5.2 -lgurobi75
import "C"
import (
	"github.com/mit-drl/goop/solvers"
)

var (
	_ solvers.Solver = NewGurobiSolver()
)
