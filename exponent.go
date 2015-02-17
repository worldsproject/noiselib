package noiselib

import "math"

type Exponent struct {
	SourceModule []Module
	Exponent     float64
}

const DefaultExp = 1.0

func (e Exponent) GetSourceModule(index int) Module {
	return e.SourceModule[index]
}

func (e Exponent) SetSourceModule(index int, source Module) {
	e.SourceModule[index] = source
}

func (e Exponent) GetValue(x, y, z float64) float64 {
	if e.SourceModule[0] == nil {
		panic("Exponent must have 1 source module.")
	}

	value := e.SourceModule[0].GetValue(z, y, z)
	return (math.Pow(math.Abs((value+1.0)/2.0), e.Exponent)*2.0 - 1.0)
}

func DefaultExponent() Exponent {
	return Exponent{make([]Module, ExponentModuleCount), DefaultExp}
}
