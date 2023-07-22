//Worada Sajai
// 650510675
// Lab03_4
// 204203 Sec 002

package main

import (
	"strings"
)

const DEBUG = false

func roundToEven(x string, bPlace uint8) string {
	if !strings.Contains(x, ".") {
		x += ".0"
	}
	// s := [...]string{}
	// s := make([]byte,len(x))
	position_dot := strings.Index(x, ".") //int
	// fmt.Printf("%T", position_dot) // type check
	// fmt.Println(position_dot)
	// result := ""
	after_dot := len(x) - position_dot - 1
	interested_position := position_dot + int(bPlace)
	// for i := 1; i < after_dot; i++ {
	// 	after_dot_str += string(x[position_dot+1] - 48)
	// 	fmt.Println(i)
	// }
	println(bPlace, "--", after_dot)
	if bPlace >= uint8(after_dot) {
		x = x + strings.Repeat("0", int(bPlace)-after_dot)
		return x

	}

	result := ""

	if bPlace < uint8(after_dot) {
		if x[interested_position+1] == '0' {
			result = (x[:interested_position+1])
		} else { // 1....
			if strings.Contains(x[(interested_position+2):], "1") { // >1/2
				x = x[:interested_position+1]
				result = sum_string(x)
			} else { // =1/2
				if x[interested_position] == '1' {
					x = x[:interested_position+1]
					result = sum_string(x)
				} else {
					result = x[:interested_position+1]
				}
			}
		}

	}

	// strings.Contains()
	// strings.Repeat("o",6)
	if DEBUG {
		println(x)
		println(result)
	}
	if bPlace == 0 {
		result = result[:len(result)-1]
	}
	// fmt.Println(((x[interested_position]) - 48))

	return result
}

func sum_string(y string) string {
	result := ""
	for i := len(y) - 1; i >= 0; i-- { // '1'
		if y[i] == '1' {
			result = "0" + result
		} else if y[i] == '0' { // '0'
			result = "1" + result
			result = y[:i] + result
			return result
		} else { // '.'
			result = "." + result
		}
	}

	return "1" + result
}
