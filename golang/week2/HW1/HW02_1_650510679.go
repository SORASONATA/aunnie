// Worada Sajai
// 650510679
// HW02_1
// 204203 Sec 002

package main

import (
	"fmt"
	"log"
	"strings"
)

// set the maximum digit length. why 36 and not 35?
const MAX = 36

func main() {
	// why are we using string?
	var n1, n2 string

	// read in two string (can be multiple lines)
	_, err := fmt.Scan(&n1, &n2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addition(n1, n2))
}

func addition(n1, n2 string) string {
	// this is just a skeleton and will give out wrong result
	result := []byte(strings.Repeat("x", MAX))
	len1 := len(n1)
	len2 := len(n2)
	// fmt.Println(result)

	// loop from the left most digit
	i, j, k := len1-1, len2-1, MAX-1

	// loop from right most position
	a := 0
	for ; i >= 0; i, j, k = i-1, j-1, k-1 {
		// current digit
		currDigit := 0 + a
		a = 0
		// add the value from the current digit of n1
		currDigit += int(n1[i]) - int('0')

		if j >= 0 {
			// add the value from the current digit of n2
			currDigit += int(n2[j]) - int('0')
		}

		if currDigit > 9 {
			a = 1
			currDigit -= 10
		}

		// convert back to byte (one char is called byte)
		result[k] = byte(currDigit + int('0')) //1='1'

		// fmt.Println(result)
	}
	for ; j >= 0; j, k = j-1, k-1 {
		currDigit := 0 + a
		a = 0
		currDigit += int(n2[j]) - int('0')

		if currDigit > 9 {
			a = 1
			currDigit -= 10
		}

		// convert back to byte (one char is called byte)
		result[k] = byte(currDigit + int('0')) //1='1'
	}
	if a == 1 {
		result[k] = '1'
		k--
	}
	// convert array of bytes to string
	return string(result[k+1:])

}
