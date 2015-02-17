package noiselib

import "math"

type Spheres struct {
	Frequency float64
}

const DefaultSphereFrequency = 1.0

func (s Spheres) SetSourceModule(index int, source Module) {
	return
}

func (s Spheres) GetSourceModule(index int) Module {
	return nil
}

func (s Spheres) GetValue(x, y, z float64) float64 {
	x *= s.Frequency
	y *= s.Frequency
	z *= s.Frequency

	distFromCenter := math.Sqrt(x*x + y*y + z*z)
	distFromSmallerSphere := distFromCenter - math.Floor(distFromCenter)
	distFromLargerSphere := 1.0 - distFromSmallerSphere
	nearestDist := math.Min(distFromSmallerSphere, distFromLargerSphere)

	return 1.0 - (nearestDist * 4.0)
}
