给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

1. Clarification:
想到用堆解决，向最大堆里面放出现元素的次数，出现最多的次数为堆顶

看到二叉搜索树，本身就想到它是一个递增的数组，因为二叉搜索树中序遍历为递增的数组

Q:是不是可以不用额外空间就可以解决这个问题呢？
A: 上一次遍历的值怎么带到下一层呢？这块代码没有想清楚

Coding:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    returnRet := make([]int,1)

    if root == nil {
        return  returnRet
    }

    ret := []int{}
    helper(root,&ret)
    currentNum := 1
    returnRet[0] = ret[0]

fmt.Println(ret)

    for i := 0;i < len(ret) - 1; {
        j := i + 1
        tmp := 1 
        for ;j < len(ret);j++ {
            if ret[i] != ret[j] {
                break
            }
            // 相等的时候 tmp++
            tmp++   
        }
        
        if tmp > currentNum {
            returnRet[0] = ret[i] 
        }

        i = j
    }

    return returnRet
}

func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    *ret = append(*ret, root.Val)
    helper(root.Right,ret)
}


哈哈哈，写了半天写错了，找到了最大的众数，其实题目里面想要找的是 BST 里面的所有众数


2. 看题解：
重复出现的数字一定是一个连续出现的
用base记录当前的数字，用count记录当前数字重复的次数，用maxCount来维护已经扫描过的出现最多的那个数字的出现次数，用anawer数组记录出现的众数

每次扫描到一个新的元素：
首先更新base和count
如果该元素和base相等，那么count自增1
否则将base更新为当前数字，count复位为1
然后更新maxCount
如果 count == maxCount, 说明当前的这个数字(base)出现的次数等于当前众数出现的次数，将base加入answer数组
如果 count > maxCount,那么说明当前的这个数字(base)出现的次数大于当前众数出现的次数，因为，我们需要将maxCount更新为count,清空answer 数组后将base加入answer数组。

说实话这段逻辑分析的很清楚
func findMode(root *TreeNode) (answer []int) {
    var base, count, maxCount int

    update := func(x int) {
        if x == base {
            count++
        } else {
            base, count = x, 1
        }
        if count == maxCount {
            answer = append(answer, base)
        } else if count > maxCount {
            maxCount = count
            answer = []int{base}
        }
    }

    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        update(node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return
}

3.复杂度分析
时间复杂度：O(n)
空间复杂度：O(n)

4.总结：
4.1: 这道题卡了一会，也说明自己的基本功不够扎实，需要多练习哈

