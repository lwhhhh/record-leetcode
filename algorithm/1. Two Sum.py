class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        mark = dict()
        result = list()

        for i in range(len(nums)):
            left = target - nums[i]
            if left in mark:
                result.append(mark[left])
                result.append(i)
                return result
            mark[nums[i]] = i  

soul = Solution()

nums = [2, 7, 11, 15]
target = 9

res = soul.twoSum(nums,target)
print(res)