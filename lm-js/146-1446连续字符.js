// 1446. 连续字符
// 给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。

// 请你返回字符串的能量。

 

// 示例 1：

// 输入：s = "leetcode"
// 输出：2
// 解释：子字符串 "ee" 长度为 2 ，只包含字符 'e' 。
// 示例 2：

// 输入：s = "abbcccddddeeeeedcba"
// 输出：5
// 解释：子字符串 "eeeee" 长度为 5 ，只包含字符 'e' 。
// 示例 3：

// 输入：s = "triplepillooooow"
// 输出：5
// 示例 4：

// 输入：s = "hooraaaaaaaaaaay"
// 输出：11
// 示例 5：

// 输入：s = "tourist"
// 输出：1
 

// 提示：

// 1 <= s.length <= 500
// s 只包含小写英文字母。
/*
    思路：首先自己想的就是一层遍历，当当前元素等于前一位时，count++,不停的更新最大值max；不想等时，则count=1
    最终返回max.

    查看题解，发现有双指针，类似于滑动窗口的概念，也试着写一写吧。

    写完之后发现我自己写的代码，虽然AC了，但是max和count的初始化都是从1开始的，比较也是从第二个元素开始的，其实当时犹豫过s为空呢，后来看到提示说非空字符串，我的初始化值不会有问题了。现在对比一下滑动窗口的代码，反而觉得更加通用和巧妙。


*/
// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
 var maxPower = function(s) {
    let max=1,count=1
    for(let i=1;i<s.length;i++){
        if(s[i]==s[i-1]){
            count++;
            max=Math.max(count,max)
        }
        else{
            count=1
        }
    }
    return max
};


//双指针-滑动弹窗
var maxPower = function(s) {
    let max=0
    let i=0;
    while(i<s.length){
        let j=i;
        while(s[i]==s[j]){
            i++;
        }
        max=Math.max(max,i-j)
    }
    return max
};


/*
    时间复杂度都为：O(N)
    空间复杂度都为：O(1)
*/