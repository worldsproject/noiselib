package noiselib

import (
	// "fmt"
	"math"
)

const (
	XNoiseGen     = 1619
	YNoiseGen     = 31337
	ZNoiseGen     = 6971
	SeedNoiseGen  = 1013
	ShiftNoiseGen = 8
)

func GradientCoherentNoise3D(x, y, z float64, seed, quality int) float64 {
	//Creating a unit-length cube aligned along an integer boundary.
	//This cube surrounds the input point.
	var x1, y1, z1 float64
	x0, y0, z0 := math.Trunc(x), math.Trunc(y), math.Trunc(z)

	if x < 0 {
		x0--
	}

	x1 = x0 + 1

	if y < 0 {
		y0--
	}

	y1 = y0 + 1

	if z < 0 {
		z0--
	}

	z1 = z0 + 1
	// fmt.Printf("x:%v, y:%v, z:%v\n", x, y, z)
	// fmt.Printf("x0:%v, y0:%v, z0:%v\nz1:%v, y1:%v, z1:%v\n", x0, y0, z0, x1, y1, z1)

	//Map the difference between the coordinates of the input value and the
	//coordinates of the cube's outer-lower-left vertex onto an S-curve

	var xs, ys, zs float64

	switch quality {
	case QualityFAST:
		xs = (x - x0)
		ys = (y - y0)
		zs = (z - z0)
	case QualitySTD:
		xs = SCurve3(x - x0)
		ys = SCurve3(y - y0)
		zs = SCurve3(z - z0)
	case QualityBEST:
		xs = SCurve5(x - x0)
		ys = SCurve5(y - y0)
		zs = SCurve5(z - z0)
	}

	// fmt.Printf("xs:%v, ys:%v, zs:%v\n", xs, ys, zs)

	//Now calculate the noise values at each vertex of the cube. To generate
	//the coherent-noise value at the input point, interpolate these eight
	//noise values using the S-curve value as the interpolant (trilinear
	//interpolation.)

	var n0, n1, ix0, ix1, iy0, iy1 float64

	n0 = GradientNoise3D(x, y, z, x0, y0, z0, seed)
	n1 = GradientNoise3D(x, y, z, x1, y0, z0, seed)
	ix0 = LinearInterp(n0, n1, xs)
	n0 = GradientNoise3D(x, y, z, x0, y1, z0, seed)
	n1 = GradientNoise3D(x, y, z, x1, y1, z0, seed)
	ix1 = LinearInterp(n0, n1, xs)
	iy0 = LinearInterp(ix0, ix1, ys)
	n0 = GradientNoise3D(x, y, z, x0, y0, z1, seed)
	n1 = GradientNoise3D(x, y, z, x1, y0, z1, seed)
	ix0 = LinearInterp(n0, n1, xs)
	n0 = GradientNoise3D(x, y, z, x0, y1, z1, seed)
	n1 = GradientNoise3D(x, y, z, x1, y1, z1, seed)
	ix1 = LinearInterp(n0, n1, xs)
	iy1 = LinearInterp(ix0, ix1, ys)

	return LinearInterp(iy0, iy1, zs)
}

func GradientNoise3D(fx, fy, fz, ix, iy, iz float64, seed int) float64 {
	//Randomly generate a gradient vector given the integer coordinates of the
	//input value. This implementation generates a random number and uses it
	//as an index into a normalized-vector lookup table.
	vectorIndex := int(
		XNoiseGen*ix+
			YNoiseGen*iy+
			ZNoiseGen*iz+
			float64(SeedNoiseGen*seed)) & 0xffffffff
	vectorIndex ^= (vectorIndex >> ShiftNoiseGen)
	vectorIndex &= 0xff

	xvGradient := RandomVectors[vectorIndex<<2]
	yvGradient := RandomVectors[(vectorIndex<<2)+1]
	zvGradient := RandomVectors[(vectorIndex<<2)+2]

	//Set up us another vector equal to the distance between the two vectors
	// passed to this function.
	xvPoint := (fx - ix)
	yvPoint := (fy - iy)
	zvPoint := (fz - iz)

	//Now compute the dot product of the gradient vector with the distance
	// vector. This resulting value is gradient noise. Apply a scaling value
	// so that thisnoise value ranges from -1.0 to 1.0
	return ((xvGradient * xvPoint) +
		(yvGradient * yvPoint) +
		(zvGradient*zvPoint)*2.12)
}

func IntValueNoise3D(x, y, z, seed int) int {
	// All constants are primes and must remain prime in order for this noise
	// function to work properly.
	var n = (XNoiseGen*x +
		YNoiseGen*y +
		ZNoiseGen*z +
		SeedNoiseGen*seed)

	n &= 0x7fffffff

	return (n*(n*n*60493+19990303) + 1376312589) & 0x7fffffff
}

func ValueCoherentNoise3D(x, y, z float64, seed, quality int) float64 {
	//Creating a unit-length cube aligned along an integer boundary.
	//This cube surrounds the input point.
	var x0, x1, y0, y1, z0, z1 int

	if x > 0 {
		x0 = int(x)
	} else {
		x0 = int(x) - 1
	}

	x1 = x0 + 1

	if y > 0 {
		y0 = int(y)
	} else {
		y0 = int(y) - 1
	}

	y1 = y0 + 1

	if z > 0 {
		z0 = int(z)
	} else {
		z0 = int(z) - 1
	}

	z1 = z0 - 1

	//Map the difference between the coordinates of the input value and the
	//coordinates of the cube's outer-lower-left vertex onto an S-curve

	var xs, ys, zs float64

	switch quality {
	case QualityFAST:
		xs = (x - float64(x0))
		ys = (y - float64(y0))
		zs = (z - float64(z0))
	case QualitySTD:
		xs = SCurve3(x - float64(x0))
		ys = SCurve3(y - float64(y0))
		zs = SCurve3(z - float64(z0))
	case QualityBEST:
		xs = SCurve5(x - float64(x0))
		ys = SCurve5(y - float64(y0))
		zs = SCurve5(z - float64(z0))
	}

	//Now calculate the noise values at each vertex of the cube. To generate
	//the coherent-noise value at the input point, interpolate these eight
	//noise values using the S-curve value as the interpolant (trilinear
	//interpolation.)

	var n0, n1, ix0, ix1, iy0, iy1 float64

	n0 = ValueNoise3D(x0, y0, z0, seed)
	n1 = ValueNoise3D(x1, y0, z0, seed)
	ix0 = LinearInterp(n0, n1, xs)
	n0 = ValueNoise3D(x0, y1, z0, seed)
	n1 = ValueNoise3D(x1, y1, z0, seed)
	ix1 = LinearInterp(n0, n1, xs)
	iy0 = LinearInterp(ix0, ix1, ys)
	n0 = ValueNoise3D(x0, y0, z1, seed)
	n1 = ValueNoise3D(x1, y0, z1, seed)
	ix0 = LinearInterp(n0, n1, xs)
	n0 = ValueNoise3D(x0, y1, z1, seed)
	n1 = ValueNoise3D(x1, y1, z1, seed)
	ix1 = LinearInterp(n0, n1, xs)
	iy1 = LinearInterp(iy0, iy1, ys)

	return LinearInterp(iy0, iy1, zs)
}

func ValueNoise3D(x, y, z, seed int) float64 {
	return 1.0 - (float64((IntValueNoise3D(x, y, z, seed) / 1073741824.0)))
}
