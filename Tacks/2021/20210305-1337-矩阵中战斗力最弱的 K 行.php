<?php
/*
 * @Descripttion: 周四补一下题
 * @Author: tacks321@qq.com
 * @Date: 2021-03-05 09:48:54
 * @LastEditTime: 2021-03-05 14:45:57
 */



/*
 * @lc app=leetcode.cn id=1337 lang=php
 *
 * [1337] 矩阵中战斗力最弱的 K 行
 *
 * https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/description/
 *
 * algorithms
 * Easy (65.32%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    11K
 * Total Submissions: 16.8K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]]\n3'
 *
 * 给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
 * 
 * 请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。
 * 
 * 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
 * 
 * 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：mat = 
 * [[1,1,0,0,0],
 * ⁠[1,1,1,1,0],
 * ⁠[1,0,0,0,0],
 * ⁠[1,1,0,0,0],
 * ⁠[1,1,1,1,1]], 
 * k = 3
 * 输出：[2,0,3]
 * 解释：
 * 每行中的军人数目：
 * 行 0 -> 2 
 * 行 1 -> 4 
 * 行 2 -> 1 
 * 行 3 -> 2 
 * 行 4 -> 5 
 * 从最弱到最强对这些行排序后得到 [2,0,3,1,4]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：mat = 
 * [[1,0,0,0],
 * [1,1,1,1],
 * [1,0,0,0],
 * [1,0,0,0]], 
 * k = 2
 * 输出：[0,2]
 * 解释： 
 * 每行中的军人数目：
 * 行 0 -> 1 
 * 行 1 -> 4 
 * 行 2 -> 1 
 * 行 3 -> 1 
 * 从最弱到最强对这些行排序后得到 [0,2,3,1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == mat.length
 * n == mat[i].length
 * 2 
 * 1 
 * matrix[i][j] 不是 0 就是 1
 * 
 * 
 */

// 【失败代码】
// 关键问题，对数组排序的话，PHP 暂时没办法比较value后在比较key）
class Solution1 {

    /**
     * @param Integer[][] $mat
     * @param Integer $k
     * @return Integer[]
     */
    function kWeakestRows($mat, $k) {
        $data = [];
        // 遍历整个二维数组
        foreach($mat as $row=>$item) {
            $cur = 0;
            foreach($item as $value) {
                if(!$value) {
                    break;
                } else {
                    $cur++;
                }
            }
            $data[$row] = $cur;
        }
        // 按照元素值 升序（关键问题出在这里，对数组排序的话，暂时没办法比较value后在比较key）
        // 排序这里要确定的是，如果战斗力相同，后面的要大
        uasort($data, [$this,'selfSort']);  
        // 然后获取所有的key 
        // array_flip 反转的时候 0竟然抛弃了 ，所以我没有用键值反转，直接获取所有的key
        $tmp = array_keys($data);
        // 取出前k个元素
        return array_slice($tmp,0,$k);
    }

    // 自定义排序
    function selfSort($a, $b) {
        if ($a == $b)  {
            return 0;
        }
        return ($a<$b) ? -1 : 1;
    }
}
 



// 代码2
class Solution2 {

    /**
     * @param Integer[][] $mat
     * @param Integer $k
     * @return Integer[]
     */
    // 执行用时：36 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：17.2 MB, 在所有 PHP 提交中击败了62.50%的用户
    function kWeakestRows($mat, $k) {
        $Forces = [];  // 武力值
        $Keys   = [];  // 对应的位置
        foreach($mat as $row => $item) {
            // 这里使用 array_sum 其实也是有时间复杂度的，可以转化为 二分。
            // 执行用时：44 ms, 在所有 PHP 提交中击败了25.00%的用户
            // 内存消耗：17.4 MB, 在所有 PHP 提交中击败了12.50%的用户
            // $Forces[] = array_sum($item);

            // 递归二分
            $Forces[] = $this->search($item, 0, count($item));
            $Keys  [] = $row;
        }
        // 多个数组排序
        array_multisort($Forces, SORT_ASC, $Keys, SORT_ASC);
        return array_slice($Keys, 0 , $k);
    }

    // 二分 (找到第一个0的位置，也即是士兵的个数)
    function search($arr, $left, $right){
        if($left == $right) {
            return $left;
        }

        $mid = $left + floor(($right - $left) / 2);

        // 来个递归二分
        // 找到第一个0的位置
        return $arr[$mid] === 0 ?
            $this->search($arr, $left, $mid) :
            $this->search($arr, $mid+1, $right);
    }
    
    
}
 


// 代码3 
class Solution3 {

    /**
     * @param Integer[][] $mat
     * @param Integer $k
     * @return Integer[]
     */
    // 执行用时：64 ms, 在所有 PHP 提交中击败了12.50%的用户
    // 内存消耗：16.8 MB, 在所有 PHP 提交中击败了100.00%的用户
    function kWeakestRows($mat, $k) {
        $m = count($mat);
        $n = count($mat[0]);

        $visited = [];
        $result = [];

        $idx = 0; // 表示
        // 纵向遍历二维数组
        for ($i = 0; $i < $n; $i++) { 
            for ($j = 0; $j < $m; $j++) { 
                // （当然如果全部武力值爆表，是不是，这个if永远进不去）
                // 如果还没有访问到当前行，并且是0 平民
                if($visited[$j] == 0 && $mat[$j][$i] == 0) {
                    // 那么把行号放进去
                    $result[$idx] = $j;
                    // 同时标记当前行号，表示访问过
                    $visited[$j] = 1;
                    // 记录满足前k个 最小战斗力的行
                    $idx++;
                    // 如果等于就可以直接退出了
                    // （当然如果前k个，有一行也都是武力值爆表，是不是也进不去）
                    if($idx == $k) {
                        return $result;
                    }
                }
            }
        }
        // === 上面的循环已经把整个 二维数组遍历一遍，并且只要是有平民的行，都会标记成1

        // 用来处理武力值爆表的情况
        // 直接一行一行遍历
        for($i=0; $i<$m; $i++) {
            // 等于0 的就是 武力值爆表的行
            if($visited[$i] == 0 ) {
                $result[$idx] = $i;    
                $idx++;
                if($idx == $k) {
                    return $result;
                }
            }
        }

    }

}
 

/*
【题目】
输入：
    一个二维数组，每一行里有士兵（用 1 表示）和平民（用 0 表示），并且士兵一定是在平民前面。

    由此得出 
        [每行战斗力即为士兵数量]：数据只有 0 和 1，并且一行的士兵数量其实就是这一行值求和的结果
        [每行的有序性]：如果一个位置是平民，那么士兵的数量一定小于这个值。即这一行的数据是有序的。
    
输出：
    要返回战斗力排名前 k 的行的序号

    由此得出
        我们需要按照每行的战斗力进行排序，而战斗力就是士兵的数量
        并且如果不同行的士兵数量一样，那么后面的行战斗力高
    
【方案】（二分计算每行的武力值从而降低时间复杂度）

    对每一行的数据求和，连同序号一起放进新的数组。
    按照要求对该数组进行排序。 (有序数列排序，可以用到二分！！！！  )
    返回前 k 个的需要。

    难道就这样结束了？？？。。。。。

【升华】（纵向遍历，直接可以找到前k个值）

    哈哈哈哈哈哈哈哈哈哈哈哈哈，当然没有

    透过现象看本质？？？？
    到底是要解决什么问题？？？？
    有没有什么优良解法？？？？

    思考
        如果一个位置出现了平民，那么它右边就不会有士兵了，
        也就是本行的战斗力已经被确定了
        也就是其实它在我们上面排序中的位置也就已经确定了
    
        我们一直都是横向看数据
        来试一试纵向看

        纵向的看一下数据，即一列一列的看。
        我们会发现，当我们在某一列遇到某行第一次出现 0 的时候，它其实就是我们目前状态下的最小战斗力。
        而我们最终需要的其实就是前 k 个这样的值

        当然如果战斗力爆表满的话，也就没有0了，所以最后还需要再处理一下这种情况。

    流程
        一列一列的遍历原始数据。
        如果遇到出现了 0，并且是在没有被访问过的行，那么把行号放进结果，并记录这一行已经被访问了。
        处理可能的战斗力全满的情况。
    


【参考题解】
    https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/solution/bao-bao-ye-neng-kan-dong-de-leetcode-ti-jie-3ge-fa/


【总结】
    当然看到这里，其实基本上也看出来我们想要优化的地方

    就是如何避免这种暴力遍历这种思路，利用二分加快到O(logn)。
    当然一个题也有多种解答方案，有时候出奇制胜，做法更简单。

    那么这里实际上是一个应用题类似的，二分的话只是思路的一种。
*/