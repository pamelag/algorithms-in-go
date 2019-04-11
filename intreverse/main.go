package main

import "fmt"
import "strconv"

func main() {
	values := "123456789"
	sValues := string(values)
	reversed := reverse(sValues)
	i, err := strconv.Atoi(reversed)
	if err != nil {
		fmt.Println("string conversion failed")
	}
	fmt.Println(i)
}

func reverse(s string) string {
	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
