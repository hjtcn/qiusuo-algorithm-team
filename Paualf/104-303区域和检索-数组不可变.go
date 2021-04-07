给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。
实现 NumArray 类：
• NumArray(int[] nums) 使用数组 nums 初始化对象
• int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）
 
示例：
输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]
解释：
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1)) 
numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
 
提示：
• 0 <= nums.length <= 104
• -105 <= nums[i] <= 105
• 0 <= i <= j < nums.length
• 最多调用 104 次 sumRange 方法
1. Clearfication:
        对 type 进行定义的类型不是很熟悉

2. Coding:
对类型的定义不是很熟悉
/*
    Clearfication:
        对 type 进行定义的类型不是很熟悉
*/
type NumArray struct {
    arr[] int
}
func Constructor(nums []int) NumArray {
    return &NumArray{
        arr:nums,
    }
}
func (this *NumArray) SumRange(left int, right int) int {
    ret := 0
    for i := left;i <= right;i++ {
        ret += *NumArray[i]
    }
    return ret
}
/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

3. 看题解：
type NumArray struct {
    sums []int
}
func Constructor(nums []int) NumArray {
    sums := make([]int,len(nums) + 1)
    
    for i,v := range nums {
        sums[i + 1] = sums[i] + v
    }
    
    return NumArray{sums}
}
func (this *NumArray) SumRange(left int, right int) int {
    return this.sums[right + 1] - this.sums[left]
}
/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
暴力法：
type NumArray struct {
    arr []int
}
func Constructor(nums []int) NumArray {
    return NumArray{nums}
}
func (this *NumArray) SumRange(left int, right int) int {
    ret := 0
    
    for i := left;i <= right;i++ {
        ret += this.arr[i]
    }
    return ret
}
/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

4. 复杂度分析：
时间复杂度：初始化O(n)，每次检索O(1), 如果不初始化的话，每次计算都是O(n) 的时间复杂度
空间复杂度：O(n) 

5. 总结：
5.1: 空间换时间，提前计算出来进行存储，然后直接获取结果

5.2: 对设计类型的题目不是很熟悉
