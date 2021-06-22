/*
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

 

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 

提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/

/*
    思路：这道题没有成功写出来，全排列写出来了，但是没有去重成功。
        记住了i>0&&nums[i]==nums[i-1]
        但是用的还是不够纯熟，尤其是结合是否使用过，以及递归的深入，就容易混乱。
        去重思路：
        借用slect记录当前数是否已经用过。
        1.用过，停止递归
        2.没用过，但是当前数等于前一个数，则停止递归

         
*/

var permuteUnique = function(nums) {
    let len=nums.length,target=[],select=[]
    nums.sort((a,b)=>a-b)
    let dfs=(arr,idx)=>{
        if(idx==len){
            target.push([...arr])
            return;
        }
        for(let i=0;i<len;i++){
            if(select[i]||(i>0&&nums[i]==nums[i-1]&&!select[i-1])){
                continue
            }
            select[i]=1
            arr.push(nums[i])
            dfs(arr,idx+1)
            select[i]=0
            arr.pop()
        }
    }

    dfs([],0)
    return target
};
/*
    时间复杂度:O(N*N!)
    空间复杂度:O(N)
*/

/*
    虽然没有写出来，但是面对全排列的时候有耐心了很多。
    慢慢成长嘛
*/