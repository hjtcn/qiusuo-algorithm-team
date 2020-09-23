# 第一周周报
## 20200914-20200919
### 本周主要刷了六道数组相关的算法题，如下：
### 本周题目
- [867][转置矩阵](https://leetcode-cn.com/problems/transpose-matrix/description/)
- [面试题 17.10.] [主要元素](https://leetcode-cn.com/problems/find-majority-element-lcci/)
- [977] [有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/description/)
- [628] [三个数的最大乘积](https://leetcode-cn.com/problems/maximum-product-of-three-numbers/description/)
- [1550] [存在连续三个奇数的数组](https://leetcode-cn.com/problems/three-consecutive-odds/description/)
- [219] [存在重复元素 II](https://leetcode-cn.com/problems/contains-duplicate-ii/description/)

### 我的题解
1.[转置矩阵](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/1-867%E8%BD%AC%E7%BD%AE%E7%9F%A9%E9%98%B5.js)

  - 这题我的思考角度就是拿出一个空数组，重新新行列交换进行赋值

2.[主要元素](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/2-%E9%9D%A2%E8%AF%95%E9%A2%981710%E4%B8%BB%E8%A6%81%E5%85%83%E7%B4%A0.js)

  1. hashMap的方法。Map对象以元素作为key,以出现次数作为value，for...of..遍历，更新value值，并记录当前出现次数最大值max，一旦max大于长度的一半就返回当前元素
  2. 摩根投票法。当一个数的重复次数超过数组长度的一半，每次将两个不相同的数删除，最终剩下的就是要找的数。

3.[有序数组的平方](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/3-977%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.js)

  1. 暴力比较两数组的平方，最小的数放入新的数组中，一方未对比完毕，合并的新的数组中
  2. 利用js数组的map方法和sort方法实现排序。
  3. 双指针。从数组的最大长度开始赋值，左右指针比较大小，left指向的值大，就将left指向的值赋予res,left加一，right指向的值大，就将right的值赋予res，right减一，直到两个指针碰面
   
4.[三个数的最大乘积](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/4-628%E4%B8%89%E4%B8%AA%E6%95%B0%E7%9A%84%E6%9C%80%E5%A4%A7%E4%B9%98%E7%A7%AF.js)

 1. 先从小到大排序，可知道三个数字最大的乘积，只可能出现在nums[0],nums[1],nums[len-1],nums[len-2],nums[len-3]之间，且nums[len-1]肯定会存在。故只需比较(nums[len-1]*nums[0]*nums[1],nums[len-1]*nums[len-2]*nums[len-3])的值谁大就行
 2. 线性判断。利用解构赋值，一层层比较，替换，然后找出最小的两个数(min1,min2)，和最大的三个数(max3,max2,max1)。三个数字最大的乘积，max1作为最大值是必定存在在三数之中，所以，对比(max1 * max2 * max3, max1 * min1 * min2)的较大值返回即可

  
5.[存在连续三个奇数的数组](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/5-1550%E5%AD%98%E5%9C%A8%E8%BF%9E%E7%BB%AD%E4%B8%89%E4%B8%AA%E5%A5%87%E6%95%B0%E7%9A%84%E6%95%B0%E7%BB%84.js)

 1. 标记记录。一层遍历，利用一个变量k记录是否存在连续的三个奇数，一旦存在不为奇数的值，k就置零，一旦k等于3，就返回true,没达到就返回false
 2. js的find方法。找不到符合的就返回undefined.遍历查询数组是否存在arr[i],arr[i+1],arr[i+2]都为奇数，此时要记得判断i不能超过数组长度减去2,防止数组溢出

6.[存在重复元素 II](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/6-219%E5%AD%98%E5%9C%A8%E9%87%8D%E5%A4%8D%E5%85%83%E7%B4%A0II.js)

 1. js的api(如slice,indexOf)。一层遍历当前元素，然后截取从i+1到i+1+k的目标数组，然后利用indexOf判断目标数组是否包含此元素，如果包含，就在长度为k的范围内出现重复值nums[i],返回true
 2. hashMap。查询map对象是否包含当前元素，没有包含就设置key为元素本身，value为当前索引，包含则判断，当前索引和以前出现的该元素的索引值的差是否小于等于k，若是就返回true,不是，就更新当前元素的value为当前索引。

### 一周总结

   这六道题的难度比较低，但是它们更代表着一个起点，一种坚持，以及不断拓展的思维。

   1. js-Api

     - 使用率比较高。map():hash判断元素是否存在，或者设置元素和位置的对应关系，sort()排序
     - 最新了解。find()：如果未找到通过筛选的元素，返回undefined
   2. 摩根投票法。
     - 原理。当一个数的重复次数超过数组长度的一半，每次将两个不相同的数删除，最终剩下的就是要找的数。
   3. 双指针查找。
     - 原理，分别从两端查找，直至left与right指针相遇。
    
### 除了坚持，别无它想
