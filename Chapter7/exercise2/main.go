package main

import (
	"fmt"
)

func main() {
	var ok bool
	var count int
	data := []string{"a", "b", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok = counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}
