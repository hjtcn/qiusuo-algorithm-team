// 剑指 Offer 34. 二叉树中和为某一值的路径
// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

 

// 示例:
// 给定如下二叉树，以及目标和 sum = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// 返回:

// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]

/*
    思路：这个类型题思路练的已经很清晰了，核心就在于记录path。
        当我再次遇到题时没思考到的：差值为0时，没有判断是叶子节点（这个错过n次了）
        从解决题目开始，我为什么总是会遗漏这个思考呢？？？？
        如果尝试从构建思路开始呢。
        发现自己太习惯，写着代码模拟着。
        构建思路能力太差

        再有更加熟练的就是递归====迭代之间的模拟联系




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
 * @param {number} sum
 * @return {number[][]}
 */


var pathSum = function(root, sum) {
    if(!root) return []
    let res=[]
    let dfs=(root,sum,path)=>{
        if(!root){
            return ;
        }
        if(sum-root.val==0&&!root.left&&!root.right){
            res.push(path)
        }
        if(root.left){
            dfs(root.left,sum-root.val,[...path,root.left.val])
        }
        if(root.right)
        {
            dfs(root.right,sum-root.val,[...path,root.right.val])
        }
    }
    dfs(root,sum,[root.val])
    return res
}

var pathSum = function(root, sum) {
    if(!root) return []
    let queue=[[root,[root.val],sum]],res=[]
    while(queue.length){
      let [curNode,path,sum]=queue.shift()
        if(!curNode.left&&!curNode.right&&sum-curNode.val==0){
            res.push(path)
        }
        if(curNode.left){
            queue.push([curNode.left,[...path,curNode.left.val],sum-curNode.val])
        }
        if(curNode.right){
            queue.push([curNode.right,[...path,curNode.right.val],sum-curNode.val])
        }
    }
    return res
};

