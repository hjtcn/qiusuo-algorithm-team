/*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
示例 1:
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数； 
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。
*/

1. Clearfication:

循环数组里面搜索它的下一个更大的数。
最近比它大的数
有什么好的解决方法么？
没有的话你会怎么处理和解决呢？
朴素的方法：
从数组这个下表向后寻找，如果下表等于最后一个元素，则将下标指向0，继续循环，循环终止条件有两个
a. 找到比它大的元素就终止循环
b. 索引下标 == 它自己的索引下标
伪代码：
func nextGreaterElements(nums []int) []int {
    ret := []int{}
    
    for i := 0;i < len(nums);i++ {
        nextGreater := findNextGreaterElements(i,nums)
        ret = append(ret, nextGreater)
    }
    
    return ret
}
func findNextGreaterElements(currentIndex int,nums []int) int {
    ret := -1
    
    for i := currentIndex + 1;i != currentIndex && i < len(nums); i++ {
        if nums[i] > nums[currentIndex] {
            ret = nums[i]
            break
        }
        
        if i == len(nums) - 1 {
            i  = 0
        }
    }
    
    return ret 
}
单调栈自己是没有想到怎么写的：
只想到了单调栈可以解决最近相似性的问题

2. Coding:
修改上面伪代码 findNextGreaterElements 添加没有加上的判断
不能加 i < len(nums) 的判断，因为如果是最后一个元素的话，它还是要执行的，不过要加元素redirect 到下标为0的元素
func findNextGreaterElements(currentIndex int,nums []int) int {
    ret := -1
    
    for i := currentIndex + 1;i != currentIndex; i++ {
        if i < len(nums) && nums[i] > nums[currentIndex] {
            ret = nums[i]
            break
        }
        
        if i >= len(nums) {
            i  = 0
        }
    }
    
    return ret 
}
又超时了，在[5,4,3,2,1] 的 case 下超时了
自己在 redirect 的代码里面写成了 
if i == len(nums) - 1 { 
    i = 0
}
然后去打印出来，一直没有 rediect 到 0, 循环以后要 i++,所以 i 应该从 -1开始加，i++ 以后从0开始
这样还是出问题了现象是
1
2
3
4
2
3
4
0
3
4
0
4
0
5
6
7
8
这样的输出，前面的循环从 i: 0 -> 3 都是正常的，当 i=4的时候异常了，也就是 currentIndex = 4,然后 i 是从 5开始的，这时候 
if i == len(nums) - 1 {
    i = -1
}
这块代码就没有起作用了，我们改成
if i >= len(nums) - 1 {
    i = -1
}
然后就好了
最后完整的代码是：
func nextGreaterElements(nums []int) []int {
    ret := []int{}
    
    if len(nums) <= 0 {
        return ret
    }
    for i := 0;i < len(nums);i++ {
        nextGreater := findNextGreaterElements(i,nums)
        ret = append(ret, nextGreater)
    }
    
    return ret
}
func findNextGreaterElements(currentIndex int,nums []int) int {
    ret := -1
    
    for i := currentIndex + 1;i != currentIndex; i++ {
        
        if i < len(nums) && nums[i] > nums[currentIndex] {
            ret = nums[i]
            break
        }
        
        if i >= len(nums) - 1 {
            i  = -1
        }
    }
    
    return ret 
}

3. 看题解:
方法1: 扩容nums,利用单调栈寻找
方法2: 使用round 变量遍历两遍

4. 复杂度分析:
时间复杂度：数组寻找：O(n*n),单调栈:O(n)
空间复杂度：O(n)

5. 总结：
5.1: 知道是知道，做到是做到，知道离做到差好远，所以多思考，多总结，多写代码，不要只是想
5.2: 如果下次遇到这种问题，单调栈的代码目前还是不能保证写出来，还是要多练
