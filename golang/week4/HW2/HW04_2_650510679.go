// Worada Sajai
// 650510679
// HW04_2
// 204203 Sec 002

package main

import (
	"strconv"
	"strings"
)

func ipv4Decode(ipUint uint32) string { //convert unit32 to string

	base_2 := strconv.FormatInt(int64(ipUint), 2)
	// println(base_2)
	// base_2 := strconv.Itoa(base_2)
	// println(base_2)
	base_2 = strings.Repeat("0", 32-len(base_2)) + base_2
	// result_arr := strings.SplitAfterN(base_2, ".", 4)

	g1, _ := strconv.ParseInt(base_2[:8], 2, 64)
	g2, _ := strconv.ParseInt(base_2[8:16], 2, 64)
	g3, _ := strconv.ParseInt(base_2[16:24], 2, 64)
	g4, _ := strconv.ParseInt(base_2[24:], 2, 64)
	// println(g1)

	return strconv.FormatInt(g1, 10) + "." + strconv.FormatInt(g2, 10) + "." + strconv.FormatInt(g3, 10) + "." + strconv.FormatInt(g4, 10)

}
func ipv4Encode(ipString string) uint32 {

	numSet := strings.Split(ipString, ".") //[10 28 4 2]
	stringNum := ""
	for i := len(numSet) - 1; i >= 0; i-- {
		intfromstr, _ := strconv.Atoi(numSet[i])
		stringNum = strings.Repeat("0", 8-len(strconv.FormatInt(int64(intfromstr), 2))) + strconv.FormatInt(int64(intfromstr), 2) + stringNum
		// println(stringNum)
	}
	intNum, _ := strconv.ParseInt(stringNum, 2, 64)
	// println(intNum)
	// g1 ,_ := strings.Repeat("0", 32-len(g1)) + g1

	return uint32(intNum)
}
