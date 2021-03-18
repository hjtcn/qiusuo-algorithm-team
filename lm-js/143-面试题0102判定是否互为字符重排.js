// 面试题 01.02. 判定是否互为字符重排
// 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

// 示例 1：

// 输入: s1 = "abc", s2 = "bca"
// 输出: true 
// 示例 2：

// 输入: s1 = "abc", s2 = "bad"
// 输出: false
// 说明：

// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100
/*
    思路：方法1:想到了使用map，记录长度较短的字符串的元素出现次数，然后长度较长的字符串遍历时开始减少次数。次数出现小于0时的情况就返回false,反之跳出循环则返回true
         方法2:题解思路：js使用api排序还是很爽滴。
         
         关于时间复杂度和空间复杂度，两种方法应该都是O(N)吧，api也是有隐式占用空间的。


*/
/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
 var CheckPermutation = function(s1, s2) {
    let map=new Map()
    let find=(longS,shortS)=>{
       for(let i=0;i<longS.length;i++){
           map.set(longS[i],(map.get(longS[i])||0)+1)
       }
       for(let i=0;i<shortS.length;i++){
           map.set(shortS[i],(map.get(shortS[i])||0)-1)
           if(map.get(shortS[i])<0){
               return false
           }
       }
       return true
    }
    return s1.length>s2.length?find(s2,s1):find(s1,s2)

};


var CheckPermutation = function(s1, s2) {
    return s1.split('').sort().join('')==s2.split('').sort().join('')
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/