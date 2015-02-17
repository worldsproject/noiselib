package noiselib

type Turbulence struct {
	Power                                          float64
	XDistortModule, YDistortModule, ZDistortModule Perlin
	SourceModules                                  []Module
}

const (
	DefaultTurbulenceFrequency = DefaultPerlinFrequency
	DefaultTurbulencePower     = 1.0
	DefaultTurbulenceRoughness = 3
	DefaultTurbulenceSeed      = DefaultPerlinSeed
)

func (t Turbulence) GetSourceModule(index int) Module {
	return t.SourceModules[index]
}

func (t Turbulence) SetSourceModule(index int, source Module) {
	t.SourceModules[index] = source
}

func (t Turbulence) GetValue(x, y, z float64) float64 {
	if t.SourceModules[0] == nil {
		panic("Turbulence must have a source.")
	}

	// Get the values from the three Perlin noise modules and
	// add each value to each occrdinate of the input value. There are also
	// some offsets added to the coordinates of the input values. This prevents
	// the distortion modules from returning zero if the (x, y, z) coordinates,
	// when multiplied by the frequency, are near the integer boundary. This is
	// due to a property of gradient coherent noise, which returns zero at
	// integer boundries.

	var x0, y0, z0 float64
	var x1, y1, z1 float64
	var x2, y2, z2 float64

	x0 = x + (12414.0 / 65536.0)
	y0 = y + (65124.0 / 65536.0)
	z0 = z + (31337.0 / 65536.0)
	x1 = x + (26519.0 / 65536.0)
	y1 = y + (18128.0 / 65536.0)
	z1 = z + (60493.0 / 65536.0)
	x2 = x + (53820.0 / 65536.0)
	y2 = y + (11213.0 / 65536.0)
	z2 = z + (44845.0 / 65536.0)

	xDistort := x + (t.XDistortModule.GetValue(x0, y0, z0) * t.Power)
	yDistort := y + (t.YDistortModule.GetValue(x1, y1, z1) * t.Power)
	zDistort := z + (t.ZDistortModule.GetValue(x2, y2, z2) * t.Power)

	return t.SourceModules[0].GetValue(xDistort, yDistort, zDistort)
}

func (t *Turbulence) SetFrequency(frequency float64) {
	t.XDistortModule.Frequency = frequency
	t.YDistortModule.Frequency = frequency
	t.ZDistortModule.Frequency = frequency
}

func (t *Turbulence) SetRoughness(roughness int) {
	t.XDistortModule.SetOctaveCount(roughness)
	t.YDistortModule.SetOctaveCount(roughness)
	t.ZDistortModule.SetOctaveCount(roughness)
}

func (t *Turbulence) SetSeed(seed int) {
	t.XDistortModule.Seed = seed
	t.XDistortModule.Seed = seed + 1
	t.XDistortModule.Seed = seed + 2
}

func DefaultTurbulence() Turbulence {
	turb := Turbulence{DefaultTurbulencePower,
		DefaultPerlin(), DefaultPerlin(), DefaultPerlin(),
		make([]Module, TurbulenceModuleCount),
	}
	turb.SetSeed(DefaultTurbulenceSeed)
	turb.SetRoughness(DefaultTurbulenceRoughness)
	turb.SetFrequency(DefaultTurbulenceFrequency)
	return turb
}
