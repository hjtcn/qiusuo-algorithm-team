给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
示例 1：
输入：root = [4,2,7,1,3], val = 5
输出：[4,2,7,1,3,5]
解释：另一个满足题目要求可以通过的树是：

示例 2：
输入：root = [40,20,60,10,30,50,70], val = 25
输出：[40,20,60,10,30,50,70,null,null,25]

示例 3：
输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
输出：[4,2,7,1,3,5]
 
提示：
给定的树上的节点数介于 0 和 10^4 之间
每个节点都有一个唯一整数值，取值范围从 0 到 10^8
-10^8 <= val <= 10^8
新值和原始二叉搜索树中的任意节点值都不同

1.Clarification:
一开始看这道题的时候有点懵，不知道怎么想，想起来前天写的那个二叉搜索树的删除，还是利用二叉搜索树找到要删除的节点，然后删除就可以了

这里我们是不是也可以类比找到二叉搜索树中要插入的地方，将元素写入进去就ok了

if root.Val > val {
insertLeft(root.Left)
}

if root.Val < val {
insertRight(root.Right)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        root = &TreeNode{
            Val:val,
        }
        return root
    }

    if root.Val > val {
        root.Left = insertIntoBST(root.Left,val)
    }

    if root.Val < val {
        root.Right = insertIntoBST(root.Right,val)
    }

    return root
}

2.看题解：
使用一个指针去找到要插入的位置
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil{
        return &TreeNode{Val:val}
    }

    node := root //一个遍历指针

    for node != nil{
        if node.Val > val{ //根节点值大于插入值，对左子树进行判断
            if node.Left == nil{ //如果空，则插入，结束循环
                node.Left = &TreeNode{Val:val}
                break
            }
            //如果不为空，则继续搜索
            node = node.Left
        }else{ //根节点值小于插入值，对右子树进行判断
            if node.Right == nil{
                node.Right = &TreeNode{Val:val} //如果为空，则插入，结束循环
                break
            }
            //如果不为空，则继续搜索
            node = node.Right
        } 
    }

    return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{val, nil, nil}
    }
    
    node := root
    for node != nil {
        if node.Left != nil && node.Val > val {
            node = node.Left
        } else if node.Right != nil && node.Val < val {
            node = node.Right
        } else if node.Val > val {
            node.Left = &TreeNode{val, nil, nil}
            break
        } else {
            node.Right = &TreeNode{val, nil, nil}
            break
        }
    }
    
    return root
}
上面两块代码逻辑差不多，第一个代码更容易理解，第二块代码更简洁，如果我选的话，第一个更好一点，因为更容易看懂，别人可以更好的维护和修改。因为写代码是给人看的

3.复杂度分析：
时间复杂度：O(n)
空间复杂度：递归O(n)，使用指针去找插入位置的话，空间复杂度O(1)

总结：
4.1: 遇到问题自己多想一下，分析一下，让自己的小脑瓜转一转，不要害怕失败，不要害怕不会
