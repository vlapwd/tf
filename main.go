package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	var ret []rune
	for _, v := range s {
		ret = append(ret, v)
		ret = append(ret, '"')
	}
	fmt.Println(string(ret))
}
