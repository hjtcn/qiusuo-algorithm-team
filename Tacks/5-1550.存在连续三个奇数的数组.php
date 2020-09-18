/*
 * @lc app=leetcode.cn id=1550 lang=php
 *
 * [1550] 存在连续三个奇数的数组
 *
 * https://leetcode-cn.com/problems/three-consecutive-odds/description/
 *
 * algorithms
 * Easy (70.10%)
 * Likes:    4
 * Dislikes: 0
 * Total Accepted:    8.4K
 * Total Submissions: 12K
 * Testcase Example:  '[2,6,4,1]'
 *
 * 给你一个整数数组 arr，请你判断数组中是否存在连续三个元素都是奇数的情况：如果存在，请返回 true ；否则，返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：arr = [2,6,4,1]
 * 输出：false
 * 解释：不存在连续三个元素都是奇数的情况。
 * 
 * 
 * 示例 2：
 * 
 * 输入：arr = [1,2,34,3,4,5,7,23,12]
 * 输出：true
 * 解释：存在连续三个元素都是奇数的情况，即 [5,7,23] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 1000
 * 1 <= arr[i] <= 1000
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $arr
     * @return Boolean
     */
    // 直观法
    // 执行用时：12 ms, 在所有 PHP 提交中击败了87.50%的用户
    // 内存消耗：14.5 MB, 在所有 PHP 提交中击败了79.17%的用户
    function threeConsecutiveOdds1($arr) {
        $size = count($arr);
        for($i=0; $i<=$size-3; $i++) {
            // 如果三个数有一个是偶数，那么判断就是0，继续循环。
            if( ($arr[$i] % 2) && ($arr[$i+1] % 2) && ($arr[$i+2] % 2) ) {
                return true;
            }
        }
        return false;
    }


    // 计数法
    // 执行用时：12 ms, 在所有 PHP 提交中击败了87.50%的用户
    // 内存消耗：14.6 MB, 在所有 PHP 提交中击败了66.67%的用户
    function threeConsecutiveOdds2($arr) {
        $oodCount = 0; // 统计奇数连续的个数
        foreach($arr as $item) {
            if( $item%2 != 0){
                $oodCount++; // 如果当前元素 是奇数，那么需要计数。
            }else{
                $oodCount = 0;
            }
            if($oodCount >= 3) {
                return true;
            }
        }
        // 实际上边界判断也不必，因为题意给的是 数组长度大于1，那么肯定是一个数组，至于是不是三个元素以上，如果不是那么直接结束foreach循环进入false
        // 另外如果遍历结束，还是找不到连续三个变量，那么就直接false。
        return false;
    }

    // 采用与判断奇偶性
    // 执行用时：12 ms, 在所有 PHP 提交中击败了87.50%的用户
    // 内存消耗：14.6 MB, 在所有 PHP 提交中击败了62.50%的用户
    function threeConsecutiveOdds3($arr) {
        $oodCount = 0; 
        foreach($arr as $item) {
            // 这里采用&计算，貌似性能没什么影响。！！！！
            if( $item&1 ){
                $oodCount++; // 如果当前元素 是奇数，那么需要计数。
            }else{
                $oodCount = 0;
            }
            if($oodCount >= 3) {
                return true;
            }
        }
        return false;
    }


    // 双指针法
    // 看起来有点复杂了，还是抛弃了
    // 执行用时：16 ms, 在所有 PHP 提交中击败了37.50%的用户
    // 内存消耗：14.6 MB, 在所有 PHP 提交中击败了66.67%的用户
    function threeConsecutiveOdds4($arr) {
        $i = 0;
        $j = 0;
        $size = count($arr);
        // 遍历数组
        for(; $j<$size; $j++){
            if($arr[$j] % 2 == 0){
                // 当前为偶数，那么判断一下 i j的位置看看是不是相隔3
                // 那么奇数位置为 $j-1
                if($j-1-$i >= 2) {
                    return true;
                } 
                // 如果不是 重置i的位置
                $i = $j+1; // 但是当前i的位置也不一定为奇数
            }
        }
        if($j-1-$i >= 2){
            return true;
        }
        return false;
    }


}
// @lc code=end


// @tacks think=start
/*
题意 meaning 
  
    一个整数数组arr,要求是否存在三个连续的奇数。
        边界判断。首先这个数组一定有至少三个元素，否则可以直接false。 

    一看到数组，二话不说肯定需要遍历。
        判断相邻三个元素是否是奇数，这里也有一个小知识点。
            如何判断一个“整数”是否是“奇数”？

                先来看一下奇偶数定义： 能够被2整除的整数是偶数，不能被2整除的是奇数。

            1>[×] return ( i%2 == 1 ) ? '奇数' ：'偶数' ;  // 取余 第一种写法显然是错误的，当i为负数就不成立。 当然你可以改成  if( i%2 ) 
            2>[√] return ( i%2 != 0 ) ? '奇数' : '偶数' ;  // 取余 根据定义推算出来的公式。
            3>[√] return ( i&1 == 1 ) ? '奇数' : '偶数' ;  // 与运算，奇数&1 结果一定为1。 

        计算机中，数据是按照补码进行二进制存储的。
            偶数的最低位一定是0
            奇数的最低位一定是1
            所以如果要判断这个数是奇数还是偶数，只要用这个数按位与1就可以了。
 
            如果结果为1，那么就是奇数，如果结果为0，那么结果就是偶数。

关键 key 

    “三个连续”，“奇数判断方法”

想法 idea


【1】 直观办法

    直接拿相邻三个变量然后做 全部为奇数判断，但是这种变通性不大。
    
        if ((arr[i] & 1) & (arr[i + 1] & 1) & (arr[i + 2] & 1)) {  return true; }   // 奇数&1 还是 1
        if ((arr[i] * arr[i + 1] * arr[i + 2]) % 2 == 1)        {  return true; }   // 奇数相乘还是奇数
    如果是要求四个连续以上，是不是你的语句会更长，不过这个方法是最好理解的。

【2】 计数判断法
    
    采用oodCount统计连续的奇数个数。
        遍历数组，然后判断当前值是否是奇数
            采用 % 取余  i%2 != 0
           
            如果是   oddCount++
            如果不是 oddCount重置为0
            判断oddCount是否等于3，那么是的话，直接返回true
        遍历结束还没找到 直接返回false


【3】 计数判断法（&判断奇偶性）

    采用 & 与    i&1 == 1


【4】 双指针

    这个是看了网友的题解
        作者：whinight
        链接：https://leetcode-cn.com/problems/three-consecutive-odds/solution/shuang-zhi-zhen-fa-tong-yong-xing-you-yu-guan-fang/

    其实双指针的方法，真的在刷题中经常遇到，所以，我在这里再练习一下。

    而且使用双指针，还能准确找到连续奇数的位置，实际上是比计数法更详细一些。（虽然本题，我个人觉的使用计数法是最佳的，但是双指针是最灵活的）
    
    双指针 起始指针i 和 结束指针j 用来标记元素位置。

    i = 0 作为奇数起始位置的指针 
    j = 0  
    循环数组
        如果当前元素j是偶数 
            那么连续奇数的结束位置是在前一位       =》  j-1 
            那么判断前面连续的奇数的间隔 是3       =》  j-1-i >=3
                如果是直接返回true
            如果判断前面连续的奇数的间隔不是3      =》  j-1-i < 3
                重置左边指针i，为j的下一位        =》  i = j+1
        如果当前元素是奇数
            那么移动j元素位置。 j++ 
    如果结束循环
        那么有可能三个连续的奇数是在数组末尾
            再次进行判断
                if (j-1-i>=2) 
    否则 
        返回 false
 
 
综上所述，时间复杂度都是O(n)
 
*/
// @tacks think=end