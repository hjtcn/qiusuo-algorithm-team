// 153. 寻找旋转排序数组中的最小值
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。

// 请找出其中最小的元素。

 

// 示例 1：

// 输入：nums = [3,4,5,1,2]
// 输出：1
// 示例 2：

// 输入：nums = [4,5,6,7,0,1,2]
// 输出：0
// 示例 3：

// 输入：nums = [1]
// 输出：1
 

// 提示：

// 1 <= nums.length <= 5000
// -5000 <= nums[i] <= 5000
// nums 中的所有整数都是 唯一 的
// nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转
/*
    思路：方法一：遍历。
                首先先老老实实遍历把这道题写一边，当出现比前一个数小的值，则直接返回，否则，返回第一个值。
         方法二：二分：我发现二分首先要确定的是中间值和谁进行比较来进行左右方向的判断。
         这道题则可以根据第一个元素(或末尾元素)和mid元素进行比较。
         如果nums[0]>nums[mid]  则左侧肯定出现过旋转。r=mid-1
            nums[0]<nums[mid]  则左侧肯定没有出现过旋转。要么首元素就是最小元素(这个提前return或者押后和nums[mid]PK都行)，要么旋转位置还在右边呢。
                                l=mid+1

*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findMin = function(nums) {
    for(let i=1;i<nums.length;i++){
        if(nums[i]<nums[i-1]){
            return nums[i]
        }
    }
    return nums[0]
};
/* 
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

var findMin = function(nums) {
    let l=0,r=nums.length-1,mid=0,res=nums[0]
    if(nums[r]>nums[l]){
        return nums[l]
    //这个判断还让我发现了，一开始我的旋转题意理解错了。
    // 举例来说 [1,2,3],[2,3,1],[3,1,2]为正确的输入
     // [2,1,3]不是正确的输入。
    }
    while(l<=r)
    {
         mid=parseInt((l+r)/2)
         if(nums[mid]<nums[mid-1]){
             return nums[mid]
         }
        if(nums[mid]<nums[0]){
            r=mid-1
        }
        else{
            l=mid+1
        }

    }
    return nums[mid]
}
/* 
    这道题我的解题思路是。1.提前判断目标元素为首元素的情况。
                      2.二分，中间元素与首元素比较，进行左右方向的二分。
                        特殊情况为：当前mid<前一位元素，则直接返回
                      3.返回中间元素。
    时间复杂度:O(logn)
    空间复杂度：O(1)

*/

