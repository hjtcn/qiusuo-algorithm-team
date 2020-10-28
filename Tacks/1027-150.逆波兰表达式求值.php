<?php

/*
 * @lc app=leetcode.cn id=150 lang=php
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
class Solution {

    /**
     * @param String[] $tokens
     * @return Integer
     */
    // 执行用时：16 ms, 在所有 PHP 提交中击败了93.10%的用户
    // 内存消耗：17.2 MB, 在所有 PHP 提交中击败了7.41%的用户
    function evalRPN1($tokens) {
        $stack = [];
        foreach($tokens as $token) {
            switch($token) {
                case "+":
                    array_push($stack, array_pop($stack)+array_pop($stack));
                    break;
                case "-":
                    array_push($stack, -array_pop($stack)+array_pop($stack));
                    break;
                case "*":
                    array_push($stack, array_pop($stack)*array_pop($stack));
                    break;
                case "/":
                    $b = array_pop($stack);
                    $a = array_pop($stack);
                    array_push($stack, intval($a/$b) );
                    break;
                default:
                    array_push($stack, intval($token));
            }
        }
        return array_pop($stack);
    }

    // 听了鹏飞之的意见，优化代码，调整后
    // 执行用时：16 ms, 在所有 PHP 提交中击败了93.10%的用户
    // 内存消耗：17 MB, 在所有 PHP 提交中击败了25.93%的用户
    function evalRPN2($tokens) {
        $stack = [];
        foreach($tokens as $k=>$token) {
            // 先判断是否是运算符
            if($token == '+' || $token == '-' || $token == '*' || $token == '/' ) {
                $b = array_pop($stack);
                $a = array_pop($stack);
                $res = 0;
                switch($token) {
                    case "+":
                        $res = $b+$a;
                        break;
                    case "-":
                        $res = -$b+$a;
                        break;
                    case "*":
                        $res = $b*$a;
                        break;
                    case "/":
                        $res = intval($a/$b);
                        break;
                    default:
                    break;
                }
                array_push($stack, $res);

            }else{
            // 不是运算符直接压入栈中
                array_push($stack, intval($token));
            }
        }
        return array_pop($stack);
    }
    
}
// @lc code=end

$obj = new Solution();

$tokens = ["2","1","+","3","*"];
$res = $obj->evalRPN2($tokens);
var_dump($res);

/*

逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
    平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 
    该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 

好处：去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果


思路:
    1、使用栈
    2、遇到数字直接进栈
    3、遇到符号，拿栈顶元素作为 运算符后面的元素，再弹出栈顶元素作为 运算符前面的元素
    4、计算后压入栈中
    5、循环以上操作

主要函数
    array_push 压入栈
    array_pop  弹出栈

*/