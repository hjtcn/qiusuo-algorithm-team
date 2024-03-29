/*
 * @lc app=leetcode.cn id=1535 lang=php
 *
 * [1535] 找出数组游戏的赢家
 *
 * https://leetcode-cn.com/problems/find-the-winner-of-an-array-game/description/
 *
 * algorithms
 * Medium (43.59%)
 * Likes:    13
 * Dislikes: 0
 * Total Accepted:    6.6K
 * Total Submissions: 15K
 * Testcase Example:  '[2,1,3,5,4,6,7]\n2'
 *
 * 给你一个由 不同 整数组成的整数数组 arr 和一个整数 k 。
 * 
 * 每回合游戏都在数组的前两个元素（即 arr[0] 和 arr[1] ）之间进行。比较 arr[0] 与 arr[1]
 * 的大小，较大的整数将会取得这一回合的胜利并保留在位置 0 ，较小的整数移至数组的末尾。当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家
 * 。
 * 
 * 返回赢得比赛的整数。
 * 
 * 题目数据 保证 游戏存在赢家。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：arr = [2,1,3,5,4,6,7], k = 2
 * 输出：5
 * 解释：一起看一下本场游戏每回合的情况：
 * 
 * 因此将进行 4 回合比赛，其中 5 是赢家，因为它连胜 2 回合。
 * 
 * 
 * 示例 2：
 * 
 * 输入：arr = [3,2,1], k = 10
 * 输出：3
 * 解释：3 将会在前 10 个回合中连续获胜。
 * 
 * 
 * 示例 3：
 * 
 * 输入：arr = [1,9,8,2,3,7,6,4,5], k = 7
 * 输出：9
 * 
 * 
 * 示例 4：
 * 
 * 输入：arr = [1,11,22,33,44,55,66,77,88,99], k = 1000000000
 * 输出：99
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= arr.length <= 10^5
 * 1 <= arr[i] <= 10^6
 * arr 所含的整数 各不相同 。
 * 1 <= k <= 10^9
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $arr
     * @param Integer $k
     * @return Integer
     */
     // 时间复杂度O(n),只遍历一次
     // 空间复杂度O(1),只需要维护几个变量。
     // 执行用时：248 ms, 在所有 PHP 提交中击败了84.21%的用户
     // 内存消耗：26.8 MB, 在所有 PHP 提交中击败了52.63%的用户
    function getWinner1($arr, $k) {
        // 数组长度至少是2，所以存在$arr[0] $arr[1]
        $prev = max($arr[0], $arr[1]); // 表示上一次回合中胜利的整数。
        // 第一回合
        if($k == 1) {
            return $prev; // 如果只用比1回合，直接返回。
        }
        // 之后的回合
        $consecutive = 1;// 胜利回合数
        $maxNum = $prev; // 数组的最大值 
        $size   = count($arr);
        // 遍历整个数组，从第二个回合开始
        for($i=2; $i<$size; $i++) {
            $curr = $arr[$i]; // 当前值
            if($prev > $curr) {
                // 如果上一回合的胜利的数，还是比当前数大，就直接增加胜利的回合数
                $consecutive++;
                if($consecutive == $k) {
                    // 如果正好连赢k个回合
                    return $prev; // 则当前就是最大值
                }
            }else{
                // 如果上一回合的胜利的数，小于当前的数，那么就重置prev,并且连赢回合数也重置为1
                $prev = $curr;
                $consecutive = 1;
            }
            $maxNum = max($maxNum, $curr); // 记录当前数组最大值
        }
        // 如果遍历结束，还是没有连赢k场，那么直接就是数组最大值，为最终赢家，因为在之后的回合数中，他总是会赢。
        return $maxNum;
    }

    // 始终维护arr[0]在部分数组中是最大的
    // 执行用时：256 ms, 在所有 PHP 提交中击败了68.42%的用户
    // 内存消耗：26.7 MB, 在所有 PHP 提交中击败了78.95%的用户
    function getWinner2($arr, $k) {
        $size = count($arr);
        $consecutive = 0 ; // 当前连续胜利的回合数
        for($i = 1; $i<$size; $i++) {
            if($arr[0] > $arr[$i]) {
                $consecutive++;
            }else{
                $consecutive = 1;   // 重置
                $arr[0] = $arr[$i]; // 保持首位元素最大
            }
            // 满足连续胜利k个回合
            if($consecutive == $k) {
                return $arr[0];
            }
        }
        return $arr[0];
    }


}
// @lc code=end


// @tacks think=start
/*
题意 meaning 
 
    此题，读完之后，并不是很清楚要如何去做，首先跟着示例，以及题意走一遍流程，摸清楚这个比赛是如何进行，并且如何胜利。

    基本上，就是每次拿数组第一个第二个值去比较，把大的留在第一个位置，小的直接扔到数组后面。
        如果arr[0] >   arr[1] 那么就是arr[0]就是本回合的胜利整数，然后连续胜利回合数+1
        如果arr[0] <=  arr[1] 那么就是arr[1]是本回合胜利整数，并且将arr[0]扔到后面，arr[1]放在前面，同时胜利数回合数重置为1
    等下次再遍历
        prev = 上一局胜利的数，也就是第一个位置的
        curr = 表示数组当前值，也就是第二个位置的
    等最后遍历结束，有可能还没满足到k的连续胜利回合数  
        那么就是本数组中最大的值就是，胜利整数。
    
    本题返回的是胜利的整数。（因为题意 不同整数组成的整数数组 arr，保证游戏存在赢家）

关键 key 
    “胜利整数”，“数组中最大的整数”，“连续胜利回合数”，“上一局胜利的整数”

想法 idea
【1】模拟法
    按题意思路
        顺着记录相关变量，就可以实现，不需要移动数组元素。
 
        只需要遍历一次数组，时间复杂度是O(n)。

【2】部分数组最大值 超简洁

    实际上，就是找部分数组中最大的值。
        不需要去管较小元素移动到数组末尾这个动作，直接跳过去就行。
        我们只需要把当前部分数组的最大值始终放在首元素即可。
    
    遍历数组
        arr[0] > arr[i]
            胜利回合数++
            i++,继续向后比较就行，不用管移动不移动。
        arr[0] <= arr[i]
            胜利回合数重置1
            重新赋值当前位置部分数组中最大的元素值，放在首位，也即是arr[0] = arr[i]
            i++，向后继续比较
        如果胜利回合数==k 
            就可以返回当前首元素的值arr[0]
    如果遍历结束还没满足k
        那么一定就是arr[0],因为始终让首位元素是最大的（部分数组中）。

*/
// @tacks think=end

