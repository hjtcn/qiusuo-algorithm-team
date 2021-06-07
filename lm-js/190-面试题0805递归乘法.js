/*
面试题 08.05. 递归乘法
递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。

示例1:

 输入：A = 1, B = 10
 输出：10
示例2:

 输入：A = 3, B = 4
 输出：12
提示:

保证乘法范围不会溢出
*/
/*
    思路：A*B的核心在于B个A相加。。。
        递归的核心则在于随着加A的过程B--
    题解： 位移啃不动呀。B>>1相当于除2    
    return ((B & 1) ? A : 0) + ((B > 1) ? multiply(A + A, B>>1) : 0);
        1）B为奇数。起始值为A。而multiply(A+A,B>>1).意味着A*2的同时，B除以2.
        6*7----->6+(6+6,3)---->6+12+(12+12,1)---->6+12+24+0

        2)B为偶数。起始值为0,
        7*6------>(7+7,3)----->14+(14+14,1)----->14+28+0
*/
var multiply = function(A, B) { 
    if(A>B){
        let tmp=A
        A=B
        B=tmp
    }
    let getPower=(A,B)=>{
        if(B<=1){
            return A
        }
        return getPower(A,B-1)+A
    }
    return getPower(A,B)
};
/*
    时间复杂度：O(min(A,B))
    空间复杂度：O(min(A,B))
*/

var multiply = function(A, B) { 
    if(A>B){
        let tmp=A
        A=B
        B=tmp
    }
    let getPower=(A,B)=>{
        return ((B&1)?A:0)+((B>1)?getPower(A+A,B>>1):0)
    }
    return getPower(A,B)
};
/*
    时间复杂度：O(log2(min(A,B)))
    空间复杂度：O(log2(min(A,B)))
    递归记录每一层的和值
*/

/*
    了解：n>>1,n/2
         n>>3,n/(2^3)
         n<<1,n*2
         n<<3,n*(2^3)
*/