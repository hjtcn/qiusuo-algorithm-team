// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

// 示例：

// 输入：n = 3
// 输出：[
//        "((()))",
//        "(()())",
//        "(())()",
//        "()(())",
//        "()()()"
//      ]


var generateParenthesis = function(n) {
    let left=right=n,res=[]
    const dfs=(left,right,str)=>{
        if(str==n*2){
            res.push(str)
            return
        }
        if(left>0){
            dfs(left-1,right,str+'(')
        }
        if(left<right){
            dfs(left,right-1,str+')')
        }
    }
    dfs(left,right,'')
    return res
};

//没思路呀，看看题解，还是有点取巧的，get到了"剪枝"新名词。
//递归，记录(左括号剩余数，右括号剩余数，之前的路径)。
//跳出递归条件：在于路径长度=n*2
//递归条件
// 1.'('可以随意加，只要剩的还有；
// 2.')'不能随意加了，需要保证'('使用的比')'多，按剩余数来说，也就是left<right时，才能加')'