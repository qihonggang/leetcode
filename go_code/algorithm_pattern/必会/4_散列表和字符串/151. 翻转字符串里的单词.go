package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
给定一个字符串，逐个翻转字符串中的每个单词。

说明：

无空格字符构成一个 单词 。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
 */

func reverseWords(s string) string {
	chs := strings.Split(s, " " )
	var strs []string
	for i := len(chs) - 1; i >= 0; i-- {
		if len(chs[i]) > 0 {
			strs = append(strs, chs[i])
		}
	}
	return strings.Join(strs, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Print(reverseWords(input))
}