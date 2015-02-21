package noiselib

import "math"

type Checkerboard struct{}

func (c Checkerboard) GetSourceModule(index int) Module {
	return nil
}

func (c Checkerboard) SetSourceModule(index int, sourceModule Module) {
	return
}

func (c Checkerboard) GetValue(x, y, z float64) float64 {
	ix := int(math.Floor(x))
	iy := int(math.Floor(y))
	iz := int(math.Floor(z))

	if (ix&1 ^ iy&1 ^ iz&1) == 0 {
		return -1.0
	} else {
		return 1.0
	}
}
