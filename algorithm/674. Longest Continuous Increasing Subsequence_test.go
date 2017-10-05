package algorithm

import (
	"fmt"
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7}
	result := findLengthOfLCIS(nums)
	fmt.Println(result)
}
