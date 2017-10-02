class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        zero_count = 0
        for i in nums:
            if i == 0:
                zero_count += 1
        for i in range(zero_count):
            nums.remove(0)
        for i in range(zero_count):
            nums.append(0)
        print(nums)
        
