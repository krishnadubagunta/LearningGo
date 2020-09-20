package tests

import (
	"github.com/krishnadubagunta/LearningGo/Slices"
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
