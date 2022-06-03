package homework

import (
	"testing"
)

const MESSAGE = "got %f, wanted %f"

func TestCalcSquareSidesTriangle(t *testing.T) {

	array := []float32{1, 2, 3, 4, 5, 6.0}
	got_ave := average(array)
	want_ave := float32(21.0 / 6)

	if got_ave != want_ave {
		t.Errorf(MESSAGE, got_ave, want_ave)
	}
}
