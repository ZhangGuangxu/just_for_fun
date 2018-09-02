package zgxutil

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// All[0] = max(A[0], A[0]+Start[1], All[1])

func maxSum(A []int) int {
    n := len(A)
    All := make([]int, n)
    Start := make([]int, n)

    Start[n-1] = A[n-1]
    All[n-1] = A[n-1]
    for i := n-2; i >= 0; i-- {
        Start[i] = max(A[i], A[i]+Start[i+1])
        All[i] = max(Start[i], All[i+1])
    }
    return All[0]
}

func maxSum2(A  []int) int {
    n := len(A)
    Start := A[n-1]
    All := A[n-1]
    for i := n-2; i >= 0; i-- {
        Start = max(A[i], A[i]+Start)
        All = max(Start, All)
    }
    return All
}

func maxSum3(A []int) int {
    n := len(A)
    Start := A[n-1]
    All := A[n-1]
    for i := n-2; i >= 0; i-- {
        if Start < 0 {
            Start = 0
        }
        Start += A[i]
        if Start > All {
            All = Start
        }
    }
    return All
}
