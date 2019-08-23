package array

// Finds the contiguous subarray within an array (containing at least one number)
// which has the largest sum.
// Finds the contiguous subarray with the largest sum.
func maxSumSubArrayMemo(arr []int) int {
	mem := make([]int, len(arr))
	mem[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		sum := arr[i] + mem[i - 1]
		if sum > arr[i] {
			mem[i] = sum
		} else {
			mem[i] = arr[i]
		}
	}

	return max(mem)
}

// Returns the largest value in a given array.
func max(arr []int) int {
	max := arr[0];

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

// Finds the contiguous subarray with the largest sum.
func maxSumSubArrayKadane(arr []int) int {
	res, sum := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		sum = cmp(arr[i] + sum, arr[i])
		res = cmp(res, sum)
	}

	return res
}

// Cmp returns the larger of x or y.
func cmp(x, y int) int {
	if x > y {
		return x
	}

	return y
}