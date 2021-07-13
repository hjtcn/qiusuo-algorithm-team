/*
377. 组合总和 Ⅳ
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。

 

示例 1：

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
示例 2：

输入：nums = [9], target = 3
输出：0
 

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000
 

进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？



*/


/*
    思路:想想排列组合有点遗忘了，就写了一下，不出意料超时了。
        每次一剪枝就会卡壳，知道直接递归返回目标组合数，但是记录过程还是有问题
    题解：递归剪枝
    let dfs(sum)=>
    {
        if(sum<0) return 0
        if(sum==0) return 1
        if(dp[sum]) return dp[sum]
        let res
        for(let i=0;i<len;i++){
            res+=dfs(sum-nums[i])
        } 
        dp[sum]=res
        return res
    }
    一开始，我是拿数组记录的，还是超时，因为查询dp[sum]是否为空的过程也是一遍遍历过程

    看题解，还是应该使用map.has(sum)

    动态规划：看完题解，再一次感受到遍历顺序的重要性

    1.dp数组的意义
    dp[j]代表和为j时的组合个数
    2.状态转移方程
    dp[j]+=dp[j-nums[i]]
    3.dp数组初始化
    dp[0]=1
    4.遍历顺序
    先遍历和，从1到target
    再遍历整数数组，从0到len
    为什么遍历整数数组放在内层？
    因为这是求排列数，顺序不同，也视作不同的组合。

    如果求组合数就是外层for循环遍历物品，内层for遍历背包。
    如果求排列数就是外层for遍历背包，内层for循环遍历物品。

    5.举例

*/

//排列组合
var combinationSum4 = function(nums, target) {
    let len=nums.length
   let count=0
    let dfs=(sum,index,arr)=>{
        if(sum<0||index>len){
            return 
        }
        if(sum==0){
            count++;
            return ;
        }
        for(let i=0;i<len;i++){
            dfs(sum-nums[i],i+1,[...arr,nums[i]])
        }
    }

    dfs(target,0,[])
    return count
};

var combinationSum4 = function(nums, target) {
    let len=nums.length
   let map=new Map()
    let dfs=(sum)=>{
        if(sum<0){
            return 0
        }
        if(sum==0){
            return 1;
        }
        if(map.has(sum)) return map.get(sum)
        let res=0
        for(let i=0;i<len;i++){
            res+=dfs(sum-nums[i])
        }
        map.set(sum,res)
        return res
    }
    return dfs(target)
};


var combinationSum4 = function(nums, target) {
    let len=nums.length
    let dp=Array(target+1).fill(0)
    dp[0]=1
    for(let i=1;i<=target;i++){
        for(let j=0;j<len;j++){
            if(i-nums[j]>=0){
                 dp[i]+=dp[i-nums[j]]
            }
        }
    }
    return dp[target]
};

/*
    时间复杂度：O(N*sum)
    空间复杂度：O(sum)
*/

/*
    完全背包还不成概念。
    但是这两天的题感受到遍历顺序的区别了
*/

