// 98. 验证二叉搜索树
// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:

// 输入:
//     2
//    / \
//   1   3
// 输出: true
// 示例 2:

// 输入:
//     5
//    / \
//   1   4
//      / \
//     3   6
// 输出: false
// 解释: 输入为: [5,1,4,null,null,3,6]。
//      根节点的值为 5 ，但是其右子节点值为 4 。



//递归
var isValidBST=function(root){
    let helper=(root,min,max)=>{
        if(!root){
            return true
        }
        if(root.val<=min||root.val>=max){
            return false
        }
        //推动递归。在于跳出前面两个坑之后，
        //min<root.val，则root.val作为根节点，寻找左子树，max值更新为root.val
        //max>root.val,root.val作为根节点，寻找右子树，min值更新为root.val
        return helper(root.left,min,root.val)&&helper(root.right,root.val,max)
    }
    return helper(root,-Infinity,Infinity)
}
/**
 * 递归，不断的更新最大值和最小值，我自己一开始的时候只写到了一个最小值，提交发现一些测试用例过不了。
 * 想了想觉得当我们尝试用暴力解决某个问题时，可以先学会构建一些测试用例，过了自己那关再去提交。
 * 后来看了题解去解决这题，前两天没有写题解，今天决定提交，再次做了这题，做的还是蛮顺手的。
 * 虽然等号这种可能忘了判断，嘻嘻嘻，下一次会更好
 *  时间复杂度：O(N)
    取决于树的节点个数，每个节点只会被入栈和出栈一次。
    空间复杂度：O(N)
    取决于树的节点个数
 */

 //中序遍历
 var isValidBST=function(root){
    let stack=[],min=-Infinity
    while(stack.length||root){
        while(root){
            stack.push(root)
            root=root.left
        }
        root=stack.pop()
        if(root.val<=min){
            return false
        }
        min=root.val
        root=root.right
    }
    return true
}


/**
 * 迭代
 * 利用中序遍历的特点：
 * 左->根->右的递增序列,然后根据这个存储到栈中,然后根据当前节点是否大于上一个中序遍历到的节点
 * 
 *  时间复杂度：O(N)
    取决于树的节点个数，每个节点只会被入栈和出栈一次。
    空间复杂度：O(N)
    取决于树的深度
 */