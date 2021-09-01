给你一棵以 root 为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。
注意，一旦删除值为 target 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。

也就是说，你需要重复此过程直到不能继续删除。

 

示例 1：

输入：root = [1,2,3,2,null,2,4], target = 2
输出：[1,null,3,null,4]
解释：
上面左边的图中，绿色节点为叶子节点，且它们的值与 target 相同（同为 2 ），它们会被删除，得到中间的图。
有一个新的节点变成了叶子节点且它的值与 target 相同，所以将再次进行删除，从而得到最右边的图。

示例 2：
输入：root = [1,3,3,3,2], target = 3
输出：[1,3,null,null,2]

示例 3：
输入：root = [1,2,null,2,null,2], target = 2
输出：[1]
解释：每一步都删除一个绿色的叶子节点（值为 2）。

示例 4：
输入：root = [1,1,1], target = 1
输出：[]

示例 5：
输入：root = [1,2,3], target = 1
输出：[1,2,3] 

提示：

1 <= target <= 1000
每一棵树最多有 3000 个节点。
每一个节点值的范围是 [1, 1000] 。

1.Clarification:
写这道题的时候写了很多 Q & A;

Q: 执行多少次这个过程呢？
A: 最多等于target 元素的个数

Q: 怎么删除元素？需要找到父节点
A: 使用map去记录父节点

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    targetNum := 0
    m := make(map[*TreeNode]*TreeNode,0)
    helper(root,target,&targetNum,m)
    fmt.Println(targetNum)
    if root.Left != nil {
        fmt.Println(m[root.Left].Val)
    }

    for i :=0 ;i < targetNum;i++{
       remove(root,target,"",m)
    }


    if root.Left == nil && root.Right == nil && root.Val == target{
        root = nil 
    }


    return root
}

// helper 统计最多等于target 元素的个数 & 使用map 记录父亲节点为了删除使用
func helper(root *TreeNode,target int,targetNum *int,m map[*TreeNode]*TreeNode) {
    if root == nil {
        return
    }
    if root.Left != nil {
        m[root.Left] = root
    }

    if root.Right != nil {
        m[root.Right] = root
    }

    if root.Val == target {
       *targetNum = *targetNum + 1    
    }

    helper(root.Left,target,targetNum,m)
    helper(root.Right,target,targetNum,m)

    return
}

// 移除等于 target 的叶子节点
// Q: 什么是叶子节点.    root.Left == nil && root.Right == nil 的节点
func remove(root *TreeNode,target int,dir string,m map[*TreeNode]*TreeNode) {
    if root == nil {
        return
    }

    // 叶子节点
    if root.Left == nil && root.Right == nil && root.Val == target{
        if _,ok := m[root];ok {
            if dir == "left" {
                m[root].Left = nil
            }

            if dir == "right" {
                m[root].Right = nil
            }
        }else {
            // 这个时候，处理中剩下一个节点的情况
            fmt.Println(123456)
            root = nil
        }
        return 
    }
    remove(root.Left,target,"left",m)
    remove(root.Right,target,"right",m)
}


里面还是有一些细节case的：

比如：
记录父亲节点的时候，怎么记录m，树的遍历的顺序自己还是想了一会的

再比如剩下一个节点的时候：
  // 这个时候，处理中剩下一个节点的情况            
fmt.Println(123456)
root = nil
我这样写不行，所以在主函数里面加了一个兜底的判断
 if root.Left == nil && root.Right == nil && root.Val == target{       
      root = nil 
}

看题解：

看了题解以后，真的是绝绝子，自己花了大半个小时做出来了，结果后序遍历一下子搞定了，本来还很高兴自己做出来了，哈哈哈，不过还是可以的

从底向上，还是构建重复子问题的能力不够

数学归纳法

3.复杂度分析：
时间复杂度：O(n) * target元素节点的个数
空间复杂度：O(n)


总结：
4.1: 工程师就是用代码实现自己想法的人，向这个方向努力，还是有很多代码和思路想不清的，用心的练习吧
