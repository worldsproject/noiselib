package noiselib

import (
	"fmt"
	"testing"
)

func TestLinearInterp(t *testing.T) {
	n0, n1, a := 1.0, 2.0, 3.0
	answer := 4.0

	test := LinearInterp(n0, n1, a)

	if test != answer {
		t.Error("For", fmt.Sprintf("n0: %v, n1: %v, a: %v", n0, n1, a),
			"Expected", answer,
			"Got", test)
	}

	a = 0.5
	answer = 1.5

	test = LinearInterp(n0, n1, a)

	if test != answer {
		t.Error("For", fmt.Sprintf("n0: %v, n1: %v, a: %v", n0, n1, a),
			"Expected", answer,
			"Got", test)
	}
}
