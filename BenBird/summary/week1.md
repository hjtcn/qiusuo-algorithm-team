## 第一周总结

#### 题目列表

* First day: [[867]转置矩阵](https://leetcode-cn.com/problems/transpose-matrix)

* The next day: [面试题17.10.主要元素](https://leetcode-cn.com/problems/find-majority-element-lcci)

* Third day: [[977]有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)

* Fourth day: [[628]三个数的最大乘积](https://leetcode-cn.com/problems/maximum-product-of-three-numbers)

* Fifth day: [[1550]存在连续三个奇数的数组](https://leetcode-cn.com/problems/three-consecutive-odds)

* Sixth day: [[219]存在重复元素II](https://leetcode-cn.com/problems/contains-duplicate-ii)

#### 我的题解

1. [转置矩阵](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/1-867%E8%BD%AC%E7%BD%AE%E7%9F%A9%E9%98%B5.go)
    
    - 转置矩阵的思路：
    
        就是行列转换，其原理，就是讲行和列坐标进行互换即 nums[i][j] = nums[j][i]
    
2. [面试题17.10.主要元素](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/2-%E9%9D%A2%E8%AF%95%E9%A2%981710%E4%B8%BB%E8%A6%81%E5%85%83%E7%B4%A0.go)
    
    -  暴力求解：
            
            创建一个map，map的key 为nums中的元素，对应的value，为该元素出现的次数，然后在循环判断map中出现次数超过nums元素总个数的一半
    
    - 摩尔算法：
    
           主要使用相邻两个数进行两两抵消的方法，因为所求数值必出现次数大于nums总元素个数的一半，所以：
           假设一定存在这个数，则一定存在相邻两个数相同 或 这种不相邻但首尾呼应{1,2,1,3,1,4,1,5,1}，这种前面的数都两两抵消，只剩下1，
           然而验证的时候刚好符合
           假设不存在这个数，则在进行两两抵消后，有两种结果，1. 全部被抵消 2. 剩下一个数字 然后经验证不成立
           
3. [有序数组的平方](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/3-977%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.go)

    - 双指针法：
    
           初始化一个新切片new_arr := make([]int, len(A), len(A))
           使用左(i)右(j)指针，刚开始，左指针指向切片最左端，有指针指向切片最右端同时遍历该有序数组
           直接将左右两端数值求平方进行对比，
           	若左边数值大，将左边求平方后数据存入new_arr，则左边指针向右移动，
           	若右边数值大，将左边求平方后数据存入new_arr，则右边指针向左移动，
           继续左右指针指向的数值，求平方后对比，循环如此，直到左右指针相遇
           最后切片new_arr就是所要获取的结果

4. [三个数的最大乘积](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/4-628%E4%B8%89%E4%B8%AA%E6%95%B0%E7%9A%84%E6%9C%80%E5%A4%A7%E4%B9%98%E7%A7%AF.go)

    - 直接求解:
    
           先对数组进行正序排列，然后根据规律可得，最大乘积之和的三个数必有排在最后一位数，另外两个数则存在第一、二位数 或 倒数第二、三位,
           然后分别对这两种可能的三个数，进行乘积比较，然后获取最大三数乘积之和
           
    - 线性扫描：
    
           线性扫描，找出最小的两位数，和最大的三位数，然后进行比较，同上排序法，将最小的两位数和最大的数的乘积 与 最大的三位数乘积比较
           获取最大三数乘积

5. [存在连续三个奇数的数组](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/5-1550%E5%AD%98%E5%9C%A8%E8%BF%9E%E7%BB%AD%E4%B8%89%E4%B8%AA%E5%A5%87%E6%95%B0%E7%9A%84%E6%95%B0%E7%BB%84.go)

    - 循环遍历：
    
          循环遍历数组arr, 连续遇到奇数的个数记为nums, 当nums == 3时返回true, 当遇到偶数就将nums 置为 0, 最后判断nums 是否等于 3

6. [存在重复元素II](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/6-219%E5%AD%98%E5%9C%A8%E9%87%8D%E5%A4%8D%E5%85%83%E7%B4%A0II.go)

    - 哈希：
    
          利用Map模拟集合，集合特点：集合中没有重复元素
          初始化Map: temp_map
          循环遍历nums
          判断每个元素是否存于temp_map,
            若存在：直接返回true
            不存在：将该元素存入Map: temp_map
          然后，判断temp_map的长度是否大于k
            若大于：释放temp_map的第一个元素
            不大于：继续遍历nums

#### 总结

本周六道题不是都为简单题目，总的来说都是解题思路大部分都是找规律，找到规律题就迎刃而解，
其中认为有点难度的就是 [面试题17.10.主要元素]，这道题的"摩尔算法" 很难想到，但是想到之后，也是很容易

加油加油，奥利给

