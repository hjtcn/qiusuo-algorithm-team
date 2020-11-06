// 剑指 Offer 59 - I. 滑动窗口的最大值
// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

// 示例:

// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7] 
// 解释: 

//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
 
// 提示：

// 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。


/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var maxSlidingWindow = function(nums, k) {
    if(nums.length<=1){
        return nums
    }
    let res=[]
    for(let i=0;i<=nums.length-k;i++){
            res.push(Math.max(...nums.slice(i,i+k)))
    }
    return res
};

/*
 暴力解决。使用了js的特定api.
 首先解构赋值将nums数组截取出相应大小的窗口，并转化为以','分割的字符串
 然后借用Math.max方法获取最大值
  时间复杂度：O(N*k)
  一层for循环，对比长度为k的窗口值的大小
  空间复杂度：O(1)
  今天才get结果数组，不算在空间复杂度之中
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var maxSlidingWindow = function(nums, k) {
    if(!nums||!nums.length||k<1){
        return []
    }
    if(k==1){
        return nums
    }
    let queue=[]//存索引喽
    let res=[]
    for(let i=0;i<nums.length;i++){
        if(i>=k){
            //要跳出窗口了，此时队列中的最大值失效，被shift掉。
            let outNode=nums[i-k]
            if(outNode==queue[0]){
                queue.shift()
            }
        }
        let curNode=nums[i]
        //  循环判断保持队列单调递减
        while(queue.length&&queue[queue.length-1]<curNode){
            queue.pop()
        }
        queue.push(curNode)
        //第k个值开始赋值结果数组
        if(i>=k-1){
            res.push(queue[0])
        }
    }
    return res
};

/*
 单调队列，这题卡了蛮久，没有思路，后来看着题解，走着逻辑，才慢慢理解的。
 保持队列单调递减。如果滑动窗口的前一个值和队列头相等，则说明跳出这个窗口，这个最大值马上要失效了。
 将队列进行shift操作，更新最大值了。
 而保持单调递减的方法，就是循环比较队列尾部和当前节点，如果当前节点大，则队列pop，直到队列为空，将当前节点push到队列中
 从第k个数开始赋值结果数组了，每次循环结尾，将队列的头部返回给结果数组即可

  时间复杂度：O(N)
  一层for循环
  空间复杂度：O(k)
  借用数组模拟队列,数组长度不会超过k的值
*/