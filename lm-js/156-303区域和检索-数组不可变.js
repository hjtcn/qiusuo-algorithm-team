// 303. 区域和检索 - 数组不可变
// 给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。

// 实现 NumArray 类：

// NumArray(int[] nums) 使用数组 nums 初始化对象
// int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）
 

// 示例：

// 输入：
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// 输出：
// [null, 1, -1, -3]

// 解释：
// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
// numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1)) 
// numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
 

// 提示：

// 0 <= nums.length <= 104
// -105 <= nums[i] <= 105
// 0 <= i <= j < nums.length
// 最多调用 104 次 sumRange 方法



/*
    思路：其实看到这种题还是比较怂的，说明构造函数还是比较不确定的。
        但是还是要试一试呀，毕竟题目也不难，只是数据赋值问题。
        看一看示例的解释。构造函数的使用方法就清晰了。
        我先看的是构造函数的sumRange方法。它包含的参数是left,right.那数组呢？需要数组来遍历呀。
        那就在NumArray构造函数中this.nums=nums，将目标数组传递过去。
        最后要返回和，那构造函数中还要有个this.sum=0呀。
        细节来了。在sumRange方法中，在遍历求和之前，要记得将原本的和清零哟。
        后来看了示例的输出，输出中第一个返回null,我以为NumArray要返回null,但是有感觉挺别扭，后来把这个return给去掉了，也是可以的。可能这个null，就是代表的无返回吧，哈哈哈
*/
/*
 * @lc app=leetcode.cn id=303 lang=javascript
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
/**
 * @param {number[]} nums
 */
 var NumArray = function(nums) {
    this.sum=0
    this.nums=nums
    // return null;
};

/** 
 * @param {number} left 
 * @param {number} right
 * @return {number}
 */
NumArray.prototype.sumRange = function(left, right) {
    this.sum=0
    for(let i=left;i<=right;i++){
        this.sum+=this.nums[i]
    }
    return this.sum
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/
/**
 * Your NumArray object will be instantiated and called as such:
 * var obj = new NumArray(nums)
 * var param_1 = obj.sumRange(left,right)
 */


/*
    题解：一直就在想着，这和动态规划有啥关系呢？
        一看题解。。没错，是我浅薄了。。。哈哈
        但是自己看了是前缀和，自己写了一写，发现数组的溢出和还是不好控制的。

       错误代码：
       var NumArray = function(nums) {
            this.sum=[nums[0]]
            for(let i=1;i<nums.length;i++){
                //这里根据往常习惯都是赋值i了。
                this.sum[i]=this.sum[i-1]+nums[i]
            }
        };

        NumArray.prototype.sumRange = function(left, right) {  
            return this.sum[right]-this.sum[left-1]
        }; 

        前缀和很好理解，但是代码控制还是不容易的，这个以后遇到的次数会很多，还是要多看看，多写写。
*/

//正确代码
var NumArray = function(nums) {
    this.sum=[nums[0]]
    for(let i=0;i<nums.length;i++){
        this.sum[i+1]=this.sum[i]+nums[i]
    }
};

/** 
 * @param {number} left 
 * @param {number} right
 * @return {number}
 */
NumArray.prototype.sumRange = function(left, right) {  
    return this.sum[right+1]-this.sum[left]
};

/*
    时间复杂度：O(N).优势在于一个数组，只计算了一次。
    空间复杂度：O(N)
*/