// 数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。

// 示例 1：

// 输入：[1,2,5,9,5,9,5,5,5]
// 输出：5
//  

// 示例 2：

// 输入：[3,2]
// 输出：-1
//  

// 示例 3：

// 输入：[2,2,1,1,1,2,2]
// 输出：2
//  

// 说明：
// 你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？
/**
 * @param {number[]} nums
 * @return {number}
 */
// 方法1.hash map
var majorityElement = function (nums) {
    if (nums.length == 1) {
         //如果数组长度为1，返回唯一元素
        return nums[0]
    }
    let max = -Infinity;
    let a = new Map()
    for (let num of nums) {
          //如果这个Map对象包含当前元素nums[i](key)，就把该元素的value值加1,如果没有，就初始化value为1
        a.set(num, a.has(num) ? a.get(num) + 1 : 1)
        //max更新最大值
        max = Math.max(max, a.get(num))
          //如果max大于一半的长度，直接返回res
        if (max > nums.length / 2) {
            return num
        }
    }
    return -1
};

/** 题解
主要元素
Map对象以元素作为key,以出现次数作为value，for...of..遍历，更新value值，并记录当前出现次数最大值max，一旦max大于长度的一半就返回当前元素
复杂度分析：
	时间复杂度：O(N)
	一次遍历，与nums的长度有关

	空间复杂度：O(N)
	定义了map对象树
 */

 //方法2.摩根投票法
 var majorityElement = function (nums) {
     //定义count，开始投票
    let num=nums[0],count=0
   for(let i=1;i<nums.length;i++){
       //如果不相同，count减一
       if(num!=nums[i]){
           count--
           if( count<0){//如果小于0，则证明还未达到一半距离，把前面的都清空，count为0,num记录当前的元素
               count=0;
               num=nums[i]
           }
       }
       else{//相同，则加一
           count++
       }
   }
   //成功获取num为主要元素，或没有寻找成功，是最后一个值
   let kk=0
   let mid=parseInt(nums.length/2+1)
   for(let i=0;i<nums.length;i++){
       if(num==nums[i]){//记录num的出现次数          
           kk++
       }
       if(kk==mid){//num的出现次数大于一半长度，就返回num
           return num
       }
   }
   //没找到，就返回-1
   return -1
};

/** 题解
主要元素
摩根投票法的原理，当一个数的重复次数超过数组长度的一半，每次将两个不相同的数删除，最终剩下的就是要找的数。
实现方案：如果当前元素不等于num，就count减一，相同就加一，如果出现小于0的情况，未达到一半距离，把count置为0，num记录为当前元素。等于0则证明所有的数字都可以相互凑成不同的数对，一起抵消。
复杂度分析：
	时间复杂度：O(N)
	一层遍历，与nums的长度有关

	空间复杂度：O(1)
	定义的都是常量
 */