package algorithm

import (
	"math"
	"sort"
)

func lengthOfLIS(nums []int) int {
	copyNums := make([]int, len(nums))
	for index, n := range nums {
		copyNums[index] = n
	}
	sort.Ints(copyNums)

	cache := make([][]int, len(nums))
	for i := 0; i <= len(nums); i++ {
		t := make([]int, len(copyNums))
		for j := 0; i <= len(copyNums); i++ {
			if i == 0 || j == 0 {
				cache[i][j] = 0
			} else if nums[i-1] == copyNums[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				cache[i][j] = int(math.Max(float64(cache[i-1][j]), float64(cache[i][j-1])))
			}
		}
	}
	return cache[len(nums)][len(nums)]
}
