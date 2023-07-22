// worada sajai
// 650510679
// Hw03_3
// 204203 Sec 002

package main

import (
	"math"
	"strings"
	// "fmt"
)

const fracLen = 7
const expLen = 8

const BASE uint8 = 2
const DEBUG = false

// from Lab01_2
const MAX_INT = 64
const DEC_PLACE = 7

func float16bitNormed(n float32) string {
	// expLen = 8, fracLen = 7
	var bias = int(pow(2, expLen-1) - 1) // bias = 127
	sign := "0"
	if DEBUG {
		println("Bias", bias)
	}

	if n < 0 {
		n = -n
		sign = "1"
	}

	var minNorm float64 = 0.0000000000000000000000000000000000000118 // dummy value
	//0.0000000000000000000000000000000000000000001
	var maxNorm float64 = 338953138925153547590470800371487866880 // dummy value

	if (float64(n) > maxNorm) || (float64(n) < minNorm) {
		if DEBUG {
			println(n, "overflow")
		}
		return ""
	}

	E := int(math.Floor(math.Log(float64(n)) / math.Log(2)))
	exp := E + bias

	n = n / float32(pow(2, E))
	// println(float64(n))
	n = n - 1

	return sign + " " + strings.Repeat("0", 8-len(posIntToBaseB(int64(exp), 2))) + posIntToBaseB(int64(exp), 2) + " " + fractionToBaseB(float64(n), 2)
}
func pow(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

//------------------------------------ Lab01_2

func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 {
		x = -x
		sign = "-"
	}
	// split at the decimal point
	front := int64(x)
	back := x - float64(front)

	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	converted := sign + frontStr + "." + backStr

	return converted
}

func fractionToBaseB(x float64, b uint8) string {

	result := []byte(strings.Repeat("x", DEC_PLACE))
	var currDigit byte

	for i := 0; i < DEC_PLACE; i++ {
		x = x * float64(b)
		front := int64(x)
		currDigit = byte((front) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		x = x - float64(front)
		result[i] = currDigit

	}

	return string(result[:])
}

func posIntToBaseB(x int64, b uint8) string {

	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1
	var currDigit byte

	for x > 0 {
		currDigit = byte((x % int64(b)) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		result[k] = currDigit
		x = x / int64(b)
		k--
	}

	return string(result[k+1:])
}
