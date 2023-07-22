// worada sajai Aun
// 650510679
// HW03_2
// 204203 Sec 002

package main

import (
	"strings"
)

const MAX = 200 // really?

func baseNAddition(r1, r2 string, base int) string {
	decPointPos1 := strings.Index(r1, ".")
	decPointPos2 := strings.Index(r2, ".")
	point := 0
	if decPointPos1 < 0 {
		r1 += ".0"
		point++
	}
	if decPointPos2 < 0 {
		r2 += ".0"
		point++
	}
	decPointPos1 = strings.Index(r1, ".")
	decPointPos2 = strings.Index(r2, ".")
	result := []byte(strings.Repeat("x", MAX))
	r1_pointDigit := len(r1[decPointPos1+1:])
	r2_pointDigit := len(r2[decPointPos2+1:])
	if r1_pointDigit > r2_pointDigit {
		r2 = r2 + strings.Repeat("0", (r1_pointDigit-r2_pointDigit))
	} else {
		r1 = r1 + strings.Repeat("0", (r2_pointDigit-r1_pointDigit))
	}
	println(r1, "--", r2)
	i, j, k := len(r1)-1, len(r2)-1, MAX-1
	a := 0
	for ; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		// current digit
		currDigit := a
		// add the value from the current digit of n1
		if i >= 0 {
			if r1[i] == '.' {
				result[k] = '.'
				continue
			}
			currDigit += int(r1[i]) - int('0')
		}
		if j >= 0 {
			// add the value from the current digit of n2
			currDigit += int(r2[j]) - int('0')
		}

		a = 0
		if currDigit >= base {
			a = 1
			currDigit -= base
		}

		// convert back to byte (one char is called byte)
		result[k] = byte(currDigit + int('0')) //1='1'
		// fmt.Println(result)

	}
	if a == 1 {
		result[k] = '1'
		k--
	}
	if point == 2 {
		result = result[:len(result)-2]
	}

	return string(result[k+1:])

}
