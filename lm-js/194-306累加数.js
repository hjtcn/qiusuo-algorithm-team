/*
306. 累加数
累加数是一个字符串，组成它的数字可以形成累加序列。

一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。

说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1:

输入: "112358"
输出: true 
解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2:

输入: "199100199"
输出: true 
解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
进阶:
你如何处理一个溢出的过大的整数输入?
*/

/*
    思路：害，敢想不会做！
    想法：1）想着把所有的可能放到一个数组中，再计算是否是累加数列
         2）js可以根据特殊符号(如逗号)切分为数组。每个数字中间都会有两种考虑，加逗号，不加逗号

    题解：看完题解觉得不难了，事后诸葛亮，哈哈哈
        大问题，拆成小的子问题。a+b=next(?)

        遍历截取字符串并转化为数字
        for(let i=0;i<len-2;i++){
            let a=Number(num.slice(0,i))
            for(let j=i+1;j<len-1;j++){
                let b=Number(num.slice(i+1,j+1))
                if(dfs(a,b,j+1)) return true;
            }
        }

        边界控制：dfs边界控制，如果i==len return true
                数字的边界控制：首位数字不以0开头。
                如i,j遍历时:if(i>0&&num[0]=='0') break;
                           if(j>0&&num[i+1]=='0')break;
*/


var isAdditiveNumber = function(num) {
    let len=num.length
    let dfs=(a,b,i)=>{
        if(i==len) return true
        for(let j=i;j<len;j++){
            if (j > i && num[i] === "0") break;
            let next=Number(num.slice(i,j+1))
            if(a+b!=next) continue;
            if(dfs(b,next,j+1)) return true
        }
        return false
    }

    for(let i=0;i<len-2;i++){
        if(i>0&&num[0]=='0') break;
        let a=Number(num.slice(0,i+1))
        for(let j=i+1;j<len-1;j++){
            if(j>i+1&&num[i+1]=='0') break;
            let b=Number(num.slice(i+1,j+1))
            if(dfs(a,b,j+1)) return true
        }
    }
    return false
};

/*
    时间复杂度：O(2^N)
    这里时间复杂度也不是太明白。。。
    空间复杂度：o(N)
*/

/*
    思路：1.问题拆分
         2.边界控制
         伪代码：    
            function backtrack(i) {
                if (满足特定条件）{
                    // 返回结果 or 退出搜索空间
                }

                for (根据i能到达的下个状态j) {
                    dosomething(i) // 做一些操作
                    dfs(j)
                    undo() // 如果有必要
                }
            }
            backtrack(0)



    我每次是写到dfs参数的时候，就开始慌了，下次试试先搭架子，需要什么参数再放什么参数

*/