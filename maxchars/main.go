package main

import "fmt"

func main() {
	s := "apple 1231111111"
	fmt.Println(maxChars(s))
}

func maxChars(s string) (int, string) {
	max := 0
	maxChar := ""
	charMap := make(map[rune]int)
	chars := []rune(s)

	for i := 0; i < len(chars); i++ {
		if _, ok := charMap[chars[i]]; ok {
			charMap[chars[i]] = charMap[chars[i]] + 1
		} else {
			charMap[chars[i]] = 1
		}
	}

	for key, value := range charMap {
		if value > max {
			max = value
			maxChar = string(key)
		}
	}
	return max, maxChar
}
