package noiselib

type TranslatePoint struct {
	SourceModule                       []Module
	XTranslate, YTranslate, ZTranslate float64
}

func (t TranslatePoint) GetSourceModule(index int) Module {
	return t.SourceModule[index]
}

func (t TranslatePoint) SetSourceModule(index int, source Module) {
	t.SourceModule[index] = source
}

func (t *TranslatePoint) SetTranslatePoints(x, y, z float64) {
	t.XTranslate = x
	t.YTranslate = y
	t.ZTranslate = z
}

func (t TranslatePoint) GetValue(x, y, z float64) float64 {
	if t.SourceModule[0] == nil {
		panic("TranslatePoint must have one source")
	}

	return t.SourceModule[0].GetValue(x+t.XTranslate, y+t.YTranslate, z+t.ZTranslate)
}
