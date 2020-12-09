// 404. 左叶子之和
// 计算给定二叉树的所有左叶子之和。

// 示例：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

var sumOfLeftLeaves = function(root) {
    let sum=0
    const helper=(root,type)=>{
        if(!root) return ;
        if(!root.left&&!root.right&&type==0){
            sum+=root.val
        }
        helper(root.left,0)
        helper(root.right,1)
    }
    helper(root,1)
    return sum
};

/** 
 * 递归，helper函数设置type参数，type为0时的叶子节点进行相加
 * 时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数
*/

var sumOfLeftLeaves = function(root) {
    if(!root) return 0
    let sum=0,queue=[root]
    while(queue.length){
        let curNode=queue.pop()
        if(curNode.left){
            curNode.left&&queue.push(curNode.left)
            if(!curNode.left.left&&!curNode.left.right){
                sum+=curNode.left.val
            }
        }
        
        curNode.right&&queue.push(curNode.right)
    }
    return sum
};

//bfs，判断当左子树没有孩子的时候，说明是左叶子节点。然后进行相加
/**
 * 时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数，队列模拟遍历树节点
 */