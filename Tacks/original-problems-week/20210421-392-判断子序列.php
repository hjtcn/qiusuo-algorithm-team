<?php
/*
 * @Descripttion: 周三的题
 * @Author: tacks321@qq.com
 * @Date: 2021-04-21 10:52:27
 * @LastEditTime: 2021-04-21 18:15:53
 */



/*
 * @lc app=leetcode.cn id=392 lang=php
 *
 * [392] 判断子序列
 *
 * https://leetcode-cn.com/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (50.75%)
 * Likes:    427
 * Dislikes: 0
 * Total Accepted:    114.9K
 * Total Submissions: 225.7K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
 * 
 * 
 * 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
 * 
 * 进阶：
 * 
 * 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T
 * 的子序列。在这种情况下，你会怎样改变代码？
 * 
 * 致谢：
 * 
 * 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abc", t = "ahbgdc"
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "axc", t = "ahbgdc"
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * 0 
 * 两个字符串都只由小写字符组成。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    // 1、双指针版
    // 执行用时：8 ms, 在所有 PHP 提交中击败了75.86%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了62.07%的用户
    function isSubsequence($s, $t) {
        $lenS = strlen($s);
        $lenT = strlen($t);
        $i = 0;
        $j = 0;
        while($i<$lenS && $j<$lenT) {
            if($s[$i] == $t[$j]) {
                $i++;
                $j++;
            }else{
                $j++;
            }
        }
        return $i == $lenS; // 是否遍历到s尾部
    }
}

class Solution2 {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    // 2、动态规划版
    // 执行用时：52 ms, 在所有 PHP 提交中击败了13.79%的用户
    // 内存消耗：28.5 MB, 在所有 PHP 提交中击败了10.34%的用户
    function isSubsequence($s, $t) {
        $lenS = strlen($s);
        $lenT = strlen($t);
        // dp数组
        // dp[i][j] 表示字符串t,从i位置开始第一次出现字符j的位置
        $dp = [];
        // 初始化边界条件 (因为是小写字母 26个)
        for ($i=0; $i <26; $i++) { 
            $dp[$lenT][$i] = $lenT;
        }
        // 倒序构建T
        for ($i=$lenT-1; $i >=0 ; $i--) { 
            for ($j=0; $j < 26 ; $j++) { 
                // 如果当前字符整好出现，那么就是本身位置
                if($t[$i] == chr($j + ord('a')))                 {
                    $dp[$i][$j] = $i;
                }else{
                    // 否则使用下一行
                    $dp[$i][$j] = $dp[$i+1][$j];
                }
            }            
        }

        // 当前位置
        $add = 0;
        // 进行匹配
        for ($i=0; $i <$lenS ; $i++) { 
            // dp中没有 s[i]; 通过dp可以O(1)的快速找到是否有对应的字符
            if ($dp[$add][ord($s[$i]) - ord('a')]  == $lenT) {
                return false;
            }
            // 如果有 就直接跳转到第一次位置之后的
            $add = $dp[$add][ord($s[$i]) - ord('a')] + 1;
        }

        return true;
    }
}


class Solution3 {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    // 3、递归版
    // 执行用时：8 ms, 在所有 PHP 提交中击败了75.86%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了24.14%的用户
    function isSubsequence($s, $t) {
        // 边界条件
        if($s == "") {
            return true;
        }
        $lenS = strlen($s);
        $lenT = strlen($t);
        if($lenS > $lenT) {
            return false;
        }

        $cur = $s[0];
        $idx = strpos($t, $cur);
        // 找不到s的字符
        if($idx === false) {
            return false;
        }
        // 如果找到的字符是t的最后一个字符，而且s不是单字符，直接返回false
        if($idx == $lenT-1 && $lenS >1) {
            return false;
        }

        // 每一次对比一个字符后， s 截取掉第一个字符
        // t需要保留从匹配的字符后面一个位置
        return $this->isSubsequence(substr($s, 1), substr($t, $idx+1));
    }
}


class Solution4 {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    // 4、栈版
    // 执行用时：8 ms, 在所有 PHP 提交中击败了75.86%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了72.41%的用户
    function isSubsequence($s, $t) {
        // 边界条件
        if($s == "") {
            return true;
        }
        $lenS = strlen($s);
        $lenT = strlen($t);
        if($lenS > $lenT) {
            return false;
        }
        if($s == $t) {
            return true;
        }

     

        $Stack = str_split($s); // 把s转化成 字符串栈
        for ($i=0; $i < $lenT; $i++) { 
            if($t[$i] == $Stack[0]) {
                array_shift($Stack); // 如果匹配弹出栈顶元素
            }
            // 如果栈为空，那么符合匹配
            if(count($Stack) === 0) {
                return true;
            }
        }
        // 外面再来一个
        if(count($Stack) === 0) {
            return true;
        }
        return false;
    }
}
// @lc code=end

/*
【题目】
    给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

    例如：s = "abc", t = "ahbgdc" ，那么s就是t的子序列
        我们可以手动模拟一下，如何对比呢？
        手动第一个字符开始  a b c 依次都在t中出现
【解析】
    1、双指针
        双指针这个就很好理解了，一个一个对比，匹配s中的字符
        伪代码
            i j 两个指针分别在 s t的初始位置
            然后贪心的对比
            while遍历字符串
                如果匹配字符成功
                    i++ j++
                如果匹配失败
                    j++
            如果最后i还没有遍历到字符串尾端，那么匹配失败
    
        时间复杂度O(n+m) 两个字符串的长度
        空间复杂度O(1)

    2、动态规划
        其实这个题吧，字符串的角度来看，貌似很难看到动态规划的影子
        
        
        但是我们通过双指针来看，我们是希望可以在t里面找到下一个 匹配的字符。
        所以如何预处理一下对于t的每一个位置，该位置开始往后每一个字符第一次出现的位置。

        通过预处理，将t这个长串。转化为 dp数组中的，然后之后如果又大量的s来进行匹配的话，都可以解决！
        dp[i][j] 表示 字符串t 从i开始往后字符j第一次出现的位置
        如果t中i的字符就是j 那么 dp[i][j] = i
        否则需要向后寻找 dp[i][j] = dp[i+1][j]

        这样 我们可以利用 dp 数组，每次  O(1) 地跳转到下一个位置，直到位置s中的每一个字符都匹配成功。

        预处理时间复杂度 O(m)，判断子序列时间复杂度 O(n)

    3、递归
        昨天那个题，看到鹏飞使用的递归，这样貌似又好像有点相同
        都是寻找到子问题
            递归通过子函数，不断调用，一层一层来解决。
            动规则是通过动态转移方程，利用方程来进行计算，通过对子问题的归纳。
        这样看来
            动规实际上是对递归的再度提升，把解题思想归纳起来了，从而得到一个方程式。
        

    4、栈
        这样再一看，递归跟栈又是相通的，好家伙，能延伸不少解题思路！！！
        
            把s转化为栈，每一次匹配，就将其弹出
            然后最后判断是否 栈为空
        时间复杂度O(len(t))
        空间复杂度O(len(s))
【总结】
    四种方法这样看下来
        如果只是解决单个s的问题，那么双指针应该是最容易理解和快速想到的，当然栈的方式也不错
        如果是要解决批量s的问题，那么动态规划应该是比较好做！


    双指针:   两个指针变量，你走一步我走一步，看谁先到头
    动态规划：找出子问题，通过归纳总结，写出动态转移方程
    递归：    找出子问题，通过子函数，进行处理，注意终止条件，递归使用的是调用栈
    栈：      凡是递归能处理的问题，栈都能解决

*/