package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 0
	checker := []string{"a", "e", "i", "o", "u"}
	s := "pamela"

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(checker); j++ {
			if strings.Contains(string(s[i]), checker[j]) {
				count++
			}
		}
	}

	fmt.Println("Count is ", count)
}
