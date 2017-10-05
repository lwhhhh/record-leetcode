package algorithm

import "fmt"

func longestConsecutive(nums []int) int {
	var res int
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			var left int
			var right int

			if v, ok := m[num-1]; ok {
				left = v
			} else {
				left = 0
			}

			if v, ok := m[num+1]; ok {
				right = v
			} else {
				right = 0
			}

			sum := left + right + 1
			fmt.Printf("num=%d   left=%d   right=%d   sum=%d\n", num, left, right, sum)
			m[num] = sum
			if sum > res {
				res = sum
			}

			m[num-left] = sum
			m[num+right] = sum
		}
	}
	return res
}
