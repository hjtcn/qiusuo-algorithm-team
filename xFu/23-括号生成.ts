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


// ============================================================
// ===                                                      ===
// ===      状态：通过,执行用时: 92 ms,内存消耗: 40.4 MB     ===
// ============================================================

function generateParenthesis(n: number): string[] {
    let matchQuote:string[] = [];

    function findMatchQuote (str:string, left:number, right:number) {
        // 括号匹配的数量足够了就可以添加结果了
        // 也可以是left和right都是0的时候结束
        if(str.length === 2 * n){
            matchQuote.push(str);
            return;
        }

        // 只要左括号存在都可以添加左括号
        if(left > 0){
            findMatchQuote(`${str}(`, left-1, right);
        }

        // 只有左括号的个数小于右括号的时候才可以添加右括号
        if(left < right){
            findMatchQuote(`${str})`, left, right-1);
        }
    }

    findMatchQuote('', n, n);
    return matchQuote;
};