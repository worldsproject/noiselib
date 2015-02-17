package noiselib

import "math"

type Min struct {
	SourceModule []Module
}

func (m Min) GetSourceModule(index int) Module {
	return m.SourceModule[index]
}

func (m Min) SetSourceModule(index int, source Module) {
	m.SourceModule[index] = source
}

func (m Min) GetValue(x, y, z float64) float64 {
	if m.SourceModule[0] == nil || m.SourceModule[1] == nil {
		panic("Min must have 2 source modules.")
	}

	return math.Min(m.SourceModule[0].GetValue(x, y, z),
		m.SourceModule[1].GetValue(x, y, z))
}
