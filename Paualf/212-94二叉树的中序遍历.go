给定一个二叉树的根节点 root ，返回它的 中序 遍历。
示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

示例 4：
输入：root = [1,2]
输出：[2,1]
示例 5：
输入：root = [1,null,2]
输出：[1,2]

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

1. Clarification:
中序遍历：
左根右
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    helper(root,&ret)
    return ret
}

func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    *ret = append(*ret,root.Val)
    helper(root.Right,ret)
}
使用迭代模拟递归实现：
如果左边一直有元素则一直向左走，将元素放入到栈中
如果左边没有元素，从栈中弹出元素，指向它的右边
循环条件为 root != nil || len(stack) > 0 {}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    if root == nil {
        return ret
    }

    stack := []*TreeNode{}
    node := root

    for node != nil || len(stack) > 0 {
        for node != nil {
            stack = append(stack,node)
            node = node.Left
        }
        top := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        ret = append(ret,top.Val)
        node = top.Right
    }

    return ret
}
对比昨天前序遍历，放入ret中元素的位置不一样，因为前序遍历，遍历到那个节点的时候就需要放入进去了

中序遍历的时候，遍历到的时候，需要放入到栈中，因为不一定是最左边的元素，从栈中弹出来的时候一定是最左边的元素

2. 看题解：
看 Morris 中序遍历：
x 无左孩子
	x加入结果
	x = x.right
x 有左孩子，找predecessor
	predecessor 右孩子为空，右孩子指向x,x = x.left
	predecessor 右孩子不为空，x加入结果，x = x.right （这个地方蛮不好理解的，因为当时已经找到x左子树上最右的节点了，所以如果它有值的话，代表已经把根节点挂上去了）

func inorderTraversal(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 遍历顺序还是需要搞懂的，数据结构本身就是维护了一种数据变化的特征，搞清楚数据变化的特征了的话，题目就比较容易理解了
4.2: 遇到问题耐心一点，不要着急，一步一步去分析，拆成小问题更容易理解
4.3: Morris 这个分析过程还是值得学习的
