# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
'''
    该题求二叉树的最大深度
    很快能想到关键词是递归,遍历
    
    递归那部分可能较难理解:
        left_max_depth和right_max_depth会在每一层堆栈帧中保存
        每一层堆栈帧return的时候会将改帧保存的left_max_depth和right_max_depth返回给上一层的left_max_depth和right_max_depth
        如此经过迭代,最终left_max_depth和right_max_depth将会是根节点的左子树和右子树最大深度,再进行一次比较大小即可

'''


class Solution(object):
    def maxDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if root is None:
            return 0

        left_max_depth = self.maxDepth(root.left)
        right_max_depth = self.maxDepth(root.right)

        if left_max_depth >= right_max_depth:
            return left_max_depth + 1
        else:
            return right_max_depth + 1


