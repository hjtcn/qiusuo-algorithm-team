// 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false



/*
思路：自己的思路首先想到了先用正则把字符串清理一下。
     这个卡了一会，因为模糊了正则的运用
     方法1：然后将字符串进行变形反转再进行对比
       strA.split('').reverse().join('')==strA
     方法2：头部shift,尾部pop,直到s的长度为1，return true
     方案3：看了一眼小涛的，才知道用双指针，不为数字或字母，提前结束，其他进行l++,r--


*/

/*
 * @lc app=leetcode.cn id=125 lang=javascript
 *
 * [125] 验证回文串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {boolean}
 */
 var isPalindrome = function(s) {
    s=s.replace(/[^A-Za-z0-9]/g,'').toLocaleLowerCase()
    while(s.length>=2){
        let sHead=s.shift(),sTail=s.pop()
        if(sHead!=sTail){
            return false
        }

    }
    return true
    
};

var isPalindrome = function(s) {
    let strA=s.replace(/[^A-Za-z0-9]/g,'').toLocaleLowerCase()
     let strB=strA.split('').reverse().join('')
    return strA==strB
    
};
// @lc code=end

/**
 * @param {string} s
 * @return {boolean}
 */
 var isPalindrome = function(s) {
    let l=0,r=s.length-1
    while(l<r){
        if(/[^A-Za-z0-9]/.test(s[l])){
            l++
            continue
        }
        if(/[^A-Za-z0-9]/.test(s[r])){
            r--
            continue
        }
        if(s[l].toLocaleLowerCase()!=s[r].toLocaleLowerCase()){
            return false
        }
        l++;
        r--;

    }
    return true
    
};

/* 
    时间复杂度：O(N)
    空间复杂度；O(N)

*/