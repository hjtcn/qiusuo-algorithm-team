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
 

// 提示：

// 节点总数 <= 10000


/** 
 思路：看到这个题思路还是蛮清晰的，知道要记录节点，路径，剩余差值
 我先尝试使用dfs做了这道题。
 虽然一开始路比较清晰，但是对准到准确的节点来说，我的第一套代码写出来的时候，发现案例输出出现重复了。
     let dfs=(root,path,sum)=>{
        if(!root){
            if(sum==0){
                res.push(path)
            }
            return ;
        }
        dfs(root.left,[...path,root.val],sum-root.val)
        dfs(root.right,[...path,root.val],sum-root.val)
    }
 这个时候我确认首先差值递减，和路径记录是没有问题，然后仔细深入思考问题，发现自己左右节点递归时，记录的路径却都是当前节点值！！！
 说明什么？左右节点的分叉路口，我根本没分开。。
 然后我就进一步调整，分开左右节点走。
 let dfs = (root, path, sum) => {
        if (sum - root.val == 0) {
            res.push(path)
            return;
        }
        root.left && (dfs(root.left, [...path, root.left.val], sum - root.val))
        root.right && (dfs(root.right, [...path, root.right.val], sum - root.val))
    }
思路走了一遍，case也过了，觉得没有问题，就提交了，然后就解答错误了。
错误案例：[1,2],1

好吧，我知道了，我又出现了没有走完整条路的问题。接着修改
 if (!root.left && !root.right && sum - root.val == 0) {//让每条路走完，当同时没有左右节点时
            res.push(path)
            return;
}

dfs打通，BFS就很顺利了，把递归的过程换为入队的过程，啦啦啦

*/
//dfs
var pathSum = function (root, sum) {
    if (!root) return []
    let res = []
    let dfs = (root, path, sum) => {
        if (!root.left && !root.right && sum - root.val == 0) {
            res.push(path)
            return;
        }
        root.left && (dfs(root.left, [...path, root.left.val], sum - root.val))
        root.right && (dfs(root.right, [...path, root.right.val], sum - root.val))

    }
    dfs(root, [root.val], sum)
    return res
}

//BFS
var pathSum = function(root, sum) {
    if(!root) return []
    let res=[],queue=[[root,[root.val],sum]]
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
}

/** 
 没有额外的空间时间需要

 两种解决方案的时间空间复杂度基本一致
 时间复杂度为:O(N)
 取决于二叉树的节点个数
 空间复杂度:O(N)
 取决于二叉树的节点个数
*/