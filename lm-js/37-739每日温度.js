// 739. 每日温度
// 请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

// 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

/**
 * @param {number[]} T
 * @return {number[]}
 */
var dailyTemperatures = function(T) {
    let res=[]
    for(let i=0;i<T.length-1;i++){
        let flag=0;
        for(let j=i+1;j<T.length;j++){
            //找到比当前值大的，返回下标差
            if(T[i]<T[j]){
                flag=1
                res.push(j-i)
                break;
            }
            //当前值直到循环结束也没有找到比它大的值，补零
            if(j==T.length-1&&!flag){
                res.push(0)
            }
        }
    }
    //长度不够，后续补零
    if(res.length<T.length){
        for(let i=res.length;i<T.length;i++){
            res.push(0)
        }
    }
    return res
};

 /**题解
  暴力解决。两层for循环，查到比当前值大的再拐回来。
  我自己做的时候还考虑了长度不够补零或直到某个值直到最后也没有找到比它大的值的情况补0
  后来看题解，了解到fill方法，可以给默认数组补充0：new Array(T.length).fill(0);
  复杂度分析：
    时间复杂度是:O(N^2)
    递归调用，直到两条链表都遍历完毕，
    空间复杂度：O(N）
    未定义额外变量
 */

/**
 * @param {number[]} T
 * @return {number[]}
 */
var dailyTemperatures = function(T) {
    //  已知长度补零的操作
    let res = new Array(T.length).fill(0);
    let stack = [];
    for(let i = 0; i < T.length; i++) {
        // 如果栈长度不为空，且当前元素大于栈顶元素所对应的温度值
        while(stack.length && T[i] > T[stack[stack.length - 1]]) {
            //弹出栈顶元素
            let topIdx = stack.pop();
            //记录栈顶元素所对应的结果值
            res[topIdx] = i - topIdx;
        }
        //当前元素入栈
        stack.push(i);
        // console.log(stack,res)
    }
    return res;
};


 /**题解
  单调栈。这个是看题解了解到的方法。
  首先栈存储的是索引，如果当前元素比栈顶元素所对应的温度要大，则弹出栈顶元素topIdx(索引)
  再记录弹出的这个栈顶元素(索引)，开始赋值res[topIdx]等于索引差,直到栈空
  复杂度分析：
    时间复杂度是:O(N)
    正向遍历温度一次，对于温度列表中的每个下标，最多有一次进栈和出栈的操作。
    空间复杂度：O(N）
    定义了stack和res数组
 */