/*
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/

/*
    思路：先排所有的组合可能
        然后剪枝去重。
        子集II去重的思路还没有忘记：
        1.先将原数组排序。
        2.如果两个数相同，且前一个数没选过，则跳过当前数值进行递归。

        两种写法的重点难点：
        1.当递归里面没有for循环，完全根据sIndex+1进行右移时：
        即可根据两个数相同，且前一个数没被选择，跳过该数进行递归。
        2.当递归里面有for循环.
        即可根据两个数相同，且j!=起始索引sIndex,可跳过bfs()
        原理：for循环不是第一次循环，并且当前元素和前一个元素相同，跳过。即回溯之后跳过相同项


*/

var combinationSum2 = function(candidates, target) {
    let arr=[],targetArr=[],len=candidates.length
    candidates.sort((a,b)=>a-b)
    let dfs=(sum,sIndex,isSelected)=>{
        if(sum<0||sIndex>len){
            return;
        }
        if(sum==0){
            targetArr.push([...arr])
            return;
        }
        if(!isSelected&&sIndex>0&&candidates[sIndex-1]==candidates[sIndex]){
            dfs(sum,sIndex+1,false)
            return;
        }
            arr.push(candidates[sIndex])
            dfs(sum-candidates[sIndex],sIndex+1,true)
            arr.pop()
            dfs(sum,sIndex+1,false)
    }
    dfs(target,0,false)
    return targetArr
};
// @lc code=end

var combinationSum2 = function(candidates, target) {
    let arr=[],targetArr=[],len=candidates.length
    candidates.sort((a,b)=>a-b)
    let dfs=(sum,sIndex)=>{
        if(sum<0){
            return;
        }
        if(sum==0){
            targetArr.push([...arr])
            return;
        }
        for(let j=sIndex;j<len;j++){
            if(j!=sIndex&&j>0&&candidates[j-1]==candidates[j]){
                continue
            }
            arr.push(candidates[j])
            dfs(sum-candidates[j],j+1)
            arr.pop()
        }
    }
    dfs(target,0)
    return targetArr
};

/*
    时间复杂度:O(N*2^N)
    空间复杂度:O(N)
*/

/*
    看到有个题解，说这题和15题三数之和(当初不会的面试题)题目类型类似，
    不知不觉我们已经可以做这种题了哦
*/