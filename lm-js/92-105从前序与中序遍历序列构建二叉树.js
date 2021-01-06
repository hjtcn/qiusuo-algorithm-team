// 105. 从前序与中序遍历序列构造二叉树
// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：

//     3
//    / \
//   9  20
//     /  \
//    15   7



/*
 * @lc app=leetcode.cn id=105 lang=javascript
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} preorder
 * @param {number[]} inorder
 * @return {TreeNode}
 */


//害，自己是完全没思路，只想到了大学考试的时候会考这个题目，但不是用代码解决的，一般就是涂涂画画，还要猜猜排除一些方案
//好吧！开始看题解


//递归，借用js的一些api，但是性能极差
var buildTree = function(preorder, inorder) {
    if(!inorder.length) return null;
    let root=new TreeNode(preorder[0])
    let mid=inorder.indexOf(preorder[0])
    //slice及其耗性能，其实思路上有分治的思想了。。。。
    root.left=buildTree(preorder.slice(1,mid+1),inorder.slice(0,mid))
    root.right=buildTree(preorder.slice(mid+1),inorder.slice(mid+1))
    return root
};

//去除slice
var buildTree=function(preorder,inorder){
    let helper=(p_start,p_end,i_start,i_end)=>{
        if(p_start>p_end) return null
        let rootVal=preorder[p_start]
        let root=new TreeNode(rootVal)
        let mid=inorder.indexOf(rootVal)
        let leftNum=mid-i_start
        root.left=helper(p_start+1,p_start+leftNum,i_start,mid-1)
        root.right=helper(p_start+leftNum+1,p_end,mid+1,i_end)
        return root
    }
    return helper(0,preorder.length-1,0,inorder.length-1)
}

//去除indexOf，使用map树存取索引值
var buildTree=function(preorder,inorder){
    let map_i=new Map()
    for(let i=0;i<inorder.length;i++){
        map_i.set(inorder[i],i)
    }
    let helper=(p_start,p_end,i_start,i_end)=>{
        if(p_start>p_end) return null
        let rootVal=preorder[p_start]
        let root=new TreeNode(rootVal)
        let mid=map_i.get(root.val)
        let leftNum=mid-i_start
        root.left=helper(p_start+1,p_start+leftNum,i_start,mid-1)
        root.right=helper(p_start+leftNum+1,p_end,mid+1,i_end)
        return root
    }
    return helper(0,preorder.length-1,0,inorder.length-1)
}


//美国大佬的解法。
// 1.构建一个结束点
// 2.前序遍历是根左右,以前序数组为底构建节点，不同的某个节点的左右子节点有为空的可能

var buildTree = function(preorder, inorder) {
    let pre = i = 0
    let build =stop=>{
        console.log(i,stop)
        if (inorder[i] != stop) {//这里不太懂
            var root = new TreeNode(preorder[pre++])
            root.left = build(root.val)
            i++
            root.right = build(stop)
            return root
        }
        return null
    }
    return build()
};

//是我太菜，代码也看不懂，啊啊啊啊，打印几个结果再缕缕思路
//打印结果：
// 0 undefined
// 0 3
// 0 9
// 1 3
// 2 undefined
// 2 20
// 2 15
// 3 20
// 4 undefined
// 4 7
// 5 undefined

// @lc code=end

