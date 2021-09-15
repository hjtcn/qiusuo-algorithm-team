给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true

示例 2：
输入：root = [1,2,3], targetSum = 5
输出：false

示例 3：

输入：root = [1,2], targetSum = 0
输出：false
 
提示：
树中节点的数目在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

1. Clarification:
第一感觉这道题做过，觉得很简单，但是真正在纸上画图，写伪代码的时候发现，其实并没有那么简单，题目里面还是有很多细节要考虑的
很多事情，想的它再难-》没啥用
它再简单你不会做-》也没啥用

去做就ok了，难与简单，会了就好，想太多，没啥用。。。
func hasPathSum(root *TreeNode, sum int) bool {
     return helper(root,0,sum)
}

func helper(root *TreeNode,currentSum,sum int)bool {
    // terminator
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        if currentSum + root.Val == sum {
            return true
        }
        return false
    }

    // process current logic
   return helper(root.Left,currentSum + root.Val,sum) || helper(root.Right, currentSum + root.Val, sum)
}
迭代写的有bug...
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    queueNode := []*TreeNode{root}
    queueVal := []int{root.Val}

    for len(queueNode) > 0 {
        size := len(queueNode)
        for size > 0 {
            size--
            node := queueNode[0]
            val := queueVal[0]
            queueNode = queueNode[1:]
            queueVal = queueVal[1:]

            // 子节点
            if node.Left == nil && node.Right == nil && node.Val + val == targetSum {
                return true
            }

            if node.Left != nil {
                queueNode = append(queueNode, node.Left)
                queueVal = append(queueVal, node.Left.Val + val)
            }

            if node.Right != nil {
                queueNode = append(queueNode, node.Right)
                queueVal = append(queueVal, node.Right.Val + val)
            }
        }
    }

    return false
}

2.看题解：
看了题解以后发现是size这个地方出错了。。。
还有判断的地方出错了，特别是有一个 continue 的地方，我觉得还是蛮容易错的。。。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    queueNode := []*TreeNode{root}
    queueVal := []int{root.Val}

    for len(queueNode) > 0 {
            node := queueNode[0]
            val := queueVal[0]
            queueNode = queueNode[1:]
            queueVal = queueVal[1:]

            // 子节点
            if node.Left == nil && node.Right == nil {
                if val == targetSum {
                    return true
                }
                continue
            }

            if node.Left != nil {
                queueNode = append(queueNode, node.Left)
                queueVal = append(queueVal, node.Left.Val + val)
            }

            if node.Right != nil {
                queueNode = append(queueNode, node.Right)
                queueVal = append(queueVal, node.Right.Val + val)
            }
        
    }

    return false
}
这种代码更像是分解子问题的代码
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 不要想太多，想看书就去看书，想做题就去做题，想休息就好好休息一下，想太多往往得不偿失
4.2: 以前觉得简单的题目没啥可做的，现在发现简单的题目可能都做不好哈，哈哈哈
