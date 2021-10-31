给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

 

示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 

提示：

树中至少有 2 个节点。
本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

1. Clarification:

找最小绝对值，因为都是整数，并且是递增的，那么中序遍历的时候比较一下就ok了。

Q:怎么比较，变量怎么放呢？
A: 没有想清楚哈

如果在生产环境或者面试的时候我这个时候会这样写，先遍历生成递增的数组然后进行比较
// 中序遍历：左根右，那么返回值放到哪里呢
func getMinimumDifference(root *TreeNode) int {
    ret := []int{}
    helper(root,&ret)

    minNum := ret[1] - ret[0]
   
    for i := 2;i < len(ret);i++{
        if ret[i] - ret[i - 1] < minNum {
            minNum = ret[i] - ret[i - 1]
        }
    }
    
    return minNum
}

func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret) 
    *ret = append(*ret, root.Val)
    helper(root.Right,ret)
}
有没有更好的方法呢？或者更简单清晰的方法

自己没有想清楚的地方，遍历的时候变量传入什么，然后会返回什么

2. 看题解：
用pre 变量保存前驱节点的值，这样即能边遍历边更新答案，不再需要显式创建数组来保存。
func getMinimumDifference(root *TreeNode) int {
    ans,pre := math.MaxInt64,-1
    var dfs func(*TreeNode)
    
    dfs = func(node *TreeNode) {
        if node == nil {
        	return
        }
        dfs(node.Left)
        if pre != -1 && node.Val - pre < ans {
        	ans = node.Val - pre
        }
        pre = node.Val
        dfs(node.Right)
    }
    
    dfs(root)
    return ans
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4. 总结：
4.1: 目前感觉难的地方在于变量的命名，怎么把逻辑写的更加清晰

4.2: 编程是一个脑力+动手的工作，要思考，同时也要动手
