<?php
/*
 * @lc app=leetcode.cn id=1438 lang=php
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 *
 * https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/
 *
 * algorithms
 * Medium (38.16%)
 * Likes:    46
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 13.8K
 * Testcase Example:  '[8,2,4,7]\n4'
 *
 * 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于
 * limit 。
 * 
 * 如果不存在满足条件的子数组，则返回 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [8,2,4,7], limit = 4
 * 输出：2 
 * 解释：所有子数组如下：
 * [8] 最大绝对差 |8-8| = 0 <= 4.
 * [8,2] 最大绝对差 |8-2| = 6 > 4. 
 * [8,2,4] 最大绝对差 |8-2| = 6 > 4.
 * [8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
 * [2] 最大绝对差 |2-2| = 0 <= 4.
 * [2,4] 最大绝对差 |2-4| = 2 <= 4.
 * [2,4,7] 最大绝对差 |2-7| = 5 > 4.
 * [4] 最大绝对差 |4-4| = 0 <= 4.
 * [4,7] 最大绝对差 |4-7| = 3 <= 4.
 * [7] 最大绝对差 |7-7| = 0 <= 4. 
 * 因此，满足题意的最长子数组的长度为 2 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [10,1,2,4,7,2], limit = 5
 * 输出：4 
 * 解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 0 <= limit <= 10^9
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $limit
     * @return Integer
     */
    // 单调增队列 
    // 容易超时
    function longestSubarray1($nums, $limit) {
       
        $size = count($nums);
        // 维护两个单调队列 
        $Q1  = new Queue(); // 单调增  队首最小值
        $Q2  = new Queue(); // 单调减  队首最大值
        // abs(q1.front() - q2.front()) > limit时，窗口缩小，j++
        // 滑动窗口 表示 (j, i)
        $result= 0;
        for($i = 0, $j = 0; $i< $size; $i++) {
            // 单调增队列 维护最小值
            while(!$Q1->isEmpty() && $Q1->back() >= $nums[$i]) {
                $Q1->pop_back();
            }
            $Q1->push_back($nums[$i]); // 放进去最小值
            // 单调减队列 维护最大值
            while(!$Q2->isEmpty() && $Q2->back() <= $nums[$i]) {
                $Q2->pop_back();
            }
            $Q2->push_back($nums[$i]); // 放进去最大值

            // 超出滑动窗口大小
            while(!$Q1->isEmpty() && !$Q2->isEmpty() && $Q2->front() - $Q1->front() >= $limit) {
                if($nums[$j] == $Q1->front()) $Q1->pop_front(); // 从队头出队
                if($nums[$j] == $Q2->front()) $Q2->pop_front();
                $j++;
            }
            $result = max($result , $i-$j+1);

        }
         
        return $result;
    }

 
    // 双指针+滑动窗口
    // 执行用时：3960 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：26.5 MB, 在所有 PHP 提交中击败100.00%的用户
    function longestSubarray2($nums, $limit) {
        $min = 0;
        $max = 0;
        $minIndex = 0;// 滑动窗口最小值的位置
        $maxIndex = 0;// 滑动窗口最大值的位置
        $i = 0;
        $number = 1; // 窗口大小
        $size = count($nums);
        while($i < $size - 1 && $j < $size) {
            $start = $i; //滑动窗口左边的值
            $j = $i + 1;
            // 符合题意要求 绝对值之差小于limit 进入循环，不断向后扩大滑动窗口大小
            if(abs($nums[$j] - $nums[$i]) <= $limit) {
                // 符号条件更新位置
                if($nums[$j] > $nums[$i]) {
                    $maxIndex = $j;
                    $minIndex = $i;
                }else{
                    $maxIndex = $i;
                    $minIndex = $j;
                }
                $max = max($nums[$j], $nums[$i]);
                $min = min($nums[$j], $nums[$i]);
                // 当前滑动窗口的最大 最小 之差小于 limit
                while(abs($max - $min) <= $limit  && $j < $size){
                    $j++; // 向后移动
                    // 如果后面的元素，大于当前滑动窗口最大值
                    if($nums[$j] >= $max) {
                        // 最大值-最小值 不符合要求 ，滑动窗口左边收缩
                        if($nums[$j] - $min > $limit) {
                            $i = $minIndex + 1;// 移动i
                        }else{
                            // 并且符合要求，那么就更新最大值的位置
                            $maxIndex = $j;
                        }
                    }
                    // 如果后面的元素，小于当前滑动窗口最小值
                    if($nums[$j] <= $min) {
                        if($max - $nums[$j] > $limit) {
                            // 不符合要求，滑动窗口收缩
                            $i = $maxIndex + 1;
                        }else{
                            // 符合要求，更新最小值的位置
                            $minIndex = $j;
                        }
                    }
                    // 得到当前滑动窗口的最大值 和 最小值
                    $max = max($nums[$j], $max);
                    $min = min($nums[$j], $min);
                }
                // 记录窗口位置
                $number = max($number , $j - $start);
            }else{
                // 不符合的直接跳过
                $i++;
            }
        }
        return $number;
        
    }

}

 
// 数组模拟队列
class Queue 
{
    private $queue;

    public function __construct()
    {
       $this->queue = [];
    }

    public function isEmpty()
    {
        return empty($this->queue);
    }

    // 移除队首元素
    public function pop_front()
    {
        if ($this->isEmpty()) {
            return false;
        } else {
            // 将数组开头的单元移出数组
            array_shift($this->queue);
        }
    }

    // 移除队尾元素
    public function pop_back()
    {
        if ($this->isEmpty()) {
            return false;
        } else {
            // 弹出数组最后一个单元 
            return array_pop($this->queue);
        }
    }

    // 入队
    public function push_back(string $item)
    {
        // 压入数组的末尾
        array_push($this->queue, $item);
    }

    // 队首元素
    public function front()
    {
        // 将数组的内部指针指向第一个单元
        return reset($this->queue);
    }

    // 队尾元素
    public function back()
    {
        return end($this->queue);
    }

}


// @lc code=end

$obj = new Solution();
// echo $obj->longestSubarray2([8,2,4,7],4);


// @tacks think=start
/*
题意 meaning 
  
    子数组！

    题意说，子数组任意的两个元素的绝对差 <= limit    也就是说， 子数组中最大值与最小值的差 <= limit

关键 key 
     
    “绝对值之差小于limit” “什么时候滑动窗口移动变化”

想法 idea
 
    这题足足花费三四个小时还是没整出来，看了题解，也没有PHP的，然后就读了好几个题解。
先暂时记录一下，之后遇到同样类似的题，在会过头看一下。

【1】双端单调队列实现
    两个单调队列，一个单调增维护最小值，一个单调减维护最大值。
    但是这个PHP容易超时。

【2】双指针+滑动窗口
    维护滑动窗口最大值最小值，并且保存其位置。
    双指针在不断移动的时候，滑动窗口的大小也在变化，同时判断max-min <= limit 
  

*/
// @tacks think=end
