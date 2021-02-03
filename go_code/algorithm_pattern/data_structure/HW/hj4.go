package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
字符串分割

题目描述
•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。

输入描述:
连续输入字符串(输入多次,每个字符串长度小于100)

输出描述:
输出到长度为8的新字符串数组
 */
func main() {
	read := bufio.NewReader(os.Stdin)
	for {
		in, _, _ := read.ReadLine()
		if len(in) == 0 {
			break
		}
		stringInput := string(in)
		for len(stringInput) >= 8 {
			fmt.Println(stringInput[0:8])
			stringInput = stringInput[8:]
		}

		if stringInput != "" {
			length := 8 - len(stringInput)
			for i := 0; i < length; i++ {
				stringInput += "0"
			}
			fmt.Println(stringInput)
		}
	}
}
