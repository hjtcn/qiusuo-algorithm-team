<?php
/*
 * @Descripttion: 周五 好好找找刷题的感觉
 * @Author: tacks321@qq.com
 * @Date: 2021-07-16 11:02:37
 * @LastEditTime: 2021-07-16 18:25:16
 */

/*
 * @lc app=leetcode.cn id=139 lang=php
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (50.35%)
 * Likes:    1055
 * Dislikes: 0
 * Total Accepted:    162.5K
 * Total Submissions: 320.5K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 * 
 * 说明：
 * 
 * 
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 * 
 * 
 * 示例 1：
 * 
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 * 
 * 
 * 示例 2：
 * 
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 * 
 * 
 * 示例 3：
 * 
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 * 
 * 
 */

// @lc code=start
class Solution {

    protected $wordDict = [];
    /**
     * @param String $s
     * @param String[] $wordDict
     * @return Boolean
     */
    // 暴力（超时）
    function wordBreak($s, $wordDict) {
        $this->wordDict = $wordDict;
        return $this->_helper($s, 0);
    }

    /**
     * 辅助递归
     *
     * @param string $str
     * @param int $start
     * @return bool
     */
    function _helper($str, $start) {
        $len = strlen($str); // 截取后的字符串的长度
        // 如果到达最后
        if($start == $len) {
            return true;
        }
        // 遍历字符串 s 
        // 这里注意end=len,是表示最后一个字符，因为substr截取的时候，end-start表示截取的长度
        // 假如 s=leetcode len=8 那么只有当start=4 end=8的时候 end-start=4，substr('leetcode', 4, 8)的时候表示截取 code
        // 因此 end并不是表示移动越界
        for ($end = $start + 1; $end <= $len ; $end++) {
            // 截取每一个单词
            $curstr = substr($str, $start, $end-$start);
            // 判断字典里面是否有这个字符串
            if(in_array($curstr, $this->wordDict)) {
                // 如果有继续递归
                if($this->_helper($str, $end)) {
                    return true;
                }
            }
        }
        return false;
    }
}


class Solution2 {

    protected $wordDict  = [];  // 字典
    protected $len       = 0;   // 字符串长度
    protected $memoryMap = [];  // 记忆化数组

    
    /**
     * @param String $s
     * @param String[] $wordDict
     * @return Boolean
     */
    // 暴力+记忆化
    // 执行用时：8 ms, 在所有 PHP 提交中击败了96.43%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了25.00%的用户
    function wordBreak($s, $wordDict) {
        $this->wordDict  = $wordDict;                         // 字典初始化
        $this->len       = strlen($s);                        // s字符串长度 初始化
        $this->memoryMap = [];  // 记忆化
        
        return $this->_helper($s, 0);
    }

    /**
     * 辅助 递归+记忆化
     *
     * @param string $str
     * @param int $start
     * @return bool
     */
    function _helper($str, $start) {
        // 递归终止条件 如果到达最后都从字典中找到，那么终止
        if($start == $this->len) {
            return true;
        }

        // Map Get
        // 已经计算过的，直接返回
        if(isset($this->memoryMap[$start])) {
            return $this->memoryMap[$start];
        }
        
        // 遍历字符串 s 
        for ($end = $start + 1; $end <= $this->len ; $end++) {
            // 截取每一个单词
            $curstr = substr($str, $start, $end-$start);
            // 判断字典里面是否有这个字符串
            if(in_array($curstr, $this->wordDict)) {
                // 如果有继续递归
                if($this->_helper($str, $end)) {
                    // Map Put 
                    // 保存计算结果，字符串可以拆分
                    $this->memoryMap[$start] = true;
                    return true;
                }
            }
        }
        // Map Put 
        // 保存计算结果，字符串不可以拆分
        $this->memoryMap[$start] = false;
        return false;
    }
}



class Solution3 {
    
    /**
     * @param String $s
     * @param String[] $wordDict
     * @return Boolean
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了75.00%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了35.71%的用户
    function wordBreak($s, $wordDict) {
        $len = strlen($s);
        $dp = array_fill(0, $len + 1, false);  // 初始化DP
        // 开始必备条件
        $dp[0] = true; // 字符串为空的时候 是存在字典中的
        
        // 遍历字符串背包
        for ($i=1; $i <= $len; $i++) { 
            // 遍历物品字典
            for ($j=0; $j < $i; $j++) { 
                // 截取当前物品，开始j,截取长度i-j
                $curstr = substr($s, $j, $i-$j);

                // 动态转移方程
                // 如果存在字典，并且是可以拆分的
                if(in_array($curstr, $wordDict) && $dp[$j]) {
                    $dp[$i] = true;
                    break; // 既然已经找到了，就提前break
                }
            }
        }
        return $dp[$len];
       
    }
}



$Obj = new Solution();
$Obj = new Solution2();
$Obj = new Solution3();

$s = "catsandog";
$wordDict = ["cats", "dog", "sand", "and", "cat"];

$s = "leetcode";
$wordDict = ["leet", "code"];

// $s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab";
// $wordDict = ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"];

var_dump($Obj->wordBreak($s, $wordDict));

// @lc code=end

/*

暴力+记忆化+动规+完全背包

【 Clarification 】
    给一个字典 wordDict 包含非空单词并且不重复 （类似物品）
        例如
            [ sky, cloud ]
    给一个非空的字符串s （类似背包）
        例如 
            skycloudsky
    
    判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词
        那么也就是转化成
            => 判断背包装的东西是否符合 物品集合

   
【 Possible-Solution 】

（1）暴力
    俗话说，万物皆可暴力求解，顶多是时间问题，给我足够内存，模拟一下题意即可
        例如：s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
            遍历一下 s = "catsandog"
                start=0 end=0 字符串是c, 看一下wordDict里面没有， continue
                start=0 end=1 字符串是ca,看一下wordDict里面没有， continue
                start=0 end=2 字符串是cat,看一下wordDict里面有，那么进行截取，截取后 sandog, 继续遍历
            遍历一下 s = "sandog" 
                start=0 end=0 字符串是s, 看一下wordDict里面没有 continue
                start=0 end=1 字符串是sa, 看一下wordDict里面没有 continue
                start=0 end=2 字符串是san, 看一下wordDict里面没有 continue
                start=0 end=3 字符串是sand, 看一下wordDict里面有,那么进行截取，截取后 og, 继续遍历
            遍历一下 s = "og" 
                start=0 end=0 字符串是o, 看一下wordDict里面没有 continue
                start=0 end=1 字符串是og, 看一下wordDict里面没有 pass
                
            最后两个字母，没有对应的字典单词，于是返回false

        复杂度
            时间 O(n³),每次截取字符串的时间
            空间 O(n) ,递归调用栈

        代码超时
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
            ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
        用PHP暴力跑了一下这个案例数据
            好家伙直接卡的没反应
        很显然，递归做了重复的操作
            可以使用数组，保存一下递归的计算结果。

（2）暴力+记忆化
    加了一个记忆化，虽然时间复杂度没有变化，但是会稍微快一些，代码竟然通过啦！

（3）动规(背包问题)
    字典就是可以选择的物品（随意使用） 
    字符串就是背包
    => 问：物品能不能把背包装满
    => 可以随意使用字典中的单词 完全背包

    动规五部曲
        ① 确定dp含义
            dp[i] 字符串长度为i，dp[i]=true的时候 表示可以拆分单词
        ② 动态转移方程
            如果
                dp[j]=true，并且[j,i]这个区间的子串出现在字典中，那么dp[i]一定true  (j<i)
                if(in_array([j,i], dict)) {
                    dp[i] = true
                }
        ③ dp初始化
            dp[i] 的状态依靠 dp[j]是否为true，那么dp[0]就是递归的根基，dp[0]一定要为true
            下标非0的dp[i]初始化为false， 只要没有被覆盖说明都是不可拆分的单词

        ④ 确定遍历顺序
            完全背包
                遍历背包放在外循环
                将遍历物品放在内循环
                    内循环从前到后
        ⑤ 举例推导  
            s = "leetcode"
            wordDict = ["leet", "code"]
            dp[0] = 1
            dp[1] = 不在物品中 0
            dp[2] = 不在物品中 0
            dp[3] = 不在物品中 0
            dp[4] = 在物品中 1
            dp[5] = 不在物品中 0
            dp[6] = 不在物品中 0
            dp[7] = 不在物品中 0
            dp[8] = 在物品中 1
            最后返回 dp[len] 即可
    
    复杂度
        时间复杂度：O(n³)。 因为substr()函数也是需要时间复杂度
        空间复杂度：O(n)。  需要dp
    

【 Code-Template 】
    暴力
        实在不行就递归，好理解，先把思路写处理
    暴力+记忆化
        递归的重复计算，采用记忆化数组进行存储
    动规完全背包
        这个地方还需要归纳总结一下，还需要多做题
        外循环
            for(i, 背包 , i++)
            内循环
                for(j, 物品, j++)
                内循环的顺序正循环
                    动规
                
    
【 Test-Cases 】
    1、代码首先要通过题目给的示例
        $s = "catsandog";
        $wordDict = ["cats", "dog", "sand", "and", "cat"];
        var_dump($Obj->wordBreak($s, $wordDict));
    2、如果提交OJ判定，出错的数据样例，也可以作为本地测试依据
        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
        ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]

【 summary 】
    1、暴力+记忆化，也是一个不错的方法，是一种比较直观的一个思路。而且竟然感觉比动规还快。
    2、完全背包还需要多看看进行总结归纳，这样理解起来会比较方便。
*/