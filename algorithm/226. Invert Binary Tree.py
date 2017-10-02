# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# 第一种解法
import copy

# Definition for a binary tree node.


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


# Create a Tree
root = TreeNode(3)
Node1 = TreeNode(6)
Node2 = TreeNode(8)
Node3 = TreeNode(2)
Node4 = TreeNode(5)
Node5 = TreeNode(7)
Node6 = TreeNode(4)


root.left = Node1
root.right = Node2
Node1.left = Node3
Node1.right = Node4
Node2.left = Node5
Node2.right = Node6


class Solution(object):
    # 解法一,自顶向下交换
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if root is None:
            return
        root.left, root.right = root.right, root.left
        self.invertTree(root.left)
        self.invertTree(root.right)
        return root

    # 解法二,自下向上交换
    def inverTree1(self, root):
        if root is None:
            return
        root.left, root.right = self.inverTree(
            self, root.right), self.inverTree(self, root.left)
        return root

    # 解法三,采用非递归方式,以栈做数据结构存放node
    def invertTree2(self, root):
        if root is None:
            return
        stack_node = []
        stack_node.append(root)
        while stack_node:
            print(stack_node)
            node = stack_node.pop()
            node.left, node.right = node.right, node.left
            if node.left is not None:
                stack_node.append(node.left)
            if node.right is not None:
                stack_node.append(node.right)
        return root

    def show(self, root):
        if root is not None:
            print(root.val)
            self.show(root.left)
            self.show(root.right)

sloution = Solution()

n = sloution.invertTree2(root)
sloution.show(root)
print(n.val)
