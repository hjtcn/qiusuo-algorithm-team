// 面试题 04.02. 最小高度树
// 给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

// 示例:
// 给定有序数组: [-10,-3,0,5,9],

// 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

//           0 
//          / \ 
//        -3   9 
//        /   / 
//      -10  5 


/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */


var sortedArrayToBST = function(nums) {
    let len=nums.length
    if(!len){
        return null
    }
    let mid=Math.floor((len-1)/2)
    let root=new TreeNode(nums[mid])
    root.left=sortedArrayToBST(nums.slice(0,mid))
    root.right=sortedArrayToBST(nums.slice(mid+1))
    return root
};

/**
 * 递归。自己做的时候没思路，不知道怎么创建树，看完题解有点崩溃了。
 * 果然思路最重要，我都已经想到需不需要中序遍历啥的了，哈哈哈哈
 * 二分构建，找不到节点的时候返回null，不停二分高度就最低。
    时间复杂度：O(N)
    每个数都被执行了一番
    
    空间复杂度：O(1)
    除了构建结果树，数组长度值，以及初始数组，没定义其它变量
 */