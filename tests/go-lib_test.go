package tests

import (
	"github.com/umjiiii/gopiah.git/lib"
	"testing"
)

func TestConvertToRupiah(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{0, "Rp. 0,00"},
		{12345, "Rp. 12.345,00"},
		{123456789, "Rp. 123.456.789,00"},
	}

	for _, tc := range testCases {
		result := lib.ConvertToRupiah(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
