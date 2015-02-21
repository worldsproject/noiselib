package noiselib

import (
	"fmt"
	"testing"
)

func TestGradientNoise3D(t *testing.T) {
	v := GradientNoise3D(0.5, 0.5, 0.5, 1.0, 1.0, 1.0, 0)

	if v != 0.51156214 {
		t.Errorf("v should be 0.5116214 but instead was %v", v)
	}
}

func TestGradientCoherentNoise3D(t *testing.T) {
	x, y, z := 0.5, 0.0, 0.5
	x0, y0, z0 := 0.0, 0.0, 0.0
	x1, y1, z1 := 1.0, 1.0, 1.0
	seed := 0
	xs, ys, zs := 0.5, 0.5, 0.5

	var n0, n1, ix0, ix1, iy0, iy1 float64

	n0 = GradientNoise3D(x, y, z, x0, y0, z0, seed)
	n1 = GradientNoise3D(x, y, z, x1, y0, z0, seed)
	ix0 = LinearInterp(n0, n1, xs)
	fmt.Printf("n0:%v, n1%v, ix0:%v\n", n0, n1, ix0)
	n0 = GradientNoise3D(x, y, z, x0, y1, z0, seed)
	n1 = GradientNoise3D(x, y, z, x1, y1, z0, seed)
	ix1 = LinearInterp(n0, n1, xs)
	iy0 = LinearInterp(ix0, ix1, ys)
	fmt.Printf("n0:%v, n1%v, ix0:%v, iy0%v\n", n0, n1, ix1, iy0)
	n0 = GradientNoise3D(x, y, z, x0, y0, z1, seed)
	n1 = GradientNoise3D(x, y, z, x1, y0, z1, seed)
	ix0 = LinearInterp(n0, n1, xs)
	fmt.Printf("n0:%v, n1%v, ix0:%v\n", n0, n1, ix0)
	n0 = GradientNoise3D(x, y, z, x0, y1, z1, seed)
	n1 = GradientNoise3D(x, y, z, x1, y1, z1, seed)
	ix1 = LinearInterp(n0, n1, xs)
	iy1 = LinearInterp(ix0, ix1, ys)
	fmt.Printf("n0:%v, n1%v, ix0:%v, iy1:%v\n", n0, n1, ix1, iy1)

	v := LinearInterp(iy0, iy1, zs)
	fmt.Printf("v:%v", v)

}
