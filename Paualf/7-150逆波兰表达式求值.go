/*
题意：
根据 逆波兰表示法，求表达式的值。
有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
说明：
整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
 
示例 1：
输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
示例 2：
输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: 该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
示例 3：
输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释: 
该算式转化为常见的中缀算术表达式为：
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

/*
    思路：如果遇到 + - * / 则从栈中弹出两个元素，然后将运算结果 push到栈中
    终止条件：
*/
func evalRPN(tokens []string) int {
    ret := 0
    stack := []int{}
    //dict := map[string]bool
    for i := 0;i < len(tokens);i++ {
        if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
            // pop two number then calcuate
            x := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            y := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            z := 0
            if tokens[i] == "+" {
                z = y + x
            }else if tokens[i] == "-" {
                z = y - x
            }else if tokens[i] == "*" {
                z = y * x
            }else if tokens[i] == "/" {
                z = y / x
            }
            stack = append(stack, z)
        }else {
            // push
            num,_ :=  strconv.Atoi(tokens[i])
            stack = append(stack, num)
        }
    }
    ret = stack[len(stack) - 1]
    return ret   
}

一开始的时候输出的是错误结果，自己想思路是没问题的哈，想到了是不是计算顺序的问题，然后就出现在减法和除法上了，一开始都是 x+y/x - y /x *y/x/y，因为减法和除法对除数和被除数是有要求的，所以还是有细节是需要注意的
可不可以进行优化，比如定义map然后然后优化if 判断是否是运算符语句呢？map 如何定义，发现自己还是不会定义和写

dict := map[string]sring {
    "+" : "+",
    "-" : "-",
    "*" : "*",
    "/" " "/",
}

map定义好了，发现并不能很好的优化代码
看了leetcode 国际站上，有使用字符串函数判断是否存在某个字符的函数
https://leetcode.com/problems/evaluate-reverse-polish-notation/discuss/314440/Go-Solution-using-stack

func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    
    var num1,num2 int
    for _,token := range tokens {
        if len(token) == 1 && strings.ContainsAny(token, "+-*/") {
            num2 = stack[len(stack) - 1]
            num1 = stack[len(stack) - 2]
            stack = stack[:len(stack) - 2]
            
            switch token {
                case "+":
                    stack = append(stack, num1 + num2)
                case "-":
                   stack = append(stack,num1 -num2)
                case "*":
                    stack = append(stack, num1 * num2)
                case "/":
                    stack = append(stack, num1 / num2)
            }
        }else {
            num,_ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }
    return stack[0]
}

优化完if判断运算符语句以后，可不可以优化判断完运算符以后，进行计算的if语句,目前看下来下面的判断语句是少不了的
可以把加减乘除这个操作拆成函数抽出来

复杂度分析：
时间复杂度: 遍历字符串数组 O(n)
空间复杂度： 使用栈存储运算:O(n)

总结：
1. 思路正确，但是有些细节，可能还是会出错

2. 写完一定要看别人的代码，因为自己写的代码不是很优雅，或者说是有好的方法自己不知道，比如使用slice实现栈的时候一下子可以弹出2个元素，自己是一个元素一个元素的弹

3. 上学的时候老师讲栈的时候一直是拿这个作为例子，现在才能写出来。。。
