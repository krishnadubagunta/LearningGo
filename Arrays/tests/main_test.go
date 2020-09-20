package tests

import (
	"github.com/krishnadubagunta/LearningGo/Arrays"
	"testing"
)

func TestArrays(t *testing.T) {
	want := [...]string{
		"a",
		"b",
		"c",
	}
	if got := Arrays.Main(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
