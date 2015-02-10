package noiselib

type Add struct {
  SourceModule []Module
}

func (a Add) GetSourceModule(index int) Module {
  return a.SourceModule[index]
}

func (a Add) GetSourceModuleCount() int {
  return 2
}

func (a Add) GetValue(x, y, z float64) float64 {
  source1 := a.SourceModule[0]
  source2 := a.SourceModule[1]

  if source1 == nil || source2 == nil {
    panic("Add Module requires two sources.")
  }

  return source1.GetValue(x, y, z) + source2.GetValue(x, y, z)
}

func (a Add) SetSourceModule(index int, source Module) {
  a.SourceModule[index] = source
}

func (a Add) NewModule() Module {
  return Add{ make([]Module, 2) }
}
