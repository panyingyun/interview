//############题目 字符串压缩 #############

package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "cccddecc"
	s1 := compress(s)
	fmt.Println("compress string = ", s1)

	s = "adef"
	s1 = compress(s)
	fmt.Println("compress string = ", s1)

	s = "pppppppp"
	s1 = compress(s)
	fmt.Println("compress string = ", s1)

	s = "abbccffffeeee"
	s1 = compress(s)
	fmt.Println("compress string = ", s1)
}

func compress(str string) string {
	orgarray := []byte(str)
	newarray := []byte("")

	//fmt.Println("array = ", orgarray)
	//fmt.Println("narray = ", newarray)
	count := 1
	for i, v := range orgarray {
		if i < len(orgarray)-1 && orgarray[i+1] == v {
			count++
		} else {
			if count == 1 {
				newarray = append(newarray, v)
				//fmt.Println("narray = ", newarray)
			} else {
				newarray = append(newarray, v)
				newarray = append(newarray, strconv.Itoa(count)...)
				//fmt.Println("narray = ", newarray)
			}
			count = 1
		}
	}
	return string(newarray)
}
