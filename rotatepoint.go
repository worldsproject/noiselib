package noiselib

import "math"

type Rotatepoint struct {
	X1Matrix, X2Matrix, X3Matrix float64
	Y1Matrix, Y2Matrix, Y3Matrix float64
	Z1Matrix, Z2Matrix, Z3Matrix float64
	XAngle, YAngle, ZAngle       float64
	SourceModule                 []Module
}

const (
	DefaultRotateX = 0.0
	DefaultRotateY = 0.0
	DefaultRotateZ = 0.0
	DegToRad       = math.Pi / 180
)

func (r Rotatepoint) GetSourceModule(index int) Module {
	return r.SourceModule[index]
}

func (r Rotatepoint) SetSourceModule(index int, source Module) {
	r.SourceModule[index] = source
}

func (r *Rotatepoint) SetAngles(xAngle, yAngle, zAngle float64) {
	xCos := math.Cos(xAngle * DegToRad)
	yCos := math.Cos(yAngle * DegToRad)
	zCos := math.Cos(zAngle * DegToRad)
	xSin := math.Sin(xAngle * DegToRad)
	ySin := math.Sin(yAngle * DegToRad)
	zSin := math.Sin(zAngle * DegToRad)

	r.X1Matrix = ySin*xSin*zSin + yCos*zCos
	r.Y1Matrix = xCos * zSin
	r.Z1Matrix = ySin*zCos - yCos*xSin*zSin

	r.X2Matrix = ySin*xSin*zCos - yCos*zSin
	r.Y2Matrix = xCos * zCos
	r.Z2Matrix = -yCos*xSin*zCos - ySin*zSin

	r.X3Matrix = -ySin * xCos
	r.Y3Matrix = xSin
	r.Z3Matrix = yCos * xCos

	r.XAngle = xAngle
	r.YAngle = yAngle
	r.ZAngle = zAngle
}

func (r *Rotatepoint) SetXAngle(xAngle float64) {
	r.SetAngles(xAngle, r.YAngle, r.ZAngle)
}

func (r *Rotatepoint) SetYAngle(yAngle float64) {
	r.SetAngles(r.XAngle, yAngle, r.ZAngle)
}

func (r *Rotatepoint) SetZAngle(zAngle float64) {
	r.SetAngles(r.XAngle, r.YAngle, zAngle)
}

func (r Rotatepoint) GetValue(x, y, z float64) float64 {
	if r.SourceModule[0] == nil {
		panic("Rotatepoint must have one source.")
	}

	nx := (r.X1Matrix * x) + (r.Y1Matrix * y) + (r.Z1Matrix * z)
	ny := (r.X2Matrix * x) + (r.Y2Matrix * y) + (r.Z2Matrix * z)
	nz := (r.X3Matrix * x) + (r.Y3Matrix * y) + (r.Z3Matrix * z)

	return r.SourceModule[0].GetValue(nx, ny, nz)
}

func DefaultRotatepoint() Rotatepoint {
	rp := Rotatepoint{SourceModule: make([]Module, RotateModuleCount)}
	rp.SetAngles(DefaultRotateX, DefaultRotateY, DefaultRotateZ)
	return rp
}
