<?php
/*
 * @lc app=leetcode.cn id=219 lang=php
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (40.38%)
 * Likes:    203
 * Dislikes: 0
 * Total Accepted:    60.9K
 * Total Submissions: 150.8K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的
 * 绝对值 至多为 k。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: nums = [1,2,3,1], k = 3
 * 输出: true
 * 
 * 示例 2:
 * 
 * 输入: nums = [1,0,1,1], k = 1
 * 输出: true
 * 
 * 示例 3:
 * 
 * 输入: nums = [1,2,3,1,2,3], k = 2
 * 输出: false
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Boolean
     */
    // 双指针-滑动窗口
    // 超出时间限制，放弃！
    function containsNearbyDuplicate1($nums, $k) {
        $size = count($nums); // 数组大小
        for($i = 0; $i<$size; $i++){ // 固定i的位置，然后看j
            // 这里采用跟之前的k个元素进行比较
            /*
            for($j = max($i-$k,0); $j<$i; $j++) { // 将每个元素与它之前的k 个元素中比较查看它们是否相等
                if($nums[$i] == $nums[$j]){
                    return true;
                }
            }
            */
            // 这里采用跟当前i位置之后的k的元素进行比较 ，其实这两种没有什么区别。
            for($j = $i+1; $j-$i<= $k; $j++) {
                // 因为j++，可能已经到数组末尾，所以需要判断 当前j位置元素是否存在。
                if(isset($nums[$j]) && ($nums[$i] == $nums[$j]) ){
                    return true;
                }
            }
        }
        return false;
    }

    // 哈希表 
    // 执行用时：28 ms, 在所有 PHP 提交中击败了90.91%的用户
    // 内存消耗：19.7 MB, 在所有 PHP 提交中击败了66.67%的用户
    function containsNearbyDuplicate2($nums, $k) {
        $hash = []; // 哈希表
        foreach($nums as $key=>$item) {
            // 如果当前hash表存在 一个key等于item
            // 并且当前元素位置，之差小于k 那么就直接符合
            if(isset($hash[$item]) && ($key-$hash[$item]) <= $k) {
                return true;
            }
            // 否则保存当前元素的值，以及位置
            $hash[$item] = $key;
        }
        return false;
    }

    // 函数法
    // 根据小马姐的思路进行编写 
    // 也是超出时间限制！
    function containsNearbyDuplicate3($nums, $k) {
        foreach($nums as $i=>$item) {
            // 当然这里不用担心，k要截取的超出nums的长度，array_slice会感知到，并且只截取到nums的末尾
            $slideArr =  array_slice($nums, $i+1, $k); // 拿到中间数组
            if(in_array($item, $slideArr)) {
                return true;
            }
        }
        return false;
    } 
}
// @lc code=end


// 测试数据
$obj = new  Solution();

$arr = [0,1,2,3,4,5,0]; // [0,1,2,3,4,5,0] 5  
$k = 3;
var_dump($obj->containsNearbyDuplicate3($arr,$k));


// @tacks think=start
/*
题意 meaning 

    拿到一个题首先分析输入输出，然后看过程。（仅代表个人习惯）
        input
            一个整数数组 nums
            一个整数为   k

        operation
            数组中任意两个索引  i j
            如果满足，nums[i] == nums[j]
            同时满足，| i - j | <= k        （只要有满足条件的这样i j就行，可能会有很多个元素值相等）


        output
            满足返回 true
            不满足则 false
  
    
关键 key 

    “i j 索引的元素值一样”  并且  “| i-j | <= k”
   

想法 idea

    一开始想到这个题，数组，肯定是需要遍历的。同时又是两个索引元素。
    这个时候可以想到是否双指针，然后去维护一个k大小的滑动窗口，但是涉及到双重循环。

    
    我们也可以用map来保存这样的相同元素的值，并且判断他们i j的位置差
        事实上我们可以固定i的位置，那么只要满足第一个j的位置有相同元素。
        并且小于k,那么我们就可以直接返回，而不需要关心他后面是否还有相等的元素。


【1】双指针（滑动窗口、线性遍历）

    第一层for循环，固定当前i的位置
        第二层for循环，移动j的位置，在一个k的大小滑动窗口中(将每个元素与它之前的 k 个元素中比较查看它们是否相等)
            比较当前j的位置是否等于i的元素的值，如果存在即返回true
    
    可以看到这里是用了双重循环，所以时间复杂度比较大，最差的情况就是k的大小等于数组大小n O(n*n) 
    空间复杂度，没有声明其他额外的变量。

    综上所述
        时间复杂度在 O(n * min(n,k))
        空间复杂度在 O(1)

【2】hash哈希法
    
    PHP的数组很强大！！！
        其中键值对数组采用了哈希表实现能够保证基本查找时间复杂度为 O(1)。 具体的原理就不多说了。

    我下面写的这个伪代码，而能不是很容易看懂，可以直接look代码，哈哈哈哈哈！

    声明一个hash数组  => 也就是类似 hash[key] = value
    遍历这个nums数组
        如果hash数组为空则给他赋值 key=当前元素item,以及 value = 当前位置i
            hash[item] = i
        如果hash数组中存在当前元素item 而且当前位置-hash数组当前元素的值 <=3 
            返回 true

    时间复杂度
        也就是遍历数组的时间O(n)
        hash数组操作元素时间O(1)
    空间复杂度
        额外申请一个hash数组，那么大小也就是k的大小。或者是整个数组。
    综上所述
        时间复杂度 O(n) 。 
        空间复杂度 O(min(n,k))。

【3】 函数法  超出时间限制

    PHP的数组很强大，而且数组的函数操作也是挺多的，非常灵活，在日常开发中也是经常使用到。
    
    遍历该数组
        然后截取当前位置i+1 到i+1+k位置之间的数组 => 中间数组 => 也就是大小为k滑动窗口
            PHP可以使用array_slice(array,offset,length)函数
                也就是array_slice(nums, i+1, k)  => slideArr 滑动数组
        再判断这个中间数组中，是否包含当前元素。
            PHP可以使用 in_array(search,array)函数进行判断
                也就是 in_array(nums[i],slideArr)  存在返回true 失败返回 false
*/
// @tacks think=end

