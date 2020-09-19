// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

//  

// 示例 1:

// 输入: nums = [1,2,3,1], k = 3
// 输出: true
// 示例 2:

// 输入: nums = [1,0,1,1], k = 1
// 输出: true
// 示例 3:

// 输入: nums = [1,2,3,1,2,3], k = 2
// 输出: false


/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function(nums, k) {
    let target=[]
    for(let i=0;i<nums.length;i++){
        //截取数组，从索引从i+1开始，到i+1+k结束
        target=nums.slice(i+1,i+1+k)
        //查找目标数组是否存在当前元素来确认是否存在重复元素
        if(target.indexOf(nums[i])!=-1){
            return true
        }
    }
    return false
};
/** 题解
219-存在重复元素II，主要使用js的api,一层遍历当前元素，然后截取从i+1到i+1+k的目标数组，然后利用indexOf判断目标数组是否包含此元素，如果包含，就在长度为k的范围内出现重复值nums[i],返回true

 复杂度分析：
    平均时间复杂度是:O(N^2)
    一层for循环变量,indexOf也会进行一次遍历查询
    空间复杂度：O(n)
    定义目标数组target
 */





/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function(nums, k) {
    //利用map对象,key存元素，value存索引
    let m=new Map()
    for(let i=0;i<nums.length;i++){
        //如果map对象已经包含当前元素
        if(m.has(nums[i])){
            //判断索引间隔是否是小于等于k，若是就返回true
            if(i-m.get(nums[i])<=k){
                return true
            }
            else{//索引间隔大于k，则更新当前元素key对应的value(位置)为i
                m.set(nums[i],i)
            }
        }
        else{//不包含当前元素，就设置当前元素的索引值为i
            m.set(nums[i],i)
        }
    }
    return false
};



/** 题解
219-存在重复元素II，一层for循环遍历，利用hashMap去查询是否包含当前元素，没有包含就设置key为元素本身，value为当前索引，包含则判断，当前索引和以前出现的该元素的索引值的差是否小于等于k，若是就返回true,不是，就更新当前元素的value为当前索引
 复杂度分析：
    平均时间复杂度是:O(N)
    一层for循环变量
    空间复杂度：O(n)
    定义map对象，以元素本身为key,以索引为value
 */