/*
94. 二叉树的中序遍历
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

 

示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[2,1]
示例 5：


输入：root = [1,null,2]
输出：[1,2]
 

提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/*
    思路：递归还是比较好写的，都成框架了。迭代还洋洋得意的想用昨天的一层while循环写。后来发现实现不了，因为中序遍历是左->根->右
    左右节点的遍历不在一层上了。仅if和else不能深层模拟。

    然后又尝试着第二种，先迭代push左节点。然后节点pop出栈作为根节点。然后再次迭代

*/
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
 var inorderTraversal = function(root) {
    let target=[]
    let dfs=(root)=>{
        if(!root) return ;
        dfs(root.left)
        target.push(root.val)
        dfs(root.right)
    }
    dfs(root)
    return target
};

var inorderTraversal = function(root) {
    let target=[]
    let stack=[]
    while(stack.length>0||root){
        while(root){
            stack.push(root)
            root=root.left
        }
        root=stack.pop()
        target.push(root.val)
        root=root.right
    }
    return target
}
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：多试试，了解的会更多一些
*/

