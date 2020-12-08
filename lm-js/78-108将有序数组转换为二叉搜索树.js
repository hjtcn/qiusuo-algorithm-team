// 108. 将有序数组转换为二叉搜索树
// 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 示例:

// 给定有序数组: [-10,-3,0,5,9],

// 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

//       0
//      / \
//    -3   9
//    /   /
//  -10  5


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
  if(!len) return null
  let mid=parseInt((len-1)/2),root=new TreeNode(nums[mid])
  root.left=sortedArrayToBST(nums.slice(0,mid))
  root.right=sortedArrayToBST(nums.slice(mid+1))
  return root
}

/*
    递归。不停二分作为根节点，再找左右节点
    时间复杂度：O(N)
    每个节点只会构建一次
    空间复杂度：O(N)
    包含树的节点数。
*/

//  还觉得就是这种解决方案了
//一看官方题解，发现还是跟第一次看到这道题的想法一致，可以依靠中序遍历的概念来解决这个题
//但当时看最小高度树这题的题解，没有中序遍历的思路，就放弃了

//其实都是二分，但是思维扩展是不一样的。而且耗时竟然少了很多，看看下面吧。

var sortedArrayToBST = function(nums) {
  const helper=(nums,left,right)=>{
      if(left>right){
           return null
      }
      //选择如果奇数选择中间位置即可，偶数选择靠左的还是靠右的，或是随机的节点作为根节点
      let mid=parseInt((left+right)/2)//靠左
      // let mid=parseInt((left+right+1)/2) 靠右
      // let mid=parseInt(left+right+Math.random(2)/2)//随机靠左右两边，嘻嘻嘻
      let root=new TreeNode(nums[mid])
      root.left=helper(nums,left,mid-1)
      root.right=helper(nums,mid+1,right)
      return root
  }
  return helper(nums,0,nums.length-1);
}

/*
    递归。额外创建一个函数，去递归，记录左右节点，愿数组也要作为参数传递，因为需要new TreeNode(nums[mid])构建节点
    时间复杂度：O(N)
    每个节点只会构建一次
    空间复杂度：O(N)
    包含树的节点数。
*/