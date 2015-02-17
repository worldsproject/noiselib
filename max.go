package noiselib

import "math"

type Max struct {
	SourceModule []Module
}

func (m Max) GetSourceModule(index int) Module {
	return m.SourceModule[index]
}

func (m Max) SetSourceModule(index int, source Module) {
	m.SourceModule[index] = source
}

func (m Max) GetValue(x, y, z float64) float64 {
	if m.SourceModule[0] == nil || m.SourceModule[1] == nil {
		panic("Max must have 2 source modules.")
	}

	return math.Max(m.SourceModule[0].GetValue(x, y, z),
		m.SourceModule[1].GetValue(x, y, z))
}
