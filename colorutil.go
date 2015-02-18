package noiselib

import (
	"image/color"
)

func BlendChannel(channel0, channel1 uint8, alpha float64) uint8 {
	c0 := float64(channel0) / 255.0
	c1 := float64(channel1) / 255.0
	return uint8(((c1 * alpha) + (c0 * (1.0 - alpha))) * 255.0)
}

func LinearInterpColor(color0, color1 color.RGBA, alpha float64) color.RGBA {
	a := BlendChannel(color0.A, color1.A, alpha)
	r := BlendChannel(color0.R, color1.R, alpha)
	g := BlendChannel(color0.G, color1.G, alpha)
	b := BlendChannel(color0.B, color1.B, alpha)

	return color.RGBA{r, g, b, a}
}
