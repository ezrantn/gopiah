// Package gopiah provides functionality to convert numbers into Indonesian Rupiah (Rp) currency format.
// It's a convenient tool for formatting monetary values in the Indonesian style, with thousands separators and decimal points.
// The return value of the conversion function is of type string.
package gopiah

import (
	"strconv"
	"strings"
)

// ConvertToRupiah converts any given number into the Indonesian Rupiah format.
// It takes an integer number as input and returns the Rupiah format as a string.
func ConvertToRupiah(number int) string {
	// Convert the number to a string and reverse it
	tempNum := reverseString(strings.Split(strconv.Itoa(number), ""))
	rupiah := ""

	// Insert thousands separators
	for i := 0; i < len(tempNum); i++ {
		if (i+1)%3 == 0 && i != len(tempNum)-1 {
			tempNum[i] = "." + tempNum[i]
		}
	}

	// Join the reversed string and add Rp. prefix and ",00" suffix for Rupiah format
	rupiah = "Rp. " + strings.Join(reverseString(tempNum), "") + ",00"
	return rupiah
}

// reverseString reverses the order of elements in a string slice.
func reverseString(str []string) []string {
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	return str
}