<?php

/*
 * @Descripttion: 原题周 来了！
 * @Author: tacks321@qq.com
 * @Date: 2021-04-19 18:33:44
 * @LastEditTime: 2021-04-20 19:07:08
 */
/*
 * @lc app=leetcode.cn id=151 lang=php
 *
 * [151] 翻转字符串里的单词
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (44.85%)
 * Likes:    300
 * Dislikes: 0
 * Total Accepted:    124.9K
 * Total Submissions: 271.1K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给定一个字符串，逐个翻转字符串中的每个单词。
 * 
 * 说明：
 * 
 * 
 * 无空格字符构成一个 单词 。
 * 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："the sky is blue"
 * 输出："blue is sky the"
 * 
 * 
 * 示例 2：
 * 
 * 输入："  hello world!  "
 * 输出："world! hello"
 * 解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
 * 
 * 
 * 示例 3：
 * 
 * 输入："a good   example"
 * 输出："example good a"
 * 解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 * 
 * 
 * 示例 4：
 * 
 * 输入：s = "  Bob    Loves  Alice   "
 * 输出："Alice Loves Bob"
 * 
 * 
 * 示例 5：
 * 
 * 输入：s = "Alice does not even like bob"
 * 输出："bob like even not does Alice"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^4
 * s 包含英文大小写字母、数字和空格 ' '
 * s 中 至少存在一个 单词
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 请尝试使用 O(1) 额外空间复杂度的原地解法。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @return String
     */
    // API
    // 执行用时：8 ms, 在所有 PHP 提交中击败了91.30%的用户
    // 内存消耗：15.7 MB, 在所有 PHP 提交中击败了5.80%的用户
    function reverseWords($s) {
        //  利用语言特性
        // explode拆分成数组 array_reverse数组翻转  implode合并  
        $s = preg_replace("/\s(?=\s)/", "\\1", $s); // 多个连续空格转化为一个空格
        $s = trim($s);
        $arr = array_reverse(explode(' ', $s));
        $arr = array_filter($arr,function ($value){
              if ($value != ' '){
                        return true;
                }
                 return false;
        });
        return implode(' ', $arr);
    }
}

class Solution2 {

    /**
     * @param String $s
     * @return String
     */
    // 利用字符串栈来解决
    // 执行用时：12 ms, 在所有 PHP 提交中击败了52.17%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了46.38%的用户
    function reverseWords($s) {
        $s = preg_replace("/\s(?=\s)/", "\\1", $s); // 多个连续空格转化为一个空格
        $s = trim($s);
        $s = $s.' '; // 最后一个空格，遍历结束，最后一个整个单词入栈 ！！！注意点

        $len = strlen($s);
        $strItem = "";  // 单词容器
        $result  = "";  // 最终结果

        for ($i=0; $i <$len; $i++) { 
            $re = $s[$i];
            // 判断当前字母是否是空格
            if($re != ' ') {
                // 单词拼接
                $strItem .= $re;
            } else {
            // 遇到空格，整个单词入栈
                $result  = $strItem.' '.$result;
                $strItem = ''; // 清空，方便下一次使用
            }
        }
        return trim($result, ' ');// 去除前后空格
    }

}



class Solution3 {

    /**
     * @param String $s
     * @return String
     */
    // 双指针
    // 执行用时：12 ms, 在所有 PHP 提交中击败了52.17%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了43.48%的用户
    function reverseWords($s) {
        $s      = trim($s);    // 移除前后空格
        $len    = strlen($s);
        $string = "";          // 处理后的字符串

        for ($i=0; $i < $len; $i++) { 
            // 如果当前是空格 ，并且下一个也是空格
            if($s[$i] == ' ' && isset($s[$i+1]) && $s[$i+1] == ' ') {
                continue;
            }else if($s[$i] == ' '){
                $string .=" ";
            }else{
                $string .=$s[$i];
            }
        }

        // 分割字符串
        $arr = explode(" ", $string);
        $left = 0;
        $right= count($arr) - 1;
        $temp = "";
        // 双指针
        while($left < $right) {
            // 前后交换两个单词
            $temp   = $arr[$left];
            $arr[$left]  = $arr[$right];
            $arr[$right] = $temp;
            $left++;
            $right--;
        }
        return implode(" ", $arr); // 用空格合并数组为字符串
    }

}

$obj = new Solution3    ();
$s = "  Bob    Loves  Alice   ";

var_dump($obj->reverseWords($s));
// @lc code=end

/*
【题目】
    一个有空格的字符串，逐个翻转字符串中的每个单词。
        字符串 =>  英文大小写字母、数字和空格
        单词   =>  无空格字符构成一个单词 
        
        输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括 => 前后空格全部删除
        如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一         => 每个单词之间的空格只保留一个
        
        O(1) 额外空间复杂度的原地解法 => 不能使用哈希数组这种

  
【解析】
    
    1、PHP现成的API处理
    2、栈
        关键词：翻转
        数据结构： 先进后出 => 栈 
    
        伪代码：
            首先移除头尾多余的空格；
            遍历字符串
                自左到右依次遍历，逐个取出单词
                一旦碰到空格，找到完整的单词，加入到栈中
                循环结束的时候，别忘记最后一个单词
            最后按照空格拼接起来。
        时间O(n) 遍历整个字符串
        空间O(1) 只是用到的了字符串变量
    3、 双指针
        时间O(n) 遍历整个字符串
        空间O(n) 用了数组
        

【思考】
    1、是否可以形成有效的算法来解决问题
    2、对字符串处理要注意边界情况防止越界
    
*/