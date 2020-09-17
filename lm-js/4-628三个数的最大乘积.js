// 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

// 示例 1:

// 输入: [1,2,3]
// 输出: 6
// 示例 2:

// 输入: [1,2,3,4]
// 输出: 24
// 注意:

// 给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
// 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。


/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumProduct = function (nums) {
    //先从小到大排序，可知道三个数字最大的乘积，只可能出现在nums[0],nums[1],nums[len-1],nums[len-2],nums[len-3]之间，且nums[len-1]肯定会存在
    nums.sort((a, b) => a - b)
    let right = nums.length - 1
    //存储最大值
    let res = nums[right]
    //记录最小两个数的乘积值，最大两个数的乘积值
    let num1 = nums[0] * nums[1], num2 = nums[right - 1] * nums[right - 2]
    //比较三个数乘积的大小。
    res = Math.max(res * num1, res * num2)
    return res
};

/** 题解
三个数的最大乘积。先sort排序（数组长度<10?'插入排序':'快排'）
先从小到大排序，可知道三个数字最大的乘积，只可能出现在nums[0],nums[1],nums[len-1],nums[len-2],nums[len-3]之间，且nums[len-1]肯定会存在
故只需比较(nums[len-1]*nums[0]*nums[1],nums[len-1]*nums[len-2]*nums[len-3])的值谁大就行

 复杂度分析：
    平均时间复杂度是:O(nlogn)
    快排的平局时间复杂度是O(nlogn)
    空间复杂度：O(1)
    定义了几个变量，存放比大小的值
 */

/**
* @param {number[]} nums
* @return {number}
*/
var maximumProduct = function (nums) {
    //max1存放最大值，max2第二大，max第三大，min1是最小值，min2是第二小的值
    let max1 = -Infinity, max2 = -Infinity, max3 = -Infinity, min1 = Infinity, min2 = Infinity
    for (let i = 0; i < nums.length; i++) {
        //如果遇到比max1还大的，说明是最大的值，就可以替换最大值，且原max1,max2往后移，变为第二大，第三大了。
        if (nums[i] > max1) {
            [max1, max2, max3] = [nums[i], max1, max2]
        }
        else if (nums[i] > max2) {//走到这里，说明没max1大，比max2大，那就还是替换第二大的值，然后原max2往后移，变为max3了
            [max2, max3] = [nums[i], max2]
        }
        else if (nums[i] > max3) {//走到这里，说明只比max3大，那就替换max3的值。
            max3 = nums[i]
        }
        if (nums[i] < min1) {//走到这里，比min1还小，就替换原min1，原min1变为min2了
            [min1, min2] = [nums[i], min1]
        }
        else if (nums[i] < min2) {//走到这里，比min2小，就只替换min2
            min2 = nums[i]
        }
    }
    //max1是必定存在在三数之中，所以，对比(max1 * max2 * max3, max1 * min1 * min2)的较大值返回即可
    return Math.max(max1 * max2 * max3, max1 * min1 * min2)
};

/** 题解
三个数的最大乘积。根据线性，利用解构赋值，一层层比较，替换，然后找出最小的两个数(min1,min2)，和最大的三个数(max3,max2,max1)
三个数字最大的乘积，max1作为最大值是必定存在在三数之中，所以，对比(max1 * max2 * max3, max1 * min1 * min2)的较大值返回即可

 复杂度分析：
    平均时间复杂度是:O(N)
    一层for循环遍历
    空间复杂度：O(1)
    定义了几个变量，存放两个最小的两个数和最大的三个数
 */