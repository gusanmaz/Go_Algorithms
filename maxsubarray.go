package main

/* O(log n) Divide and Conquer Implementation of Maximum Sub Array Problem
See https://en.wikipedia.org/wiki/Maximum_subarray_problem for details
of the problem.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

type MaxMiner interface {
	Max(...int) int
	Min(...int) int
}

type SuperInt int
type Slice []SuperInt

func (slice Slice) Max() int {
	maxVal := slice[0]
	for _, k := range slice {
		if k > maxVal {
			maxVal = k
		}
	}
	return int(maxVal)
}

func (slice Slice) Min() int {
	minVal := slice[0]
	for _, k := range slice {
		if k < minVal {
			minVal = k
		}
	}
	return int(minVal)
}

func generateRandArr(minVal, maxVal, len int) []int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	rang := maxVal - minVal + 1
	randArr := make([]int, len)
	for i := 0; i < len; i++ {
		randArr[i] = minVal + r.Intn(rang)
	}
	return randArr
}

func findLeftMaxSubArr(arr []int, minInd, maxInd int) (ind, val int) {
	var arrSum = arr[maxInd]
	var maxSum = arrSum
	var maxSumInd = maxInd
	for i := maxInd - 1; i >= minInd; i-- {
		arrSum += arr[i]
		if arrSum > maxSum {
			maxSumInd = i
			maxSum = arrSum
		}
	}
	return maxSumInd, maxSum
}

func findRightMaxSubArr(arr []int, minInd, maxInd int) (ind, val int) {
	var arrSum = arr[minInd]
	var maxSum = arrSum
	var maxSumInd = minInd
	for i := minInd + 1; i <= maxInd; i++ {
		arrSum += arr[i]
		if arrSum > maxSum {
			maxSumInd = i
			maxSum = arrSum
		}
	}
	return maxSumInd, maxSum
}

func maxSubArray(arr []int, beginInd, endInd int) (minInd, maxInd, val int) {
	if (endInd - beginInd) == 0 {
		return beginInd, endInd, arr[beginInd]
	} else {
		minInd = beginInd
		maxInd = endInd
		midInd := (maxInd + minInd) / 2
		lMinInd, lMaxInd, lVal := maxSubArray(arr, minInd, midInd)
		rMinInd, rMaxInd, rVal := maxSubArray(arr, midInd+1, maxInd)
		lInd, lMinVal := findLeftMaxSubArr(arr, minInd, midInd)
		rInd, rMaxVal := findRightMaxSubArr(arr, midInd+1, maxInd)

		mVal := rMaxVal + lMinVal
		values := Slice{SuperInt(lVal), SuperInt(rVal), SuperInt(mVal)}
		maxVal := values.Max()

		switch maxVal {
		case lVal:
			return lMinInd, lMaxInd, lVal
		case rVal:
			return rMinInd, rMaxInd, rVal
		case mVal:
			return lInd, rInd, mVal
		default:
			// Program counter should never reach here
			return -1, -1, -999
		}
	}
}

func main() {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15 - 4, 7}
	arr = generateRandArr(-20, 20, 25)
	a, b, c := maxSubArray(arr, 0, len(arr)-1)
	fmt.Println("Our randomly generated array: ", arr)
	fmt.Println(fmt.Sprintf("Start Ind: %d End Ind: %d Sum: %d of maximum subarray of our array.", a, b, c))
}
