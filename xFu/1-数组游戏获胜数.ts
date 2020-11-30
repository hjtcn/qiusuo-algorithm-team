
// 1535. 找出数组游戏的赢家

// 给你一个由 不同 整数组成的整数数组 arr 和一个整数 k 。

// 每回合游戏都在数组的前两个元素（即 arr[0] 和 arr[1] ）之间进行。
// 比较 arr[0] 与 arr[1] 的大小，较大的整数将会取得这一回合的胜利并保留在位置 0 ，较小的整数移至数组的末尾。
// 当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家 。

// 返回赢得比赛的整数。

// 题目数据 保证 游戏存在赢家。


function getWinner(arr: number[], k: number): number {
    // 当前最大值连续获胜次数
    let currentRound = 1;
    
    // 获取第一个和第二个元素用于判断
    let firstElement = arr.shift() || 0;
    
    // 取得的最大值
    let maxNum = firstElement;

    let round = arr.length;

    for(let i=1; i<round; i++){
        let secondElement = arr[i];
        if(maxNum < secondElement){
            maxNum = secondElement;
            currentRound = 1;
        }else{
            currentRound++;
            // 循环次数够了，直接返回最大值
            if(currentRound === k){
                return maxNum;
            }
        }
    }

    // 一轮后的最大值，由于数组中这个值已经是最大的了，因此后续遍历没有意义
    return maxNum;
};