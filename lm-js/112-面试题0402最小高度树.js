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

/* 
    有序数组。中序遍历？左根右？
    递归：
    二分查找根，即左右子节点

    迭代：尝试模拟队列，但是没有成功，索引切割左右数组的边界值控制不成功，有时间再试试。。。。
*/

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
    let mid=Math.floor(nums.length/2)
    if(nums.length==0){
        return null;
    }
    let root=new TreeNode(nums[mid])
    root.left=sortedArrayToBST(nums.slice(0,mid))
    root.right=sortedArrayToBST(nums.slice(mid+1))
    return root
}

// 时间复杂度：O(N)
// 取决于数组长度递归每个节点

// 空间复杂度：O(N)
// 取决于节点个数