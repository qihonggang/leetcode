package main

import "strconv"

/**
根据 逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

 

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation
 */

func evalRPN(tokens []string) int {
	number := []int{}
	for _, v := range tokens {
		l := len(number)
		switch v {
		case "+":
			number = append(number[:l-2], number[l-2] + number[l-1])
		case "-":
			number = append(number[:l-2], number[l-2] - number[l-1])
		case "*":
			number = append(number[:l-2], number[l-2] * number[l-1])
		case "/":
			number = append(number[:l-2], number[l-2] / number[l-1])
		default:
			num, _ := strconv.Atoi(v)
			number = append(number, num)
		}
	}
	return number[0]
}