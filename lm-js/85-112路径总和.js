// 112. 路径总和
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

// 说明: 叶子节点是指没有子节点的节点。

// 示例: 
// 给定如下二叉树，以及目标和 sum = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \      \
//         7    2      1
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

    var hasPathSum = function(root, sum) {
        let flag=0
        let helper=(root,sum)=>{
            if(!root) return false
            if(!root.left&&!root.right&&sum==root.val){
                    flag=1
                    return true
            }
            helper(root.left,sum-root.val)
            helper(root.right,sum-root.val)
        }
        helper(root,sum)
        return flag
    };
    //这是自己一开始写的代码，控制返回的时候有点乱，就觉得不用再另写一个递归函数，然后又尝试尝试精简代码

    var hasPathSum = function(root, sum) {
        if(!root) return false
        if(!root.left&&!root.right&&sum==root.val){
                return true
        }
        return hasPathSum(root.left,sum-root.val)||hasPathSum(root.right,sum-root.val)
    };


    //BFS
    var hasPathSum=function(root,sum){
        if(!root) return false
        let queue=[[root,sum]]
        while(queue.length){
            let [curNode,sum]=queue.shift()
            if(!curNode.left&&!curNode.right&&sum==curNode.val){
                return true
            }
            curNode.left&&(queue.push([curNode.left,sum-curNode.val]))
            curNode.right&&(queue.push([curNode.right,sum-curNode.val]))
        }
        return false
    }


    //虽然这道题比较简单，但是自己写出来两种方案也证明了熟练度确实有加深，棒棒哒
    //尤其是写广搜的时候，也学会了同时记录节点和差值。

    // 时间复杂度：O(N)
    // 取决于树的节点数
    // 空间复杂度：O(N)
    // 取决于树的节点数，队列模拟遍历树节点
