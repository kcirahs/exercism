package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Portable Network Graphics"
	sp := strings.Split(s, " ")
	var acrbyte []byte
	for i, _ := range sp {
		v := []byte(sp[i])[0]
		acrbyte = append(acrbyte, v)
	}
	fmt.Println(string(acrbyte))
	fmt.Println(sp)
}
