<?php
/*
 * @lc app=leetcode.cn id=840 lang=php
 *
 * [840] 矩阵中的幻方
 *
 * https://leetcode-cn.com/problems/magic-squares-in-grid/description/
 *
 * algorithms
 * Medium (34.90%)
 * Likes:    38
 * Dislikes: 0
 * Total Accepted:    7.1K
 * Total Submissions: 20.5K
 * Testcase Example:  '[[4,3,8,4],[9,5,1,9],[2,7,6,2]]'
 *
 * 3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。
 * 
 * 给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入: [[4,3,8,4],
 * ⁠     [9,5,1,9],
 * ⁠     [2,7,6,2]]
 * 输出: 1
 * 解释: 
 * 下面的子矩阵是一个 3 x 3 的幻方：
 * 438
 * 951
 * 276
 * 
 * 而这一个不是：
 * 384
 * 519
 * 762
 * 
 * 总的来说，在本示例所给定的矩阵中只有一个 3 x 3 的幻方子矩阵。
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= grid.length <= 10
 * 1 <= grid[0].length <= 10
 * 0 <= grid[i][j] <= 15
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    // 执行用时：20 ms, 在所有 PHP 提交中击败了66.67%的用户
    // 内存消耗：14.7 MB, 在所有 PHP 提交中击败了33.33%的用户
    function numMagicSquaresInside($grid) {
        $row = count($grid);   // 矩阵行
        $col = count($grid[0]);// 矩阵列 
        $res = 0;
        $skip= []; // 标记要跳过的元素1
        foreach($grid as $i=>$R){
            foreach($R as $j=>$C) {
                if($C > 9) {
                    $skip[$i-1][$j-1] = 1;
                    $skip[$i-1][$j] = 1;
                    $skip[$i-1][$j+1] = 1;
                    $skip[$i][$j-1] = 1;
                    $skip[$i][$j] = 1;
                    $skip[$i][$j+1] = 1;
                    $skip[$i+1][$j-1] = 1;
                    $skip[$i+1][$j] = 1;
                    $skip[$i+1][$j+1] = 1;
                }
            }
        }
        // 遍历寻找3*3矩阵
        for($r=0; $r<$row-2; $r++) {
            for($c=0; $c<$col-2; $c++) {
                // 如果被标记直接跳过
                if($skip[$r+1][$c+1]) continue;
                // 如果中心不是5 直接跳过
                if($grid[$r+1][$c+1] != 5) continue;
                if($this->magic($grid, $r+1, $c+1)) {
                    $res++;
                }
            }
        }
        return $res;
    }

    // 判断是否符合要求 8条线和都是15
    function magic($grid, $i, $j)
    {
        // 判断9个数是否一样
        $count = []; // 计数
        for ($m=-1; $m <2 ; $m++) { 
            for ($n=-1; $n <2 ; $n++) { 
                $count[$grid[$i+$m][$j+$n]]++;
            }
        }
        foreach($count as $item) {
            // 每个数字都是1个
            if($item != 1) return false;
        }
        // 三行
        if($grid[$i-1][$j-1]+$grid[$i-1][$j]+$grid[$i-1][$j+1] != 15) return false;
        if($grid[$i][$j-1]+$grid[$i][$j]+$grid[$i][$j+1] != 15) return false;
        if($grid[$i+1][$j-1]+$grid[$i+1][$j]+$grid[$i+1][$j+1] != 15) return false;
        // 三列
        if($grid[$i-1][$j-1]+$grid[$i][$j-1]+$grid[$i+1][$j-1] != 15) return false;
        if($grid[$i-1][$j]+$grid[$i][$j]+$grid[$i+1][$j] != 15) return false;
        if($grid[$i-1][$j+1]+$grid[$i][$j+1]+$grid[$i+1][$j+1] != 15) return false;
        // 两个对角线
        if($grid[$i-1][$j-1]+$grid[$i][$j]+$grid[$i+1][$j+1] != 15) return false;
        if($grid[$i-1][$j+1]+$grid[$i][$j]+$grid[$i+1][$j-1] != 15) return false;
        return true;
    }
    
}
// @lc code=end

// @tacks think=start
/*
题意 meaning 
  
    需要找到3*3的矩阵，而且这个矩阵满足每列以及两条对角线上的各数之和都相等。另外就是这个9个值都不一样。

    可以得出性质： 
        所有元素之和    => 45（因为每个位置元素都不一样，也就是1,2,3,4,5,6,7,8,9的排列）   =>  0 <= grid[i][j] <= 15 => 如果当前元素不是0~9，可以直接跳过
        行列对角线都一样=> 15（八条线，这样刚好满足15*3=45）
        穿过中心的四条线=> 15*4=60 其中中心元素被重复计算三次，所以是（60-45）/ = 5     =>   所以中心只能是5
    

关键 key 

    “确定中心元素”     
    “九个数字都不一样并且范围保证0~9”
    “八条线的和都是15”

想法 idea

计数暴力法

- skip记录不能成为中心元素
- count记录所有元素出现的次数
- magic() 验证中心元素周围的元素每个和是不是15

复杂度
   - 时间复杂度 O(n*m)
   - 空间复杂度 O(n*m) 用到skip来跳过不符合的中心元素
   
     
  
    

*/
// @tacks think=end