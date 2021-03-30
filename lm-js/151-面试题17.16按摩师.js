// 面试题 17.16. 按摩师
// 一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

// 注意：本题相对原题稍作改动

 

// 示例 1：

// 输入： [1,2,3,1]
// 输出： 4
// 解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
// 示例 2：

// 输入： [2,7,9,3,1]
// 输出： 12
// 解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
// 示例 3：

// 输入： [2,1,4,5,3,1,1,3]
// 输出： 12
// 解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。



/*
    思路：选一个用例，开始画状态转移方程
         2  1  4  5  3  1  1  3
 总时长   2
            2   
               6   
                  7  
                     9   
                        9  
                           10
                              12
    考虑：f(x)=Math.max(f(x-2)+nums[i],f(x-1))的方式去更新最长总时长
    静下心来找状态转移方程的规律，也不难嘛，膨胀膨胀！！！继续加油，坚持
*/


var massage = function(nums) {
    let dp=[],max=0,len=nums.length
    //由于状态方程中有dp(x-2)，因此我提前作了数组长度的判断。
    if(!len)
     return 0
    if(len==1)
     return nums[0]
    if(len==2)
     return Math.max(nums[0],nums[1])
    
    dp[0]=nums[0],dp[1]=Math.max(nums[0],nums[1])
    for(let i=2;i<nums.length;i++){
        dp[i]=Math.max(dp[i-2]+nums[i],dp[i-1])
    }
    return dp[len-1]
};


/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/


//进一步优化可以不用dp存储所有曾经出现的最大值了。
//用prev存储前两位的值，将空间复杂度从O(N)将为O(1)
var massage = function(nums) {
    let len=nums.length
    //由于状态方程中有dp(x-2)，因此我提前作了数组长度的判断。
    if(!len)
     return 0
    if(len==1)
     return nums[0]
    if(len==2)
     return Math.max(nums[0],nums[1])
    
    let prev=nums[0],cur=Math.max(prev,nums[1])
    for(let i=2;i<nums.length;i++){
        let tmp=cur
        cur=Math.max(prev+nums[i],cur)
        prev=tmp
    }
    return cur
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/