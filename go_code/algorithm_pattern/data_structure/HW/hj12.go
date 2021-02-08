package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
/**
字符串反转

题目描述
接受一个只包含小写字母的字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

输入描述:
输入一行，为一个只包含小写字母的字符串。

输出描述:
输出该字符串反转后的字符串。
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")
	if len(line) == 0 || len(line) > 1000 {
		return
	}
	for i := len(line) - 1; i >= 0; i-- {
		fmt.Print(string(line[i]))
	}
}
