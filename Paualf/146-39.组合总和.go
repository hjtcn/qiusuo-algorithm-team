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

1. Clearfication:
     target
            可以无限制的重复的使用
        [2,3,6,7] => target : 7
        分解子问题：
            target = 7
            你要去找 7 - 2 = 5,这个时候 target 切换成5，然后在 [2,3,6,7] 找出=5的组合
            最后放入到结果数组中的时候，需要进行去重，对 path进行 mapKey 一下，然后判断是否存在
        忽然感觉分解子问题是一个逆人性的思维

2. Coding:
func combinationSum(candidates []int, target int) [][]int {
    paths := [][]int{}
    path := []int{}
    //sort.Ints(candidates)
    m := make(map[string]int,0)
    helper(candidates, target,&paths,path,m)
    
    return paths
}

func helper(candidates []int, target int,paths *[][]int,path []int,m map[string]int) {
    // terminator
    if target == 0 {
        
    tmp := make([]int,len(path))
        copy(tmp,path)
        sort.Ints(tmp)
        mapKey := getMapKey(tmp)
        //fmt.Println(mapKey)
        
        if _,ok := m[mapKey];ok {
            return
        }
        m[mapKey] = 1
      
        *paths = append(*paths, tmp)
        return
    }

    // process current logic
    for i := 0;i < len(candidates);i++ {
        if target - candidates[i] >= 0{
            helper(candidates,target - candidates[i],paths,append(path, candidates[i]),m)
        }
    }
    // dirll down
}

func getMapKey(path []int) string {
    res := ""
    
    for _,v := range path {
        res += strconv.Itoa(v)
    }
    return res
}

3. 看题解：
看了题解的这个更像是回溯，感觉自己写的是纯递归
func combinationSum(candidates []int, target int) (ans [][]int) {
    comb := []int{}
    var dfs func(target, idx int)
    dfs = func(target, idx int) {
        if idx == len(candidates) {
            return
        }
        if target == 0 {
            ans = append(ans, append([]int(nil), comb...))
            return
        }
        // 直接跳过
        dfs(target, idx+1)
        // 选择当前数
        if target-candidates[idx] >= 0 {
            comb = append(comb, candidates[idx])
            dfs(target-candidates[idx], idx)
            comb = comb[:len(comb)-1]
        }
    }
    dfs(target, 0)
    return
}

4. 复杂度分析：
时间复杂度：O(n * n ^k）
空间复杂度：O(n)

5. 总结：
5.1: 将问题分解子问题的这个思路，感觉有点逆人性的感觉，人本身是想要顺着思路想问题，但是如果你想解决这种复杂的问题，你就要去想办法，去把复杂的问题去进行分解，这个时候就需要你的注意力集中了
