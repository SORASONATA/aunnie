// worada sajai
// 650510679
// HW01_1
// Problem A
// 204203 Sec 002

package main

func factorial(num1 int8) int64 {
	var result int64 = 1
	// if num1 < 2{
	// 	return result
	// }
	// else{
	// 	result*=num1
	// 	return result + factorial(num-1)
	// }

	for i := 1; i < int(num1)+1; i++ {
		result *= int64(i)
	}
	return result
}
