package util

import (
	"fmt"
	"strconv"
)

func tansUtil(){
	// byte转数字
	s := "12345" // s[0] 类型是byte
	num := int(s[0] - '0') // 1
	str := string(s[0]) // "1"
	b := byte(num + '0') // '1'
	fmt.Printf("%d%s%n", num, str, b) // 111

	// 字符串转数字
	num2, _ := strconv.Atoi(s)
	str2 := strconv.Itoa(num)

}
