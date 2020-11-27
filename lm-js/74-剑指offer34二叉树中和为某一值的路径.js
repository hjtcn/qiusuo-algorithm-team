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
    let floorArr=[],res=[]
    //count为上一次的剩余差值
    let dfs=(root,count)=>{
        if(!root) {
            return;
        }
        floorArr.push(root.val)  
        if(!root.left&&!root.right){
            if(count-root.val==0){
                res.push([...floorArr])
                //这里很重要。深拷贝，否则会随着更改，res中元素都指向最后一个floorArr
            }
        }
        dfs(root.left,count-root.val)
        dfs(root.right,count-root.val)
        floorArr.pop()
    }
    dfs(root,sum)
    return res
};

/**
 * 递归。又读错题了，以为是路径中出现相等就行，没有必须查到叶子节点为止。
 * 最后遇到[1,2]和为1的情况报错，才及时醒悟，必须查到叶子节点才是一条完整路径
 *  时间复杂度：O(N)
    取决于树的节点数
    
    空间复杂度：O(N)
    取决于树节点数
 * 
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
    if(!root){
        return [];
    }
    let queue=[[root,sum,[root.val]]],res=[]
    while(queue.length){
        let [node,num,path]=queue.pop()
        if(!node.left&&!node.right&&num==node.val){
            res.push(path)
        }
        if(node.left){
            queue.push([node.left,num-node.val,[...path,node.left.val]])
        }
         if(node.right){
            queue.push([node.right,num-node.val,[...path,node.right.val]])
        }

    }
    return res
};


/**
 * BFS,和以往不同的是，不止需要记录节点，还要需要记录path，和差值
 * 当查询到叶子节点时，如果差值正好等于当前节点，则将路径push入结果数组
 * 
     时间复杂度：O(N)
    取决于树的节点数
    
    空间复杂度：O(N*M)
    取决于树节点数,队列不仅存储节点，还要存储路径，路径最长为树的深度
 */