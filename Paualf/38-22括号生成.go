/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例：
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/

Clearfication:
     从最小子问题开始分析：
a. 生成所有的括号组
b. 判断是否合法
c. 将合法的加到结果数组里面
细节点：
1. 开辟存储返回的结果数组
2. 构造递归生成的字符串
3. 判断是否吃是否合法
继续拆分：
判断是否合法的逻辑：
如果是左括号放入栈中，如果是右括号，判断栈顶元素是否为左括号，如果为左括号则将栈顶元素弹出，否则 返回 false, 最后判断栈是否为空
写出了一个半成品，卡在了 tmpStr 赋值的时候

 func generateParenthesis(n int) []string {
    ret := []string{}
    res := []string{}
    str := ""
    generateAllParenthesis(2 * n,0,&res,str)
    for _,v := range res {
        fmt.Println(v)
        if isValid(v) {
            ret = append(ret, v)
        }
    }
    return ret
}
func generateAllParenthesis(num,level int,res *[]string,str string){
    // terminator
    if level == num {
        *res = append(*res,str)
        return
    }
    // process current logic
    str += "("
    generateAllParenthesis(num,level + 1,res,str)
    str += ")"
    generateAllParenthesis(num, level + 1,res,str)
}
func isValid(s string)bool {
    stack := []rune{}
    for _,v := range s {
        if v == '(' {
            stack = append(stack, v)
        }else if v == ')' {
            top := stack[len(stack) - 1]
            if top == '(' {
                // pop
                stack = stack[:len(stack) - 1]
            }else {
                return false
            }
        }
    }
    return len(stack) == 0
}


上面的代码出问题感觉主要在
// terminator
 if level == num {
    tmpStr := str
    *res = append(*res,tmpStr)
    return
 }
因为str 还是会在后面使用的，所以这里也有点疑惑，append 到 *res 的值会超过长度6
搞了差不多半个小时，没有搞出来，换个思路，去看看回溯法
func backtrack(left int, right int, str string, res *[]string) {
    if left == 0 && right == 0 {
        *res = append(*res, str)
        return
    }
    if left > 0 {
        backtrack(left-1, right, str+"(", res)
    }
    if right > left {
        backtrack(left, right-1, str+")", res)
    }
}
func generateParenthesis(n int) []string {
    var res []string
    backtrack(n, n, "", &res)
    return res
}

下午看了上面的回溯法，感觉应该是自己写的地方有问题，然后经过调试找到了问题所在
 func generateParenthesis(n int) []string {
    ret := []string{}
    res := []string{}
    str := ""
    generateAllParenthesis(2 * n,0,&res,str)
    for _,v := range res {
        fmt.Println(v)
        if isValid(v) {
            ret = append(ret, v)
        }
    }
    return ret
}
func generateAllParenthesis(num,level int,res *[]string,str string){
    // terminator
    if level == num {
        *res = append(*res,str)
        return
    }
    // process current logic
    generateAllParenthesis(num,level + 1,res,str+"(")
    generateAllParenthesis(num, level + 1,res,str+")")
}
func isValid(s string)bool {    
    balance := 0
    for _,v := range s {
        if v == '(' {
            balance++
        }else if v == ')' {
            balance--
        }
        
        if balance < 0 {
            return false
        }
    }
    return balance == 0
}

自己原来写的是
str += "("
generateAllParenthesis(num,level + 1,res,str)
str += ")"
generateAllParenthesis(num, level + 1,res,str)
这样写的话，第1个str 的修改回影响后面的str 的修改的，这样写是会有问题的，我们可以写成
generateAllParenthesis(num,level + 1,res,str+"(")
generateAllParenthesis(num, level + 1,res,str+")")
这样第一层str的修改，不会影响第二层str的值
还有 判断字符串是否合法，我们可以通过一个变量的增加和删除，当变量小于0的时候 return false 
上面是最基础的思路，然后我们发现，我们在遍历的过程中是可以知道左括号的数量和右括号的数量的，当左右括号数量>n，这个时候是不合法的，同时当右括号数量大于左括号数量的时候，这个时候也是不合法的，我们可以提前终止这些不合法的遍历

func backtrack(left,right,num int, str string,ret *[]string) {
    if left == num && right == num {
        *ret = append(*ret, str)
        return
    }
    if left < num {
        backtrack(left + 1,right,num,str + "(",ret)
    }
    if right < left {
        backtrack(left, right + 1,num,str + ")", ret)
    }
} 

func generateParenthesis(n int) []string {
    ret := []string{}
    str := ""
    backtrack(0,0,n,str,&ret)
    return ret
}

复杂度分析：
时间复杂度：暴力解法：O(2^n)，这里还是有点理不清的感觉，遍历二叉树的时候时间复杂度为O(n), 递归抽象出来也是一颗树型结构，时间复杂度就变成了 2^n
空间复杂度：O(n) 递归栈的深度
如果是剪枝的话时间复杂度和空间复杂度是多少呢？发现自己不会分析了，对递归的空间复杂度和时间复杂度也不是很熟悉
后来想了一下，时间复杂度还是O(2^n), 空间复杂度是O(n)

总结：
递归的复杂度分析后面需要多多注意一下
