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

	test := grad.GetColor(-1.0)

	if test.R != 0 || test.G != 0 || test.B != 0 {
		t.Error("For", "grad.GetColor(-1.0)",
			"Expected", black,
			"Got", test)
	}

	test = grad.GetColor(1.0)

	if test.R != 255 || test.G != 255 || test.B != 255 {
		t.Error("For", "grad.GetColor(1.0)",
			"Expected", white,
			"Got", test)
	}

	test = grad.GetColor(0.0)

	if test.R != 127 || test.G != 127 || test.B != 127 {
		t.Error("For", "grad.GetColor(0.0)",
			"Expected", "{128, 128, 128, 155}",
			"Got", test)
	}
}
