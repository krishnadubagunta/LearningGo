package tests

import (
	"github.com/krishnadubagunta/LearningGo/Hello"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}