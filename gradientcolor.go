package noiselib

import (
	// "fmt"
	"image/color"
)

type GradientColor struct {
	GradientPoints map[float64]color.RGBA
}

func (g *GradientColor) AddGradientPoint(position float64, color color.RGBA) {
	g.GradientPoints[position] = color
}

func (g *GradientColor) ClearGradient() {
	g.GradientPoints = make(map[float64]color.RGBA)
}

func (g *GradientColor) GetColor(position float64) color.RGBA {
	keys := []float64{}
	for k, _ := range g.GradientPoints {
		keys = append(keys, k)
	}

	indexPos := 0

	for _, k := range keys {
		if position < k {
			break
		}
		indexPos++
	}

	index0 := ClampValue(indexPos-1, 0, len(g.GradientPoints)-1)
	index1 := ClampValue(indexPos, 0, len(g.GradientPoints)-1)

	if index0 == index1 {
		return g.GradientPoints[keys[index1]]
	}

	input0 := keys[index0]
	input1 := keys[index1]
	alpha := (position - input0) / (input1 - input0)

	color0 := g.GradientPoints[keys[index0]]
	color1 := g.GradientPoints[keys[index1]]

	return LinearInterpColor(color0, color1, alpha)
}
