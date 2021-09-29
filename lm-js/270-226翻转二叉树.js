/*
226. 翻转二叉树
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
*/

/*
    思路:递归，也就是左右节点交换。
        一开始直接root.left=dfs(root.right)
                root.right=dfs(root.left)
        后来发现root.left已经被修改，这样是不合理的。
        因此提前记录left=dfs(root.left),right=dfs(root.right)
        root.left=right
        root.right=left
        这种方法写的还挺顺手，习惯了感觉递归真是太棒了。

        迭代:递归的思路如何用迭代模拟出来。
            迭代的思路可以用递归模拟出来吗？
            原理：左右节点互换
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
 var invertTree = function(root) {
    let dfs=(root)=>{
        if(!root) return null
        let left=dfs(root.left),right=dfs(root.right)
        root.right=left
        root.left=right
        return root
    }
    return dfs(root)
}


var invertTree = function(root) {
    if(!root) return null
    let stack=[root]
    while(stack.length>0){
        let curNode=stack.shift()
        //交换过程
        let left=curNode.left,right=curNode.right
        curNode.left=right
        curNode.right=left
        if(curNode.left){
            stack.push(curNode.left)
        }
        if(curNode.right){
            stack.push(curNode.right)
        }
    }
    return root
}
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    现在看到这种题写起来是有点爽哈。
*/