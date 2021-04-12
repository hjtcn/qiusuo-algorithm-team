<?php
/*
 * @Descripttion: 上周三的题，这次补一下(2021/03/31)
 * @Author: tacks321@qq.com
 * @Date: 2021-04-09 17:03:47
 * @LastEditTime: 2021-04-09 17:43:41
 */



/*
 * @lc app=leetcode.cn id=1025 lang=php
 *
 * [1025] 除数博弈
 *
 * https://leetcode-cn.com/problems/divisor-game/description/
 *
 * algorithms
 * Easy (71.48%)
 * Likes:    290
 * Dislikes: 0
 * Total Accepted:    68.4K
 * Total Submissions: 96.1K
 * Testcase Example:  '2'
 *
 * 爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
 * 
 * 最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
 * 
 * 
 * 选出任一 x，满足 0 < x < N 且 N % x == 0 。
 * 用 N - x 替换黑板上的数字 N 。
 * 
 * 
 * 如果玩家无法执行这些操作，就会输掉游戏。
 * 
 * 只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 False。假设两个玩家都以最佳状态参与游戏。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：2
 * 输出：true
 * 解释：爱丽丝选择 1，鲍勃无法进行操作。
 * 
 * 
 * 示例 2：
 * 
 * 输入：3
 * 输出：false
 * 解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= N <= 1000
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer $n
     * @return Boolean
     */
    // 数学思维
    // 执行用时：8 ms, 在所有 PHP 提交中击败了33.33%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了11.11%的用户
    function divisorGame($n) {
        return $n % 2 == 0;
    }
}


class Solution2 {

    /**
     * @param Integer $n
     * @return Boolean
     */
    // 动态规划
    public function divisorGame($n) {
        // 初始化
        $dp = array_fill(0, $n + 1, false);
        // dp为输入为N时的输赢情况
        // dp[1] 必败
        // dp[2] 必赢
        $dp[1] = false;
        $dp[2] = true;
        // i 为 N
        for ($i=3; $i <= $n ; $i++) { 
            // j为x  0<x<n
            for ($j=1; $j < $i ; $j++) { 
                // Bob N = N -x 
                // 可以 根据前面的推断结果得知 ans[N-x]的输赢情况
                if($i % $j === 0 && $dp[$i-$j] === false) {
                    $dp[$i] = true;
                }
            }
        }
        return $dp[$n];

    }
}


$obj = new Solution2();
var_dump($obj->divisorGame(3));



// @lc code=end

/*
【题目】
    博弈，两个人互相操作

        选出任一 x，满足 0 < x < N 且 N % x == 0 。
        用 N - x 替换黑板上的数字 N 。
        直到无法操作的时候，游戏结束


【解析】

    【数学思维】
        N 为奇数的时候 Alice（先手）必败，N 为偶数的时候 Alice 必胜。

参见网友@YangJiaqi的解释

    因为至少可以取1 ,所以最后决胜负的时候 就是谁取到1 就输了
    拿到奇数时 : 只有一种选择 就是 减去 奇数 , 
    给出偶数 拿到偶数的时候: 
        有 2 中选择 
            1. 减去奇数返回奇数; 
            2. 减去偶数 返回偶数 可见 拿到偶数的人 有能力选择保持另一个人 奇偶性 永远是 奇数
    由于拿到1的人就输 所以 偶数能一直保持 另一个人是奇数 , 且值一定不停的在减少趋向于 1 所以 开始手里是偶数的人能赢



    【动态规划】


*/