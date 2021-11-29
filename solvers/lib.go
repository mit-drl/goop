package solvers

// #cgo CXXFLAGS: --std=c++11 -I/Library/gurobi912/mac64/include -I/Users/kwesirutledge/Documents/Michigan/Research/goop/include
// #cgo CXXFLAGS: -I/usr/local/opt/lp_solve/include
// #cgo LDFLAGS: -L/usr/local/opt/lp_solve/lib -llpsolve55
// #cgo LDFLAGS: -L/Library/gurobi912/mac64/lib -lgurobi_c++ -lgurobi91
import "C"
