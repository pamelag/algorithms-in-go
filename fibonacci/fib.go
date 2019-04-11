package main

import (
	"fmt"
	"math/big"
)

var memoize map[int]*big.Int

func init() {
	memoize = make(map[int]*big.Int)
}

func fib_mem(n int) *big.Int {
	if n < 0 {
		return nil
	}

	// base case
	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	// check if we have stored it before
	if val, ok := memoize[n]; ok {
		return val
	}

	// initialize map then add 2 previous fib values
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], fib_mem(n-1))
	memoize[n].Add(memoize[n], fib_mem(n-2))

	return memoize[n]

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", fib_mem(i))
	}
	fmt.Println()
}
