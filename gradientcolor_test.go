package noiselib

import (
	"image/color"
	"testing"
)

func TestGradientColor(t *testing.T) {
	grad := GradientColor{}
	grad.ClearGradient()
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	grad.AddGradientPoint(-1.0, black)
	grad.AddGradientPoint(1.0, white)

	cases := []struct {
		in  float64
		out color.RGBA
	}{
		{-1.0, color.RGBA{0, 0, 0, 255}},
		{1.0, color.RGBA{255, 255, 255, 255}},
		{0.5, color.RGBA{191, 191, 191, 255}},
		{0.0, color.RGBA{127, 127, 127, 255}},
		{-5.0, color.RGBA{0, 0, 0, 255}},
		{6.0, color.RGBA{255, 255, 255, 255}},
	}

	for _, c := range cases {
		got := grad.GetColor(c.in)

		if got != c.out {
			t.Errorf("Value from %v should be %v but got %v", c.in, c.out, got)
		}
	}
}
