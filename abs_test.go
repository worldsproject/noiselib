package noiselib

import ("testing")

func TestAbs(t *testing.T) {
  one := Constant{1.0}

  two := Constant{Value:-1.5}

  abs := Abs{make([]Module, AbsModuleCount)}
  abs.SetSourceModule(0, one)

  if abs.GetValue(0,0,0) != 1.0 {
    t.Error("For", one.GetValue(0,0,0),
            "Expected", 1.0,
            "Got", abs.GetValue(0,0,0))
  }

  abs.SetSourceModule(0, two)

  if abs.GetValue(0,0,0) != 1.5 {
    t.Error("For", one.GetValue(0,0,0),
            "Expected", 1.5,
            "Got", abs.GetValue(0,0,0))
  }
}
