package goop_test

import (
	"testing"

	"github.com/mit-drl/goop"
)

func TestLinearExprCoeffsAndConstant(t *testing.T) {
	m := goop.NewModel()
	x := m.AddBinaryVar()
	y := m.AddBinaryVar()

	// 2 * x + 4 * y - 5
	coeffs := []float64{2, 4}
	constant := -5.0
	expr := goop.Sum(x.Mult(coeffs[0]), y.Mult(coeffs[1]), goop.K(constant))

	for i, coeff := range expr.Coeffs() {
		if coeffs[i] != coeff {
			t.Errorf("Coeff mismatch: %v != %v", coeff, coeffs[i])
		}
	}

	if expr.Constant() != constant {
		t.Errorf("Constant mismatch: %v != %v", expr.Constant(), constant)
	}
}
