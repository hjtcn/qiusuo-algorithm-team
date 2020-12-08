/*
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
示例：
输入：[1,2,3,4,5,null,7,8]
        1
       /  \ 
      2    3
     / \    \ 
    4   5    7
   /
  8
输出：[[1],[2,3],[4,5,7],[8]]
*/

Clarfication：
题目变了形式：感觉还是二叉树的层次遍历，将每一层放到一个数组里面，上一次做这个类型的题目的时候卡在了 len(queue) 上面
BFS,主干代码是可以写出来的，返回值的时候有点困惑 []*ListNode 看着像是一维链表, 题目中的输出看着是 二维数组 [[1],[2,3],[4,5,7],[8]] 有疑惑的地方
DFS,使用level 记录该层节点，然后放入到同一个level 数组中去，这个代码看能不能写出来
主干代码ok，返回值的地方有问题

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
     ret := make([]*ListNode, 0)
     if tree == nil {
         return ret
     }
     queue := make([]*TreeNode, 0)
     queue = append(queue, tree)
    for len(queue) > 0 {
        size := len(queue)
        tmp := make([]*ListNode, 0)
        for size > 0 {
            size--
            node := queue[0]
            queue = queue[1:]
            tmp = append(tmp, node)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        ret = append(ret, tmp)
    }
     return ret
}

题目中已经提示了 定义了
Definition for singly-linked list.
  type ListNode struct {
      Val int
      Next *ListNode
 }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
     nodeLists := make([]*ListNode, 0)
     if tree == nil {
         return nodeLists
     }
     queue := make([]*TreeNode, 0)
     queue = append(queue, tree)
    for len(queue) > 0 {
        size := len(queue)
        link := &ListNode{Val:0}
        currentNode := link
        for size > 0 {
            size--
            node := queue[0]
            queue = queue[1:]
            currentNode.Next = &ListNode{Val:node.Val}
            currentNode = currentNode.Next
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        nodeLists = append(nodeLists, link.Next)
    }
     return nodeLists
}

如果用dfs如何实现和存储呢？
dfs想到这里，每层函数的结果如何存储，又卡壳了。。。
func listOfDepth(tree *TreeNode) []*ListNode {
     nodeLists := make([]*ListNode, 0)
     if tree == nil {
         return nodeLists
     }
    
     helper(tree,1)
     return nodeLists
}

func helper(tree *TreeNode,level int)[]*ListNode {
    
    // process current login
    // drill down
}

看了题解中的代码
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 声明为全局变量，下面函数处理的时候就不需要传递该参数了
var (
    listNodes []*ListNode
)
func listOfDepth(tree *TreeNode) []*ListNode {
    listNodes = make([]*ListNode, 0)
    doListOfDepth(tree, 1)
    return listNodes
}
func doListOfDepth(TreeNode *TreeNode,depth int) {
    if TreeNode == nil {
        return
    }
    listNode := &ListNode {
        Val:TreeNode.Val,
    }
    // process current login
    if len(listNodes) < depth {
        //该层链表没有元素的时候
        listNodes = append(listNodes, listNode)
    }else {
        // 该层链表有元素的情况下 depth - 1 也很巧妙
        headListNode := listNodes[depth - 1]
        cur := headListNode
        for cur.Next != nil {
            cur = cur.Next
        }
        cur.Next = listNode
    }
    // drill down
    doListOfDepth(TreeNode.Left, depth + 1)
    doListOfDepth(TreeNode.Right, depth + 1)
}
使用指针传递参数：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    depth := 0
    // new 返回的是指针
    list := new ([]*ListNode)
    dfsListOfDepth(tree,list,depth)
    return *list
}
func dfsListOfDepth(tree *TreeNode,list *[]*ListNode,depth int) {
    if tree == nil {
        return 
    }
    if depth >= len(*list) {
        *list = append(*list, &ListNode{tree.Val, nil})
    }else {
        head := (*list)[depth]
        for head.Next != nil {
            head = head.Next
        }
        head.Next = &ListNode{tree.Val, nil}
    }
    dfsListOfDepth(tree.Left,list,depth + 1)
    dfsListOfDepth(tree.Right,list,depth + 1)
}

总结：
1. 对返回值没有搞清楚题意，说明对链表还不是很熟悉
2. 树的大部分题目应该都是两种解法的，多熟悉不同的解法
3. 细节还是蛮重要的，一个 = 号或者一个判断都会引起程序出现bug或者不正常运行，所以还是要多注意细节
