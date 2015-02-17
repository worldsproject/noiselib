package noiselib

import "math"

type Power struct {
	SourceModule []Module
}

func (m Power) GetSourceModule(index int) Module {
	return m.SourceModule[index]
}

func (m Power) SetSourceModule(index int, source Module) {
	m.SourceModule[index] = source
}

func (m Power) GetValue(x, y, z float64) float64 {
	if m.SourceModule[0] == nil || m.SourceModule[1] == nil {
		panic("Power must have 2 source modules.")
	}

	return math.Pow(m.SourceModule[0].GetValue(x, y, z),
		m.SourceModule[1].GetValue(x, y, z))
}
