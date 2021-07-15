<?php
/*
 * @Descripttion: 周四，好久没刷题了
 * @Author: tacks321@qq.com
 * @Date: 2021-07-15 10:45:44
 * @LastEditTime: 2021-07-15 17:41:02
 */


/*
 * @lc app=leetcode.cn id=279 lang=php
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (60.70%)
 * Likes:    1011
 * Dislikes: 0
 * Total Accepted:    175.8K
 * Total Submissions: 280.8K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 * 
 * 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
 * 
 * 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11
 * 不是。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 12
 * 输出：3 
 * 解释：12 = 4 + 4 + 4
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param int $n
     * @return int
     */
    // 暴力递归（超时）
    function numSquares($n) {
        if($n == 0) {
            return 0;
        }
        // 初始化
        $count = $n;
        // 依次减去一个完全平方数
        for ($i=1; $i*$i <= $n; $i++) { 
            // 选择最小的
            $count = min($count ,$this->numSquares($n - $i*$i) + 1);
        }
        return $count;
    }
}

class Solution2 {

    /**
     * @param int $n
     * @return int
     */
    // 暴力递归+记忆化数组 （PHP貌似还是超时）
    function numSquares($n) {
        return $this->_helper($n, []);
    }

    private function _helper($num, $memoryMap) {
        // Map Get 增加记忆化数组，防止重复计算
        if(isset($memoryMap[$num])) {
            return $memoryMap[$num];
        }
        if($num == 0) {
            return 0;
        }
        // 初始化
        $count = $num;
        // 依次减去一个完全平方数
        for ($i=1; $i*$i <= $num; $i++) { 
            // 选择最小的
            $count = min($count ,$this->numSquares($num - $i*$i) + 1);
        }
        // Map Put 计算过的就保存起来
        $memoryMap[$num] = $count;
        return $count;
    }
}


class Solution3 {

    /**
     * @param int $n
     * @return int
     */
    // 动规
    // 执行用时：824 ms, 在所有 PHP 提交中击败了77.03%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了77.03%的用户
    function numSquares($n) {
        if($n == 0) {
            return 0;
        }
        // 初始化
        $dp = array_fill(0, $n+1, 0);
        $dp[0] = 0;

        // 外循环: 完全平方数 物品
        for ($i=1; $i <= $n ; $i++) { 
            // 注意这里初始化的问题！！！！！！！！！！！！要给一个最大值，或者本身
            $dp[$i] = $i; // 当前i的完全平方数的个数最多是i个，也就是 i个1

            // 内循环：背包正序
            for ($j=1; $j * $j <= $i ; $j++) { 
                // $j * $j 当前的完全平方数
                // 选择最小的个数，最不济选择 i个1
                $dp[$i] = min($dp[$i], $dp[$i - $j * $j] + 1);
            }
        }
        return $dp[$n];
    }
}


class Solution4 {

    /**
     * @param int $n
     * @return int
     */
    // BFS 广度优先遍历 + visited剪枝
    // 执行用时：4476 ms, 在所有 PHP 提交中击败了5.40%的用户
    // 内存消耗：15.8 MB, 在所有 PHP 提交中击败了14.86%的用户
    function numSquares($n) {
        $queue   = [];// 队列
        $visited = [];// 访问过的
        $level   = 0; // 顶点

        // 入队
        array_push($queue, $n);
        while(!empty($queue)) {
            $size = count($queue);
            $level++; // 层级

            for ($i=0; $i < $size ; $i++) { 
                $cur  = array_shift($queue); // 弹出队首
                ////////////////////////////【核心判断】//////////////////////////////////
                // 依次减去 完全平方数，生成下一层
                for ($j=1; $j * $j <= $cur; $j++) { 
                    $next = $cur - $j*$j;
                    // 遍历到值为空的时候
                    if($next == 0) {
                        return $level;
                    }
                    // 如果我们遍历的之前已经计算过，这里进行 剪枝，断掉后面的路

                    // 如果我们没有遍历过这个，需要加入队列，在下一层进行遍历，同时加入一个记录访问过的节点
                    if(!in_array($next, $visited)) {
                        array_push($queue, $next);
                        array_push($visited, $next);
                    }
                }
                ////////////////////////////////////////////////////////////////////////

            }
            
        }
        return -1;
    }
}



class Solution5 {

    /**
     * @param int $n
     * @return int
     */
    // 四平方和定理
    // 4 ms, 在所有 PHP 提交中击败了98.65%的用户
    // 15.1 MB, 在所有 PHP 提交中击败了97.30%的用户
    function numSquares($n) {
        // ① 自身完全平方数
        if($this->_isPerfectSquareNum($n)) {
            return 1;
        }
        // ② 四个平方数之和
        if($this->_isCheckAnswer4($n)) {
            return 4;
        }

        // ③ 两个完全平方数之和
        for ($a=1; $a*$a <= $n ; $a++) { 
            $b = $n - $a*$a;
            if($this->_isPerfectSquareNum($b)) {
                return 2;
            }
        }
        // ④ 最后排除法 三个平方数之和
        return 3;
    }


    /**
     * 判断是否是 完全平方数
     *
    * @param int $x
     * @return bool
     */
    private function _isPerfectSquareNum($x) {
        $y = (int)sqrt($x);
        return $y*$y == $x;
    }

    /**
     * 判断是否 四平方和
     *
     * @param int $x
     * @return bool
     */
    private function _isCheckAnswer4($x) {
        // n = (4)^k * (8m+7) 
        while($x % 4 == 0) {
            $x /= 4;
        }
        return $x % 8 == 7;
    }
}



// @lc code=end


$Obj = new Solution5();
var_dump($Obj->numSquares(46));


/*
【题目】
    完全平方数：
        完全平方指用一个整数乘以自己例如1*1，2*2，3*3等，依此类推。若一个数能表示成某个整数的平方的形式，则称这个数为完全平方数

        如果一个正整数 a 是某一个整数 b 的平方，那么这个正整数 a 叫做完全平方数。零也可称为完全平方数
        
        完全平方数是非负数

    题意：
        给一个整数n，需要找到  完全平方数的和 = n ，并且完全平方数的个数最少。
        
【参考】
    https://leetcode-cn.com/problems/perfect-squares/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--51/
【解析】

完全背包
    完全背包和01背包问题唯一不同的地方就是，每种物品有无限件，可以重复选择。

    完全背包：暴力递归+记忆化数组+动规+BFS+四平方和定理

    1、暴力递归
            模拟题目
                假设 n = 12    
                步骤就是，先将n-完全平方数，然后剩余的值，再分解成完全平方数和所需要的最小的个数
                
                在 12 范围内的完全平方数，可以选择 [1,4,9]

                    ① 方案N1
                        12-1 = 11 
                        然后剩下的值 再分解成完全平方数和所需要的最小的个数 
                    ② 方案N2
                        12-4 = 8
                        然后剩下的值 再分解成完全平方数和所需要的最小的个数
                    ③ 方案N3
                        12-9 = 3
                        然后剩下的值 再分解成完全平方数和所需要的最小的个数

                    于是我们需要从
                        (N1+1) (N2+1) (N3+1) 三种方案中选择最小个数的
                
                下一步
                    就是求 11、8、3 如何分解
                    
                直至变成0
 
            时间复杂度
                
    2、暴力递归+记忆化数组memory 
            对于递归中重复出现的计算，可以增加记忆化数组


    3、动规，完全背包的应用
            递归相当于，先进递归栈，入栈入栈。   后出递归栈，出栈出栈
        动态规划的转移方程对应递归的过程，
        动态规划的初始条件对应递归的入口。

        递归 => 改写 => 动规
        
            ① 初始化dp
                dp数组长度 n+1
                dp[i] 表示当前i数字需要的完全平方数的个数
            ② 动态转移方程
                dp[i] = min(dp[i], dp[i-j*j] + 1) 
                j*j 是完全平方数 （j*j < i）
            ③ 时间复杂度
                O(n * √n) 。由于内层遍历所有平方数小于i的数j，遍历区间[1,√n+1]
            
        完全背包的解释
            dp[i] 在 m 件商品里面选择 ，使得容量为 n 的背包可以装满的最小价值量
            
            ① 初始状态
                背包容量为0 只能装 0 件物品，价值为 0
                但是这里注意，要初始化为 dp[i] 设置为 Integer.MAX_VALUE 或者 n（背包最多装 n 个物品）。
            ② 动态转移方程
                物品装进去就是   dp[i-num] + 1
                dp[i] = min($dp[i-num] + 1, $dp[$i])
            
            模板
                for(外循环：遍历物品)
                    for(内循环：背包的剩余容量)
                        动规


    4、BFS
        一层一层的进行广度计算，第一层减去一个完全平方数得到第二层，自顶向下，直到某层出现了0，这样就得到了最小个数
        https://pic.leetcode-cn.com/641036134c746b8bba3be26299a9cbd8493950dba92edcf50408031903c4f37c.jpg

        层次遍历，然后这样自顶向下，可以比DFS更快的得到最少的层数。
        递归处理，需要一个一个值向下减到0，然后再回溯判断，比较慢。

        同时BFS中也会由很多节点会重复，加入 visited数组来进行去重，加快遍历速度。
        
    5、数学思想 四平方和定理
        四平方和定理证明了任意一个正整数都可以被表示为至多四个正整数的平方和
            少于四个平方数的，像 12 这种，可以补一个 0 也可以看成四个平方数，12 = 4 + 4 + 4 + 0。
            知道了这个定理，对于题目要找的解，其实只可能是 1, 2, 3, 4 其中某个数。
        
        当前仅当 n≠ (4)^k * (8m+7) ；  至多可以是三个正整数的和
            如果是1个正整数的话，直接判断本身是否完全平方数 √n * √n = n
            如果是2个正整数的话，则有n= a² + b²，枚举 a的值[1, √n], 判断 n-a² 是否是完全平方数
            如果是3个正整数的话，排除法
        当前仅当 n= (4)^k * (8m+7) ；  只能是4个正整数的和


总结
    1、完全背包问题，物品可以重复拿。
    2、可以动规来进行处理。
 
        

*/