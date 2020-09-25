/*
 * @lc app=leetcode.cn id=969 lang=php
 *
 * [969] 煎饼排序
 *
 * https://leetcode-cn.com/problems/pancake-sorting/description/
 *
 * algorithms
 * Medium (64.63%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 11.9K
 * Testcase Example:  '[3,2,4,1]'
 *
 * 给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k
 * 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行）以完成对数组 A 的排序。
 * 
 * 返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[3,2,4,1]
 * 输出：[4,2,4,3]
 * 解释：
 * 我们执行 4 次煎饼翻转，k 值分别为 4，2，4，和 3。
 * 初始状态 A = [3, 2, 4, 1]
 * 第一次翻转后 (k=4): A = [1, 4, 2, 3]
 * 第二次翻转后 (k=2): A = [4, 1, 2, 3]
 * 第三次翻转后 (k=4): A = [3, 2, 1, 4]
 * 第四次翻转后 (k=3): A = [1, 2, 3, 4]，此时已完成排序。 
 * 
 * 
 * 示例 2：
 * 
 * 输入：[1,2,3]
 * 输出：[]
 * 解释：
 * 输入已经排序，因此不需要翻转任何内容。
 * 请注意，其他可能的答案，如[3，3]，也将被接受。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= A.length <= 100
 * A[i] 是 [1, 2, ..., A.length] 的排列
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $arr
     * @return Integer[]
     */
     // 执行用时：24 ms 在所有 PHP 提交中击败了50.00%的用户
     // 内存消耗：14.5 MB, 在所有 PHP 提交中击败了100.00%的用户
    function pancakeSort($arr) {
        $res = [];
        $size= count($arr);
        // 这个遍历的的元素的在第几个位置。方便后面进行截取数组
        for($i=$size; $i>1; $i--) {
            // 找到最大值的位置
            $maxIndex = $this->findMaxIndex($arr, 0, $i-1);
            $res[] = $maxIndex + 1; // 反转的位置,第几个元素
            $res[] = $i;  // 然后再全部向后反一次
            // array_slice 截取从0到最大值 截取前面的元素 ，进行array_reverse反转，然后再将array_merge后面的数组不动拼接起来

            /*
              示例
             [3, 2, 4, 1] => 找到最大元素位置是第三个 3 => 也就是将前面三个元素进行 [4,2,3]  
                    （这里就采用array_slice先将前面 0 ~ 3个元素进行截取，然后反转）
             后面没有参数反转的元素也就是 [1]
                    （这里就采用array_slice将后面的元素开始，一直到结尾截取下来）
             最终两个数组拼接到一起
            [4,2,3,1]
             */
            // 现将最大值翻到前面
            $arr = array_merge(array_reverse(array_slice($arr, 0, $maxIndex + 1)), array_slice($arr, $maxIndex + 1));
            // 然后再对最大元素那一部分数组进行反转
            $arr = array_merge(array_reverse(array_slice($arr,0,$i)), array_slice($arr,$i));
        }

        return $res;
    }

    // 递归寻找最大值的位置
    function findMaxIndex($arr, $start, $end){
        // 终止条件
        if($start == $end){
            return $start;
        }
        $mid = floor(($start + $end) / 2);
        $left= $this->findMaxIndex($arr, $start, $mid); // 左边最大值
        $right= $this->findMaxIndex($arr, $mid+1, $end); // 右边最大值
        if($arr[$left] > $arr[$right]) {
            return $left;
        }else{
            return $right;
        }

    }
}
// @lc code=end


// @tacks think=start
/*
题意 meaning 
给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k
 * 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行）以完成对数组 A 的排序。
 * 
 * 返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。

    一开始没太看懂题意，后来慢慢捋一下思路，其实就是翻大饼。

- 一开始我们[3, 2, 4, 1]，找到最大的数，翻动前面k=3个元素，把他翻到前面 [4,2,3,1]
- 然后再翻动k=4的元素全部翻过去 [1,3,2,4]，这样4就是最大元素，下次不用翻了
- 然后[1,3,2,4]中最大值3，翻动前面k=2 这样就是[3,1,2,4]
- 然后再全部翻动前面k=3的元素全部翻过去 [2,1,3,4]，这样3就是最大元素，下次不用翻了
- 然在反动k=2 这样[1,2,3,4] 就变成有序的
- 所以k的次序也可以 [3,4,2,3,2] 也可以是正确答案。

关键 key 

    “找到前面模块的最大值的位置”
    “一般需要两次操作，第一次将最大值翻到前面，第二次将最大值翻到后面，完成有序。”

想法 idea

【最大值反转法】
- 从后面向前面遍历
    - 二分查找，找到当前数组的最大值的位置
    - 然后将前面到最大值位置的元素进行一次反转，也就是reverse，我调用的内置函数。
    - 然后再将前面的翻到后面，这样完成一次最大值的反转。
    - 慢慢的后面的元素都会有序，这样即可成功。

 
   - 时间复杂度，这里采用了内置函数，用于数组切割反转合并，另外使用二分求最大值下标    

*/
// @tacks think=end