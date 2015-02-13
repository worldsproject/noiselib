package noiselib

type Blend struct {
  SourceModule []Module
}

func (b Blend) GetSourceModule(index int) Module {
  return b.SourceModule[index]
}

func (b Blend) SetSourceModule(index int, sourceModule Module) {
  b.SourceModule[index] = sourceModule
}

func (b Blend) GetValue(x, y, z float64) float64 {
  if b.SourceModule[0] == nil || b.SourceModule[1] == nil || b.SourceModule[2] == nil {
    panic("Blend must have 3 sources.")
  }

  v0 := b.SourceModule[0].GetValue(x, y, z)
  v1 := b.SourceModule[1].GetValue(x, y, z)
  alpha := b.SourceModule[2].GetValue(x, y, z)
  return LinearInterp(v0, v1, alpha)
}
