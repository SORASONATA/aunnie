// WOrada Sajai
// 650510679
// Lab01_2
// 204203 Sec 002
package main

import (
	"strings"
)

const MAX_INT = 64
const DEC_PLACE = 6

func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 { // turn negative numbers to positive
		x = -x
		sign = "-"
	}
	// split at the decimal point
	front := int64(x)
	back := x - float64(front)

	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	// putting every part together
	converted := sign + frontStr + "." + backStr

	return converted

}

func fractionToBaseB(x float64, b uint8) string {
	// fmt.Println(x)
	// only need to implement this function
	//var result string
	res := ""
	for i := 0; i < 6; i++ {
		// fmt.Println(i)
		x *= float64(b)
		// fmt.Println(int(x))
		result := "0"
		if x >= 1 {
			if int(x) > 9 {
				result = string('A' + byte(int(x)-10))
			} else {
				result = string(byte(int(x) + '0'))
				// fmt.Println(string(byte(int(x) + '0')))
			}

			// res += result
			// x = float64(x) - float64(int64(x))
			// fmt.Println(x)
		}
		res += result
		x = float64(x) - float64(int64(x))
	}
	// fmt.Println(res)
	return res
}

func posIntToBaseB(x int64, b uint8) string {
	// this function is working correctly
	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1
	var currDigit byte

	for x > 0 {
		// calculate and convert back to char
		currDigit = byte((x % int64(b)) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		result[k] = currDigit
		x = x / int64(b)
		k--
	}
	// convert from byte array to string
	return string(result[k+1:])
}
