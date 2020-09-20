package tests

import (
	"github.com/krishnadubagunta/LearningGo/Slices"
	"reflect"
	"testing"
)

func TestSlicesCompositeLiterals(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := Slices.ValuesOfSameType()
	for i, wantV := range want {
		if wantV != got[i] {
			t.Errorf("Not equal, expected to be equal")
		}
	}
}

func TestSlicingASlice(t *testing.T) {
	want := []int{4, 5, 7, 8, 42}
	for i := 0; i < 5; i++ {
		wantI := want[i:]
		gotI := Slices.SlicingASlice(i, want)
		for j, wantJ := range wantI {
			if wantJ != gotI[j] {
				t.Errorf("Failed, Expected to be equal")
			}
		}
	}
}

func TestAppendSlices(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if got := Slices.AppendToSlice(slice1, slice2...); reflect.TypeOf(got) != reflect.TypeOf(want) && len(got) != len(want) {
		t.Errorf("Not Equal Arrays")
	}

	if got := Slices.ConcatTwoSlices(slice1, slice2); reflect.TypeOf(got) != reflect.TypeOf(want) && len(got) != len(want) {
		t.Errorf("Not Equal Arrays")
	}
}
