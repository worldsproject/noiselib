package noiselib

type Const struct {
  Value float64
}

func (c Const) GetSourceModule(index int) Module {
  return nil
}

func (c Const) GetSourceModuleCount() int {
  return 0
}

func (c Const) GetValue(x, y, z float64) float64 {
  return c.Value
}

func (c Const) SetSourceModule(index int, sourceModule Module) {
  return
}

func (c Const) NewModule() Module {
  return Const { 0.0 }
}
