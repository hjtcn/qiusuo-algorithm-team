// 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]



 

// 示例 1:

// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6 
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。
// 示例 2:

// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// 输出: 2
// 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 

// 说明:

// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。



//*********************************************************
//读题！读题！读题！二叉搜索树
//二叉搜索树：左<根<右。搞清楚，题就贼简单了


//读题：二叉搜索树：左<根<右
var lowestCommonAncestor = function(root, p, q) {
    if(root.val<p.val&&root.val<q.val){
        return lowestCommonAncestor(root.right,p,q)
    }
    else if(root.val>p.val&&root.val>q.val){
        return lowestCommonAncestor(root.left,p,q)
    }
    else{
        return root
    }
}
/** 
    递归逻辑：
    如果当前节点root皆小于p,q的值，就只能从root.right递归寻找
    如果当前节点root皆大于p,q的值，就只能从root.left递归寻找
    如果当前节点root大于其中一个，小于其中另一个，则返回root
    // 时间复杂度：O(N)
    // 取决于树的节点数
    // 空间复杂度：O(N)
    // 取决于树的节点数
*/



var lowestCommonAncestor = function(root, p, q) {
    let queue=[root]
    while(queue.length){
        let curNode=queue.shift()
        if(curNode.val<p.val&&curNode.val<q.val){
            queue.push(curNode.right)
        }
        else if(curNode.val>p.val&&curNode.val>q.val){
            queue.push(curNode.left)
        }
        else{
            return curNode
        }
    }
}

/*
    思路基本和上层一致，突然感觉有点摸清套路
    dfs和bfs的区别在于，深搜递归的位置在于广搜进行push入队的过程
    // 时间复杂度：O(N)
    // 取决于树的节点数
    // 空间复杂度：O(N)
    // 取决于树的节点数，队列模拟遍历树节点
*/


