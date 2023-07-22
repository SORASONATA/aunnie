//Worada Sajai
//650510679
//Lab02_3
//204203
package main

import (
	"math"
)

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
