// 101. 对称二叉树
// 给定一个二叉树，检查它是否是镜像对称的。

 

// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
 

// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

//     1
//    / \
//   2   2
//    \   \
//    3    3
 

// 进阶：

// 你可以运用递归和迭代两种方法解决这个问题吗？

/*
思路
 这个题重点在于左右节点的构建与对比
 递归：
 如左右节点left,right
 三种情况：if都为空，则返回true
         else if 一个为空或者两节点val不同，返回false
         else 继续递归dfs(left.left,right.right)&&dfs(left.right,right.left)

迭代:
把root当作为左右节点放在队列中[[root,root]]
队列的循环中。左右节点对比：
    基本情况和递归类似：if都为空，则continue。//因为不是递归了，如果return就直接返回结果了
                     else if 一个为空或两个节点值不相等，返回false //提前结束
                    else  (left.left,right.right)入队，（left.right,right.left)入队，//构建左右对比的队列

*/


var isSymmetric = function(root) {
    if(!root) return true
    let dfs=(left,right)=>{
        if(!left&&!right) return true
        else if(!left||!right) return false
        else if(left.val!=right.val) return false
        return dfs(left.left,right.right)&&dfs(left.right,right.left)
    }
    dfs(root.left,root.right)
}



var isSymmetric = function(root) {
    if(!root) return true
    let queue=[[root,root]]
    while(queue.length){
        let [left,right]=queue.shift()
        if(!left&&!right) continue
        if(!left||!right||left.val!=right.val) return false
        queue.push([left.left,right.right])
        queue.push([left.right,right.left])      
    }
    return true
}

// 时间复杂度：O(N)
// 取决于节点个数

// 空间复杂度：O(N)
// 取决于节点数

/* 
    今天一下做了两题，还是比较顺手的，明天有时间把前两道题再刷一遍，熟练熟练
*/