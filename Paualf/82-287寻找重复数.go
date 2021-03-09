/*
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。
示例 1：
输入：nums = [1,3,4,2,2]
输出：2
示例 2：
输入：nums = [3,1,3,4,2]
输出：3
示例 3：
输入：nums = [1,1]
输出：1
示例 4：
输入：nums = [1,1,2]
输出：1
 
提示：
2 <= n <= 3 * 104
nums.length == n + 1
1 <= nums[i] <= n
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
 
进阶：
如何证明 nums 中至少存在一个重复的数字?
你可以在不修改数组 nums 的情况下解决这个问题吗？
你可以只用常量级 O(1) 的额外空间解决这个问题吗？
你可以设计一个时间复杂度小于 O(n2) 的解决方案吗？
*/

1. Clearfication:
nums 中只有一个整数出现两次或多次，其余整数均只出现一次
暴力解法：遍历两遍数组，寻找出现两次或多次的值
for i:[0-> len(nums)
   for j [i + 1,len(nums) )_
if nums[i] == nums[j] {
return nums[i]
}    
return -1
hashMap: 使用hashMap 记录数字出现次数，第一遍历记录出现次数，第二次遍历找出hashMap 中 > 1的数据
m := make([]int, len(nums))
for i : [0 -> len(nums) - 1] 
index : = nums[i]
m[index]++
for i : [0 -> len(nums) - 1]
index := nums[i]
if m[index] > 1 {
return nums[index]
}
return -1

2. Coding:
暴力两遍搜索数组：
func findDuplicate(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    for i := 0;i < len(nums);i++ {
        for j := i + 1;j < len(nums);j++ {
            if nums[i] == nums[j] {
                return nums[i]
            }
        }
    }
    return -1
}

HashMap:
func findDuplicate(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    m := make(map[int]int, len(nums))
    for i := 0;i < len(nums);i++ {
        index := nums[i]
        m[index]++
    }
    for i := 0;i < len(nums);i++ {
        index := nums[i]
        
        if m[index] > 1 {
            return nums[i]
        }
    }
    return -1
}

优化hashMap 遍历方式
func findDuplicate(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    m := make(map[int]int, len(nums))
    for i := 0;i < len(nums);i++ {
        index := nums[i]
        m[index]++
    }
    for k,v := range m {
        if v > 1 {
            return k
        }
    }
 
    return -1
}

3. 看题解:
抽屉原理：
9个抽屉，10个苹果，肯定有一个抽屉放了两个苹果，找到放了两个苹果的抽屉
自己没有注意到的细节，初始化的时候 l = 1
然后 mid 的变化是 l,r 的变化和mid的变化
func findDuplicate(nums []int) int {
    n := len(nums)
    l,r := 1,n-1
    ans = -1
    
    for l <= r {
        mid := (l + r) >> 1
        cnt := 0
        
        for i := 0;i < n;i++ {
            if nums[i] <= mid {
                cnt++
            }
        }
        
        if cnt <= mid {
            l = mid + 1
        }else {
            r = mid - 1
            ans = mid
        }
    }
    
    return ans
}

还可以联系到快慢指针，我的个神勒
参考：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/qian-duan-ling-hun-hua-shi-tu-jie-kuai-man-zhi-z-3/
[1,3,4,2,2]
把索引0的元素的值当成链表的头节点
再下一个节点的索引为3，即为第一个2
再下一个节点的索引为2，即为4
再下一个节点的索引为4，即为第二个2
再下一个节点的索引为2，即为4
此时形成了一个环
而形成环的原因是下一个节点的索引一致，即为重复的数字

func findDuplicate(nums []int) int {
    slow,fast := 0,0
    for slow,fast = nums[slow],nums[nums[fast]];slow != fast;slow,fast = nums[slow],nums[nums[fast]] {
    
    }
    slow = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    
    return slow
}

4. 复杂度分析： 
暴力搜索：
时间复杂度：O(n*n) 遍历两遍数组
空间复杂度：O(1)

hashMap:
时间复杂度：O(n) 遍历数组
空间复杂度：O(n) 使用hashmap 记录出现次数

二分查找：
时间复杂度：O(n) 二分收缩查找空间的时候的复杂度是O(logN), 收缩完空间以后，要进行计数遍历，这个时候复杂度是O(n)
空间复杂度：O(1）常量空间

快慢指针：
时间复杂度：O(n)  
空间复杂度：O(1) 

5. 总结：
5.1: 算法对于解决问题是有明确的定义和解决步骤的
