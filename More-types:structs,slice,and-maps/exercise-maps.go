package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	str := strings.Fields(s)

	for _, v := range str {
		m[v]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
