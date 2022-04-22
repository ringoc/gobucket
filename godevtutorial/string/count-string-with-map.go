package main

import (
	"fmt"
)

func main() {

	var str string
	fmt.Print("Please input your string: ")
	fmt.Scanf("%s", &str)
	m := make(map[string]int)

	for pos, char := range str {
		if val, found := m[string(char)]; found {
			m[string(char)] = val + 1
		} else {
			m[string(char)] = 1
		}
		fmt.Printf("pos: %d / char: %c / count: %d\n", pos, char, m[string(char)])
	}
	fmt.Println("map:", m)
}
