package main

import (
	"testing"

	"github.com/abhilashr1/linearequations/compute"
)

func TestSolution(t *testing.T) {
	// 8x + 5y = 2
	// 2x - 4y = -10
	d := [2][2]float32{
		{8, 5},
		{2, -4},
	}
	d1 := [2][2]float32{
		{2, 5},
		{-10, -4},
	}
	d2 := [2][2]float32{
		{8, 2},
		{2, -10},
	}
	val1, val2 := compute.Solution(d, d1, d2)
	if val1 != -1 && val2 != 2 {
		t.Errorf("Result is incorrect")
	}

}

func TestSolution2(t *testing.T) {
	// 5x + 2y = 4
	// 3x + 2y = 8
	d := [2][2]float32{
		{5, 2},
		{3, 2},
	}
	d1 := [2][2]float32{
		{4, 2},
		{8, 2},
	}
	d2 := [2][2]float32{
		{5, 4},
		{3, 8},
	}
	val1, val2 := compute.Solution(d, d1, d2)
	if val1 != -2 && val2 != 7 {
		t.Errorf("Result is incorrect")
	}

}

func TestDeterminant(t *testing.T) {
	d := [2][2]float32{
		{5, 2},
		{3, 2},
	}
	d1 := [2][2]float32{
		{4, 2},
		{8, 2},
	}
	d2 := [2][2]float32{
		{5, 4},
		{3, 8},
	}
	val1 := compute.Determinant(d)
	val2 := compute.Determinant(d1)
	val3 := compute.Determinant(d2)

	if val1 != 4 && val2 != -8 && val3 != 28 {
		t.Errorf("Result is incorrect")
	}

}
