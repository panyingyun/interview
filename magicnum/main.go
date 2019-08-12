//############题目 回文数 ##############
//#判断一个整数是否是回文数
//#回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数
//#举例：
//#121 是回文数  1221 是回文数   12321是回文数
//#
//#返回 true

package main

import (
	"fmt"
)

func main() {
	n1 := 123
	ret1 := isMagicNum(n1)
	fmt.Println("ret1 = ", ret1)

	n2 := 1221
	ret2 := isMagicNum(n2)
	fmt.Println("ret2 = ", ret2)

	n3 := 12321
	ret3 := isMagicNum(n3)
	fmt.Println("ret3 = ", ret3)
}

//将输入的数据反过来 123 --> 321，然后和原来的值进行比较即可

func isMagicNum(n int) bool {
	if n < 0 {
		return false
	}
	pre := n
	rev := 0
	for {
		rev = rev*10 + n%10 //取最后一位
		n = n / 10
		//fmt.Println("rev = ", rev)
		//fmt.Println("n = ", n)
		if n == 0 {
			break
		}
	}
	//fmt.Println("n = ", n)
	return pre == rev
}
