package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a = 0
	fmt.Println(a+80)
	a = 1
	var b = "hello!"
	var kr = "한글"

	fmt.Println(a,"+", len(b), "=", a+len(b))
	fmt.Println(utf8.RuneCountInString(kr))
}