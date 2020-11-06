<?php


// 剑指 Offer 59 - I. 滑动窗口的最大值
// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

// 示例:

// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7] 
// 解释: 

//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

// 提示：

// 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。


// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    // 函数法，可见非常慢
    // 执行用时：2208 ms, 在所有 PHP 提交中击败了6.25%的用户
    // 内存消耗：22.9 MB, 在所有 PHP 提交中击败了9.38%的用户
    function maxSlidingWindow1($nums, $k) {
        $data = []; // 结果数组
        foreach($nums as $key=>$value) {
            $arr = array_slice($nums, $key, $k); // 截取k长度的滑动窗口
            if(count($arr) < $k) break;
            // 查找最大的值
            $data[] = max($arr);
        }
        return $data;
    }


    // 执行用时：48 ms, 在所有 PHP 提交中击败了100.00%的用户                                                                                             
    // 内存消耗：23.3 MB, 在所有 PHP 提交中击败了6.25%的用户
    function maxSlidingWindow($nums, $k) {
        // 边界判断
        $size = count($nums);
        if($size < 1) {
            return $nums;
        }
        // 队列下标
        $queue = [];

        // 结果数组
        $data = [];

        for($i=0; $i<$size; $i++) {
            // 如果队列有值，进行比较
            while($queue && $nums[end($queue)] < $nums[$i]) {
                // 队尾元素小于 当前值
                // 就移除队尾元素，直到队尾元素大于当前元素进行停止
                array_pop($queue);
            }

            // 向队尾添加元素 下标
            $queue[] = $i;

            if($queue[0] < $i+1-$k) {
                // 如果当前队首，也即是队列最大值，小于当前窗口左边界，那么队首就是无效的，移除队首
                array_shift($queue);
            }

            // 如果窗口形成，就把窗口最大值放入结果集
            if($i+1 >= $k) {
                $data[] = $nums[$queue[0]];// 队首元素就是最大的，单调递减队列
            }
        }
        return $data;

    }

}
// @lc code=end


/*

【1】函数法
    遍历整个数组，依次截取k大小的滑动窗口，并且使用max查找数组中最大的值

   max查找数组效率很慢，但是也是比较容易想到

【2】双端单调队列
     如何在每次窗口滑动后，将 “获取窗口内最大值” 的时间复杂度从 O(k) 降低至 O(1)O。


    单调队列还是比较晦涩难懂，害，还是多看，多画图理解，不然这代码都看不懂了。
*/