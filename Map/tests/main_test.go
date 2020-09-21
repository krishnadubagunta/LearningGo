package tests

import (
	"github.com/krishnadubagunta/LearningGo/Map"
	"testing"
)

func TestStringIntegerGetValue(t *testing.T) {
	currentStrIntVal := Map.StringInteger{
		"James": 1,
		"Bond":  2,
	}
	want := 1;
	if v := currentStrIntVal.GetValueAdd("James"); *v != want {
		t.Errorf("Expected to return %d, but returned %d", want, *v)
	}

	if v := currentStrIntVal.GetValueAdd("Barnabas"); v != nil {
		t.Errorf("Expected to return nil, but got %d instead", *v)
	}
}

func TestIntegerStringGetValue(t *testing.T) {
	currentIntStrVal := Map.IntegerString{
		1: "James",
		2: "Bond",
	}

	want := "James"
	if v := currentIntStrVal.GetValueAdd(1); *v != want {
		t.Errorf("Expected to return %s, but returned %s", want, *v)
	}

	if v := currentIntStrVal.GetValueAdd(3); v != nil {
		t.Errorf("Expected to return nil, but got %s instead", *v)
	}
}