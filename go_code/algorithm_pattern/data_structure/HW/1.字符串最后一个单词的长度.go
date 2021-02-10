package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	stringData := string(data)
	arrData := strings.Fields(stringData)
	arrLen := len(arrData)
	if arrLen > 0 {
		fmt.Println(len(arrData[arrLen-1]))
	} else {
		fmt.Println(0)
	}

}
