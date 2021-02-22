/*
给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。
示例：
输入: root = [4,2,6,1,3,null,null]
输出: 1
解释:
注意，root是树节点对象(TreeNode object)，而不是数组。
给定的树 [4,2,6,1,3,null,null] 可表示为下图:
          4
        /   \
      2      6
     / \    
    1   3  
最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
*/
1. Clearfication:
树是二叉搜索树：二叉搜索树的特点，左子树 < 根节点 < 右子树
看了这道题，一开始有一个想法：两节点差的最小值是不是在树的最左边，自己想的是二叉搜索树中序遍历以后的数组是递增的，所以想差值是不是在最右边
后来仔细分析了一下，二叉搜索树的特点是中序遍历是递增的，和元素之间的差值是没有关系的，所以我们要找差值最小的，可以中序遍历以后，然后去遍历比较递增数组的差值（Gap),去找到最小值

伪代码：
中序遍历
ret := []int{}
helper(root, &ret)
gap := ret[1] - ret[0]
for i := 1;i < len(ret) - 1;i++ {
    currentGap = ret[i + 1] - ret[i]
    if currentGap < gap {
        gap = currentGap
    }
}
return gap

// 返回递增数组 中序遍历 左根右
func helper(root *TreeNode, ret *[] int) {
    if root == nil {
        return 
    }
    helper(root.Left, ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}
写伪代码想到，如果树的节点元素是空需要怎么处理呢？
如果树的节点元素个数只有1个，我们怎么处理呢？
边界条件需要考虑清楚
[4,2,6,1,3,null,null]

2. Coding：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    if root == nil {
        return 0
    }
    ret := []int{}
    helper(root, &ret)
    if len(ret) == 1 {
        return ret[0]
    }
    gap := ret[1] - ret[0]
    for i := 1;i < len(ret) - 1;i++ {
        currentGap := ret[i + 1] - ret[i]
        if currentGap < gap {
            gap = currentGap
        }
    }
    return gap
}

// 返回递增数组 中序遍历 左根右
func helper(root *TreeNode, ret *[] int) {
    if root == nil {
        return 
    }
    helper(root.Left, ret)
    *ret = append(*ret, root.Val)
    helper(root.Right, ret)
}

3. 看题解：
看了题解以后了解到了，Go里面最大值是 math.MaxInt32，自己一开始想去初始化，发现自己不知道怎么去写
还看到可以在中序遍历的过程中计算差值然后去比较, 需要定义 prevVal，然后去更新处理prevVal 的值
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    min := math.MaxInt32
    prevVal := math.MinInt32
    
    helper(root, &min, prevVal)
    return min
}

func helper(root *TreeNode, min *int,prevVal int) {
    if root == nil {
        return 
    }
    
    helper(root.Left, min,preVal)
    
    if prevVal != math.MinInt32 {
        gap := root.Val - prevVal
        if gap < min {
            min = gap   
        }
    }
    
    prevVal = root.Val
    
    helper(root.Right, min,preVal)
}

上面是自己根据题解模拟写出来的思路：里面有一个和细节点在与 prevVal int 是传形式参数还是指针,上面传的是形式参数，然后提交了以后发现报错了，是因为它在运行的过程中，prevVal 已经改变了，它返回给上层数据的时候，并没有把这个改变的值传过去，上面还是使用的原有的形式参数，所以运行结果就对不上了。
我们将 prevVal int 修改为指针就可以了 

func minDiffInBST(root *TreeNode) int {
    min := math.MaxInt32
    prevVal := math.MinInt32
    
    helper(root, &min, &prevVal)
    return min
}

func helper(root *TreeNode, min *int,prevVal *int) {
    if root == nil {
        return 
    }
    
    helper(root.Left, min, prevVal)
    
    if *prevVal != math.MinInt32 {
        gap := root.Val - *prevVal
        if gap < *min {
            *min = gap   
        }
    }
    
    *prevVal = root.Val
    
    helper(root.Right, min,prevVal)
}

4. 复杂度分析：
时间复杂度： O(n) : 遍历每个节点

空间复杂度： O(n)：二叉树中序遍历使用栈空间或者使用数组存储的空间

5. 总结
5.1. 这道题自己考虑了边界条件还是有一丢丢进步的

5.2. 通过这道题，自己也看到了一道简单的题目里面如果你去认真分析理解，也可以找到自己原来不会的东西，在这道题中就是递归调用的时候我们传参数是使用形参还是指针还是根据我们需求想清楚的
