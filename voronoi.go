package noiselib

import "math"

type Voronoi struct {
	Displacement, Frequency float64
	EnableDistance          bool
	Seed                    int
}

const (
	DefaultVoronoiDisplacement = 1.0
	DefaultVoronoiFrequency    = 1.0
	DefaultVoronoiSeed         = 0
)

func (v Voronoi) GetSourceModule(index int) Module {
	return nil
}

func (v Voronoi) SetSourceModule(index int, source Module) {
	return
}

func (v Voronoi) GetValue(x, y, z float64) float64 {
	x *= v.Frequency
	y *= v.Frequency
	z *= v.Frequency

	var xInt, yInt, zInt int

	if x > 0.0 {
		xInt = int(x)
	} else {
		xInt = int(x) - 1
	}

	if y > 0.0 {
		yInt = int(y)
	} else {
		yInt = int(y) - 1
	}

	if z > 0.0 {
		zInt = int(z)
	} else {
		zInt = int(z) - 1
	}

	minDist := 2147483647.0
	var xCandidate, yCandidate, zCandidate float64

	// Inside each unit cube, there is a seed point at a random position. Go
	// through each of the nearby cubes until we find a buce with a seed point
	// that is closest to the specified position.
	for zCur := zInt - 2; zCur <= zInt+2; zCur++ {
		for yCur := yInt - 2; yCur <= yInt+2; yCur++ {
			for xCur := xInt - 2; xCur <= xInt+2; xCur++ {
				xPos := float64(xCur) + ValueNoise3D(xCur, yCur, zCur, v.Seed)
				yPos := float64(yCur) + ValueNoise3D(xCur, yCur, zCur, v.Seed+1)
				zPos := float64(zCur) + ValueNoise3D(xCur, yCur, zCur, v.Seed+2)
				xDist := xPos - x
				yDist := yPos - y
				zDist := zPos - z
				dist := xDist*xDist + yDist*yDist + zDist*zDist

				if dist < minDist {
					minDist = dist
					xCandidate = xPos
					yCandidate = yPos
					zCandidate = zPos
				}
			}
		}
	}

	var value float64

	if v.EnableDistance {
		// Determine distance to nearest seed point.
		xDist := xCandidate - x
		yDist := yCandidate - y
		zDist := zCandidate - z

		value = math.Sqrt(xDist*xDist+yDist*yDist+zDist*zDist)*SQRT_3 - 1.0
	} else {
		value = 0.0
	}

	// Return the calculated distance with the displacement value applied.
	return value + (v.Displacement * ValueNoise3D(
		int(math.Floor(xCandidate)),
		int(math.Floor(yCandidate)),
		int(math.Floor(zCandidate)), 0))
}

func DefaultVoronoi() Voronoi {
	return Voronoi{DefaultVoronoiDisplacement, DefaultVoronoiFrequency, true, DefaultVoronoiSeed}
}
