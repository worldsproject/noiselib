package noiselib

import (
	"testing"
)

func TestGradientNoise3D(t *testing.T) {
	v := GradientNoise3D(0.5, 0.5, 0.5, 1.0, 1.0, 1.0, 0)

	if v != 0.51156214 {
		t.Errorf("v should be 0.5116214 but instead was %v", v)
	}
}
