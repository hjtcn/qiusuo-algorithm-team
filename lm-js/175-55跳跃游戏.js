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
    思路：自己有两个思路。
        方法1.倒着走，count记录走到当前位置需要几步，当前值是否足够走，如果够，则flag=true;如果不够，则flag=false,用count++,计算下一个位置是否足够走。
        遍历结束，没有找到足够走的，返回false,第一个元素足够走，就返回true.

        方法2.记录0的出现次数count，因为只有0的时候走不动，跳不到下一步.
        1)如果非0，更新count=max(count,nums[i])-1.（减一是因为走下一步呀）
        2)如果为0，更新count--;
        如果count<0,说明跨不过去了，返回false
        遍历结束，返回true
        方法2遇到的坑，差点都放弃这个思路了，刚开始没想到要更新最大值，果然还是不够贪心呀，哈哈哈哈

    题解：拿下标i来更新能跳的最远的位置k
        k=max(k,i+nums[i])

        看完题解又想起方法2的一个点,不关注最后一个值为多少了，因此length-1.
        如果for(let i=0;i<len;i++),[0]就不合理了。
*/


var canJump = function(nums) {
    let len=nums.length
    let count=0,flag=true
    for(let i=len-1;i>=0;i--){
        if(nums[i]>=count){
            flag=true
            count=1
        }
        else{
            count++;
            flag=false
        }
    }
    return flag
};

var canJump = function(nums) {
    let len=nums.length,count=0
    for(let i=0;i<len-1;i++){
        if(nums[i]){
            count=Math.max(count,nums[i])-1
        }
        else{
            count--;
        }
        if(count<0){
            return false
        }
    }
    return true
}


var canJump = function(nums) {
    let len=nums.length,k=0
    for(let i=0;i<len;i++){
        if(i>k) return false
        k=Math.max(k,i+nums[i])
    }
    return true
}
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    写了快两周的贪心题了，好难系统化呀，我们还是没有找到模型，是不是可以从书上借鉴借鉴题目类型。
*/