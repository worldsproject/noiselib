package noiselib

type Constant struct {
  Value float64
}

func (c Constant) GetSourceModule(index int) Module {
  return nil
}

func (c Constant) GetSourceModuleCount() int {
  return 0
}

func (c Constant) GetValue(x, y, z float64) float64 {
  return c.Value
}

func (c Constant) SetSourceModule(index int, sourceModule Module) {
  return
}

func (c Constant) NewModule() Constant {
  return Constant { 0.0 }
}
