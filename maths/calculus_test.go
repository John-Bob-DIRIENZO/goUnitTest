package maths

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, -3)
	want := 2
	if result != want {
		t.Errorf("Add(5, -3) FAILED. Expected %d, got %d \n", want, result)
	} else {
		t.Logf("Add(5, -3) PASSED. Expected %d, got %d \n", want, result)
	}
}

func TestAdd2(t *testing.T) {
	result := Add(-5, -3)
	want := -8
	if result != want {
		t.Errorf("Add(-5, -3) FAILED. Expected %d, got %d \n", want, result)
	} else {
		t.Logf("Add(-5, -3) PASSED. Expected %d, got %d \n", want, result)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(5, -3)
	want := 8
	if result != want {
		t.Errorf("Substract(5, -3) FAILED. Expected %d, got %d \n", want, result)
	} else {
		t.Logf("Substract(5, -3) PASSED. Expected %d, got %d \n", want, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, -3)
	want := -15
	if result != want {
		t.Errorf("Multiply(5, -3) FAILED. Expected %d, got %d \n", want, result)
	} else {
		t.Logf("Multiply(5, -3) PASSED. Expected %d, got %d \n", want, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(-5, 0)
	want := float64(0)
	if result != want {
		t.Errorf("Divide(-5, 0) FAILED. Expected %f, got %f \n", want, result)
	} else {
		t.Logf("Divide(-5, 0) PASSED. Expected %f, got %f \n", want, result)
	}
}

func TestDivide2(t *testing.T) {
	result := Divide(-15, 3)
	want := float64(-5)
	if result != want {
		t.Errorf("Divide(-15, 3) FAILED. Expected %f, got %f \n", want, result)
	} else {
		t.Logf("Divide(-15, 3) PASSED. Expected %f, got %f \n", want, result)
	}
}
