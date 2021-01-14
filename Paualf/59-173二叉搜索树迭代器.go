/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
调用 next() 将返回二叉搜索树中的下一个最小的数。
示例：
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false
 
提示：
next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
*/

1. Clearfication:
调用next() 将返回二叉搜索树中的下一个最小的数
第一时间想到的是使用数组存储二叉搜索树的遍历结果，这样数组中就是有序的了，然后获取下一个最小的数
自己还是动手去写，想了一下，然后去看了题解。。。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    nums []int
    root *TreeNode
}
func Constructor(root *TreeNode) BSTIterator {
    nums := []int{}
    inorder(root, &nums)
    
    return BSTIterator {
        nums:nums,
        root:root,
    }
}
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    smallest := this.nums[0]
    this.nums = this.nums[1:]
    return smallest
}
/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    if len(this.nums) > 0 {
        return true
    }
    return false
}
func inorder(root *TreeNode,nums *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, nums)
    *nums = append(*nums, root.Val)
    inorder(root.Right, nums)
}
/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
看到题解里面这个代码感觉还是比较好的，不全部中序遍历，只是遍历左子树，将最小的元素找到，然后将节点存储到栈中，然后回退，然后遍历右节点
type BSTIterator struct {
    stack []*TreeNode
}
func Constructor(root *TreeNode) BSTIterator {
    result :=  BSTIterator{stack: []*TreeNode{}}
    for root != nil {
        result.stack = append(result.stack, root)
        root = root.Left
    }
    return result
}
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    top := this.stack[len(this.stack)-1]
    result := top.Val
    this.stack = this.stack[:len(this.stack)-1]
    if top.Right != nil {
        top = top.Right
        for top != nil {
            this.stack = append(this.stack, top)
            top = top.Left
        }
    }
    return result
}
/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    if len(this.stack) == 0 {
        return false
    } else {
        return true
    }
}

复杂度分析：
时间复杂度：next和hasNext 的操纵时间复杂度是O(1), 第一种使用中序遍历构建有序数组的时间复杂度是O(n),第二种方法使用栈来构建部分数据，有可能是O(1),如果树是斜树的话，时间复杂度可能是O(n), 平均下来是O(1)
空间复杂度：O(n),第一种方法使用数组存储有序序列，第二种方法使用栈进行存储

总结：
1. 自己害怕这种使用type 结构体的题目

2. 对这种设计以及要自己设计数据结构的题目不熟悉，感觉本质上还是对数据结构没有掌握，没有养成分析数据结构的习惯
