给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
 
示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
 
提示：
树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000
1. Clearfication:
    最小深度：
            BFS:
                使用队列存储，如果 
                root.Left == nil && root.Right == nil {
                    return depth
                }
            DFS：
忘掉了，还是没有掌握哈。。。

2. Coding:
思路没有分析清楚
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{root}
    depth := 1
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[:len(queue) - 1]
        if node.Left == nil && node.Right == nil {
            return depth
        }
         if node.Right != nil {
            queue = append(queue, node.Right)
        }
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
      
        depth++ 
    }
    return depth
}

3. 看题解：
深度优先搜索：后序遍历
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    if root.Left == nil && root.Right == nil {
        return 1
    }
    
    minD := math.MaxInt32
    if root.Left != nil {
        minD = min(minDepth(root.Left), minD)
    }
    
    if root.Right != nil {
        minD = min(minDepth(root.Right),minD)
    }
    
    return minD + 1
}

func min(x,y int) int {
    if x < y {
        return x
    }
    return y
}

func minDepth(root *TreeNode)int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{}
    count := []int{}
    queue = append(queue, root)
    count = append(count, 1)
    
    for i := 0;i < len(queue);i++ {
        node := queue[i]
        depth := count[i]
        
        if node.Left == nil && node.Right == nil {
            return depth
        }
        
        if node.Left != nil {
            queue = append(queue, node.Left)
            count = append(count, depth + 1)
        }
        
        if node.Right != nil {
            queue = append(queue, node.Right)
            count = append(count,depth + 1)
        }
    }
    
    return 0
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

5. 总结：
5.1: 后序遍历还是想起来，感觉挺容易出错的

5.2: 使用queue 的时候，要遍历当前层也是出错了。。。

5.3: 培养分解子问题的思维是真的有点难
