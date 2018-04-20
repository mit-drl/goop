package mip_test

import (
	"fmt"
	"testing"

	"github.com/mit-drl/goop/mip"
	"github.com/mit-drl/goop/mip/solvers"
)

func TestSimpleMIP(t *testing.T) {
	m := mip.NewModel()
	m.ShowLog(true)
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()
	z := m.AddBinaryVar()

	m.AddConstr(mip.Sum(x, y.Mult(2), z.Mult(3)).LessEq(mip.K(4)))
	m.AddConstr(mip.Sum(x, y).GreaterEq(mip.One))

	obj := mip.Sum(x, y, z.Mult(2))
	m.SetObjective(obj, mip.SenseMaximize)
	sol, err := m.Optimize(solvers.Gurobi)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("x =", sol.Value(x))
	t.Log("y =", sol.Value(y))
	t.Log("z =", sol.Value(z))
}

func TestSumRowsCols(t *testing.T) {
	m := mip.NewModel()
	m.ShowLog(true)
	rows := 4
	cols := 4
	vs := m.AddBinaryVarMatrix(rows, cols)

	for i := 0; i < cols; i++ {
		m.AddConstr(mip.SumCol(vs, i).Eq(mip.One))
	}

	for i := 0; i < rows; i++ {
		m.AddConstr(mip.SumRow(vs, i).Eq(mip.One))
	}

	sol, err := m.Optimize(solvers.Gurobi)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prettyPrintVarMatrix(vs, sol))
}

func prettyPrintVarMatrix(vs [][]*mip.Var, sol *mip.Solution) string {
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
