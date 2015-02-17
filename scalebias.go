package noiselib

type ScaleBias struct {
	SourceModule []Module
	Bias, Scale  float64
}

const (
	DefaultBias  = 0.0
	DefaultScale = 1.0
)

func (a ScaleBias) GetSourceModule(index int) Module {
	return a.SourceModule[index]
}

func (a ScaleBias) GetValue(x, y, z float64) float64 {
	if a.SourceModule[0] == nil {
		panic("ScaleBias requires a source module.")
	}

	return a.SourceModule[0].GetValue(x, y, z)*a.Scale + a.Bias
}

func (a ScaleBias) SetSourceModule(index int, source Module) {
	a.SourceModule[index] = source
}

func DefaultScaleBias() ScaleBias {
	return ScaleBias{make([]Module, ScaleBiasModuleCount), DefaultBias, DefaultScale}
}
