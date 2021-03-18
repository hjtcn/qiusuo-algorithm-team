<?php

/*
 * @Descripttion:  周四了
 * @Author: tacks321@qq.com
 * @Date: 2021-03-18 11:13:49
 * @LastEditTime: 2021-03-18 11:44:49
 */

/*

面试题 01.02. 判定是否互为字符重排

给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true 
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100

*/

class Solution1 {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    // 执行用时：4 ms, 在所有 PHP 提交中击败了94.44%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了40.28%的用户
    function CheckPermutation($s1, $s2) {
        // Some judgments
        if(strlen($s1) != strlen($s2)) {
            return false;
        }
        // Converts a string to an array
        $arr1 = str_split($s1);
        $arr2 = str_split($s2);
        // This function sorts an array
        sort($arr1);
        sort($arr2);

        return $arr1 == $arr2;
    }
}

class Solutio2 {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了55.56%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了12.50%的用户
    function CheckPermutation($s1, $s2) {
        // Counts all the values of an array
        return array_count_values(str_split($s1)) == array_count_values(str_split($s2));
    }
}
 

class Solutio3 {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了8.33%的用户
    // 内存消耗：15 MB, 在所有 PHP 提交中击败了97.22%的用户
    function CheckPermutation($s1, $s2) {
        // [count_chars] Return information about characters used in a string
        // [array_diff_assoc]  Computes the difference of arrays with additional index check
        return empty(array_diff_assoc(count_chars($s1) , count_chars($s2)));
    }
}
 

class Solutio4 {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    // 笨办法
    // 执行用时：12 ms, 在所有 PHP 提交中击败了8.33%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了19.44%的用户
    function CheckPermutation($s1, $s2) {
        $len1 = strlen($s1);
        $len2 = strlen($s2);
        if ($len1 != $len2) {
            return false;
        }
        $map = [];
        // 把 s1 对应生成一个 map
        for ($i = 0; $i < $len1; $i++) {
            $map[$s1[$i]] = $map[$s1[$i]] ? ($map[$s1[$i]] + 1) : 1;
        }

        // 遍历 s2
        for ($j = 0; $j < $len2; $j++) {
            // 如果 s2 出现与 s1不同的字符， 直接结束
            if (!$map[$s2[$j]]) {
                return false;
            } else {
                // 出现的话，就看出现的次数是否一致
                $map[$s2[$j]] = $map[$s2[$j]] - 1;
                
                if ($map[$s2[$j]] < 0) {
                    return false;
                }
            }
        }
        return true;
    }
}
 


/*
【题目】

    给定两个字符串 s1 和 s2，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

    貌似好像，只要两个字符串的字母以及字母对应的个数相等，就可以。

    感觉主要是PHP 函数使用
【解析】
    1、字符排序
        字符串分割成为数组，
        然后按照字典重排序，    [sort]排序函数 时间复杂度 O(n*logn) 空间复杂度 O(n)
        然后进行比较两个数组是否相等

    2、统计每个字符的数量
        字符串分割称数组，
        统计数组中每个字符的数量，[array_count_values]也是用到数组，挺耗时的，占内存
        然后比较这字符数量是否相等

    3、差集法
        字符串直接统计每个字符的数量 返回数组 [count_chars] 返回数组，每个字符对应的个数
        比较两个数组的差集是否为空

    4、笨办法 map

    感觉基本上没有什么最优的解法， 相对来说倾向于第一种，使用sort排序后对比。

*/