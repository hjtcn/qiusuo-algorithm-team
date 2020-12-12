// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

// 示例:

// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]


// ============================================================
// ===                                                      ===
// ===      状态：通过,执行用时: 88 ms,内存消耗: 41.9 MB     ===
// ============================================================

function permute(nums: number[]): number[][] {
    let result:number[][] = [];

    function todo(newList){

        // 遍历数组，找到没有处理过的数，依次递归处理
        nums.forEach(num => {
            if(newList.indexOf(num) > -1) return;
            todo([ ...newList, num ]);
        });

        // 过滤掉不正确的组合
        if(newList.length === nums.length) {
            result.push(newList);
        }

    }

    // 从每一个数据开始一次递归
    nums.forEach(num => {
        todo([num]);
    })
    
    return result;

};