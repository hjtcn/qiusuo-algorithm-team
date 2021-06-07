二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。
返回转换后的单向链表的头节点。
注意：本题相对原题稍作改动
示例：
输入： [4,2,5,1,3,null,6,0]
输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]

 1. Clearfication:
        重点：把二叉搜索树转换为单向链表，并且依然符合二叉搜索树的性质
            找重复子问题：
                重复子问题是什么？
2. Coding:

3. 看题解：
递归：超时
中序遍历：
因为要保证二叉搜索树的性质，所以找到最小的节点开始进行遍历，中序遍历比较适合
里面要维护一个链表的头，这是和原来中序遍历不一样的地方，引入一个变量，你的思维和难度就会增加一些

func convertBiNode(root *TreeNode) *TreeNode {
    newTree := &TreeNode{}
    inOrder(root, newTree)
    return newTree.Right
}

func inOrder(root *TreeNode, pos *TreeNode) *TreeNode {
    if root == nil {
        return pos
    }
    pos = inOrder(root.Left, pos)
    pos.Right = root
    pos.Left = nil
    pos = root
    pos = inOrder(root.Right, pos)
    return pos
}

迭代：
定义一个链表头，每次出栈都加入到当前链表头的右子树，并将当前链表头左子树置为nil，然后移动链表头
func convertBiNode(root *TreeNode) *TreeNode {
    newList := &TreeNode{}
    head := newList
    var stack []*TreeNode
    
    for root != nil || len(stack) > 0 {
        if root != nil {
            // 入栈
            stack = append(stack, root)
            root = root.Left
        }else {
            // 出栈
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            // 将pop加入到单向链表
            pop.Left = nil
            newList.Right = pop
            newList = pop
            //继续遍历右子树
            root = pop.Right
        }
    }
    
    return head.Right
}   

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) : 递归调用栈空间 or 迭代开辟栈空间

5. 总结：
5.1:中序遍历的变形，加了一个需要维护的链表头节点，你对问题的理解就要多想一点，这个时候发现自己不会了，说明自己对基础的数据结构如树和链表的使用还是没有掌握

5.2: 递归：分而治之，就想当前结果对不对，然后使用数学归纳法去判断是否正确，不要人肉递归
