package noiselib

import "math"

//Performs a cubic interpolation between two values bound between two
//other values.
func CubicInterp(n0, n1, n2, n3, a float64) float64 {
  p := (n3 - n2) - (n0 - n1)
  q := (n0 - n1) - p
  r := n2 - n0

  return (p * math.Pow(a, 3)) + (q * math.Pow(a, 2)) + (r * a) + n1
}

//Performs a linear interpolation between two values
func LinearInterp(n0, n1, a float64) float64 {
  return((1 - a) * n0) + (a * n1)
}

//Maps a value onto a cubic S-curve
func SCurve3(a float64) float64 {
  return (a * a * (3-2 * a))
}

//Maps a value onto a quintic S-curve.
func SCurve5(a float64) float64 {
  return (6 * math.Pow(a, 5)) - (15 * math.Pow(a, 4)) + (10 * math.Pow(a, 3))
}
