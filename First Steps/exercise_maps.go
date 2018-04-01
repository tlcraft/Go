package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m map[string]int = make(map[string]int)
	var fields []string = strings.Fields(s)

	for _, value := range fields {
		m[value] = m[value] + 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
