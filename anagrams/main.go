package main

import "fmt"

func main() {
	s1 := "rail safety"
	s2 := "fairy tales"
	fmt.Println(anagrams(s1, s2))
}

func anagrams(s1 string, s2 string) bool {
	chars1 := []rune(s1)
	chars2 := []rune(s2)
	charMap1 := make(map[rune]int)
	charMap2 := make(map[rune]int)

	for i := 0; i < len(chars1); i++ {
		if _, ok := charMap1[chars1[i]]; ok {
			charMap1[chars1[i]] = charMap1[chars1[i]] + 1
		} else {
			charMap1[chars1[i]] = 1
		}
	}

	for i := 0; i < len(chars2); i++ {
		if _, ok := charMap2[chars2[i]]; ok {
			charMap2[chars2[i]] = charMap2[chars2[i]] + 1
		} else {
			charMap2[chars2[i]] = 1
		}
	}

	for key, value := range charMap1 {
		if value != charMap2[key] {
			return false
		}
	}

	return true
}
