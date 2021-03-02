<?php
/*
 * @Descripttion: 周二来了！
 * @Author: tacks321@qq.com
 * @Date: 2021-03-02 10:08:33
 * @LastEditTime: 2021-03-02 14:08:39
 */

/*
 * @lc app=leetcode.cn id=287 lang=php
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (66.15%)
 * Likes:    1099
 * Dislikes: 0
 * Total Accepted:    123.6K
 * Total Submissions: 186.3K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
 * 
 * 假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,3,4,2,2]
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,1,3,4,2]
 * 输出：3
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,1]
 * 输出：1
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：nums = [1,1,2]
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * nums.length == n + 1
 * 1 
 * nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 如何证明 nums 中至少存在一个重复的数字?
 * 你可以在不修改数组 nums 的情况下解决这个问题吗？
 * 你可以只用常量级 O(1) 的额外空间解决这个问题吗？
 * 你可以设计一个时间复杂度小于 O(n^2) 的解决方案吗？
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // [1]哈希法 
    // 时间复杂度O(n) 执行用时：24 ms, 在所有 PHP 提交中击败了54.76%的用户
    // 空间复杂度O(n) 内存消耗：18 MB, 在所有 PHP 提交中击败了12.82%的用户
    function findDuplicate($nums) {
        $hash = []; // 额外空间 哈希
        foreach($nums as $v) {
            if(isset($hash[$v])) {
                return $v;
            }
            // 复制到hash
            $hash[$v] = $v;
        }
    }

}

class Solution2 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // [2]函数法
    // 这个方法拖拖拉拉太麻烦了
    // 执行用时：24 ms, 在所有 PHP 提交中击败了54.76%的用户
    // 内存消耗：18 MB, 在所有 PHP 提交中击败了12.82%的用户
    function findDuplicate($nums) {
        // 查询出来对应值的重复个数，除了重复值，其他都是1
        // array_count_values() 统计数组中所有的值。返回的是 关联数组（值 => 出现的次数）
        $nums = array_count_values($nums);

            // 另一个函数
            // return array_search(max($nums),$nums);

        // 键值反转
        $nums = array_flip($nums);
        // 删除key=1的数组，也就是
        unset($nums[1]);
        // 重建索引
        $nums = array_values($nums);
        // 返回
        return $nums[0];
    }
}



class Solution3 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // [3]二分
    // 时间复杂度 O(nlogn) 执行用时：24 ms, 在所有 PHP 提交中击败了54.76%的用户
    // 空间复杂度 O(1)     内存消耗：17.3 MB, 在所有 PHP 提交中击败了30.77%的用户
    function findDuplicate($nums) {
        $len  = count($nums);
        $left = 1;
        $right= $len - 1;
        $ans  = -1; // 最终答案
        while($left <= $right) {
            // 数组中间位置
            $mid = $left + floor( ($right - $left) / 2);

            // 当前元素出现次数
            $cnt = 0;
            // ????????????? 二分还是有点迷
            foreach($nums as $num) {
                if($num <= $mid) {
                    $cnt++;
                }
            }
            
            // 二分搜索
            if($cnt <= $mid) {
                // 目标在右边
                $left = $mid + 1;
            } else {
                // 目标在左边
                $right= $mid - 1;
                $ans  = $mid;
            }
        }
        return $ans;
    }
}



class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // [4]快慢指针?????????????
    // 时间复杂度O(n) 执行用时：20 ms, 在所有 PHP 提交中击败了85.71%的用户
    // 空间复杂度O(1) 内存消耗：17.2 MB, 在所有 PHP 提交中击败了53.85%的用户
    function findDuplicate($nums) {
        $slow = $fast = 0;
        // 第一阶段
        do {
            $slow = $nums[$slow];
            $fast = $nums[$nums[$fast]];
        }while($slow != $fast);

        $res = 0;

        // 第二阶段
        while($slow != $res) {
            $slow = $nums[$slow];
            $res = $nums[$res];
        }
        return $res;
    }
}


// @lc code=end

/*
【题目】
    其实这个题，貌似在面试的时候被问到过，差不多一样意思的题。

    本题指明，nums数组中只有一个整数出现多次，其他都是一次。找到这个重复的整数。nums中整数都是 1~n

    当然这种题目，我们如果要解决，一般业务代码都会采用hash来进行判断

    想要用到二分，本题貌似并不是有序数组

    重点：nums数组大小是 ( 1~ n+1 ), 包含整数范围是 ( 1~n )
        假定重复数为 target , cnt[i]表示的是小于等于i的次数
        [1] 如果 target 出现两次，那么整好其他的整数个出现一次
            因为小于 \textit{target}target 的数 ii 满足 \textit{cnt}[i]=icnt[i]=i，大于等于 \textit{target}target 的数 jj 满足 cnt[j]=j+1cnt[j]=j+1
            => 小于target的数，一定是cnt[i] = i
            => 大于等于 target的数， 一定是 cnt[j] = j + 1
        [2] 如果 target 出现两次以上，那么其他整数有可能会缺失。



【解法】
    1、采用hash
        遍历一遍，利用对应的hash是否存在进行判断
    2、函数法 
        PHP的函数是坠屌的，但是时间复杂度并不会少
    3、二分（今天的主题）  
        稍微有点不一样了
            我们采用cnt[i] 来表示小于等于 i 的元素，那么cnt[i]是单调递增式的。

            我们需要找到第一个 cnt[i]>i 的位置（采用二分搜索）


            二分开始（二分搜索核心：缩小搜索空间）
                再次遍历nums
                    cnt 统计小于当前元素的元素个数
                如果cnt <= mid
                    那么就缩小左边的 left = mid + 1
                如果cnt > mid
                    那么缩小右边的 right = mid -1
                    // 记录下第一个cnt>mid的位置
                    ans = mid 
    4、快慢指针


*/