// 557. 反转字符串中的单词 III
// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

 

// 示例：

// 输入："Let's take LeetCode contest"
// 输出："s'teL ekat edoCteeL tsetnoc"
 

// 提示：

// 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

/*
    思路：一开始想用的是js的数组api,reverse()反转数组
        因此：1.字符串变为数组
             2.数组遍历item。item接着split变为数组，reverse()反转，join()变为字符串。然后每一项push进入目标数组
             3.数组变为字符串返回

        看js某位大佬题解。
        var reverseWords = function(s) {
            let arr=s.split("").reverse().join("")
            return arr.split(" ").reverse().join(' ')
        };

        看官方题解。不过对它的时间复杂度O(N)存在质疑。
        遍历过程，遇到空格，首位交换之前的单词
        而且用js还遇到一个小坑。人家用c语言，卡卡交换，直接字符串某一位就可以交换了，但是我们的还是要换成数组进行对应下标进行交换。果然不能语言之间有些东西是不能照搬滴。

*/

/*
 * @lc app=leetcode.cn id=557 lang=javascript
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
 var reverseWords = function(s) {
    let words=s.split(" "),res=[]
    words.forEach(item=>{
        let reverseStr=item.split('').reverse().join('')
        res.push(reverseStr)
    })
    let resStr=res.join(" ")
    return resStr
};
// @lc code=end

var reverseWords = function(s) {
    let arr=s.split("").reverse().join("")
    return arr.split(" ").reverse().join(' ')
};
/*
    前两个方法：时间复杂度：O(N)
              空间复杂度：O(N)
*/

var reverseWords = function(s) {
    s=s.split('')
    let i=0,len=s.length
    while(i<len){
        let start=i
        while(i<len&&s[i]!=' '){
            i++;
        }
        let left=start,right=i-1
        while(left<right){
            let tmp=s[left]
            s[left]=s[right]
            s[right]=tmp;
            left++;
            right--;
        }
        while(i<len&&s[i]==' '){
          i++;
        }
    }
    s=s.join('')
    return s
  };

  /*
    时间复杂度：O(N)。题解说每个字符是O(1)的时间交换。
    但是我觉得，while(left<right)相当于又回去遍历交换了一次呀。这个位置有点懵
    空间复杂度：O(1)
  */