给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：
输入：nums = [0]
输出：[[],[0]]

1. Clearfication:
  树的分支：
        选 or 不选
        terminator
        process current logic
        drill down
        reverse current status

2. Coding:
func subsets(nums []int) [][]int {
    paths := [][]int{}
    path := []int{}
    index := 0
    length := len(nums)
    helper(nums,&paths,index,length,path)
    return paths
}
func helper(nums []int,paths *[][]int,index,length int,path []int) {
    // termiantor
    if index == length {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *paths = append(*paths, tmp)
        return
    }
    //选 or 不选
    // 不选
    helper(nums,paths,index + 1,length,path)
    // 选
    helper(nums,paths,index+1,length,append(path, nums[index]))
    // 需不需要 restore current status ? 
}   

3. 看题解：
看了题解里面有二进制的，使用二进制判断序列
没有看懂。。。
mask >> i & 1 > 0?
func subsets(nums []int) (ans [][]int) {
    n := len(nums)
    for mask := 0; mask < 1<<n; mask++ {
        set := []int{}
        for i, v := range nums {
            if mask>>i&1 > 0 {
                set = append(set, v)
            }
        }
        ans = append(ans, append([]int(nil), set...))
    }
    return
}

4. 复杂度分析：
时间复杂度：O(n*2^n): 一共 2^n个状态，每种状态需要O(n) 的时间来构造子集
空间复杂度：O(n)

5. 总结：
5.1: 如果了解整体框架的话不难的，如果不知道思路的话有可能就凉凉了, 学东西也是如此，有些东西不知道学会了，或者请教别人知道了，就可能有思路了，如果自己一直闷头想的话或许就不知道怎么解决了
