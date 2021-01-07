<?php
/*
 * @Descripttion: 每日一题，冲啊
 * @Author: tacks321@qq.com
 * @Date: 2021-01-04 18:08:19
 * @LastEditTime: 2021-01-04 19:40:25
 */


/*
 * @lc app=leetcode.cn id=509 lang=php
 *
 * [509] 斐波那契数
 *
 * https://leetcode-cn.com/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (66.14%)
 * Likes:    217
 * Dislikes: 0
 * Total Accepted:    125.2K
 * Total Submissions: 184.3K
 * Testcase Example:  '2'
 *
 * 斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
 * 
 * 
 * F(0) = 0，F(1) = 1
 * F(n) = F(n - 1) + F(n - 2)，其中 n > 1
 * 
 * 
 * 给你 n ，请计算 F(n) 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：2
 * 输出：1
 * 解释：F(2) = F(1) + F(0) = 1 + 0 = 1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：3
 * 输出：2
 * 解释：F(3) = F(2) + F(1) = 1 + 1 = 2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：4
 * 输出：3
 * 解释：F(4) = F(3) + F(2) = 2 + 1 = 3
 * 
 * 
 * 
 * 提示：
 * 
 * 0 <= n <= 30
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
     // 递归处理
    // 执行用时：4 ms, 在所有 PHP 提交中击败了95.45%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了10.91%的用户
    static $cache = []; // 缓存数组，用来计算计算过的值
    function fib($n) {
        if(isset($this->cache[$n])) {
            return $this->cache[$n];
        } else {
            if($n < 2) {
                return $n;
            } else {
                // 递归处理
                $res =  $this->fib($n-1) + $this->fib($n-2);
                $this->cache[$n] = $res;
            }
        }
        return $res;
    }

    /**
     * @param Integer $n
     * @return Integer
     */
    // 动态规划
    // 执行用时：0 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了38.18%的用户
    function fib2($n) {
        if($n < 2) {
            return $n;
        }
        $a = 0;
        $b = 1;
        for($i=2; $i<=$n; $i++) {
            // 根据公式计算当前值
            $c = $a + $b;
            // 动态规划
            $a = $b; // 替换下一个a的值
            $b = $c; // 替换下一个b的值
        }
        return $b;
    }


    // 根据题意得 0 <= n <= 30
    // 执行用时：4 ms, 在所有 PHP 提交中击败了95.45%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了37.27%的用户
    function fib3($n)
    {
        // 投机取巧
        $arr = [0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711,28657,46368,75025,121393,196418,317811,514229,832040];
        return $arr[$n];
    }
    


}
// @lc code=end




/*
印象中好像和爬楼梯的题目有点相似

看懂一题，通过两题

[70][爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/description/)

推出来公式就好

斐波那契数的边界条件是 F(0)=0  和 F(1)=1。当 n>1 时，每一项的和都等于前两项的和，因此有如下递推关系：
    f(x) = f(x-1) + f(x-2)
 
    
【递归】
    采用缓存数组进行优化，对重复的值不再进行计算

【动态规划】
    动态规划看似很难，其实就是不断重复利用前面两个变量。 => 滚动数组思想
    f(x) = f(x-1) + f(x-2) =》 c=a+b
    始终都用这三个变量来，不断向前推移，计算出来下一个值。
    
    时间复杂度O(n)。 一共有多少个数，就需要迭代多少次
    空间复杂度O(1)。

=====》 抽离解法思路

【递归】
    缓存结果。递归到底，回溯遇相同的计算过程 不再递归，直接读取缓存数组中的值

【动态规划】
    动态规划问题的一般形式就是求最值 => 核心问题是穷举

    动态规划特点：「重叠子问题」=> 「最优子结构」=> 「状态转移方程」

    

*/