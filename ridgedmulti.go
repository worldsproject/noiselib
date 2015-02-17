package noiselib

import "math"

type Ridgedmulti struct {
	Frequency, Lacunarity, SpectralWeight, Offset, Gain float64
	OctaveCount, Quality, Seed                          int
	SpectralWeights                                     []float64
}

const (
	DefaultRidgedFrequency      = 1.0
	DefaultRidgedLacunarity     = 2.0
	DefaultRidgedOctaveCount    = 6
	DefaultRidgedQuality        = QualitySTD
	DefaultRidgedSeed           = 0
	RidgedMaxOctave             = 30
	DefaultRidgedSpectralWeight = 1.0
	DefaultRidgedOffset         = 1.0
	DefaultRidgedGain           = 2.0
)

// SetOctaveCount sets the number of octaves performed in the noise calculation.
// Automatically clamps the given value to [1, PerlinMaxOctave].
func (r *Ridgedmulti) SetOctaveCount(octave int) {
	r.OctaveCount = ClampValue(octave, 1, RidgedMaxOctave)
}

func (r Ridgedmulti) GetSourceModule(index int) Module {
	return nil
}

func (r Ridgedmulti) SetSourceModule(index int, source Module) {
	return
}

func (r Ridgedmulti) CalcSpectralWeights() {
	h := r.SpectralWeight

	frequency := 1.0

	for i := 0; i < RidgedMaxOctave; i++ {
		r.SpectralWeights[i] = math.Pow(frequency, -h)
		frequency *= r.Lacunarity
	}
}

func (r Ridgedmulti) GetValue(x, y, z float64) float64 {
	x *= r.Frequency
	y *= r.Frequency
	z *= r.Frequency

	signal := 0.0
	value := 0.0
	weight := 1.0

	for curOctave := 0; curOctave < r.OctaveCount; curOctave++ {
		// Get the coherent-noise value.
		seed := (r.Seed + curOctave) & 0x7fffffff
		signal = GradientCoherentNoise3D(x, y, z, seed, r.Quality)

		// Make the ridges
		signal = math.Abs(signal)
		signal = r.Offset - signal

		// Square the signal to increase the sharpness of the ridges.
		signal *= signal

		// The weighting from the previous octave is applied to the signa.
		// larger values have higher weights, producing sharp points along the
		// ridges.
		signal *= weight

		// Weight successive contributions by the previous signal.
		weight = signal * r.Gain
		if weight > 1.0 {
			weight = 1.0
		}

		if weight < 0.0 {
			weight = 0.0
		}

		// Add the signal to the output value
		value += (signal * r.SpectralWeights[curOctave])

		// Go to the next octave
		x *= r.Lacunarity
		y *= r.Lacunarity
		z *= r.Lacunarity
	}

	return (value * 1.35) - 1.0
}

func DefaultRidgedmulti() Ridgedmulti {
	rm := Ridgedmulti{
		Frequency:       DefaultRidgedFrequency,
		Lacunarity:      DefaultRidgedLacunarity,
		SpectralWeight:  DefaultRidgedSpectralWeight,
		Offset:          DefaultRidgedOffset,
		Gain:            DefaultRidgedGain,
		OctaveCount:     DefaultRidgedOctaveCount,
		Quality:         DefaultRidgedQuality,
		Seed:            DefaultRidgedSeed,
		SpectralWeights: make([]float64, RidgedMaxOctave)}

	rm.CalcSpectralWeights()
	return rm
}
