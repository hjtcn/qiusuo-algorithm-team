输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

 为了让您更好地理解问题，以下面的二叉搜索树为例：

 


我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。


下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

1. Clarification:

看到这种多种数据结构的题目还是有点害怕的，不知道怎么下手

只分析出来，因为是有序的，二叉搜索树想到了中序遍历

后面又遇到：怎么构造循环双向链表呢？这里就懵逼了


2. 看题解：

排序链表：中序遍历，从小到大访问树的节点
双向链表，构建相邻节点引用关系时，设前驱节点pre和当前节点cur,pre.right = cur,cur.left = pre
循环链表：设链表头节点head和尾节点tail，则应构建head.left = tail 和 tail.right = head

var head,pre *Node

func treeToDoublyList(root *Node)*Node {
    if root == nil {
    	return
    }
    
    dfs(root)
    head.Left = pre
    pre.Right = head
    return head
}

func dfs(cur *Node) {
    if cur == nil {
    	return
    }
    
    dfs(cur.Left)
    if pre != nil {
    	pre.Right = cur
    }else {
    	head = cur
    }
    cur.Left = pre
    pre = cur
    dfs(cur.Right)
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) :使用栈空间

4. 总结：
4.1: 忽然发现现在看代码不难了，难的是设计方案，怎么把方案设计的合理？这里为什么要这样处理？而不是这样处理，感觉比去写代码更难

4.2: 因为出方案，是要做很多调研，要去做 trade-off,要去不断思考的
