/*
* @Author: kk
* @Date:   2022-06-26 21:47:44
* @Last Modified by:   kk
* @Last Modified time: 2022-06-26 23:00:10
 */
package main

import "math"

func additiveInverse(x string) (string, int64) {
	result := ""
	i := len(x) - 1
	for i >= 0 {
		if x[i] == '1' {
			result = "1" + result
			i--
			for i >= 0 {
				if x[i] == '0' {
					result = "1" + result
				} else {
					result = "0" + result
				}
				i--
			}
		} else {
			result += "0"
			i--
		}

	}

	return result, twosComplToInt(result)
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
