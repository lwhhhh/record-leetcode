package algorithm

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastOne := nums[0]
	result := 1
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > lastOne {
			temp++
		} else {
			temp = 1
		}
		//fmt.Println(lastOne, nums[i], temp)

		if temp > result {
			result = temp
		}
		lastOne = nums[i]
	}
	return result
}
