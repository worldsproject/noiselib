package noiselib

import "sort"

type Curve struct {
  ControlPoints map[float64]float64
  SourceModule []Module
}

func (c Curve) GetSourceModule(index int) Module {
  return c.SourceModule[index]
}

func (c Curve) SetSourceModule(index int, sourceModule Module) {
  c.SourceModule[index] = sourceModule
}

func (c *Curve) AddControlPoint(inputValue, outputValue float64) {
  c.ControlPoints[inputValue] = outputValue
}

func (c *Curve) ClearAllControlPoints() {
  c.ControlPoints = make(map[float64]float64)
}

func (c Curve) GetValue(x, y, z float64) float64 {
  //Get the ordering of the keys.
  var keys []float64

  for k := range c.ControlPoints {
    keys = append(keys, k)
  }
  sort.Float64s(keys)

  //Ensure that we have a source module.
  if c.SourceModule[0] == nil {
    panic("Curve must have a single source.")
  }

  //Curve must have 4 interpolation points.
  if len(keys) < 4 {
    panic("There must be at least 4 control points in a Curve module.")
  }

  //Get the output vluae from the soruce module.
  sourceModuleValue := c.SourceModule[0].GetValue(x, y, z)

  // Find the first element in the control point arrray that has an input value
  // larger than the output value from the source module.
  indexS := 0

  for _, k := range keys {
    if sourceModuleValue < c.ControlPoints[k] {
      break
    }

    indexS++
  }

  // Find the four nearest control points so that we can perform buvic
  // interpolation.
  index0 := ClampValue(indexS - 2, 0, len(keys))
  index1 := ClampValue(indexS - 1, 0, len(keys))
  index2 := ClampValue(indexS,     0, len(keys))
  index3 := ClampValue(indexS + 1, 0, len(keys))

  // If some control points are missing (which occurs if the value from the
  // source module is greater than the largest input value or less than the
  // smallest input value of the control point array), get the corresponding
  // output value of the nearest control point and exit now.
  if index1 == index2 {
    return c.ControlPoints[keys[index1]]
  }

  // Compute the alpha value used for cubic interpolation.
  input0 := c.ControlPoints[keys[index1]]
  input1 := c.ControlPoints[keys[index2]]
  alpha := (sourceModuleValue - input0) / (input1 - input0)

  return CubicInterp(
    c.ControlPoints[keys[index0]],
    c.ControlPoints[keys[index1]],
    c.ControlPoints[keys[index2]],
    c.ControlPoints[keys[index3]],
    alpha)
}
