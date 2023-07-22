package main

import (
	"math"
	"strings"
)

// import "strings"

func addNSubtract(n1, n2 string, bitLen uint8) (int64, int64) {

	n1 = addIndex(n1, bitLen)
	n2 = addIndex(n2, bitLen)
	result_add := sum2stringbin(n1, n2)
	result_sub := sum2stringbin(n1, twocom(n2))

	return twosComplToInt(result_add), twosComplToInt(result_sub)
}

func twocom(n1 string) string {
	i := len(n1) - 1
	result := ""
	for i >= 0 {
		if n1[i] == '0' {
			result = "0" + result
		} else {
			result = "1" + result
			i--
			for i >= 0 {
				if n1[i] == '1' {
					result = "0" + result
				} else {
					result = "1" + result
				}
				i--
			}
		}
		i--
	}
	return result
}

func addIndex(x string, bitLen uint8) string {

	if len(x) < int(bitLen) {
		if x[0] == '1' {
			x = strings.Repeat("1", int(bitLen)-len(x)) + x
		} else {
			x = strings.Repeat("0", int(bitLen)-len(x)) + x
		}
	}

	return x[len(x)-int(bitLen):]
}

func sum2stringbin(n1, n2 string) string {
	carry := 0
	result := ""
	for i := len(n1) - 1; i >= 0; i-- {
		currdigit := carry
		currdigit += int(n1[i] - '0')
		currdigit += int(n2[i] - '0')
		carry = currdigit / 2
		result = string(byte(currdigit%2)+'0') + result
	}
	return result
}

func twosComplToInt(x string) int64 {

	var result int64 = 0
	// fmt.Println(x[0] - '0')
	var Ans int64
	for i := 0; i < len(string(x)); i++ {
		a := int64(x[i] - 48)
		if i == 0 {
			Ans = -int64(math.Pow(2, float64(len(x)-1))) * a
		} else {
			Ans = int64(math.Pow(2, float64(len(x)-i-1))) * a
		}
		result += Ans
	}

	return result
}
