/*
416. 分割等和子集
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

 

示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
 

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

/*
    思路：
    01背包：脑子还有着昨天目标和01背包的记忆，今天的题就不是那么难了。
    首先和要分为两个相等的子集。
    因此可以判断数组和是不是一个偶数。
    奇数直接返回false
    偶数则可以sum/2求出目标和target
    再对比01背包，求出元素中是否存在目标和target.
    for(let i=0;i<len;i++){
        for(let j=target;j>=nums[i];j--){
            dp[j]=dp[j]+dp[j-nums[i]]
        }
    }

    题解：
    递归：枚举路径中是否存在sum/2的目标和
    边界判断:
    if(i==nums.length||curSum>target)
    return false
    if(curSum==target)
    return true
    //加或不加
    return dfs(i+1,curSum+nums[i])||dfs(i+1,curSum)

    仅仅这些，最后会超时，因此需要记忆化递归
    if(map.has(key)){
        return map.get(key)
    }
    let res=dfs(i+1,curSum+nums[i])||dfs(i+1,curSum)
    map.set(key,res)
    return res
    
*/


var canPartition = function (nums) {
    let len = nums.length, sum = 0;
    for (let i = 0; i < len; i++) {
        sum += nums[i]
    }
    if (sum % 2) return false
    let target = sum / 2
    let dp = Array(target + 1).fill(0)
    dp[0] = 1
    for (let i = 0; i < len; i++) {
        for (let j = target; j >= nums[i]; j--) {
            dp[j] = dp[j] + dp[j - nums[i]]
        }
    }
    return dp[target]
};


var canPartition = function (nums) {
    let len = nums.length, sum = 0;
    for (let i = 0; i < len; i++) {
        sum += nums[i]
    }
    if (sum % 2) return false
    let target = sum / 2,map=new Map()
    let dfs = (i, curSum) => {
        if(i==nums.length||curSum>target){
            return false
        }
        if (curSum == target) {
            return true
        }
        //记忆化：
        let key=curSum+'&'+i
        if(map.has(key)){
            return map.get(key)
        }
       let res= dfs(i + 1, curSum + nums[i])||dfs(i + 1, curSum)
       map.set(key,res)
        return res
    }
    return dfs(0,0)
};


/*
    时间复杂度：O(n*target)
    空间复杂度：O(target)
*/

/*
    思考：递归要考虑边界，剪枝，记忆化。
         01背包要进行变形，向模版靠拢。
*/