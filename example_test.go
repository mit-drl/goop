package goop_test

import (
	"fmt"

	"github.com/mit-drl/goop"
	"github.com/mit-drl/goop/solvers"
)

// This example shows how goop can be used to solve a simple MIP:
//  // maximize    x +   y + 2 z
//  // subject to  x + 2 y + 3 z <= 4
//  //             x +   y       >= 1
//  // x, y, z binary
// MIP being modelled is the same as in http://www.gurobi.com/documentation/7.5/examples/mip1_cpp_cpp.html
func ExampleModel_simple() {
	// Instantiate a new model
	m := goop.NewModel()

	// Add your variables to the model
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()
	z := m.AddBinaryVar()

	// Add your constraints
	m.AddConstr(goop.Sum(x, y.Mult(2), z.Mult(3)).LessEq(goop.K(4)))
	m.AddConstr(goop.Sum(x, y).GreaterEq(goop.One))

	// Set a linear objective using your variables
	obj := goop.Sum(x, y, z.Mult(2))
	m.SetObjective(obj, goop.SenseMaximize)

	// Optimize the variables according to the model
	sol, err := m.Optimize(solvers.NewLPSolveSolver())

	// Check if there is an error from the solver. No error should be returned
	// for this model
	if err != nil {
		panic("Should not have an error")
	}

	// Print out the solution
	fmt.Println("x =", sol.Value(x))
	fmt.Println("y =", sol.Value(y))
	fmt.Println("z =", sol.Value(z))

	// Output:
	// x = 1
	// y = 0
	// z = 1
}
