/*
 * @lc app=leetcode.cn id=867 lang=php
 *
 * [867] 转置矩阵
 *
 * https://leetcode-cn.com/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (67.80%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    25K
 * Total Submissions: 36.9K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个矩阵 A， 返回 A 的转置矩阵。
 * 
 * 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[1,4,7],[2,5,8],[3,6,9]]
 * 
 * 
 * 示例 2：
 * 
 * 输入：[[1,2,3],[4,5,6]]
 * 输出：[[1,4],[2,5],[3,6]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= A.length <= 1000
 * 1 <= A[0].length <= 1000
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $A
     * @return Integer[][]
     */
    function transpose1($A) {
        $row = count($A);
        $col = count($A[0]);
        $res = [];
        // PHP for循环
        for($i=0; $i<$row; $i++) {
            for($j=0; $j<$col; $j++) {
                $res[$j][$i] = $A[$i][$j]; // 数组转置赋值
            }
        }
        return $res;
    }

    /**
     * @param Integer[][] $A
     * @return Integer[][]
     */
    function transpose($A) {
       $res = [];
       // PHP 最常用的foreach遍历数组
       foreach($A as $i=>$item) {
           foreach($item as $j=>$val) {
               $res[$j][$i] = $val;
           }
       }
       return $res;
    }
}
// @lc code=end

/*

思路：
    1、采用双重循环，拿到二维数组也就是矩阵的某个元素。
    2、然后重新赋值给一个新的数组，只是在赋值的时候，复制到斜对角的位置。
    3、那么斜对角其实就是 i、j对换位置。

时间复杂度：
    跟矩阵的m、n有关，如m行n列的矩阵，那么时间复杂度就是 O(m*n)

解法：
【1】采用for循环遍历二维数组
绝大部分的编程语言都会支持for循环，当然PHP也是支持的，我这里只是做个示例，之后我基本都会使用foreach。

【2】采用foreach遍历二维数组
相关语法了解：https://www.php.net/manual/zh/control-structures.foreach.php

例如：
foreach (array_expression as $key => $value){
    statement
}
格式遍历给定的 array_expression 数组。
每次循环中，当前单元的值被赋给 $value 并且数组内部的指针向前移一步（因此下一次循环中将会得到下一个单元）。
另外遍历的时候当前单元的键名也会在每次循环中被赋给变量 $key。

PHP的数组很强大！


*/

