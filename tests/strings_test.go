package main

import (
	"testing"

	"github.com/abhilashr1/linearequations/stringfilter"
)

func TestWithCoefficientInput(t *testing.T) {
	text1 := "5x+3y=1"
	text2 := "5x-9y=5"
	result, error := stringfilter.Filter(text1, text2)
	if error != nil {
		t.Errorf("Encountered an error: " + error.Error())
	}
	if len(result) == 0 {
		t.Errorf("Result is nil")
	}
}

func TestNoCoefficientInput(t *testing.T) {
	text1 := "x+y=1"
	text2 := "x-y=5"
	result, error := stringfilter.Filter(text1, text2)
	if error != nil {
		t.Errorf("Encountered an error: " + error.Error())
	}
	if len(result) == 0 {
		t.Errorf("Result is nil")
	}
}

func TestAddOneBefore(t *testing.T) {
	text1 := "x+y=1"
	result := stringfilter.AddOneBefore(text1, 0)
	if result == "" {
		t.Errorf("Result is empty")
	}
}

func TestFilterUtil(t *testing.T) {
	text1 := "x+y=1"
	result := stringfilter.FilterUtil(text1)
	if result == "" {
		t.Errorf("Result is empty")
	}
}

func TestGetVariables(t *testing.T) {
	text1 := "x+y=1"
	result := stringfilter.GetVariables(text1)
	if len(result) != 2 {
		t.Errorf("Result is incorrect")
	}
}
