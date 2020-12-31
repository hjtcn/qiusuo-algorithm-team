/*
 * @lc app=leetcode.cn id=1424 lang=php
 *
 * [1424] 对角线遍历 II
 *
 * https://leetcode-cn.com/problems/diagonal-traverse-ii/description/
 *
 * algorithms
 * Medium (37.32%)
 * Likes:    18
 * Dislikes: 0
 * Total Accepted:    3.9K
 * Total Submissions: 10.5K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个列表 nums ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，按顺序返回 nums 中对角线上的整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：nums = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,4,2,7,5,3,8,6,9]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
 * 输出：[1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
 * 输出：[1,4,2,5,3,8,6,9,7,10,11]
 * 
 * 
 * 示例 4：
 * 
 * 输入：nums = [[1,2,3,4,5,6]]
 * 输出：[1,2,3,4,5,6]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i].length <= 10^5
 * 1 <= nums[i][j] <= 10^9
 * nums 中最多有 10^5 个数字。
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[][] $nums
     * @return Integer[]
     */
    // 广度优先遍历
    function findDiagonalOrder($nums) {
        $node= [[0,0]]; // 根节点
        $res = [];    // 结果数组
        $skip= [];    // 已经访问过，需要跳过的节点
        while(!empty($node)) {
            $item = array_shift($node); // 弹出数组开头元素
            // 坐标
            $x = $item[0];
            $y = $item[1];
            $skip[$x][$y] = 1; // 已经访问过，需要跳过的节点
            // 判断下面节点没有被访问，是否存在
            if(isset($nums[$x+1][$y]) && !isset($skip[$x+1][$y])) {
                $skip[$x+1][$y] = 1; //标记
                $node[] = [$x+1, $y]; // push新的节点
            }
            // 判断下面节点没有被访问，是否存在
            if(isset($nums[$x][$y+1]) && !isset($skip[$x][$y+1])) {
                $skip[$x][$y+1] = 1; //标记
                $node[] = [$x, $y+1]; // push新的节点
            }
            $res[] = $nums[$x][$y]; // 追加元素
        }
        return $res;
    }
}
// @lc code=end


// @tacks think=start
/*
题意 meaning 

    对角线遍历，实际上你看左上角的元素，可以理解为根节点

    可以转化为广度优先遍历，进行访问元素，同时判断元素是否存在

关键 key 

    “广度优先遍历”
    “判断元素是否存在”

想法 idea

    按照斜角就是广度优先的遍历过程，在队列中按下和右的顺序把点加入队列即可
 
    声明一个skip数组用于标记是否需要跳过



*/
// @tacks think=end