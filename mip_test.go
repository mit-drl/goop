package goop_test

import (
	"fmt"
	"testing"

	"github.com/mit-drl/goop"
	"github.com/mit-drl/goop/solvers"
)

func solveSimpleMIPModel(t *testing.T, solver solvers.Solver) {
	m := goop.NewModel()
	m.ShowLog(false)
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()
	z := m.AddBinaryVar()

	m.AddConstr(goop.Sum(x, y.Mult(2), z.Mult(3)).LessEq(goop.K(4)))
	m.AddConstr(goop.Sum(x, y).GreaterEq(goop.One))

	obj := goop.Sum(x, y, z.Mult(2))
	m.SetObjective(obj, goop.SenseMaximize)
	sol, err := m.Optimize(solver)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("x =", sol.Value(x))
	t.Log("y =", sol.Value(y))
	t.Log("z =", sol.Value(z))
}

func solveSumRowsColsModel(t *testing.T, solver solvers.Solver) {
	m := goop.NewModel()
	m.ShowLog(false)
	rows := 4
	cols := 4
	vs := m.AddBinaryVarMatrix(rows, cols)

	for i := 0; i < cols; i++ {
		m.AddConstr(goop.SumCol(vs, i).Eq(goop.One))
	}

	for i := 0; i < rows; i++ {
		m.AddConstr(goop.SumRow(vs, i).Eq(goop.One))
	}

	sol, err := m.Optimize(solver)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prettyPrintVarMatrix(vs, sol))
}

func prettyPrintVarMatrix(vs [][]*goop.Var, sol *goop.Solution) string {
	rows := len(vs)
	cols := len(vs[0])

	matStr := ""
	for i := 0; i < rows; i++ {
		rowStr := ""
		for j := 0; j < cols; j++ {
			if sol.Value(vs[i][j]) > 0.1 {
				rowStr += "1 "
			} else {
				rowStr += "0 "
			}
		}
		matStr += rowStr + "\n"
	}

	return matStr
}
