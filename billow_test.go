package noiselib

import "testing"

func TestBillow(t *testing.T) {
  billow := DefaultBillow()

  if billow.GetValue(0, 0, 0) < -2 {
    t.Errorf("Invalid Value")
  }
}
