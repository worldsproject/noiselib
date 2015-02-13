package noiselib

const DefaultClampLowerBound = -1.0
const DefaultClampUpperBound = 1.0

type Clamp struct {
  SourceModule []Module
  LowerBound, UpperBound float64
}

func (c Clamp) GetSourceModule(index int) Module {
  return c.SourceModule[index]
}

func (c Clamp) SetSourceModule(index int, sourceModule Module) {
  c.SourceModule[index] = sourceModule
}

func (c Clamp) GetValue(x, y, z float64) float64 {
  if c.SourceModule[0] == nil {
    panic("Clamp must have one source")
  }

  value := c.SourceModule[0].GetValue(x, y, z)

  if value < c.LowerBound {
    return c.LowerBound
  } else if value > c.UpperBound {
    return c.UpperBound
  } else {
    return value
  }
}

func (c *Clamp) SetBounds(lowerBound, upperBound float64) {
  if lowerBound > upperBound {
    panic("Clamps LowerBound must be <= UpperBound")
  }

  c.LowerBound = lowerBound
  c.UpperBound = upperBound
}

func DefaultClamp(module Module) *Clamp {
  clamp := Clamp{SourceModule: make([]Module, ClampModuleCount),
    UpperBound:DefaultClampUpperBound,
    LowerBound: DefaultClampLowerBound}
  clamp.SourceModule[0] = module

  return &clamp
}
