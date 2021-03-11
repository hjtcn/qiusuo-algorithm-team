<?php
/*
 * @Descripttion: 明天休个假 今天先做了 哈哈哈哈哈
 * @Author: tacks321@qq.com
 * @Date: 2021-03-11 14:15:14
 * @LastEditTime: 2021-03-11 15:23:42
 */


/*
 * @lc app=leetcode.cn id=74 lang=php
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (39.71%)
 * Likes:    329
 * Dislikes: 0
 * Total Accepted:    85.4K
 * Total Submissions: 209.5K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 * 
 * 
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 
 * -10^4 
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer[][] $matrix
     * @param Integer $target
     * @return Boolean
     */
    // 二分思路
    // 执行用时：12 ms, 在所有 PHP 提交中击败了89.29%的用户
    // 内存消耗：16 MB, 在所有 PHP 提交中击败了25.00%的用户
    function searchMatrix($matrix, $target) {
        $m = count($matrix);
        $n = count($matrix[0]);
        if($m == 0) {
            return false;
        }
        // 
        $left = 0;
        $right= $m*$n - 1;
        // 二分框架
        while($left <= $right) {
            $mIdx = ($left + $right) >> 1;
            // 二维数组 的坐标 转化为一维数组的 mIdx
            $mItem= $matrix[floor($mIdx / $n)][$mIdx % $n];
            if($mItem == $target) {
                return true;
            } elseif($mItem > $target){
                // 缩小右边界
                $right = $mIdx - 1;
            } else {
                // 缩小左边界
                $left  = $mIdx  + 1;
            }
        }
        return false;
    }

}


class Solution2 {

    /**
     * @param Integer[][] $matrix
     * @param Integer $target
     * @return Boolean
     */
    // 缩小领域法
    // 执行用时：16 ms, 在所有 PHP 提交中击败了39.29%的用户
    // 内存消耗：15.9 MB, 在所有 PHP 提交中击败了75.00%的用户
    function searchMatrix($matrix, $target) {
        $row = count($matrix);
        $col = count($matrix[0]);
        // 从左下角开始
        $left= $row - 1;
        $right= 0;

        // 边界范围
        while($left >=0 && $right<$col) {
            if($matrix[$left][$right] > $target) {
                // 向上移动
                $left--;
            }elseif($matrix[$left][$right] < $target) {
                // 向右移动
                $right++;
            }else{
                // 找到目标值
                return true;
            }
        }
        return false;
       
    }

}

// @lc code=end

/*
【题目】
    二维数组
        每行中的整数从左到右按升序排列。
        每行的第一个整数大于前一行的最后一个整数。
    => 这两句话，告诉我们，是一个递增的数列

    那么如何将二维数组转化成 一个递增的数组？

    =》 其实可以利用下标来做一些转化
        一个m*n 的二维数组
        那么一个位置的坐标 （row,col） => (idx/n , idx%n)

        从而将二维数组转化为 一维数组 idx表示数组下标
【题解】
    ①二分
        left = 0
        right= m*n - 1
        while left < right 
            读取一维数组的mid = (left+right)>>1
            那么对应的序号就是 row = mid/n  ; col= mid%n 
    
    时间复杂度O(log(n*m))

    ②缩小领域法
        因为每一行递增，每一列递增
            =>
            我们可以从右上角往左下角找 (或者从左下角往右上角找)。
            每次比较可以排除一行或者一列，时间复杂度为O(m+n)

*/