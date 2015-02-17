package noiselib

type Invert struct {
	SourceModule []Module
}

func (i Invert) GetSourceModule(index int) Module {
	return i.SourceModule[index]
}

func (i Invert) SetSourceModule(index int, source Module) {
	i.SourceModule[index] = source
}

func (i Invert) GetValue(x, y, z float64) float64 {
	if i.SourceModule[0] == nil {
		panic("Invert must have 1 source module.")
	}

	return -i.GetValue(x, y, z)
}
