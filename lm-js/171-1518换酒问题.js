/*
1518. 换酒问题
小区便利店正在促销，用 numExchange 个空酒瓶可以兑换一瓶新酒。你购入了 numBottles 瓶酒。

如果喝掉了酒瓶中的酒，那么酒瓶就会变成空的。

请你计算 最多 能喝到多少瓶酒。

 

示例 1：



输入：numBottles = 9, numExchange = 3
输出：13
解释：你可以用 3 个空酒瓶兑换 1 瓶酒。
所以最多能喝到 9 + 3 + 1 = 13 瓶酒。
示例 2：



输入：numBottles = 15, numExchange = 4
输出：19
解释：你可以用 4 个空酒瓶兑换 1 瓶酒。
所以最多能喝到 15 + 3 + 1 = 19 瓶酒。
示例 3：

输入：numBottles = 5, numExchange = 5
输出：6
示例 4：

输入：numBottles = 2, numExchange = 3
输出：2
 

提示：

1 <= numBottles <= 100
2 <= numExchange <= 100
*/

/*
    思路：暴力模拟。循环拿瓶盖去兑换酒，直到瓶盖不足以兑换1瓶酒时
        构建sum记录酒数,在循环中，sum+=parseInt(numBottles/numExchange)
        兑换过一次后，numBottles更新
    题解：数学思路：一共有b瓶酒，e个瓶盖兑换一瓶酒
                兑换一次，失去的瓶子数e-1
                结束时机，不足以兑换一瓶，也就是剩余瓶盖小于e
                一共b瓶，兑换n次，最后最后瓶盖小于e
                b-n(e-1)<e
                        n>(b-e)/e-1
                总个数为b+(b-e)/e-1+1
                 
        
*/

/**
 * @param {number} numBottles
 * @param {number} numExchange
 * @return {number}
 */
 var numWaterBottles = function(numBottles, numExchange) {
    let sum=numBottles
    while(parseInt(numBottles/numExchange)){
        let exchange=parseInt(numBottles/numExchange)
        sum+=exchange
        numBottles=exchange+numBottles%numExchange
    }
    return sum
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

var numWaterBottles = function(numBottles, numExchange) {
    return numBottles>=numExchange?parseInt((numBottles-numExchange)/(numExchange-1))+numBottles+1:numBottles
}

/*
    时间复杂度：O(1)
    空间复杂度：O(1)
*/