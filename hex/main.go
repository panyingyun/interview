//############To Hex ##############

package main

import (
	"fmt"
)

func main() {
	a := 26
	fmt.Printf("%x \n", a)
	fmt.Printf("%v \n", toHex(a))

	a = 1000
	fmt.Printf("%x \n", a)
	fmt.Printf("%v \n", toHex(a))
}

func toHex(num int) string {
	s := "0123456789abcdef"
	if num == 0 {
		return "0"
	}
	ret := ""
	for num != 0 {
		high := num % 16
		ret = s[high:high+1] + ret
		num = num / 16
	}
	return ret
}
