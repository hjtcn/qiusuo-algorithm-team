// 235. 二叉搜索树的最近公共祖先
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
};

// 时间复杂度：O(N)
// 取决于树的节点数
// 空间复杂度：O(N)
// 取决于树的节点数，队列模拟遍历树节点

//题目的核心在于二叉搜索树：左子树的节点都小于根节点，右子树的节点都大于根节点

//弄清这个递归方法就非常简单了，当然我一开始不知道这些是完全没有思路的。
//知道这道题是剑指offer原题，但是上次因为没有js的格式就跳过了，这次这里有了，这周五也要把之前那道题补上来


var lowestCommonAncestor = function(root, p, q) {
    let queue=[root]
    while(queue.length){
        let curNode=queue.shift()
        if(curNode.val>p.val&&curNode.val>q.val){
            queue.push(curNode.left)
        }
        else if(curNode.val<p.val&&curNode.val<q.val){
            queue.push(curNode.right)
        }
        else{
            return curNode
        }
    }
    return root
};


//BFS。弄清规律后思路确实蛮简单的，手撕一遍过，啦啦啦，二叉搜索树规律要记牢


// 时间复杂度：O(N)
// 取决于树的节点数
// 空间复杂度：O(N)
// 取决于树的节点数，队列模拟遍历树节点


//还在题解中看到迭代的解法，和BFS很相似，但是没有利用数组模拟队列。
//而是直接利用root节点指针的跳转，最终返回root节点
//空间复杂度肯定是要比BFS要优化一些的
var lowestCommonAncestor = function(root, p, q) {
    while(root){
        if(root.val>p.val&&root.val>q.val){
            root=root.left
        }
        else if(root.val<p.val&&root.val<q.val){
            root=root.right
        }
        else{
           break;
        }
    }
    return root
};
