package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
句子逆序

题目描述
将一个英文语句以单词为单位逆序排放。例如“I am a boy”，逆序排放后为“boy a am I”
所有单词之间用一个空格隔开，语句中除了英文字母外，不再包含其他字符

输入描述:
输入一个英文语句，每个单词用空格隔开。保证输入只包含空格和字母。

输出描述:
得到逆序的句子
 */
func main() {
	results := []string{}
	reader :=  bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	if len(line) == 0 {
		return
	}
	line = strings.Trim(line, "\r\n")
	words := strings.Split(line, " ")
	for i := len(words) - 1; i >= 0; i-- {
		results = append(results, words[i])
	}
	fmt.Print(strings.Join(results, " "))
}
