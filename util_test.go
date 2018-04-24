package goop_test

import (
	"testing"

	"github.com/mit-drl/goop"
)

func TestDot(t *testing.T) {
	N := 10
	m := goop.NewModel()
	xs := m.AddBinaryVarVector(N)
	coeffs := make([]float64, N)

	for i := 0; i < N; i++ {
		coeffs[i] = float64(i + 1)
	}

	expr := goop.Dot(xs, coeffs)

	for i, coeff := range expr.Coeffs() {
		if coeffs[i] != coeff {
			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
		}
	}

	if expr.Constant() != 0 {
		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
	}
}

func TestDotPanic(t *testing.T) {
	N := 10
	m := goop.NewModel()
	xs := m.AddBinaryVarVector(N)
	coeffs := make([]float64, N-1)

	for i := 0; i < N-1; i++ {
		coeffs[i] = float64(i + 1)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Coeff size mismatch: Code did not panic")
		}
	}()

	goop.Dot(xs, coeffs)
}

func TestSumVars(t *testing.T) {
	numVars := 3
	m := goop.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()
	z := m.AddBinaryVar()
	expr := goop.SumVars(x, y, z)

	for _, coeff := range expr.Coeffs() {
		if coeff != 1 {
			t.Errorf("Coeff mismatch: %v != 1", coeff)
		}
	}

	if expr.NumVars() != numVars {
		t.Errorf("NumVars mismatch: %v != %v", expr.NumVars(), numVars)
	}

	if expr.Constant() != 0 {
		t.Errorf("Constant mismatch: %v != 0", expr.Constant())
	}
}
