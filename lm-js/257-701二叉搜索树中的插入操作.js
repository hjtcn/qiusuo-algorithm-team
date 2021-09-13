/*
701. 二叉搜索树中的插入操作
给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。

 

示例 1：


输入：root = [4,2,7,1,3], val = 5
输出：[4,2,7,1,3,5]
解释：另一个满足题目要求可以通过的树是：

示例 2：

输入：root = [40,20,60,10,30,50,70], val = 25
输出：[40,20,60,10,30,50,70,null,null,25]
示例 3：

输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
输出：[4,2,7,1,3,5]
 

 

提示：

给定的树上的节点数介于 0 和 10^4 之间
每个节点都有一个唯一整数值，取值范围从 0 到 10^8
-10^8 <= val <= 10^8
新值和原始二叉搜索树中的任意节点值都不同

*/

/*
思路：陷到坑里了。也想到了如果root.val>val就往左走，如果root.val<val就往右走。
     但是我用的递归遍历，出不来了，返回值把自己整懵了。
题解：1.如果root.val>val,向左走
     2.如果root.val<val,向右走
     迭代方法：如果向左走的过程中找不到值了，root.left=newNode
             同理，向右走的过程中也找不到值了，root.right=newNode
     递归方式:从根节点出发。如果root.val>val,root.left=dfs(root.left)
                          root.val<val,root.right=dfs(root.right)
            返回值：如果root为空，则返回新节点即可。反之，将root返回。

*/

var insertIntoBST = function(root, val) {
    if(!root) return new TreeNode(val)
    let target=root
    while(root){
        if(root.val>val){
            if(!root.left){
                root.left=new TreeNode(val)
                break;
            }
            root=root.left
        }
        else{
            if(!root.right){
                root.right=new TreeNode(val)
                break;
            }
            root=root.right
        }
    }
    return target
};
// @lc code=end

var insertIntoBST = function(root, val) {
    if(!root) return new TreeNode(val)
    let dfs=(root)=>{
        if(!root) return new TreeNode(val)
        if(root.val<val){
            root.right=dfs(root.right)
        }
        else{
            root.left=dfs(root.left)
        }
        return root
    }
    return dfs(root)
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：递归的返回值和执行过程总是弄懵
*/