class Solution:

    def convert2tup(self,s):
        return tuple(map(int,s))

    def openLock(self, deadends, target):
        """
        :type deadends: List[str]
        :type target: str
        :rtype: int
        """
        initial = '0000'
        result = 0
        deadends = set(map(self.convert2tup,deadends))
        curr = self.convert2tup(initial)

        while curr:
            pass
        return result

soul = Solution()

deadends = ['0100','1234']
target = ['0200']

soul.openLock(deadends,target)