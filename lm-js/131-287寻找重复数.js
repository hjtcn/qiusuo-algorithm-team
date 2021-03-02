// 287. 寻找重复数
// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

// 假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

 

// 示例 1：

// 输入：nums = [1,3,4,2,2]
// 输出：2
// 示例 2：

// 输入：nums = [3,1,3,4,2]
// 输出：3
// 示例 3：

// 输入：nums = [1,1]
// 输出：1
// 示例 4：

// 输入：nums = [1,1,2]
// 输出：1
 

// 提示：

// 2 <= n <= 3 * 104
// nums.length == n + 1
// 1 <= nums[i] <= n
// nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
 

// 进阶：

// 如何证明 nums 中至少存在一个重复的数字?
// 你可以在不修改数组 nums 的情况下解决这个问题吗？
// 你可以只用常量级 O(1) 的额外空间解决这个问题吗？
// 你可以设计一个时间复杂度小于 O(n2) 的解决方案吗？
/*
    思路：如果是根据之前的思路：我会使用map，来记录出现次数。空间复杂度为O(N)
    找出重复数,最简单的遍历一次，则时间复杂度为O(N)
    由于我们为了攻克二分，我想到了时间复杂度的优化，可以优化为O(logN)
    AC过后，我看到了进阶思考。
    如何空间复杂度为O(1),时间复杂度为O(N^2)以下？
    好吧！没有想出来，去扒了扒题解，很巧妙，根据该题特殊的数据类型可考虑的思路。
    [1,2,3,4] 正常排列，没有重复  <=当前元素的个数为下标
    如[1,3,5,2,6,2,4]   
    第一次二分：mid=3,num(小于中间索引的值的个数)=4
              mid<num,r=mid-1=2,l=0，res=3
    第二次二分：mid=1,num=1
              mid==num,l=mid+1=2,r=2
    第三次二分：mid=2,num=3
              mid<num,r=1,l=2,res=2
    二分结束.res=2


    好吧，我虽然模拟了很多遍，还是不是很理解下标的个数的特殊性。。。。



*/

/**
 * @param {number[]} nums
 * @return {number}
 */
//这是自己写的方法。
var findDuplicate = function(nums) {
    let r=nums.length-1,hasNum=0,map=new Map(),res=0
    let dfs=(l,r)=>{
        if(l<=r)
        {let mid=parseInt((l+r)/2)
        hasNum=map.get(nums[mid])
        if(!hasNum){
            map.set(nums[mid],1)
            dfs(l,mid-1)
            dfs(mid+1,r)
        }
        else{
            res=nums[mid]
        }
        }
    }
    dfs(0,r)
    return res
};


//这是看题解get的方法，还是不是很懂。。。。
var findDuplicate = function(nums) {
    let l=0,r=nums.length-1,res
    while(l<=r){
        let mid=parseInt((l+r)/2)
        let num=0
        for(let i=0;i<nums.length;i++){
            if(nums[i]<=mid){
                //记录个数
                num+=1
            }

        }
        if(num<=mid){//说明目标在右边
            l=mid+1
        }
        else{ //目标在左边
            r=mid-1
            res=mid
        }
    }
    return res
}

  