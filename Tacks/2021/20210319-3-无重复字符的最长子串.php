<?php
/*
 * @Descripttion: 周五了 冲冲冲
 * @Author: tacks321@qq.com
 * @Date: 2021-03-19 09:50:23
 * @LastEditTime: 2021-03-19 14:12:55
 */



/*
 * @lc app=leetcode.cn id=3 lang=php
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (36.16%)
 * Likes:    5148
 * Dislikes: 0
 * Total Accepted:    887.2K
 * Total Submissions: 2.4M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: s = "abcabcbb"
 * 输出: 3 
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 * 
 * 示例 4:
 * 
 * 
 * 输入: s = ""
 * 输出: 0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * s 由英文字母、数字、符号和空格组成
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @return Integer
     */
    // 模拟法
    // 执行用时：12 ms, 在所有 PHP 提交中击败了97.22%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了32.15%的用户
    function lengthOfLongestSubstring($s) {
        $len = strlen($s); // 总的字符的长度
        if($len <= 1) {
            return $s;
        }
        
        $sub = '';         // 存储子串 当前里面不会有重复的字符
        $maxlen = 0;       // 最长子串的长度
        for($i=0; $i<$len; $i++) {
            $pos = strpos($sub, $s[$i]);
            if($pos !== false) {
                // 当前字符 有重复
                $sub .= $s[$i]; // 先追加上去，例如pww
                $sub  = substr($sub, $pos+1); // 从重复的位置开始截取 pww => w
            } else {
                // 当前字符 无重复
                $sub .= $s[$i];
            }
            // 每一次这样后，比较当前的 sublen 与 当前最大子串maxlen 谁更大
            $sublen = strlen($sub); // 当前子串长度
            $maxlen =  $sublen > $maxlen ? $sublen : $maxlen;
        }
        return $maxlen; // 得到的长度，不一定是sub的
    }
}

class Solution2 {

    /**
     * @param String $s
     * @return Integer
     */
    // 哈希法
    // 执行用时：20 ms, 在所有 PHP 提交中击败了80.13%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了42.53%的用户
    function lengthOfLongestSubstring($s) {
        $len = strlen($s);
        if($len <= 1) {
            return $len;
        }

        $left = 0;   // 左指针
        $hash = [];  // 记录哈希表    键为 字母，值为 索引
        $maxlen= 0;  // 全局 max 记录窗口最大长度
        #  以 pwwkew 为例，遍历整个字符串
        #  left=i=0, 窗口内只有一个字母 p, length=1, hash=['p' => 0]
        #  left=0, i=1, 索引 1 对应的字母 w 不在 hash 内，窗口向右滑动, length=2, hash=['p' => 0, 'w' => 1]
        #  left=0, right=2, 索引 2 对应的字母 w 在 hash 内,窗口起点要更新为上一个 w 的索引的下一个位置,left=hash['w'] + 1.  // 移动左指针
        # 同时， hash['w'] 更新为新的索引 ，指向最后一个 w, hash=['p' => 0, 'w' => 2]
 
        for ($i=0; $i < $len; $i++) { 
            $char = $s[$i]; // 当前字符
            if(isset($hash[$char])) {
                // 如果有重复的
                $left = max($left , $hash[$char] + 1);
            }
            // 提前结束
            if($left + $maxlen >= $len) {
                break; 
            }
            $hash[$char] = $i; // 记录一下字符
            $maxlen  = max($maxlen, $i-$left+1);
        }        

        return $maxlen;
    }
}
// @lc code=end



/*
【题目】
    找到字符串中，无重复字符的最长子串

    子串的问题，还是挺有趣

    当然这里的字符串都是 字母 数字这种单字节的，不包含中文。
【解析】
    1、模拟法
        顺着题目的意思，一步一步遍历字符串
            每次拿到子串，都进行查重，判断是否有重复的字符
                如果有重复的字符
                    就从重复的字符处截断，从新计算
                如果没有重复的字符
                    就在后面进行追加字符
            每次都比较当前子串的长度与当前最大子串的长度
        随着遍历完毕，可以得到一个最长子串的长度

        这里主要是 strpos(查找是字符的位置)  substr(字符串截取) 比较费时

    2、滑动窗口
        参考了别人的解法，其实就是把 strpos的作用换成 hash来模拟

        

*/