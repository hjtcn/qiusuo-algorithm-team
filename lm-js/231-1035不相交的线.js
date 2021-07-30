/*
1035. 不相交的线
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：

 nums1[i] == nums2[j]
且绘制的直线不与任何其他连线（非水平线）相交。
请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

以这种方法绘制线条，并返回可以绘制的最大连线数。

 

示例 1：


输入：nums1 = [1,4,2], nums2 = [1,2,4]
输出：2
解释：可以画出两条不交叉的线，如上图所示。 
但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
示例 2：

输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
输出：3
示例 3：

输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
输出：2
 

提示：

1 <= nums1.length <= 500
1 <= nums2.length <= 500
1 <= nums1[i], nums2[i] <= 2000
 
*/

/*
    思路：觉得这道题可以和周三的最长重复子数组的题放到一起对比着来。
    这道题不相交，就是最长重复子数序列，且不断添加的过程，区别就是断开了，之前的连线数还是要加上的。

    1.dp数组的含义
    dp[i][j]:从0～i，0～j的最大连线数
    2.状态转移方程
    if(nums[i]==nums[j]){
        dp[i][j]=dp[i-1][j-1]+1
    }
    else{
        dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
    }
    3.初始化
    dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    且为了防止数组溢出，依然对比
    nums[i-1]和nums[j-1]
    当然此时遍历的i,j均从1开始遍历
    4.遍历顺序。两个数组比较，无所谓先后。
    5.举例
    
*/

var maxUncrossedLines = function(nums1, nums2) {
    let len1=nums1.length
    let len2=nums2.length
    let dp=Array.from(Array(len1+1),()=>Array(len2+1).fill(0))
    for(let i=1;i<=len1;i++){
        for(let j=1;j<=len2;j++){
            if(nums1[i-1]==nums2[j-1]){
                dp[i][j]=dp[i-1][j-1]+1
            }
            else{
                dp[i][j]=Math.max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[len1][len2]
};
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N^2)
*/

//降维滚动优化，重点还是在于dp[i-1][j-1]如何表示出来
// dp[i-1][j-1]   dp[j]
// dp[j-1]        目标dp[j]
 var maxUncrossedLines = function(nums1, nums2) {
    let len1=nums1.length
    let len2=nums2.length
    let dp=Array(len2+1).fill(0)
    for(let i=1;i<=len1;i++){
        let cur=0
        for(let j=1;j<=len2;j++){
            let prev=cur
            cur=dp[j]
            if(nums1[i-1]==nums2[j-1]){
                dp[j]=prev+1
            }
            else{
                dp[j]=Math.max(dp[j],dp[j-1])
            }
        }
    }
    return dp[len2]
};
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N)
*/

/*  
    熟能生巧
*/