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

1. Clearfication:
第一道题是可以重复使用元素，这道题，每个元素组合中只能使用一次，所以需要加一个map判断这个元素是否已经使用过

2. Coding:
递归的时候有些地方，选 or 不选的时候没有想清楚，漏掉了一些情况
func combinationSum2(candidates []int, target int) [][]int {
    paths := [][]int{}
    path := []int{}
    m := make(map[string]int)
    hasUsed := make(map[int]int)
    helper(candidates,&paths,path,target,0,m,hasUsed)
    return paths
}

func helper(candidates []int,paths *[][]int,path []int,target,index int,m map[string]int,hasUsed map[int]int) {
    if index >= len(candidates) {
        return
    }
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp,path)
        sort.Ints(tmp)
        mapKey := getMapKey(tmp)
        if _,ok := m[mapKey];ok {
            return
        }
        m[mapKey] = 1
        fmt.Println(tmp)
        *paths = append(*paths, tmp)
        return
    }

    for i := index;i < len(candidates);i++ {
        
        // 选，首先判断是否使用过，如果使用过的话就不能再次使用了
        if _,ok := hasUsed[i];ok {
            continue
        }
        if target - candidates[i] >= 0 {
            fmt.Println(path)
            hasUsed[i] = 1
            path = append(path, candidates[i])
            helper(candidates, paths,append(path, candidates[i]),target - candidates[i],index + 1,m,hasUsed)
            path = path[:len(path) -1]
            hasUsed[i] = 0
            fmt.Println(path)
              // 不选index i
            helper(candidates,paths,path,target,index + 1,m,hasUsed)
        }
    }
}
func getMapKey(path []int) string {
    res := ""
    for _,v := range path {
        res += strconv.Itoa(v)
    }
    return res
}

3. 看题解：
i - 1 >= start && candidates[i - 1] == candidates[i] 这一句我没有看懂。。。
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := [][]int{}
    var dfs func(start int, temp []int, sum int)
    dfs = func(start int, temp []int, sum int) {
        if sum >= target {
            if sum == target {
                t := make([]int, len(temp))
                copy(t, temp)
                res = append(res, t)
            }
            return
        }
        for i := start; i < len(candidates); i++ {
            if i-1 >= start && candidates[i-1] == candidates[i] {
                continue
            }
            temp = append(temp, candidates[i])
            dfs(i+1, temp, sum+candidates[i])
            temp = temp[:len(temp)-1]
        }
    }
    dfs(0, []int{}, 0)
    return res
}

看题解思路挺清晰的，应该是那里我没有想清楚。。。
var res [][]int
var candis []int
func combinationSum2(candidates []int, target int) [][]int {
    res = nil
    candis = candidates
    sort.Ints(candidates)
    backtrack(target, nil, 0)
    return res
}

func backtrack(target int, cur []int, index int) {
    if target == 0 {
        res = append(res, append([]int{}, cur...))
        return
    }
    if target < 0 {
        return
    }
    for i := index; i < len(candis); i++ {
        if i > index && candis[i] == candis[i-1] {
            continue 
        }
        backtrack(target-candis[i], append(cur, candis[i]), i+1)
    }
}

4. 复杂度分析：
时间复杂度：O(n*2^n)
空间复杂度：O(n) : 递归深度

5. 总结：
5.1: 感觉算法和数据结构还是需要理论知识的，理论知识是复杂和枯燥的，同时这些也是故去大神智慧的结晶和总结。。。
