package algorithm

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	nums1 := make([]int, 8)
	nums2 := []int{2, 3, 5, 8}

	nums1[0] = 3
	nums1[1] = 4
	nums1[2] = 7
	nums1[3] = 8

	merge(nums1, 4, nums2, len(nums2))

	fmt.Println(nums1)
}
