package main

import "fmt"

func main() {
	fmt.Println(palindrome(reverse("abba")))
}

func palindrome(s string, _s string) bool {
	return s == _s
}

func reverse(s string) (string, string) {
	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return s, string(chars)
}
