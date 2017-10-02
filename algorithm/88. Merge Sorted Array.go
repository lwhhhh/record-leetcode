package algorithm

func merge(nums1 []int, m int, nums2 []int, n int) {
	//fmt.Println(nums1, nums2)
	indexM := m - 1
	indexN := n - 1
	indexMerged := m + n - 1
	indexCurr := indexMerged

	for indexM >= 0 && indexN >= 0 {
		//fmt.Printf("nums1[%d]=%d nums2[%d]=%d\n", indexM, nums1[indexM], indexN, nums2[indexN])
		n1 := nums1[indexM]
		n2 := nums2[indexN]
		if n1 >= n2 {
			nums1[indexCurr] = n1
			indexM--
		} else {
			nums1[indexCurr] = n2
			indexN--
		}
		indexCurr--
	}

	var indexLast int
	var nums []int
	if indexM > 0 {
		indexLast = indexM
		nums = nums1
	} else {
		indexLast = indexN
		nums = nums2
	}

	for ; indexLast >= 0; indexLast-- {
		nums1[indexCurr] = nums[indexCurr]
		indexCurr--
	}
}
