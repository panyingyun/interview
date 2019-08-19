//############题目 SUM2 ##############
//#给定N个整数的数组S,在S中找到2个整数，使之和等于一个给定的数A（A=Si+Sj)
//#举例：
//#S=[12,7,11,15], target = 18
//#因为 S[1] + S[2] =7 + 11 = 18
//#返回 [1,2]

package main

import (
	"fmt"
)

func main() {
	s := []int{0, 2, 1, 3}
	target := 4
	ret1 := find1(s, target)
	fmt.Println("find1 ret1 = ", ret1)

	ret2 := find2(s, target)
	fmt.Println("find2 ret2 = ", ret2)
}

//暴力破解，时间复杂度O(n^2)
func find1(s []int, target int) []int {
	for k1, v1 := range s {
		for k2, v2 := range s {
			if v1+v2 == target && k1 != k2 {
				return []int{k1, k2}
			}
		}
	}
	return nil
}

//使用HashMap建立映射，然后使用r = target-s1, 然后查询r是否在HashMap中存在
//HashMap Key为s1 Value为 1
//(Key, Value) 即为 (12,0) (7,1)  (11, 2) (15,3)
//时间复杂度 O(n)
func find2(s []int, target int) []int {
	//建立HashMap
	m := make(map[int]int)
	for i, v := range s {
		m[v] = i
	}
	fmt.Println(m)
	//遍历和查询
	for i, v := range s {
		r := target - v
		k, ok := m[r]

		//当ℹ和k相同时，该配对不符合要求，自身和自身相加等于target的情况排除掉
		if ok && i != k {
			return []int{i, k}
		}
	}
	return nil
}
