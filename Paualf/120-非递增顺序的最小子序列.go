1. Clearfication:
感觉这道题条件挺多的：
抽取一个子序列，满足该子序列的元素之和 严格 大于未包含在该子序列中各元素之和
Q: 怎么找到这个子序列呢？
A:
Q: 子序列需要长度最小
A:
Q: 子序列降序排列
A:

2. Coding:

3. 看题解：
排序后的贪心
func minSubsequence(nums []int) []int {
    Max := 0
    n := 0
    num := make([]int,0)
    for i := 0;i < len(nums);i++ {
        for j := i + 1;j < len(nums);j++ {
            if nums[i] > nums[j] {
                nums[i],nums[j] = nums[j],nums[i]
            }
        }
        Max += nums[i]
    }
    
    for i := len(nums) - 1;i >= 0;i++ {
        n += nums[i]
        Max -= nums[i]
        
        num = append(num, nums[i])
        if n > Max {
            return num
        }
    }   
    
    return num
}

func minSubsequence(nums []int)[]int {
    var res[]int
    sort.Ints(nums)
    
    sum := 0
    for _,n := range nums {
        sum += n
    }
    counter := 0
    i := len(nums)
    half := sum / 2
    
    for counter <= half {
        i--
        counter += nums[i]
        res = append(res, nums[i])
    }
    
    return res
}

4. 复杂度分析：
时间复杂度：O(n*n) or O(nlogn) 取决于排序算法的时间复杂度
空间复杂度：O(1)

5. 总结：
5.1：写代码是一个动脑子的活，还是要多动脑多动手
