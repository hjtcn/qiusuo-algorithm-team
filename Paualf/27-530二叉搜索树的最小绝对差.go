/*
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
*/

Clearfication:
题目中给出树节点的值都为非负数，绝对值最小的应该在 最小的两个值之间
又因为二叉搜索树的性质，最小的两个数可以通过中序遍历，获取到前面两个值, 然后进行求解
然后写出代码：

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    arr := []int{}
    helper(root, &arr)
    //fmt.Println(arr)
    return arr[1] - arr[0]
}

func helper(root *TreeNode,arr *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,arr)
    *arr = append(*arr, root.Val)
    helper(root.Right, arr)
}

提交发现是错的，打印了一下数组，发现绝对值最小的不一定是最前面的两个值的差，也又可能是后面相邻的值
然后又想是不是比较数组相邻的值就可以找到了，因为数组值递增的，相邻的两个数之间的差值是最小的

func getMinimumDifference(root *TreeNode) int {
    arr := []int{}
    helper(root, &arr)
    minDistance := arr[1] - arr[0]
    for i := 2;i < len(arr);i++ {
        currentDistance := arr[i] - arr[i - 1]
        if currentDistance < minDistance {
            minDistance = currentDistance
        }
    }
    return minDistance
}

func helper(root *TreeNode,arr *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,arr)
    *arr = append(*arr, root.Val)
    helper(root.Right, arr)
}

这样就 ok 了
BFS里面的 helper 是怎么写的呢？
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    arr := []int{}
    helper(root, &arr)
    minDistance := arr[1] - arr[0]
    for i := 2;i < len(arr);i++ {
        currentDistance := arr[i] - arr[i - 1]
        if currentDistance < minDistance {
            minDistance = currentDistance
        }
    }
    return minDistance
}

func helper(root *TreeNode,arr *[]int) {
    stack := []*TreeNode{}
    
    for root != nil || len(stack)  > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        *arr = append(*arr,node.Val)
        root = node.Right
    }
}

题解中直接使用两个变量进行比较这样就不用开辟数组存储每个数组的元素了

func getMinimumDifference(root *TreeNode) int {
    ans,pre := math.MaxInt64, -1
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

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归栈深度或者是使用栈的空间

总结：
1. 还是没有想清楚，二叉搜索树中最小的值在那个地方 -> 原因在于没有找例子去模拟，所以我们有思路以后，要有例子去模拟，感觉也明白了我们在工程中为什么要写单元测试，可以避免我们没有想到的一些case，同时让我们可以更多的思考

2. 多写代码，多思考
