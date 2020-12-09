/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
示例 1:
输入: 
	Tree 1                     Tree 2                  
          1                         2                             
         / \                       / \                            
        3   2                     1   3                        
       /                           \   \                      
      5                             4   7                  
输出: 
合并后的树:
      3
     / \
    4   5
   / \   \ 
  5   4   7
注意: 合并必须从两个树的根节点开始。
*/

Clearfiaction:
看到这道题和之前做过的一道，也是两个节点 p,q 两个节点，去进行遍历的有点像，等下去找一下那道题
思路：
DFS：同时遍历两个树，如果树2和树1节点值重复，将树2的节点的值加到树1上面，如果不重复，则以树1为准
BFS:   使用两个队列记录然后进行遍历，如果树2的节点的值和树1的节点的值重复，将树2的节点的值加到树1上面，否则以树1的节点为准
// it is my think not right 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
      if t1 == nil && t2 == nil {
          return nil
      }
      if t1 == nil && t2 != nil {
          t1 = t2
     }
     if t1 != nil && t2 != nil {
         t1.Val += t2.Val
     }
   
    mergeTrees(t1.Left,t2.Left)
    mergeTrees(t2.Right,t2.Right)
    return t1 
}
看了题解，知道了自己错的地方，合并以后两个节点是一个新的节点，然后这个新的节点的左右节点都是要变的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
      if t1 == nil {
          return t2
      }
      if t2 == nil {
          return t1
      }
    t1.Val += t2.Val
    
    t1.Left =  mergeTrees(t1.Left,t2.Left)
    t1.Right = mergeTrees(t1.Right,t2.Right)
    return t1 
}
今天先不看BFS了，有点累，休息下

总结：
1. 如何思考和解决比最后的代码和答案更重要
2. 好好工作的时候也要好好休息
