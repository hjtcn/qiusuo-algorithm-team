// 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。

// 如果不存在满足条件的子数组，则返回 0 。

//  

// 示例 1：

// 输入：nums = [8,2,4,7], limit = 4
// 输出：2 
// 解释：所有子数组如下：
// [8] 最大绝对差 |8-8| = 0 <= 4.
// [8,2] 最大绝对差 |8-2| = 6 > 4. 
// [8,2,4] 最大绝对差 |8-2| = 6 > 4.
// [8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
// [2] 最大绝对差 |2-2| = 0 <= 4.
// [2,4] 最大绝对差 |2-4| = 2 <= 4.
// [2,4,7] 最大绝对差 |2-7| = 5 > 4.
// [4] 最大绝对差 |4-4| = 0 <= 4.
// [4,7] 最大绝对差 |4-7| = 3 <= 4.
// [7] 最大绝对差 |7-7| = 0 <= 4. 
// 因此，满足题意的最长子数组的长度为 2 。
// 示例 2：

// 输入：nums = [10,1,2,4,7,2], limit = 5
// 输出：4 
// 解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
// 示例 3：

// 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
// 输出：3
//  

// 提示：

// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^9
// 0 <= limit <= 10^9

/**
 * @param {number[]} nums
 * @param {number} limit
 * @return {number}
 */
var longestSubarray = function(nums, limit) {
    let res = 1, n = nums.length,
        left = 0, right = 1,
        min = [ nums[0] ], max = [ nums[0] ];
    //不停维护最小值的数组及最大值的数组
    while (right < n) {
        //如果右指针的当前元素>最大值，就把max数组pop掉这个最大值
      while (min.length > 0 && min[ min.length - 1 ] > nums[right]) min.pop();
      //如果左指针的当前元素<最小值，就把min数组pop掉这个最小值
      while (max.length > 0 && max[ max.length - 1 ] < nums[right]) max.pop();
      //将当前右指针指向的元素，push到min数组和max数组中。
      min.push( nums[right] );
      max.push( nums[right] );
      
      while (Math.abs( max[0] - min[0] ) > limit) {//比较max数组和min数组的第一个元素，当绝对差大于limit，判断当前左指针是等于min[0]还是max[0],将min从头部删除，left++
        if (min[0] === nums[left]) min.shift();
        if (max[0] === nums[left]) max.shift();
        left ++;
      }
      
      res = Math.max(res, right - left + 1);
      right ++;
    }
    
    return res;
  };

//这道题并没有自己解决出来，想到了左右指针以及最大值和最小值的绝对差去控制，但没想到判断绝对差大于limit的情况下用while循环进行左指针的++，以及以数组的形式存储最大值和最小值
  /** 题解
1438-绝对差不超过限制的最长连续子数组。利用滑动窗口的形式。
最长连续子数组，考虑到需要找到满足任意两数的绝对差不大于limit的情况，即记录right-left+1,并不断更新较大的窗口长度。
如何确定right及left的值？首先当不出现窗口元素的最大值和最小值的绝对差不超过limit，right会在不超过原数组长度下一直加一，反之，当出现绝对差大于limit，就要循环left加一了。
而整个过程比较重要的就是如何维护min数组(递增数组，保证第一个元素为最小值)和max数组(递减元素，保证第一个元素为最大值)。
拿min数组举例：当遇到当前right指向的元素，小于当前min数组的最后一个元素，说明min的最后一个元素，随着left加加进行缩圈，不可能成为最小值了，需要把min.pop(),然后push当前right指向的元素。
然后考虑left加加的缩圈行为，循环条件为最大值和最小值的绝对差大于limit.如果min的首个元素等于当前元素，就把第一个元素从头部去掉，left可以直接加1.
max数组也是如此。
最后每一次右指针右移之前，先更新窗口的最大长度。
 复杂度分析：
    平均时间复杂度是:O(N)
    一层for循环变量，让right是否超过数组长度进行循环控制，第一次出现不符合情况，left跟上来，如果找到符合情况的，right继续向前走。
    空间复杂度：O(N)
    声明两个数组，放置最大值和最小值。
 */
