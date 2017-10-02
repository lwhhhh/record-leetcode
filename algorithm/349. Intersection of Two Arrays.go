package algorithm

func intersection(nums1 []int, nums2 []int) []int {
	var result = []int{}

	m := make(map[int]bool)

	for _, num := range nums1 {
		if _, ok := m[num]; !ok {
			m[num] = false
		}
	}

	for _, num := range nums2 {
		if value, ok := m[num]; ok && value == false {
			m[num] = true
			result = append(result, num)
		}
	}

	return result
}
