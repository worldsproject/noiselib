package noiselib

import ("testing"
        "math/rand"
        )

func TestClamp(t *testing.T) {
  billow := DefaultBillow()

  clamp := DefaultClamp(billow)
  clamp.SetBounds(0.4, 0.6)

  for i := 0; i < 100; i++ {
    x := rand.Float64() * 5 + 5
    y := rand.Float64() * 5 + 5
    z := rand.Float64() * 5 + 5

    got := clamp.GetValue(x, y, z)

    if got < 0.4 || got > 0.6 {
      t.Errorf("For (%v, %v, %v) the value should be between 0.4 and 0.6, but instead got %v", x, y, z, got)
    }
  }
}
