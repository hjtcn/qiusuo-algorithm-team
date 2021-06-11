/*
90. 子集 II
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

 

示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
 

提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10

*/

/*
    思路：去重。
        鉴于昨天的题是无重复的所有子集
        我就先把架子打好了。可选可不选，筛选出所有的可能。
        当然不同点也在于去重，刚开始想把所有可能在再次筛选，筛元素好筛，但是筛数组，，，emmmmmm,每一个数组都是不相等的。
        看题解
    题解：剪枝。1.原数组排序
              2.找规律：如[1,2,2]
              选第1，2数和选第1，3会得到相同的子集
             即，如果当前数(2)和前一个数相同(2),
             同时没有选择前一个数(2)的集合中，接下来只要一选择当前数(2)就会导致重复。
             结论：当前值和前面的数相等且前一个数没被选择，则跳过选择这一步。

*/
var subsetsWithDup = function(nums) {
    let target=[],arr=[],len=nums.length
    nums.sort((a,b)=>a-b)
    let dfs=(isChoose,i)=>{
        if(i==len){         
           target.push([...arr])
            return;
        }
        if(!isChoose&&i>0&&nums[i-1]==nums[i]){
            //只继续走不选择当前值的递归。
            dfs(false,i+1)
            return;
        }
        arr.push(nums[i])
        dfs(true,i+1)
        arr.pop()
        dfs(false,i+1)
    }
    dfs(false,0)
    return target
};

/*
    时间复杂度:O(N*2^N)
    空间复杂度:O(N)
*/

/*
    递归中的剪枝操作还是不够深入
*/