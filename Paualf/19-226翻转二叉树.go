/*
翻转一棵二叉树。
示例：
输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

思路：
找重复性：
头节点为4
tmp = 4->left
4 -> left = 4->right
4->right  = tmp
递归重复这个操作就可以将树进行翻转了

1. terminator
2. process current login
3. drill down
4. reverse current status

1. root == nil
return result

2. tmp = root.Left
     root.Left = root.Right
     root.Right = root.Left

3. drill down
if root.Left != nil {
invertTree(root.Left)
 }
if root.Right != nil {
invertTree(root.Right)
 }

4. no reverse current status

func invertTree(root *TreeNode) *TreeNode {
     if root == nil {
         return nil
     }
         
     if root.Left != nil || root.Right != nil {
         tmp := root.Left
         root.Left = root.Right
         root.Right = tmp
     }
     
     if root.Left != nil {
         invertTree(root.Left)
     }
     if root.Right != nil {
         invertTree(root.Right)
     }
     return root
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(height) 递归时使用到的栈

代码可以更简单：
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}

总结：
  找问题的重复性，将一个大问题分解为多个自问题，然后把子问题解决了，然后组合解决大问题
