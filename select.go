package noiselib

type Select struct {
	SourceModule                        []Module
	LowerBound, UpperBound, EdgeFalloff float64
}

const DefaultSelectEdgeFalloff = 0.0
const DefaultSelectLowerBound = -1.0
const DefaultSelectUpperBound = 1.0

func (s Select) GetSourceModule(index int) Module {
	return s.SourceModule[index]
}

func (s Select) SetSourceModule(index int, source Module) {
	s.SourceModule[index] = source
}

func (s Select) GetValue(x, y, z float64) float64 {
	if s.SourceModule[0] == nil || s.SourceModule[1] == nil || s.SourceModule[2] == nil {
		panic("Select requires 3 source modules.")
	}

	controlValue := s.SourceModule[2].GetValue(x, y, z)

	if s.EdgeFalloff > 0.0 {
		if controlValue < (s.LowerBound - s.EdgeFalloff) {
			// The output value from the control module is below the selector
			// threshold; return the output value from the first source module.
			return s.SourceModule[0].GetValue(x, y, z)
		} else if controlValue < (s.LowerBound + s.EdgeFalloff) {
			// The output value from the control module is near the lower end of the
			// selector threshold and within the smooth curve. Interpolate between
			// the output values from the first and second source modules.
			lowerCurve := (s.LowerBound - s.EdgeFalloff)
			upperCurve := (s.LowerBound + s.EdgeFalloff)
			alpha := SCurve3((controlValue - lowerCurve) / (upperCurve - lowerCurve))
			return LinearInterp(s.SourceModule[0].GetValue(x, y, z), s.SourceModule[1].GetValue(x, y, z), alpha)
		} else if controlValue < s.UpperBound-s.EdgeFalloff {
			//The output value from the contorl module is within the selector
			// threshold; return the output value from the second source module.
			return s.SourceModule[1].GetValue(x, y, z)
		} else if controlValue < s.UpperBound+s.EdgeFalloff {
			//The output value from the control module is near the upper end of the
			// selector threshold and within the smooth curve. Interpolate between
			// the output values from the first and second source modules.
			lowerCurve := (s.UpperBound - s.EdgeFalloff)
			upperCurve := (s.UpperBound + s.EdgeFalloff)
			alpha := SCurve3((controlValue - lowerCurve) / (upperCurve - lowerCurve))
			return LinearInterp(s.SourceModule[1].GetValue(x, y, z), s.SourceModule[0].GetValue(x, y, z), alpha)
		} else {
			// Output value from the contorl module is above the selector threshold;
			// return the output value from the first soruce module
			return s.SourceModule[0].GetValue(x, y, z)
		}
	} else {
		if controlValue < s.LowerBound || controlValue > s.UpperBound {
			return s.SourceModule[0].GetValue(x, y, z)
		} else {
			return s.SourceModule[1].GetValue(x, y, z)
		}
	}
}

func (s *Select) SetBounds(lowerBound, upperBound float64) {
	if lowerBound > upperBound {
		panic("LowerBound in Select must be lower than UpperBound")
	}

	s.LowerBound = lowerBound
	s.UpperBound = upperBound
	s.SetEdgeFalloff(s.EdgeFalloff)
}

func (s *Select) SetEdgeFalloff(edgeFalloff float64) {
	boundSize := s.UpperBound - s.LowerBound

	if edgeFalloff > (boundSize / 2) {
		s.EdgeFalloff = (boundSize / 2)
	} else {
		s.EdgeFalloff = edgeFalloff
	}
}

func DefaultSelect() Select {
	return Select{make([]Module, SelectModuleCount), DefaultSelectLowerBound, DefaultSelectUpperBound, DefaultSelectEdgeFalloff}
}
