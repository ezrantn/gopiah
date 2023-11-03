/**
 * Author: Ezra Natanael
 * File: go-lib.go
 */

package lib

import (
	"fmt"
	"strconv"
	"strings"
)

func reverseString(str []string) []string {
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	return str
}

func ConvertToRupiah(number int) string {
	tempNum := reverseString(strings.Split(fmt.Sprintf(strconv.Itoa(number)), ""))
	rupiah := ""

	for i := 0; i < len(tempNum); i++ {
		if (i+1)%3 == 0 && i != len(tempNum)-1 {
			tempNum[i] = "." + tempNum[i]
		}
	}

	rupiah = "Rp. " + strings.Join(reverseString(tempNum), "") + ",00"
	return rupiah
}
