package noiselib

import (
	"math"
)

func MapToCylinder(angle, height float64, source Module) float64 {
	var x, y, z float64

	x = math.Cos(angle * DegToRad)
	y = height
	z = math.Sin(angle * DegToRad)

	return source.GetValue(x, y, z)
}

func MapToLine(sx, sy, sz, ex, ey, ez, p float64, attenuate bool, source Module) float64 {
	x := (ex-sx)*p + sx
	y := (ey-sy)*p + sy
	z := (ez-sz)*p + sz
	value := source.GetValue(x, y, z)

	if attenuate {
		return p * (1.0 - p) * 4 * value
	} else {
		return value
	}
}

func MapToPlane(x, y float64, source Module) float64 {
	return source.GetValue(x, 0, y)
}

func MapToSphere(lat, lon float64, source Module) float64 {
	r := math.Cos(DegToRad * lat)
	x := r * math.Cos(DegToRad*lon)
	y := math.Sin(DegToRad * lat)
	z := r * math.Sin(DegToRad*lon)

	return source.GetValue(x, y, z)
}

func NoiseMapCylinder(lowerAngleBound, upperAngleBound,
	lowerHeightBound, upperHeightBound float64,
	width, height int, source Module) [][]float64 {

	m := MakeMap(width, height)

	angleExtent := upperAngleBound - lowerAngleBound
	heightExtent := upperHeightBound - lowerHeightBound
	xDelta := angleExtent / float64(width)
	yDelta := heightExtent / float64(height)
	curAngle := lowerAngleBound
	curHeight := lowerHeightBound

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			m[x][y] = MapToCylinder(curAngle, curHeight, source)
			curAngle += xDelta
		}
		curAngle = lowerAngleBound
		curHeight += yDelta
	}
	return m
}

func NoiseMapPlane(lowerXBound, upperXBound, lowerZBound, upperZBound float64,
	width, height int, seamless bool, source Module) [][]float64 {

	m := MakeMap(width, height)

	xExtent := upperXBound - lowerXBound
	zExtent := upperZBound - lowerZBound
	xDelta := xExtent / float64(width)
	zDelta := zExtent / float64(height)
	xCur := lowerXBound
	zCur := lowerZBound

	for z := 0; z < height; z++ {
		for x := 0; x < width; x++ {
			if seamless {

				swValue := MapToPlane(xCur, zCur, source)
				seValue := MapToPlane(xCur+xExtent, zCur, source)
				nwValue := MapToPlane(xCur, zCur+zExtent, source)
				neValue := MapToPlane(xCur+xExtent, zCur+zExtent, source)

				xBlend := 1.0 - ((xCur - lowerXBound) / xExtent)
				zBlend := 1.0 - ((zCur - lowerZBound) / zExtent)

				z0 := LinearInterp(swValue, seValue, xBlend)
				z1 := LinearInterp(nwValue, neValue, xBlend)

				m[x][z] = LinearInterp(z0, z1, zBlend)
			} else {
				m[x][z] = MapToPlane(xCur, zCur, source)
			}

			xCur += xDelta
		}
		xCur = lowerXBound
		zCur += zDelta
	}

	return m
}

func NoiseMapSphere(eastBound, westBound, northBound, southBound float64,
	width, height int, source Module) [][]float64 {

	m := MakeMap(width, height)

	lonExtent := eastBound - westBound
	latExtent := northBound - southBound
	xDelta := lonExtent / float64(width)
	yDelta := latExtent / float64(height)
	curLon := westBound
	curLat := southBound

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			m[x][y] = MapToSphere(curLat, curLon, source)
			curLon += xDelta
		}
		curLon = westBound
		curLat += yDelta
	}

	return m
}

func MakeMap(width, height int) [][]float64 {
	m := make([][]float64, width)
	for i := range m {
		m[i] = make([]float64, height)
	}

	return m
}
