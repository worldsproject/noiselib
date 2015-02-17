package noiselib

type ScalePoint struct {
	SourceModule           []Module
	XScale, YScale, ZScale float64
}

const (
	DefaultScalePointX = 1.0
	DefaultScalePointY = 1.0
	DefaultScalePointZ = 1.0
)

func (a ScalePoint) GetSourceModule(index int) Module {
	return a.SourceModule[index]
}

func (a ScalePoint) GetValue(x, y, z float64) float64 {
	if a.SourceModule[0] == nil {
		panic("ScalePoint requires a source module.")
	}

	return a.SourceModule[0].GetValue(x*a.XScale, y*a.YScale, z*a.ZScale)
}

func (a ScalePoint) SetSourceModule(index int, source Module) {
	a.SourceModule[index] = source
}

func DefaultScalePoint() ScalePoint {
	return ScalePoint{make([]Module, ScalePointModuleCount), DefaultScalePointX, DefaultScalePointY, DefaultScalePointZ}
}
