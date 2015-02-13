package noiselib

import "testing"

func TestAdd(t *testing.T) {
  one := Constant{1.0}
  two := Constant{2.0}

  add := Add{make([]Module, AddModuleCount)}
  add.SetSourceModule(0, one)
  add.SetSourceModule(1, two)

  if add.GetValue(0,0,0) != 3.0 {
    t.Error("For", "add.GetValue(0,0,0)",
            "Expected", 3.0,
            "Got", add.GetValue(0,0,0))
  }
}
