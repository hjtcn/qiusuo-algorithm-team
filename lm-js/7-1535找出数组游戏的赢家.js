// 给你一个由 不同 整数组成的整数数组 arr 和一个整数 k 。

// 每回合游戏都在数组的前两个元素（即 arr[0] 和 arr[1] ）之间进行。比较 arr[0] 与 arr[1] 的大小，较大的整数将会取得这一回合的胜利并保留在位置 0 ，较小的整数移至数组的末尾。当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家 。

// 返回赢得比赛的整数。

// 题目数据 保证 游戏存在赢家。
// 提示：

// 2 <= arr.length <= 10^5
// 1 <= arr[i] <= 10^6
// arr 所含的整数 各不相同 。
// 1 <= k <= 10^9

//暴力模拟一直超时，还需要时间实验。。。。。。。。

/**
 * @param {number[]} arr
 * @param {number} k
 * @return {number}
 */
var getWinner = function (arr, k) {
    let max = arr[0], count = 0
    for (let i = 1; i < arr.length; i++) {
        if (max < arr[i] && count < k) {
            //更新max，且count设置为1，最大值赢的回合数重新计数
            max = arr[i]
            count = 1
        }
        else {//连胜
            count++;
        }
        if (count == k) {//返回当前最大值
            return max
        }
    }
    return max
};

/** 题解
1535-找出数组游戏的赢家。一层for循环遍历，记录最大的值，如果arr[i]>max,则更新max，且count设置为1，最大值赢的回合数重新计数，如果小于max则说明连胜，count++，
然后判断如果所赢回合数等于k，则返回当前最大值
 复杂度分析：
    平均时间复杂度是:O(N)
    一层for循环变量
    空间复杂度：O(1)
    声明几个变量，存储最大值
 */


