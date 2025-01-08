package main

import (
	"fmt"
	"strings"
)

func main() {
	r := IsPalindrome("a")
	fmt.Println(r)
}

func IsPalindrome(str string) bool {
	word := strings.Split(str, "")
	var backword string
	for i := len(word) - 1; i >= 0; i-- {
		fmt.Println(i)
		backword += word[i]
	}
	fmt.Println(str)
	fmt.Println(backword)
	return str == backword
}
