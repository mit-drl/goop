package goop_test

import (
	"testing"

	"github.com/mit-drl/goop"
)

func TestConstantPlus(t *testing.T) {
	two := goop.K(2)
	three := goop.K(3)
	five := goop.Sum(two, three)

	if five.Constant() != 5 {
		t.Errorf("five (%v) not equal to 5", five)
	}
}

func TestConstantMult(t *testing.T) {
	two := goop.K(2)
	six := two.Mult(3)

	if six.Constant() != 6 {
		t.Errorf("six (%v) not equal to 6", six)
	}
}
