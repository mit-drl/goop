// +build !pre_xenial

package solvers

// #cgo CXXFLAGS: --std=c++11 -I${SRCDIR}/../.third_party/gurobi/include
// #cgo CXXFLAGS: -I${SRCDIR}/../.third_party/lpsolve
// #cgo LDFLAGS: -L${SRCDIR}/../.third_party/gurobi/lib -lgurobi_g++5.2 -lgurobi75
// #cgo LDFLAGS: -L${SRCDIR}/../.third_party/lpsolve -llpsolve55
import "C"
