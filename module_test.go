package noiselib

import "testing"

func TestClampValue(t *testing.T) {
	lower, upper := 2, 5

	test := 1

	answer := 2

	if ClampValue(test, lower, upper) != answer {
		t.Error("For", test,
			"Expected", answer,
			"Got", ClampValue(test, lower, upper))
	}

	test = 4
	answer = 4

	if ClampValue(test, lower, upper) != answer {
		t.Error("For", test,
			"Expected", answer,
			"Got", ClampValue(test, lower, upper))
	}

	test = 10
	answer = 5

	if ClampValue(test, lower, upper) != answer {
		t.Error("For", test,
			"Expected", answer,
			"Got", ClampValue(test, lower, upper))
	}
}
