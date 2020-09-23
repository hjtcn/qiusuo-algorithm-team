/*
 * @lc app=leetcode.cn id=1267 lang=php
 *
 * [1267] 统计参与通信的服务器 https://leetcode-cn.com/problems/count-servers-that-communicate/
  
    这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。

    如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。

    请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。

 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
     // 计数法
     // 执行用时：204 ms, 在所有 PHP 提交中击败了100.00%的用户
     // 内存消耗：21.8 MB, 在所有 PHP 提交中击败了100.00%的用户
    function countServers1($grid) {
        $count_m = []; // 统计行位置上 值=1的 可以通信的服务器
        $count_n = []; // 统计列位置上 值=1的 可以通信的服务器
        foreach($grid as $i=>$item) {
            foreach($item as $j=>$val) {
                if($val == 1) { // 计数
                    $count_m[$i]++;
                    $count_n[$j]++;
                } 
            }
        }
        $res = 0;
        foreach($grid as $i=>$item) {
            foreach($item as $j=>$val) {
                // 判断 只要在i行或者j列，满足超过一台服务器即可进行通信。
                if($val == 1 && ($count_m[$i] > 1 || $count_n[$j] >1) ) {
                   $res++;
                } 
            }
        }
        return $res;
    }


   
}
// @lc code=end

