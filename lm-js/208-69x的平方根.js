/*
69. x 的平方根
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
由于返回类型是整数，小数部分将被舍去。
*/

/*
    思路：模版还是有的
    while(l<=r){
        let mid=parseInt((l+r)/2)
        if(mid*mid>x){
            r=mid-1
        }
        else if(mid*mid<x){
            l=mid+1
        }
        else{
            return mid
        }
    }
    模糊点:
    1.返回值是l?r?mid?
    2.和while条件有关吗？
    3.l直接=mid+1,会有错过的吗？


*/

var mySqrt = function(x) {
    let l=1,r=x
    while(l<=r){
        let mid=parseInt((l+r)/2)
        if(mid*mid>x){
            r=mid-1
        }
        else if(mid*mid==x){
            return mid
        }
        else{
            l=mid+1
        }
    }
    return r
}

//高数忘完了，切线方程;牛顿迭代法
var mySqrt = function(x) {
    if(x==0) return 0
    let mid=x/2
    let th=0.01
    while(Math.abs(mid*mid-x)>th){
        mid=mid*0.5+x*0.5/mid
    }
    return parseInt(mid)
}

/*
    时间复杂度：O(logN)
    空间复杂度：O(N)

*/

/*
    思考：斜率还知道，切线方程是真的忘了。
*/