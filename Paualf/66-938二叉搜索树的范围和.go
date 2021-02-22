/*
给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
示例 1：
输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32
示例 2：
输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23
 
提示：
树中节点数目在范围 [1, 2 * 104] 内
1 <= Node.val <= 105
1 <= low <= high <= 105
所有 Node.val 互不相同
*/

1. Clearfication:
 二叉搜索树：中序遍历是递增数组， 找到 low和high 之间的位置，进行求和
如示例：中序遍历：[3,5,7,10,15,18]
然后找到循环数组，找到 low=7 和 hight = 15 的之间的元素进行求和

伪代码：
ret := []int{}
sum := 0
for _,val range ret {
    if val >= low && val <= high {
        sum += val
    }
}
return sum
helper(root *TreeNode, ret *[] int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}

Coding:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {                                                                                                                         ret := []int{}
     sum := 0
    helper(root,&ret)
    
    for _,val := range ret {
        if val >= low && val <= high {
            sum += val
        }
    }
    return sum                                                       
}
func helper(root *TreeNode, ret *[] int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}

看了题解以后:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    dfs(root,low,high,&sum)
    return sum
}
func dfs(root *TreeNode,low,high int,sum *int) {
    // terminator
    if root == nil {
        return
    }
    
    // process current logic
    if root.Val >= low && root.Val <= high {
        *sum = *sum + root.Val
    }
    
    // drill down
    if root.Val > low {
        dfs(root.Left,low,high,sum)
    }
    
    if root.Val < high {
        dfs(root.Right, low, high,sum)
    }
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    
    if root == nil {
        return sum
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node != nil {
            if node.Val >= low && node.Val <= high {
                sum += node.Val
            }
            
            if node.Val > low {
                queue = append(queue,node.Left)
            }
            
            if node.Val < high {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return sum
}

Other think:
如果数组数据量贼大的话，我们怎么在大的数组里面找到 low 和 high 之间的元素呢，我们可以使用 二分查找去降低时间复杂度，从O(n) 的时间复杂度降到 O(logn)

复杂度分析：
时间复杂度：O(n) : 二叉搜索树的中序遍历，在有序数组中进行求和
空间复杂度：O(n) : 二叉搜索树的中序遍历，递归调用栈空间

总结：
1. 先分析清楚再去写代码的感觉好舒服哈，好爽

2. 看了题解以后发现自己并没有用到二叉搜索树的性质=》 自己对二叉搜索树还是没有理解和掌握

3. 对递归思想还是不是很熟悉
