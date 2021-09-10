/*
654. 最大二叉树
给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

 

示例 1：


输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
    - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
        - 空数组，无子节点。
        - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
            - 空数组，无子节点。
            - 只有一个元素，所以子节点是一个值为 1 的节点。
    - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
        - 只有一个元素，所以子节点是一个值为 0 的节点。
        - 空数组，无子节点。
示例 2：


输入：nums = [3,2,1]
输出：[3,null,2,null,1]
 

提示：

1 <= nums.length <= 1000
0 <= nums[i] <= 1000
nums 中的所有整数 互不相同
*/

/*
    思路：递归中有排序，时间复杂度不太优。
        一个是原始数组nums，一个是排序的数组copyNums，找到最大值node作为根节点,位置为mid。
        node.left=dfs(nums.slice(0,mid))
        node.right=dfs(nums.slice(mid+1))
        return node
    题解：记录左右边界，去找到根节点(也就是最大值)
*/
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
var constructMaximumBinaryTree = function(nums) {
    if(!nums.length){
        return null
    }
    let copyNums=[...nums].sort((a,b)=>a-b)
    let len=nums.length
    let node=new TreeNode(copyNums[len-1])
    let mid=nums.indexOf(copyNums[len-1])
    let leftNum=nums.slice(0,mid),rightNum=nums.slice(mid+1)
    node.left=constructMaximumBinaryTree(leftNum)
    node.right=constructMaximumBinaryTree(rightNum)
    return node
};

var constructMaximumBinaryTree = function(nums) {
    let dfs=(left,right)=>{
        if(left>right) return null
        let max=-Infinity,index=-1
        for(let i=left;i<=right;i++){
            if(nums[i]>max){
                max=nums[i]
                index=i
            }
        }
        let node=new TreeNode(max)
        node.left=dfs(left,index-1)
        node.right=dfs(index+1,right)
        return node
    }
    return dfs(0,nums.length-1)
}

/*
    时间复杂度：O(NlogN)
    空间复杂度：O(N)
*/
