package noiselib

import "sort"

type Terrace struct {
	InvertTerraces bool
	ControlPoints  []float64
	SourceModule   []Module
}

func (t Terrace) GetSourceModule(index int) Module {
	return t.SourceModule[index]
}

func (t Terrace) SetSourceModule(index int, source Module) {
	t.SourceModule[index] = source
}

func (t Terrace) GetValue(x, y, z float64) float64 {
	if t.SourceModule[0] == nil {
		panic("Terrace must have a source.")
	}

	if len(t.ControlPoints) < 2 {
		panic("Terrace must have at least 2 control points.")
	}

	sourceValue := t.SourceModule[0].GetValue(x, y, z)

	indexPos := 0

	for indexPos = range t.ControlPoints {
		if sourceValue < t.ControlPoints[indexPos] {
			break
		}
	}

	// Find the two nearest control points so that we can map their values
	// onto a quadratic curve.
	index0 := ClampValue(indexPos-1, 0, len(t.ControlPoints))
	index1 := ClampValue(indexPos, 0, len(t.ControlPoints))

	// If some control points are missing (which offurcs if the output value from
	// the source module is greater than the largest value or less than the
	// smallest value of the control point array), get the value of the nearest
	// control point and exit now.
	if index0 == index1 {
		return t.ControlPoints[index1]
	}

	// Compute the alpha value used for linear interpolation.
	value0 := t.ControlPoints[index0]
	value1 := t.ControlPoints[index1]
	alpha := (sourceValue - value0) / (value1 - value0)

	if t.InvertTerraces {
		alpha = 1.0 - alpha
		t := value0
		value0 = value1
		value1 = t
	}

	// Squaring the alpha produces the terrace effect.
	alpha *= alpha

	// Now perform the linear interpolation given the alpha value.
	return LinearInterp(value0, value1, alpha)
}

func (t Terrace) AddControlPoint(value float64) {
	t.ControlPoints = append(t.ControlPoints, value)
	sort.Float64s(t.ControlPoints)
}

func (t Terrace) ClearAllControlPoints() {
	t.ControlPoints = make([]float64, 2)
}

func (t Terrace) MakeControlPoints(count int) {
	if count < 2 {
		panic("ControlPoint count must be at least 2.")
	}

	t.ClearAllControlPoints()

	terraceStep := 2.0 / (float64(count) - 1)
	value := -1.0

	for i := 0; i < count; i++ {
		t.AddControlPoint(value)
		value += terraceStep
	}
}

func DefaultTerrace() Terrace {
	t := Terrace{false, make([]float64, 2), make([]Module, TerraceModuleCount)}
	t.MakeControlPoints(2)
	return t
}
