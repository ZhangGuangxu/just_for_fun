package zgxutil

import (
    "testing"
)

func TestMaxSumOfArray(t *testing.T) {
    a := []int{1, -2, 3, 5, -3, 2}
    if maxSum(a) != 8 {
        t.Error("should be 8")
    }
    a = []int{0, -2, 3, 5, -1, 2}
    if maxSum(a) != 9 {
        t.Error("should be 9")
    }
    a = []int{-9, -2, -3, -5, -3}
    if maxSum(a) != -2 {
        t.Error("should be -2")
    }
}

func TestMaxSumOfArray2(t *testing.T) {
    a := []int{1, -2, 3, 5, -3, 2}
    if maxSum2(a) != 8 {
        t.Error("should be 8")
    }
    a = []int{0, -2, 3, 5, -1, 2}
    if maxSum2(a) != 9 {
        t.Error("should be 9")
    }
    a = []int{-9, -2, -3, -5, -3}
    if maxSum2(a) != -2 {
        t.Error("should be -2")
    }
}

func TestMaxSumOfArray3(t *testing.T) {
    a := []int{1, -2, 3, 5, -3, 2}
    if maxSum3(a) != 8 {
        t.Error("should be 8")
    }
    a = []int{0, -2, 3, 5, -1, 2}
    if maxSum3(a) != 9 {
        t.Error("should be 9")
    }
    a = []int{-9, -2, -3, -5, -3}
    if maxSum3(a) != -2 {
        t.Error("should be -2")
    }
}
