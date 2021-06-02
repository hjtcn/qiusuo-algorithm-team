/*
55. 跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

 

示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 

提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105

*/

/*
    思路：更新能跨越的最远距离，如果下标超出最远距离，说明走不到这个下标处
*/

var canJump = function(nums) {
    let len=nums.length,max=0;
    for(let i=0;i<len;i++){
        if(i>max){
            return false
        }
        max=Math.max(max,i+nums[i])
    }
    return true
}

var canJump = function(nums) {
    let len=nums.length,count=0,flag=true;
    for(let i=len-1;i>=0;i--){
        if(nums[i]>=count){
            flag=true
            count=1
        }
        else{
            flag=false
            count++;
        }
    }
    return flag
}
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：记忆点和灵感真的很神奇，这次上来就知道先去更新最远距离。
    当初想到的倒序，反而很多细节控制都没弄好，可能是脑子里想的多了。
*/