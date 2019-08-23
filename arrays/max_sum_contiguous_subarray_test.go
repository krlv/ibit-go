package array

import (
	"testing"
)

func TestMaxSumSubArrayMemo(t *testing.T) {
	max := maxSumSubArrayMemo([]int{-2, 1,-3, 4,-1, 2, 1, -5, 4})
	if max != 6 {
		t.Error("6 is a cmp sum of contiguous subarray [4,-1, 2, 1] of [-2, 1,-3, 4,-1, 2, 1, -5, 4] array")
	}

	max = maxSumSubArrayMemo([]int{-5, 36, 89, -194, -140, -418, 63, 77, 77, -197, -211, -425, 58, -25, -107, 93, -183})
	if max != 217 {
		t.Error("217 is a cmp sum of contiguous subarray [63, 77, 77] of [-5, 36, 89, -194, -140, -418, 63, 77, 77, -197, -211, -425, 58, -25, -107, 93, -183] array")
	}

	max = maxSumSubArrayMemo([]int{-120, -20})
	if max != -20 {
		t.Error("-20 is a cmp sum of contiguous subarray [-20] of [-120, -20] array")
	}
}

func TestCmp(t *testing.T) {
	if cmp(1, 2) != 2 {
		t.Error("2 is a cmp of (1, 2) pair")
	}

	if cmp(-1, -2) != -1 {
		t.Error("-1 is a cmp of (-1, -2) pair")
	}
}

func TestMaxSumSubArrayKadane(t *testing.T)  {
	max := maxSumSubArrayKadane([]int{-2, 1,-3, 4,-1, 2, 1, -5, 4})
	if max != 6 {
		t.Error("6 is a cmp sum of contiguous subarray [4,-1, 2, 1] of [-2, 1,-3, 4,-1, 2, 1, -5, 4] array")
	}

	max = maxSumSubArrayKadane([]int{-5, 36, 89, -194, -140, -418, 63, 77, 77, -197, -211, -425, 58, -25, -107, 93, -183})
	if max != 217 {
		t.Error("217 is a cmp sum of contiguous subarray [63, 77, 77] of [-5, 36, 89, -194, -140, -418, 63, 77, 77, -197, -211, -425, 58, -25, -107, 93, -183] array")
	}

	max = maxSumSubArrayKadane([]int{-120, -20})
	if max != -20 {
		t.Error("-20 is a cmp sum of contiguous subarray [-20] of [-120, -20] array")
	}
}
