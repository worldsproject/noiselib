package noiselib

import ("math")

type Abs struct {
  SourceModule []Module
}

func (a Abs) GetSourceModule(index int) Module {
  return a.SourceModule[index]
}

func (a Abs) GetSourceModuleCount() int {
  return 1
}

func (a Abs) GetValue(x, y, z float64) float64 {
  source := a.SourceModule[0]

  if source == nil {
    panic("Abs requires a source module.")
  }

  return math.Abs(source.GetValue(x, y, z))
}

func (a Abs) SetSourceModule(index int, source Module) {
  a.SourceModule[index] = source
}
