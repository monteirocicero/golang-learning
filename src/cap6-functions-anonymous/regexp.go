package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	expr := regexp.MustCompile("\\b\\w")

	transformer := func(s string) string {
		return strings.ToUpper(s)
	}

	text := "alterto santos dumont"
	fmt.Println(transformer(text))
	fmt.Println(expr.ReplaceAllStringFunc(text, transformer))
}
