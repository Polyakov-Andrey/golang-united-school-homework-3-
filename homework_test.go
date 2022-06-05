package homework

import (
	"testing"
)

const MESSAGE = "%+v, %+v"

func TestAverage(t *testing.T) {
	array := [15]float32{1, 2, 3, 4, 5, 6.0}
	got_ave := average(array)
	want_ave := float32(21.0 / 6)

	if got_ave != want_ave {
		t.Errorf(MESSAGE, got_ave, want_ave)
	}
}

func TestReverse(t *testing.T) {
	array := []int64{1, 2, 5, 15}
	got_ave := reverse(array)
	want_ave := []int64{15, 5, 2, 1}

	if !equal(got_ave, want_ave) {
		t.Errorf(MESSAGE, got_ave, want_ave)
	}
}

func TestSortMapValues(t *testing.T) {
	mapData := map[int]string{2: "a", 0: "b", 1: "c"}
	got := sortMapValues(mapData)
	want := []string{"b", "c", "a"}

	if !sameStringSlice(got, want) {
		t.Errorf(MESSAGE, got, want)
	}
}
func TestSortMapValues_2(t *testing.T) {
	mapData := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	got := sortMapValues(mapData)
	want := []string{"bb", "aa", "cc"}

	if !sameStringSlice(got, want) {
		t.Errorf(MESSAGE, got, want)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
