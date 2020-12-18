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

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function(nums) {
    let res=[]
    let helper=(arr)=>{
        if(nums.length==arr.length){
            //如果直接push arr,arr都是空数组,而slice返回新数组
            res.push(arr.slice())
            return;
        }
        for(let i=0;i<nums.length;i++){
            if(arr.indexOf(nums[i])==-1)
            {
                //核心。标记后清除标记
                arr.push(nums[i])
                helper(arr)
                arr.pop()
            }
        }
    }
    helper([])
    return res
};


//回溯。回溯的核心即走的位置添加标记，发现路走不通，清除标记，然后从上一步试着发散找路。
//讲真，还是比较怵这种类型题，走着走着发现走不通了，都有点绕，只能多刷几遍，熟能生巧
// 时间复杂度：O(N^2)
// 一层for循环以及js封装的数组查询api，最差时间复杂度为O(N^2)
// 空间复杂度：O(N)
// 主要是结果数组和递归函数中转换的arr数组