package noiselib

import "testing"

func TestClampValue(t *testing.T) {
	lower, upper := 2, 5

	cases := []struct {
		in, out int
	}{
		{1, 2},
		{4, 4},
		{10, 5},
	}

	for _, c := range cases {
		got := ClampValue(c.in, lower, upper)

		if got != c.out {
			t.Errorf("Value from %+v should be %v but got %v", c.in, c.out, got)
		}
	}
}
