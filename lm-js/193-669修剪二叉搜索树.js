/*
669. 修剪二叉搜索树
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

 

示例 1：


输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]
示例 2：


输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]
示例 3：

输入：root = [1], low = 1, high = 2
输出：[1]
示例 4：

输入：root = [1,null,2], low = 1, high = 3
输出：[1,null,2]
示例 5：

输入：root = [1,null,2], low = 2, high = 4
输出：[2]
 

提示：

树中节点数在范围 [1, 104] 内
0 <= Node.val <= 104
树中每个节点的值都是唯一的
题目数据保证输入是一棵有效的二叉搜索树
0 <= low <= high <= 104

*/

/*
    思路：想叉劈了。一开始想的中序排列，如果不在边界内， 就把这个节点清空。
        看了题解之后才醒过来劲。
        如果node小于边界，需要找右子节点接着遍历。
        如果node大于边界，需要找左节点进行遍历。
        如果处于边界内，则正常遍历左子节点和右子节点，并返回node
*/

var trimBST = function(root, low, high) {
    let dfs=(node)=>{
        if(!node) return null;
        if(node.val<low){
            return dfs(node.right)
        }
        if(node.val>high){
            return dfs(node.left)
        }
        else{
        node.left=dfs(node.left)
        node.right=dfs(node.right)
        return node
        }

    }
    return dfs(root)
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：递归方法的返回和执行有点薄弱了
*/