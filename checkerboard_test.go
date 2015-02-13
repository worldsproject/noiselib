package noiselib

import "testing"

func TestCheckerboard(t *testing.T) {
  one := Checkerboard{}

  cases := []struct {
    in []float64
    out float64
  } {
    {[]float64{0, 0, 0}, -1.0},
    {[]float64{1, 1, 0}, -1.0},
    {[]float64{3.5, 1.2, 1.0}, 1.0},
  }

  for _, c := range cases {
    got := one.GetValue(c.in[0], c.in[1], c.in[2])

    if got != c.out {
      t.Errorf("Value from %+v should be %v but got %v", c.in, c.out, got)
    }
  }
}
