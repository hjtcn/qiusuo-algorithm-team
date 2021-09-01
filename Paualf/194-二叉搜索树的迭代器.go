实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。

你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。

 

示例：


输入
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
输出
[null, 3, 7, true, 9, true, 15, true, 20, false]

解释
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // 返回 3
bSTIterator.next();    // 返回 7
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 9
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 15
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 20
bSTIterator.hasNext(); // 返回 False
 

提示：

树中节点的数目在范围 [1, 105] 内
0 <= Node.val <= 106
最多调用 105 次 hasNext 和 next 操作
 

进阶：

你可以设计一个满足下述条件的解决方案吗？next() 和 hasNext() 操作均摊时间复杂度为 O(1) ，并使用 O(h) 内存。其中 h 是树的高度。

1. Clarification:
自己的思路：
将二叉搜索树初始化到数组中
hasNext 右侧遍历是否存在数字，数组下标 * 2  是左节点，数组下标*2+1是右节点
获取左节点或者右节点的值

想着不是很难，遇到数组下标的时候怎么初始化把我搞晕了。。。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    arr[] int
}


func Constructor(root *TreeNode) BSTIterator {
    tree := &BSTIterator{
    }

    if root == nil {
        return 0
    }
    
    level := 0
    helperConstructor(root, tree.arr,level,'left')

    return tree
}

func helperConstructor(root *TreeNode,arr []int,level int,dir string) {
    if root == nil {
        return
    }

    if dir == 'left' {
        arr[level * 2] = append(arr, root.Val)
    }else {
        arr[level * 2 + 1] = append(arr,root.Val)
    }
    
    helperConstructor(root,arr,level + 1,'left')
    helperConstructor(root,arr,level + 1,'right')
}


func (this *BSTIterator) Next() int {
    
}


func (this *BSTIterator) HasNext() bool {
     
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */


2.看题解：
type BSTIterator struct {
    arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
    it.inorder(root)
    return
}

func (it *BSTIterator) inorder(node *TreeNode) {
    if node == nil {
        return
    }
    it.inorder(node.Left)
    it.arr = append(it.arr, node.Val)
    it.inorder(node.Right)
}

func (it *BSTIterator) Next() int {
    val := it.arr[0]
    it.arr = it.arr[1:]
    return val
}

func (it *BSTIterator) HasNext() bool {
    return len(it.arr) > 0
}
中序遍历将二叉树的遍历结果放到内存中，获取的时候直接用就可以了

看了代码以后发现自己对题目的题解理解错了。。。

type BSTIterator struct {
    stack []*TreeNode
    cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{cur: root}
}

func (it *BSTIterator) Next() int {
    for node := it.cur; node != nil; node = node.Left {
        it.stack = append(it.stack, node)
    }
    it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
    val := it.cur.Val
    it.cur = it.cur.Right
    return val
}

func (it *BSTIterator) HasNext() bool {
    return it.cur != nil || len(it.stack) > 0
}
Next 的方法有点巧妙，stack用的说实话还有一丢丢没理解

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) or O(1)

4. 总结：
4.1: 一定一定不要害怕失败，勇敢一点点，不会做也要自己去尝试哈
