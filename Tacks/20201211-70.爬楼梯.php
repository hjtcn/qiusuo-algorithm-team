<?php
/*
 * @Descripttion: 补一下周四的题
 * @Author: tacks321@qq.com
 * @Date: 2020-12-11 09:49:53
 * @LastEditTime: 2020-12-11 10:24:47
 */

/*
 * @lc app=leetcode.cn id=70 lang=php
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (50.76%)
 * Likes:    1371
 * Dislikes: 0
 * Total Accepted:    330.9K
 * Total Submissions: 648.9K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 注意：给定 n 是一个正整数。
 * 
 * 示例 1：
 * 
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 * 
 * 示例 2：
 * 
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了53.85%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了18.89%的用户
    function climbStairs1($n) {
        if($n <= 2) {
            return $n;
        }
        $a = 1;
        $b = 2;
        for($i=3; $i<=$n; $i++) {
            $c = $a+$b;// 当前值  f(x) = f(x-1) + f(x-2)
            $a = $b;// 替换下一个a值
            $b = $c;// 替换下一个b值
        }
        return $b;
    }


    static $cache = []; // 缓存数组
    //执行用时：12 ms, 在所有 PHP 提交中击败了53.85%的用户
    //内存消耗：15.5 MB, 在所有 PHP 提交中击败了5.21%的用户
    function climbStairs($n) {
        if(isset($this->cache[$n])) {
            return $this->cache[$n];
        } else{
            if($n <= 2 ) {
                return $n;
            } else {
                // 递归处理
                $res = $this->climbStairs($n-1) + $this->climbStairs($n-2) ;
                $this->cache[$n] = $res;
            }
        }
        return $res;

    }
}
// @lc code=end


$obj = new Solution();
var_dump($obj->climbStairs(50));
/*


这个感觉跟斐波那契数列有点类似，推出来公式就好了
    f(x) = f(x-1) + f(x-2)

【动态规划】
    时间复杂度O(n)

【递归】
    采用缓存数组进行优化，对重复的值不再进行计算


然后看到题解中有 矩阵快速幂 
感叹 真谛牛逼
数学真牛逼 但是看起来好头疼



*/