package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var str string
	var char1 uint8
	fmt.Scanf("%s", &str)
	fmt.Scanf("%c", &char1)
	count, err := StringCountWord(str, char1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	return
}

func StringCountWord(s string, c uint8) (int, error){
	var count int = 0
	if len(s) == 0 {
		return 0, errors.New("nil string")
	}
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count, nil
}
