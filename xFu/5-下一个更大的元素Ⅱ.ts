// 503. 下一个更大元素 II
// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

// 示例 1:

// 输入: [1,2,1]
// 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数； 
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
// 注意: 输入数组的长度不会超过 10000。


// 执行结果：
// 通过
// 执行用时：408 ms, 在所有 typescript 提交中击败了16.67%的用户
// 内存消耗：45 MB, 在所有 typescript 提交中击败了16.67%的用户

function nextGreaterElements(nums: number[]): number[] {
    let total = nums.length;
    return nums.map((num, index) => {

        for(let i = index + 1; i !== index; i){

            // 找到一个大的数就直接结束循环
            if(nums[i] > num){
                return nums[i];
            }else{
                // 超出最大长度就回到0,没超过就+1
                if(i >= total){
                    i = 0;
                }else{
                    i++;
                }
            }
        }

        return -1;
    });
};


// 代码虽然写的比较精简，但是执行时间却比写的第一版不精简的版本多了几百毫秒，JS的内部方法比基础方法运行效率差好多
// 下方是用基础for循环写的，同一个逻辑思想，多耗费了0.3mb的内存，执行用时少了170ms左右

// 执行结果：通过
// 执行用时：232 ms, 在所有 typescript 提交中击败了16.67%的用户
// 内存消耗：45.3 MB, 在所有 typescript 提交中击败了16.67%的用户

function nextGreaterElements1(nums: number[]): number[] {
    let list = [];
    let total = nums.length;

    for(let index = 0; index < total;index++){
        let num = nums[index];
        for(let i = index + 1; i !== index; i){

            // 找到一个大的数就直接结束循环
            if(nums[i] > num){
                list[index] = nums[i];
                break;
            }else{
                // 超出最大长度就回到0,没超过就+1
                if(i >= total){
                    i = 0;
                }else{
                    i++;
                }
            }
        }
        if(list[index] === undefined){
            list[index] = -1;
        }
    }

    return list;
};
