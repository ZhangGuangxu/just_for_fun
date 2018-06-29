package main

import (
	"errors"
	"fmt"
)

// 一个整型数组，里面有n个数，包括正数、0、负数。找出和最大的子数组。要求算法时间复杂度O(n)

var errEmptyArray = errors.New("empty array")

func main() {
	//s := []int{1, 2, 3}
	//fmt.Println(s[0 : len(s)+1]) // error: slice bounds out of range
	//fmt.Println(s[0:len(s)])

	s := []int{1, 2, 3, -1, 0, 2, -5, 10, 20, 100, -20}
	i, j, _ := FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{}
	i, j, err := FindSubArrayWithMaxSum(s)
	if err != nil {
		fmt.Println(err)
	}

	s = []int{-100}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{1, 2, 3, 4, 5}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{-1, -2, -3, -4, -5}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{1, 2, 3, -1, 10000}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{1, 2, 3, -10, 10000}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{1, 0, 0, 0, 0}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])

	s = []int{0, 0, 0, 0, 0}
	i, j, _ = FindSubArrayWithMaxSum(s)
	fmt.Printf("%d, %d, %v\n", i, j, s[i:j+1])
}

// FindSubArrayWithMaxSum 的返回值指定了子数组的范围[i, j]
func FindSubArrayWithMaxSum(s []int) (int, int, error) {
	sLen := len(s)
	if sLen == 0 {
		return 0, 0, errEmptyArray
	} else if sLen == 1 {
		return 0, 0, nil
	}

	sum := make([]int, sLen)
	sum[0] = s[0]

	maxSum := s[0]
	maxSumIdx := 0
	minSum := s[0]
	minSumIdx := 0

	for i := 1; i < sLen; i++ {
		sum[i] = s[i-1] + s[i]
		if sum[i] > maxSum {
			maxSum = sum[i]
			maxSumIdx = i
		} else if sum[i] < minSum {
			minSum = sum[i]
			minSumIdx = i
		}
	}

	if maxSumIdx > minSumIdx {
		if minSum < 0 {
			return minSumIdx + 1, maxSumIdx, nil
		}
	}

	return 0, maxSumIdx, nil
}
