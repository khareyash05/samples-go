package utils

import (
	"testing"
)

// Test generated using Keploy

// Test generated using Keploy

func TestDivide_ValidInputsAndZeroDivision_004(t *testing.T) {
	result := Divide(10, 2)
	expected := 5
	if result != expected {
		t.Errorf("Divide(10, 2) = %d; want %d", result, expected)
	}

	result = Divide(10, 0)
	expected = 0
	if result != expected {
		t.Errorf("Divide(10, 0) = %d; want %d", result, expected)
	}
}

// Test generated using Keploy

func TestIsPalindrome_ValidInputs_006(t *testing.T) {
	result := IsPalindrome("racecar")
	if !result {
		t.Errorf("IsPalindrome(\"racecar\") = %v; want true", result)
	}

	result = IsPalindrome("hello")
	if result {
		t.Errorf("IsPalindrome(\"hello\") = %v; want false", result)
	}
}

func TestAdd_ValidInputs_001(t *testing.T) {
	result := Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Add(3, 5) = %d; want %d", result, expected)
	}
}

// Test generated using Keploy

func TestSubtract_ValidInputs_002(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; want %d", result, expected)
	}
}

// Test generated using Keploy

func TestMultiply_ValidInputs_003(t *testing.T) {
	result := Multiply(3, 7)
	expected := 21
	if result != expected {
		t.Errorf("Multiply(3, 7) = %d; want %d", result, expected)
	}
}

// Test generated using Keploy
