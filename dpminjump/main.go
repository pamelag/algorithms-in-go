package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 1, 3, 2, 3, 4, 5, 1, 2, 8}
	x := 1000 // mimics a large value close to infinity
	minJumpArray := [10]int{x, x, x, x, x, x, x, x, x, x}
	jumpPathArray := [10]int{0, x, x, x, x, x, x, x, x, x}
	minJumpArray[0] = 0

	for i := 1; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			// check if i can be reached using the value at j-th index in numbers array "from" the "current j-th index"
			if i <= j+numbers[j] {

				// As the computation proceeds the infinite large value would be
				// replaced by the jump index value at the first instance the inifintely
				// large value (1000 in this case) is encountered. However, once replaced
				// in the later iterations it would be compared as like against like
				// with another jump index to get the minimum

				// there is always 1 jump needed to get from j to i, and there would have
				// been a jump required to get to the j-th index, which would now be retreived
				// from the memoized data in the minJumpArr (stored in the j-th index)
				// hence minJumpArr[j]+1 is compared with the current value of minJumpArr[i]
				// which could either be a large number if its the first time a j is able
				// to reach i or a relatively smaller value if there was another j which
				// could reach i at an earlier iteration.

				//checking which one is minimum
				if minJumpArray[i] > minJumpArray[j]+1 {
					minJumpArray[i] = minJumpArray[j] + 1
					jumpPathArray[i] = j

				}

			}
		}
	}

	fmt.Println("minJumpArr ", minJumpArray)
	fmt.Println("jumpPathArr ", jumpPathArray)
}
