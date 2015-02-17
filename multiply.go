package noiselib

type Multiply struct {
	SourceModule []Module
}

func (m Multiply) GetSourceModule(index int) Module {
	return m.SourceModule[index]
}

func (m Multiply) SetSourceModule(index int, source Module) {
	m.SourceModule[index] = source
}

func (m Multiply) GetValue(x, y, z float64) float64 {
	if m.SourceModule[0] == nil || m.SourceModule[1] == nil {
		panic("Multiply must have 2 source modules.")
	}

	return m.SourceModule[0].GetValue(x, y, z) * m.SourceModule[1].GetValue(x, y, z)
}
