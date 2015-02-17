package noiselib

type Perlin struct {
	Frequency, Lacunarity, Persistence float64
	OctaveCount, Quality, Seed         int
}

const (
	DefaultPerlinFrequency   = 1.0
	DefaultPerlinLacunarity  = 2.0
	DefaultPerlinOctaveCount = 6
	DefaultPerlinPersistence = 0.5
	DefaultPerlinQuality     = QualitySTD
	DefaultPerlinSeed        = 0
	PerlinMaxOctave          = 30
)

func (p Perlin) GetSourceModule(index int) Module {
	return nil
}

func (p Perlin) SetSourceModule(index int, source Module) {
	return
}

// SetOctaveCount sets the number of octaves performed in the noise calculation.
// Automatically clamps the given value to [1, PerlinMaxOctave].
func (p *Perlin) SetOctaveCount(octave int) {
	p.OctaveCount = ClampValue(octave, 1, PerlinMaxOctave)
}

func (p Perlin) GetValue(x, y, z float64) float64 {
	value := 0.0
	signal := 0.0
	curPersistence := 1.0
	seed := 0

	x *= p.Frequency
	y *= p.Frequency
	z *= p.Frequency

	for curOctave := 0; curOctave < p.OctaveCount; curOctave++ {
		// Get the coherent-noise value from the input value and add it to the
		// final result.
		seed = (p.Seed + curOctave) & 0xffffffff
		signal = GradientCoherentNoise3D(x, y, z, seed, p.Quality)
		value += signal * curPersistence

		//Prepare the next octave.
		x *= p.Lacunarity
		y *= p.Lacunarity
		z *= p.Lacunarity

		curPersistence *= p.Persistence
	}

	return value
}

func DefaultPerlin() Perlin {
	return Perlin{
		Frequency:   DefaultPerlinFrequency,
		Lacunarity:  DefaultPerlinLacunarity,
		Persistence: DefaultPerlinPersistence,
		OctaveCount: DefaultPerlinOctaveCount,
		Quality:     DefaultPerlinQuality,
		Seed:        DefaultPerlinSeed}
}
