/*
面试题 17.12. BiNode
二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。

返回转换后的单向链表的头节点。

注意：本题相对原题稍作改动

 

示例：

输入： [4,2,5,1,3,null,6,0]
输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]
提示：

节点数量不会超过 100000。
*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */

/*
    思路：想到了排序过程，但是没想到中序遍历。
        中序遍历模版:
        let dfs=(root)=>{
            if(!root) return;
            dfs(root.left)
            res.push(root.val)
            dfs(root.right)
        }

        目的要一个排好序的二叉搜索树。
        中序遍历：左--->根---->右
        目标二叉树的left节点都为null,right节点要链接父级的根节点

*/
 var convertBiNode = function(root) {
   let prevNode=new TreeNode(0)
   let target=prevNode
   let dfs=(root)=>{
       if(!root){
           return ;
       }
       dfs(root.left)
       root.left=null
       prevNode.right=root
       prevNode=root
       dfs(root.right)
   }
    dfs(root)
    return target.right
}; 