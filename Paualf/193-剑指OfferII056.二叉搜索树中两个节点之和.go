给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。假设二叉搜索树中节点的值均唯一。 

示例 1：
输入: root = [8,6,10,5,7,9,11], k = 12
输出: true
解释: 节点 5 和节点 7 之和等于 12
示例 2：

输入: root = [8,6,10,5,7,9,11], k = 22
输出: false
解释: 不存在两个节点值之和为 22 的节点
 

提示：

二叉树的节点个数的范围是  [1, 104].
-104 <= Node.val <= 104
root 为二叉搜索树
-105 <= k <= 105
 
1. Clarification:

二叉搜索树：
 	左子树 < 根节点 < 右子树

二叉搜索树中序遍历的话是一个有序的数组，然后在数组中进行双指针寻找是否有 target == k 的，如果有则返回true，没有返回false

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    ret := []int{}
    helper(root,&ret)

    i := 0
    j := len(ret) - 1

    for i < j {
        if ret[i] + ret[j] == k {
            return true
        }else if ret[i] + ret[j] > k {
            j--
        }else if ret[i] + ret[j] < k {
            i++
        }
    }

    return false
}

// 中序遍历，左中右
func helper(root *TreeNode,ret *[]int){
    if root == nil {
        return
    }

    helper(root.Left,ret)
    *ret = append(*ret, root.Val)
    helper(root.Right,ret)
}
感觉应该有更简单的解法，我没想出来，看下题解了

看题解：
遍历一遍，判断是否在hash里面
func findTarget(root *TreeNode, k int) bool {
    
    set := make(map[int]int)
    return find(root, k, set)
    
}

func find(root *TreeNode, k int, set map[int]int) bool {
    
    if root == nil {
        return false
    }
    if _, ok := set[k - root.Val]; ok {
        return ok
    }
    
    set[root.Val]++
    return find(root.Left, k, set) || find(root.Right, k, set)
    
}

这种方法也挺好的

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, root, k)
}

func dfs(root, cur *TreeNode, k int) bool {
	if cur == nil {
		return false
	}

	return search(root, cur, k-cur.Val) || dfs(root, cur.Left, k) || dfs(root, cur.Right, k)
}

func search(root, curr *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	return (root.Val == val) && (root != curr) || (root.Val < val) && search(root.Right, curr, val) || (root.Val > val) && search(root.Left, curr, val)
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 使用map or set or 递归都需要占用空间


4. 总结：
4.1: 知识要变成自己的东西灵活运用的话，还是要自己用心去体会，去实践的

4.2: 先了解理论，然后去实践
