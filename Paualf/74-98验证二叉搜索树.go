/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
• 节点的左子树只包含小于当前节点的数。
• 节点的右子树只包含大于当前节点的数。
• 所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

1. Clearfication:
左子树节点 < 根节点 < 右子树节点 && 整体符合这个规则
伪代码：
return root.Left.Val < root.Val && root.Right.Val > root.Val && isValidBST(root.Left) && isValidBST(root.Right)
发现有点懵逼不知道怎么分析了，连验证二叉搜索树是升序序列都忘记了
看了题解以后思路是：更新它的上下界

2. Coding：
func isValidBST(root *TreeNode) bool {
    return helper(root,math.MinInt64,math.MaxInt64)
}
func helper(root *TreeNode,lower,upper int) bool {
    if root == nil {
        return true
    }
    
    if root.Val <= lower || root.Val >= upper {
        return false
    }
    
    return helper(root.Left,lower,root.Val) && helepr(root.Right,root.Val,upper)
}

func isValidBST(root *TreeNode) bool {
    ret := []int{}
    helper(root,&ret)
    
    for i := 0;i < len(ret) - 1;i++ {
        if ret[i] >= ret[i+1] {
            return false
        }
    }
    
    return true
}
func helper(root *TreeNode,ret *[]int) {
    if root == nil {
        return
    }
    helper(root.Left,ret)
    *ret = append(*ret, root.Val)
    helper(root.Right,ret)
}

如何使用队列进行中序遍历：
func isValidBST(root *TreeNode)bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    
    return true
}
改写了一下
没有改写出来，发现自己对函数返回值没有想清楚,下面是自己改写的半成品
func isValidBST(root *TreeNode) bool {
    min := math.MinInt64
    return helper(root,&min,)
}
func helper(root *TreeNode,min *int) bool{
    if root == nil {
        return true
    }
    helper(root.Left,ret)
    if root.Val <= min {
        return false
    }
    *min = root.Val
    helper(root.Right,ret)
}

3. 看题解：
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    _,_,valid := helper(root)
    return valid
}

func helper(root *TreeNode) (int, int, bool) {
    // curMax,curMin means max/min node values of current subtree
    curMax,curMin := root.Val,root.Val
    
    if root.Left != nil {
        max,min,valid := helper(root.Left)
        if !valid || root.Val <= max {
            return -1,-1,false
        }
        curMin = min
    }
    
    if root.Right != nil {
        max,min,valid := helper(root.Right)
        if !valid || root.Val >= min {
            return -1,-1,false
        }
        curMax = max
    }
    
    return curMax,curMin,true
}
多返回值情况下，自己没有想清楚了，对后序遍历去解决问题不熟练的

4. 复杂度分析：
时间复杂度：O(n) 遍历每个节点

空间复杂度：O(n) 递归调用栈空间或者开辟队列长度

5. 总结：
5.1：感觉对二叉树的前序中序后序遍历的扩展去解决问题还是有点分析不清楚，没有理解到本质
