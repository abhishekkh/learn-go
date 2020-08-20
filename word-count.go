package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m:= make(map[string]int)
	words:= strings.Fields(s)
	for _, word := range words{
		elem := m[word]
		m[word] = elem + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}