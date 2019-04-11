package main

func main() {
	totalLength := 5
	lengthOfPieces := [4]int{1, 2, 3, 4}
	costOfPieces := [4]int{2, 5, 7, 8}

	var memoization [5][6]int

	memoization[0][0] = 0
	memoization[0][1] = 0
	memoization[0][2] = 0
	memoization[0][3] = 0
	memoization[0][4] = 0
	memoization[0][5] = 0

	memoization[0][0] = 0
	memoization[1][0] = 0
	memoization[2][0] = 0
	memoization[3][0] = 0
	memoization[4][0] = 0
	profit := 0

	for i := 1; i < totalLength; i++ {
		for j, k := 1, 1; j < len(memoization[i]); j++ {
			if lengthOfPieces[k] == j {
				if i >= j {
					division := i / j
					remainder := i % j

					profit = costOfPieces[j] * ((division * j) + remainder)
					memoization[i][j] = profit

				} else {
					unaffectedProfit := memoization[i-1][j]
					memoization[i][j] = unaffectedProfit
				}
				k++
			}
		}
	}
}
