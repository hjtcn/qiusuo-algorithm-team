/*
给定一棵二叉搜索树，请找出其中第k大的节点。
 
示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
限制：
1 ≤ k ≤ 二叉搜索树元素个数
*/

Clarfication:
二叉搜索树特点：左子树 < 根节点 < 右子树，中序遍历为一个有序数组
然后返回第k个节点  返回有序数组中 k-1 下标所在的元素

复杂度预估：
时间复杂度：O(n), 遍历节点
空间复杂度：O(n), 存储中序遍历的二叉树节点

Coding:
一开始的时候写的是返回的是 ret[k-1]，运行了发现结果不对，题目要求是第k大，我们中序列遍历是递增序列，应该返回的是 length - k 才可以

递归：
func kthLargest(root *TreeNode, k int) int {
    ret := []int{}
    helper(root, &ret)
    length := len(ret)
    return ret[length - k]
}

func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left, ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}

BFS:
写BFS的时候想要不要判断树是否为空呢？题目中是有提示长度大小的 
限制：
1 <= k <= 二叉搜索树元素个数

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
    ret := []int{}
    helper(root, &ret)
    length := len(ret)
    return ret[length - k]
}

func helper(root *TreeNode,ret *[]int) {
    stack := []*TreeNode{}
    
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        } 
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        *ret = append(*ret, node.Val)
        root = node.Right
    }
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点将节点放入到数组中
空间复杂度：O(n) 存储遍历的每个节点
看了题解我们是可以不用遍历完的找到那个数返回就可以了
直接先遍历右节点 -> 根节点 -> 左节点，遍历出来的顺序是递减的序列
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
    num := 0
    ret := 0
    helper(root,&ret,&num,k)
    return ret
}

func helper(root *TreeNode,ret *int,num *int,k int){
    if root == nil {
        return
    }
    helper(root.Right, ret,num, k)
    (*num)+=1
    
    if *num == k {
        *ret = root.Val
        return
    }
    helper(root.Left, ret,num, k)
    
    return
}

迭代，从 k-- 到0，先放右节点，然后放左节点，最后输出
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
    res := 0
   stack := []*TreeNode{}
    
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Right
        } 
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        k -= 1
        if k == 0 {
            res = node.Val
            break
        }
       
        root = node.Left
    }
    return res
}

总结：
1. 题目还是没有想清楚就开始做了，返回数组的地方有问题

2. 迭代遍历二叉树的时候还是花了一点时间，一开始写成了层序遍历二叉树，还是没有熟练的掌握

3. 改变二叉搜索树也可以输出递减的序列的，理解为什么这样做比做出来更重要，知识点和算法要灵活应用，实际项目中变化也会非常多
