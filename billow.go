package noiselib

import "math"

const DefaultBillowFrequency = float64(1)
const DefaultBillowLacunarity = float64(2)
const DefaultBillowOctaveCount = 6
const DefaultBillowPersistence = float64(0.5)
const DefaultBillowQuality = QualitySTD
const DefaultBillowSeed = 0
const DefaultBillowMaxOctave = 30

type Billow struct {
  SourceModule []Module
  Frequency, Lacunarity, Persistence float64
  Seed, Quality, OctaveCount int
}

func (b Billow) GetSourceModule(index int) Module {
  return b.SourceModule[index]
}

func (b Billow) GetValue(x, y, z float64) float64 {
  value, signal, curPersistence := float64(0), float64(0), float64(1)

  var seed int

  x *= b.Frequency
  y *= b.Frequency
  z *= b.Frequency

  for curOctave := 0; curOctave < b.OctaveCount; curOctave++ {
    // Get the coherent-noise value from the input value and add it to the
    // final result.
    seed = (b.Seed + curOctave) & 0xffffffff
    signal = GradientCoherentNoise3D(x, y, z, seed, b.Quality)
    signal = 2.0 * math.Abs(signal) - 1.0
    value += signal * curPersistence

    //Prepare the next octave.
    x *= b.Lacunarity
    y *= b.Lacunarity
    z *= b.Lacunarity
  }

  value += 0.5

  return value
}

func (b Billow) SetSourceModule(index int, module Module) {
  b.SourceModule[index] = module
}

func DefaultBillow() Billow {
  return Billow {
    make([]Module, 0),
    DefaultBillowFrequency,
    DefaultBillowLacunarity,
    DefaultBillowPersistence,
    DefaultBillowSeed,
    DefaultBillowQuality,
    DefaultBillowOctaveCount,
  }
}
