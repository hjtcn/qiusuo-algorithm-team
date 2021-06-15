/*
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
 

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500

*/

/*
    思路：这个思路挺清晰的。
        递归的模版写出来
        let dfs=(sum)=>{
            if(sum=0){
                targetArr.push([...arr])
                return;
            }
            //减去所有可能的值
            //能无限的重复选取
            for(let i=0;i<len;i++){
                arr.push(num[i])
                dfs(sum-num[i])
                arr.pop()
            }
        }
        接下来就是进一步补全思路的坑点：
        1.if(sum<0) return;
        2.因为解集中中不能包含重复的组合
            如2，3，6，7
            选到3的时候就不能再往前选择2了.
         所有dfs函数需要加一个参数，记录当前值选到那个位置了。
         随后在for循环遍历的时候，i的起始位置应该为sIndex
         即：1)for(let i=sIndex;i<len;i++)
            2)dfs(sum-num[i],i)
    题解：官方题解去除了for循环，记录当前位置进行右移。
         此时递归的边界控制也变为了
         if(sIndex>=len) return;
         大致模版：
         let dfs=(sum,sIndex)=>{
             if(sIndex==len){
                 return;
             }
             if(sum==0){
                 targetArr.push([...arr])
                 return;
             }
             //这里意思是跳过选择当前数字
             dfs(sum,sIndex+1)
             if(sum-num[sIndex]>=0){
                 arr.push(num[sIndex])
                 //这个是重复选择当前数字，因此sIndex不变
                 dfs(sum-num[sIndex],sIndex)
                 arr.pop()
             }
         }
    

            
*/

var combinationSum = function(candidates, target) {
    let targetArr=[],len=candidates.length,arr=[]
    let dfs=(sum,sIndex)=>{
        if(sum==0){
            targetArr.push([...arr])
            return;
        }
        if(sum<0){
            return;
        }
        for(let i=sIndex;i<len;i++){
            arr.push(candidates[i])
            dfs(sum-candidates[i],i)
            arr.pop()
        }
    }
    dfs(target,0)
    return targetArr
};


var combinationSum = function(candidates, target) {
    let targetArr=[],len=candidates.length,arr=[]
    let dfs=(sum,sIndex)=>{
        if(sIndex==len){
            return;
        }
        if(sum==0){
            targetArr.push([...arr])
            return;
        }
        dfs(sum,sIndex+1)
        if(sum-candidates[sIndex]>=0){
            arr.push(candidates[sIndex])
            dfs(sum-candidates[sIndex],sIndex)
            arr.pop()
        }
    }
    dfs(target,0)
    return targetArr
}

/*
    时间复杂度：O(S) S为targetArr的长度。官方题解说的是实际运行情况远远小于(N*2^N) 
    空间复杂度：O(target)
*/

/*
    思路：我的思路竟然和01背包时候的目标和对上了，感觉这种题越来越顺畅了。
*/