// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

 

// 示例 1：

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2：

// 输入：nums = [1]
// 输出：1
// 示例 3：

// 输入：nums = [0]
// 输出：0
// 示例 4：

// 输入：nums = [-1]
// 输出：-1
// 示例 5：

// 输入：nums = [-100000]
// 输出：-100000
 

// 提示：

// 1 <= nums.length <= 3 * 104
// -105 <= nums[i] <= 105
 

// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

/*
    思路。贪心：如果当前和>0，则值得加之前的值。
                 当前和<0,则直接等于当前值即可
             在遍历中不断更新最大和
        暴力：两层循环。
            for(let i=0;i<len;i++){
                let sum=0
                for(let j=i;j<len;j++){
                    //算出每一个区间的和，并更新最大和
                    sum+=nums[i]
                    max=Math.max(sum,max)
                }
            }

*/

 //贪心
 var maxSubArray = function(nums) {
    let sum=0
    let max=nums[0]
    for(let i=0;i<nums.length;i++){
        if(sum>0){
            sum+=nums[i]
        }
        else{
            sum=nums[i]
        }
        max=Math.max(max,sum)
    }
    return max
};
/*
时间复杂度：O(n)
空间复杂度：O(1)
*/
// @lc code=end
//暴力
var maxSubArray = function(nums) {
    let max=-Infinity
    for(let i=0;i<nums.length;i++){
        sum=0;
        for(let j=i;j<nums.length;j++){
            sum+=nums[j]
            if(sum>max){
                max=sum
            }
        }
    }
    return max
}
/*
    时间复杂度：O(N^2)
    空间复杂度：O(1)
*/


//动态规划。dp(x-1)的更新也是前一位的最大和。因此其求dp(x)?// 当前的的最大和dp，可能等于前一位dp(x-1)加上当前值或直接等于当前值，谁大等于谁，哈哈哈。
// 状态转移方程：dp(x)=Math.max(dp(x-1)+nums[i],nums[i]).
var maxSubArray = function(nums) {
    let dp=0,max=-Infinity
    for(let i=0;i<nums.length;i++){
        //核心方程
        dp=Math.max(dp+nums[i],nums[i])
        max=Math.max(dp,max)
    }
    return max
}
/*
时间复杂度：O(n)
空间复杂度：O(1)
*/


//分治：线段树。临摹了一边没错哦，不太好理解，跟着思路和代码，走了一边，pushUp有点懵。
function Status(l,r,m,i){
    this.lSum=l
    this.rSum=r
    this.mSum=m
    this.iSum=i
}
const pushUp=(l,r)=>{
    const iSum=l.iSum+r.iSum
    const lSum=Math.max(l.lSum,r.lSum+l.iSum)
    const rSum=Math.max(r.rSum,l.rSum+r.iSum)
    const mSum=Math.max(Math.max(l.mSum,r.mSum),l.rSum+r.lSum)
    return new Status(lSum,rSum,mSum,iSum)
}
function getInfo(a,l,r){
    if(l==r){
        return new Status(a[l],a[l],a[l],a[l])
    }
    const mid=(l+r)>>>1
    const lSub=getInfo(a,l,mid),rSub=getInfo(a,mid+1,r)
    return pushUp(lSub,rSub)
}
var maxSubArray = function(nums) {
    return getInfo(nums,0,nums.length-1).mSum
}

/*
时间复杂度：O(n)
空间复杂度：O(nlogn)
*/


/*
 锻炼构建状态转移的过程，以小见大
 我确实出的是简单题，哈哈哈，继续加油。。。。。，
*/