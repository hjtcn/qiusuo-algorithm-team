/*
718. 最长重复子数组
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

 

示例：

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。
 

提示：

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100

*/

/*
    思路：掉坑里了。
         1.不仅处理了相同的字符。
         dp[i][j]=dp[i-1][j-1]+1
         2.还处理了不同的字符。
          错误代码：dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
          也知道，当不等的时候应该直接断开，但是没想到咋处理。
         一看题解懵逼了，要么不处理，要么就赋值为0就断开了。
         3.也想了滑动弹窗，但是两把尺子同时控制，把我有点吓到了。
    题解：看了代码随想录里面的降维。滚动降维，里面的图还是比较好理解的。
    和上面不同的是遍历方法：
    由于滚动降维需要复用上面一层的dp[j-1]
    因此需要从后往前，避免重复覆盖
*/

var findLength = function(nums1, nums2) {
    let len1=nums1.length,len2=nums2.length
    let dp=Array.from(Array(len1+1),Array(len2+1).fill(0))
    let max=0
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
            if(nums1[i-1]==nums2[j-1]){
                    dp[i][j]=dp[i-1][j-1]+1
            }
            max=Math.max(dp[i][j],max)
        }
    }
    return max
};
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N^2)
*/
//滚动降维：此时就必须处理不等时候赋值为0，因为覆盖一遍，需要重新赋值为0.
var findLength = function(nums1, nums2) {
    let len1=nums1.length,len2=nums2.length
    let dp=Array(len1+1).fill(0)
    let max=0
    for(let i=1;i<=len1;i++){
        for(let j=len2;j>=1;j--){
            if(nums1[i-1]==nums2[j-1]){
                    dp[j]=dp[j-1]+1
            }
            else{
                dp[j]=0
            }
            max=Math.max(dp[j],max)
        }
    }
    return max
};
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N)
*/
