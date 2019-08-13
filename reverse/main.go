//############Reverse##############

package main

import (
	"fmt"
)

func main() {
	a := 123
	fmt.Println("reverse  = ", reverse(a)) // 321

	a = -123
	fmt.Println("reverse  = ", reverse(a)) // -321
}

func reverse(n int) int {
	if n == 0 {
		return 0
	}

	ret := 0

	for n != 0 {
		low := n % 10
		n = n / 10
		ret = ret*10 + low
	}
	return ret
}
