package homework

import (
	"testing"
)

const MESSAGE = "got %f, wanted %f"

func TestAverage(t *testing.T) {

	array := [15]float32{1, 2, 3, 4, 5, 6.0}

	// array1 := new([15]float32)[0:15]
	// println(reflect.TypeOf(array1))
	// copy(array1, array[])
	// println(len(array))
	got_ave := average(array)
	want_ave := float32(21.0 / 6)

	if got_ave != want_ave {
		t.Errorf(MESSAGE, got_ave, want_ave)
	}
}
