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

1. Clearfication:
  重复的子集，如何进行去重？
       
        值构建一个唯一的 hash map的key，然后判断这个key是否存在？
        对slice进行md5 然后判断这个md5后的key是否存在，如果不存在则加入，存在则contine
        写着写着想到为什么slice不能做string的key呢？ 真的需要md5么，如果sort以后，排序一下，然后将 slice格式化为字符串然后不就行了？
        为什么不用slice作为map的key，因为slice是不可以比较的哈

2. Coding
func subsetsWithDup(nums []int) [][]int {
    paths := [][]int{}
    path := []int{}
    start := 0
    end := len(nums)
    sort.Ints(nums)
    m := make(map[string]int, end)
    helper(nums,&paths,path,start,end,m)
    return paths
}
func helper(nums []int,paths *[][]int,path []int,start,end int,m map[string]int) {
    // terminator
    if start == end {
        mapKey := getMapKey(path)
        //fmt.println(path)
        //fmt.Println(mapKey)
        if _,ok := m[mapKey];ok {
            return    
        }
        m[mapKey] = 1
        tmp := make([]int, len(path))
        copy(tmp,path)
        *paths = append(*paths, tmp)
        return
    }
    // process current logic
    // 不选
    helper(nums,paths,path,start + 1,end,m)
    // 选
    helper(nums,paths,append(path,nums[start]),start + 1,end,m)
}
func getMapKey(path []int) string {
    res := ""
    
    for _,v := range path {
        res = res + strconv.Itoa(v)
    }
    return res
}

3. 看题解：
递归的时候，发现如果没有选择上一个数，且当前数字与上一个数相同，则可以跳过当前生成的子集s
使用一个变量表示加判断进行剪枝
if !choosePre && cur > 0 && nums[cur - 1] == nums[cur] {
    return
}
迭代思维：
func makeSubsets(nums []int, result *[][]int, start int, subset []int) {
    *result = append(*result, subset)
    
    for i := start; i < len(nums); i++ {
        if i == start || nums[i] != nums[i-1] {
            nextSubset := append([]int{}, subset...)
            makeSubsets(nums, result, i+1, append(nextSubset, nums[i]))
        }
    }
}

func subsetsWithDup(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    makeSubsets(nums, &result, 0, []int{})
    
    return result
}

4. 复杂度分析：
时间复杂度：O(n*2^n) n个节点，每个节点都可以选or不选
空间复杂度：O(n)：递归栈空间

5. 总结：
5.1: 一开始也没有思路，怎么去重呢？ 常用的方法是hash，hash的话怎么防止mapKey冲突呢？想到了 md5,md5有碰撞，不是很好，有没有更好的方法， 自己去构造string,怎么保证构造的string key 不重复，可以对切片一开始进行排序

5.2: 解决问题和做事的时候会遇到各种各样的问题，运用我们的智慧去解决它
