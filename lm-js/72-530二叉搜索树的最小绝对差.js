// 530. 二叉搜索树的最小绝对差
// 给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

 

// 示例：

// 输入：

//    1
//     \
//      3
//     /
//    2

// 输出：
// 1

// 解释：
// 最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 

// 提示：

// 树中至少有 2 个节点。

var getMinimumDifference = function(root) {
    let arr=[],minVal=Infinity
    let dfs=(root)=>{
    if(!root) return;
     dfs(root.left)
     arr.push(root.val)
     dfs(root.right)
    }
    dfs(root)
    for(let i=1;i<arr.length;i++){
        if(Math.abs(arr[i]-arr[i-1])<minVal){
            minVal=Math.abs(arr[i]-arr[i-1])
        }
    }
    return minVal
};

/**
    //递归。趁热打铁，依然是中序遍历是递增序列的原理
    一开始又没读懂题，以为是相邻两个节点的最小绝对差，哈哈哈，做了一会这个思路
    将递增序列存入数组后，一次for循环，更新最小值

    时间复杂度：O(N)
    取决于树的节点数。
    空间复杂度：O(N)
    取决于树节点数，不会超过n
 */


 //迭代
 var getMinimumDifference = function(root) {
    let lastNode=Infinity,minVal=Infinity,queue=[]
    while(root||queue.length){
        while(root){
            queue.push(root)
            root=root.left
        }
        root=queue.pop()
        if(Math.abs(root.val-lastNode)<minVal){
            minVal=Math.abs(root.val-lastNode)
        }
        lastNode=root.val
        root=root.right
    }
    return minVal
};

/**
    //迭代。不断更新队列的形式去中序遍历。
    后续想，push的过程就可以更新最小值，是不是算是一种优化呢
    最后提交的时候，空间复杂度是优化了一些，不过时间复杂度也多了一层判断。
    时间复杂度：O(N)
    取决于树的节点数。
    空间复杂度：O(N)
    取决于树节点数，不会超过n
 */