package noiselib

import "math"

type Cylinders struct {
	Frequency float64
}

func (c Cylinders) GetSourceModule(index int) Module {
	return nil
}

func (c Cylinders) SetSourceModule(index int, sourceModule Module) {
	return
}

func (c Cylinders) GetValue(x, y, z float64) float64 {
	x *= c.Frequency
	z *= c.Frequency

	distFromCenter := math.Sqrt(x*x + z*z)
	distFromSmallerSphere := distFromCenter - math.Floor(distFromCenter)
	distFromLargerSphere := 1.0 - distFromSmallerSphere
	nearestDist := math.Min(distFromSmallerSphere, distFromLargerSphere)

	return 1.0 - (nearestDist * 4.0)
}

const DefaultCylinderFrequency = 1.0

func DefaultCylinders() Cylinders {
	return Cylinders{DefaultCylinderFrequency}
}
