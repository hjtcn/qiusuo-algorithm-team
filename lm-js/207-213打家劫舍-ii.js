/*
213. 打家劫舍 II
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

 

示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：

输入：nums = [0]
输出：0
 

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/

/*
    思路：鉴于昨天的题，可以知道偷盗不相邻的房屋。状态转移方程为:dp[i]=Math.max(dp[i-1],dp[i-2]+nums[i])
    而今天升级了，第一个和最后一个是相邻的。
    那怎么处理呢？看了题解了解到：
    第一个和最后一个是互斥的。
    因此求两个偷盗范围(0,len-1)和（1,len)的最大值。

    封装求dp[len-1]最大值的函数。
    分别调用。

    降维还是和之前一样
    prev1代表dp[i-2]
    prev2代表dp[i-1]
    遍历更新这两个值,遍历结束后返回prev2

*/
var rob = function(nums) {
    let len=nums.length
    if(len<2)
    return nums[0]
    if(len==2) return Math.max(nums[0],nums[1])  
    let getMax=(arr)=>{
        let dp=[],len2=arr.length
        if(len2.length<2) return arr[0]
        dp[0]=arr[0]
        dp[1]=Math.max(arr[0],arr[1])
        for(let i=2;i<len2;i++){
            dp[i]=Math.max(dp[i-1],dp[i-2]+arr[i])
        }
        return dp[len2-1]
    }
    let max1=getMax(nums.slice(1))
    let max2=getMax(nums.slice(0,len-1))
    return Math.max(max1,max2)
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
// 降维 
var rob = function(nums) {
    let len=nums.length
    if(len<2)
    return nums[0]
    if(len==2) return Math.max(nums[0],nums[1])  
    let getMax=(i,j)=>{
        if(j-i<2) return nums[i]
        let prev1=nums[i]
        let prev2=Math.max(nums[i],nums[i+1])
        for(let x=i+2;x<j;x++){
            let kk=prev2
            prev2=Math.max(prev2,prev1+nums[x])
            prev1=kk

        }
        return prev2
    }
    let max1=getMax(1,len)
    let max2=getMax(0,len-1)
    return Math.max(max1,max2)
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/


/*
    思考：之前这个思路一直琢磨不透。
        比如提前去除第一个房屋，寻找(2,n)的房屋最大值。
        那要是(2,n)没选中n的话，第一个房屋可不可以加上呢？为什么这个思路就毫无破绽了呢？
        后来想通了，如果在(2,n)范围内没选中n,说明n的价值不够，不值得选。那么(1,n-1)中肯定能够找到相对比较大的金额。

        (2,n)和（1,n-1)两种互斥的情况，反而将值得选1，还是值得选n，甚至两种都不值得选的情况都给处理过了。

*/