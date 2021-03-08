// 69. x 的平方根
// 实现 int sqrt(int x) 函数。

// 计算并返回 x 的平方根，其中 x 是非负整数。

// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

// 示例 1:

// 输入: 4
// 输出: 2
// 示例 2:

// 输入: 8
// 输出: 2
// 说明: 8 的平方根是 2.82842..., 
//      由于返回类型是整数，小数部分将被舍去。

/*
    思路：这道题对我来说难度点在于left,right的更新，以及最终返回值的确定。
         首先我对于二分的架子表现在这道题可以怎么搭建还是比较清晰的。
         很快就写出了搭好框架，根据每次更新得到的mid*mid与目标x去对比更新。
         
         在while循环之外，我原本返回的是mid，后来发现，当x==8，预期值为2，而我却返回3.
         后来我尝试返回right提交了下,此时我对于为什么返回right其实有些模模糊糊的。
         仔细思考了一下跳出循环的最后一次mid是没有计算的，不能直接返回。


         看题解了解到还有牛顿迭代法，跟线性方程有关，确实有些牛啊，死记住吧，一方面排斥是数学方向，一方面也是知识的积累嘛

*/

/**
 * @param {number} x
 * @return {number}
 */
 var mySqrt = function(x) {
    let left=1,right=x
    while(left<=right){
        let mid=parseInt((left+right)/2)
        if(mid*mid>x){
            right=mid-1
        }
        else if(mid*mid==x){
            return mid
        }
        else{
            left=mid+1
        }
    }
    return right
};



/**
 * @param {number} x
 * @return {number}
 */
 var mySqrt = function(x) {
    if(x === 0) {
        return 0
    }
    return Math.floor(sqrt(x)(x))
};
function sqrt(x) {
    let origin = x;
    return function inner(x) {
        (x+a/x)/2
        let xn1 = x - (x*x-origin) / (2*x)
        if(x - xn1 <= 0.1) { // 精度问题
            return xn1
        }
        return inner(xn1)
    }
}

/*
    时间复杂度：O(logN).牛顿迭代法要收敛的快一些
    空间复杂度:O(1)
*/