// Package gopiah_test provides test cases for the gopiah package.
// It tests the functionality of converting numbers into Indonesian Rupiah format.
package gopiah

import (
	"github.com/ezrantn/gopiah"
	"testing"
)

// TestConvertToRupiah tests the ConvertToRupiah function in the gopiah package.
// It verifies whether the function correctly converts numbers to Indonesian Rupiah format.
func TestConvertToRupiah(t *testing.T) {
	// Define test cases with input numbers and their expected Rupiah format strings
	testCases := []struct {
		input    int
		expected string
	}{
		{0, "Rp. 0,00"},
		{12345, "Rp. 12.345,00"},
		{123456789, "Rp. 123.456.789,00"},
	}

	for _, tc := range testCases {
		result := gopiah.ConvertToRupiah(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected %s, but got %s", tc.input, tc.expected, result)
		}
	}

	// Output:
	// 0 = Rp. 0,00
	// 12345  = Rp. 12.345,00
	// 123456789 = Rp. 123.456.789,00
}
