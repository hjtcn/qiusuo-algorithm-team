/*
162. 寻找峰值
峰值元素是指其值大于左右相邻值的元素。

给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

 

示例 1：

输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例 2：

输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5 
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
 

提示：

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]


*/

/*
    思路：这次就没有像上次一样，去理解趋势，再小心向峰值靠拢。
        可能题做的多了，敢大胆发散了
        let mid=(l+r)/2
        1.首先确定目标范围，直接返回mid
        if(nums[mid]>nums[mid-1]&&nums[mid]>nums[mid+!]){
            return mid
        }
        2.如果不在上述范围，那那种范围可以l向右靠拢呢？
        else if(nums[mid]<nums[mid+1]){
            l=mid+1
        }
        3.那种范围需要r向右靠拢呢？
        else if(nums[mid]>nums[mid+1]){
            r=mid
        }

        但最后最后我还是踩了坑。
        1.while循环的准确条件到底是什么呢？
        while(l<r)  or   while(l<=r)
        区别在于l可需不需要等于r或者说可不可以等于r?
        不可以!!l=r,mid=l=r,一旦遇到第3种范围，r=mid，就会陷入死循环。

        2.返回值为l，mid还是r？？？？
        while循环条件为while(l<r)
        跳出循环为l>=r
        它的上一步肯定是跳到了第2种范围
        l=mid+1。此时的nums[mid+1]就是比较大的值呀，
        至少nums[mid+1]>nums[mid]
        因此，最后应该返回l
*/

var findPeakElement = function(nums) {
    let l=0,r=nums.length-1
    let mid=parseInt((l+r)/2)
    while(l<r){
        mid=parseInt((l+r)/2)
        if(nums[mid]>nums[mid-1]&&nums[mid]>nums[mid+1]){
            return mid
        }
        else if(nums[mid]<nums[mid+1]){
            l=mid+1
        }
       else if(nums[mid]>nums[mid+1]){
            r=mid
        }

    }
    return l
};

/*
    时间复杂度:O(logN)
    空间复杂度:O(1)
*/

/*
    思考：可以详细的把二分的各种条件，边界推一遍，下次尝试完整的考虑一下，不能继续蒙下去了。
        周五要团建，提前把题刷了，开心心。
*/