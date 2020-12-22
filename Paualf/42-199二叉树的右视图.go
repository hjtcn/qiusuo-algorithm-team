/*
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
示例:
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/

1. Clearfiction:
二叉树的右视图：
从右边看这个二叉树看到的数据元素，感觉还是挺有意思的
开始想，如果根节点的右子树一直右元素，一直向右遍历是不是就可以得到结果了
仔细想了一下，如果右子树为空， 左子树的右子树不为空， 还是可以看到元素的
仔细想了一下，目前是没有思路的：
看了题解发现，
DFS 没有看懂，就去看了BFS,BFS 思路还是蛮清晰的，获取每一层最右边的节点

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    queue := []*TreeNode{root}
    result := []int{}
    
    for len(queue) != 0 {
        l := len(queue)
        result = append(result, queue[l-1].Val)
        
        for i := 0;i < l;i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        queue = queue[l:]
    }
    
    return result
}

感觉代码里面的
queue = queue[l:]

还是蛮巧妙和细节的，感觉是比较容易出错的地方

对于DFS的话看这个代码还是懂了，怎么写了

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    list := make([]int, 0)
    dfs(root, &list, 0)
    return list
}
func dfs(node *TreeNode, list *[]int, level int) {
    if node == nil {
        return
    }
    
    if len(*list) == level {
        *list = append(*list, node.Val)
    }
    
    dfs(node.Right, list, level + 1)
    dfs(node,Left, list, level + 1)
}

每一层右节点第一个访问的放入数组中，判断条件使用的还是蛮巧妙的

复杂度分析：
时间复杂度：O(n) 遍历每个节点，所以时间复杂度为O(n)
空间复杂度：O(n) 递归使用栈空间或者队列开辟的空间

总结：
1. 题目还是比较有意思的，从不同的角度理解二叉树

2. 题目本身是不难的，还是要仔细分析

3. git branch --show-current 查看当前分支
