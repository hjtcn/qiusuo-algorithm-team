// 647. 回文子串
// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

 

// 示例 1：

// 输入："abc"
// 输出：3
// 解释：三个回文子串: "a", "b", "c"
// 示例 2：

// 输入："aaa"
// 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 

// 提示：

// 输入的字符串长度不会超过 1000 。

/*
    暴力。判断当前子串是否是回文，然后不断记录

*/
// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
 var countSubstrings = function(s) {
    let count=0
    //用api反转会超时，有点刺激了，以后自己写判断回文的方法吧。
    // let isPalindrome=(s)=>{
    //     return s==s.split('').reverse().join('')
    // }
    let isPalindrome=(s)=>{
        let i=0;j=s.length-1
        while(i<j){
            if(s[i]!==s[j]){
             return false
           }
          i++;
          j--;
        }
        return true;
     }
    for(let i=0;i<s.length;i++){
        for(let j=0;j<s.length;j++){
            if(isPalindrome(s.slice(i,j))){
                count++;
            }
        }
    }
    return count
};

/*
    时间复杂度：O(N^3)
    空间复杂度：O(1)
*/
// @lc code=end


//方案2:找回文中心，向两端发散，遇到两端不等的停止发散。
/*
    如何寻找回文中心呢。回文中心可以有一个或两个组成
    如'ab'中心点可以为a,b,ab，回文中心为ab 左位置为为0，右位置1
      'abc'中心点可以为a,b,c,ab,bc
      'abcd'中心点可以为a,b,c,d,ab,bc,cd
      'abcde'中心点可以为a    0   left=0,right=0
                       ab   1   left=0,right=1
                       b    2   left=1,right=1
                       bc   3   left=1,right=2
                       c    4   left=2,right=2
                       cd   5   left=2,right=3
                       d    6   left=3,right=3
                       de   7   left=3,right=4
                       e    8   left=4,right=4
        9个
    来吧，找规律。回文中心个数为2n-1个
    回文中心左起始位置
*/

/**
 * @param {string} s
 * @return {number}
 */
 var countSubstrings = function(s) {
    let centerLen=s.length*2-1,count=0;
    for(let i=0;i<centerLen;i++){
        let left=parseInt(i/2),right=left+i%2
        while(left>=0&&right<s.length&&s[left]==s[right]){
            left--;
            right++;
            count++;
        }
    }
    return count
}

/*
    时间复杂度：O(N^2)
    空间复杂度：O(1)
*/
//动态规划.理解完思路后自己尝试做了一遍，没想到又踩坑了。
// i,j的遍历，要保持i一直小于j，这样才不会导致判断dp[i+1][j-1]的时候，可能它还没赋过值。
var countSubstrings = function(s) {
    let len=s.length,dp=[],count=0
    for(let i=0;i<len;i++){
        dp[i]=new Array(len).fill(0)
    }
    for(let j=0;j<len;j++){
        for(let i=0;i<=j;i++){
            if(j==i){
                dp[i][j]=1
                count++;
            }
            else if(s[i]==s[j]&&j-i==1){
                dp[i][j]=1
                count++;
            }
            else if(j-i>1&&s[i]==s[j]&&dp[i+1][j-1]){
                dp[i][j]=1
                count++;
            }
        }
    }
    return count
}
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N^2)
*/

// 一维压缩优化 
var countSubstrings = function(s) {
    let len=s.length,count=0
    let dp=new Array(len).fill(0)
    for(let j=0;j<len;j++){
        for(let i=0;i<=j;i++){
            //三种类型可合并
            if(j==i){
                dp[i]=1
                count++;
            }
            else if(s[i]==s[j]&&j-i==1){
                dp[i]=1
                count++;
            }
            else if(j-i>1&&s[i]==s[j]&&dp[i+1]){
                dp[i]=1
                count++;
            }
            else{
                //本来想着自己初始化为0了，就没加这一步，后来发现会出错
                //因为需要一种每一次遍历，需要一次更新
                dp[i]=0
            }
        }
    }
    return count
}
/*
    时间复杂度：O(N^2)
    空间复杂度：O(N)
*/


//代码优化
var countSubstrings = function(s) {
    let len=s.length,count=0
    let dp=new Array(len).fill(0)
    for(let j=0;j<len;j++){
        for(let i=0;i<=j;i++){
           if(j==i||(s[i]==s[j]&&j-i==1)||(j-i>1&&s[i]==s[j]&&dp[i+1])){
                dp[i]=1
                count++;
           }
            else{
                dp[i]=0
            }
        }
    }
    return count
}


//题解中还看到一个蛮巧妙的解法，正反向记录字符串
//抛弃算法，也算是暴力的一种吧。
var countSubstrings = function(s) {
    let count=0,len=s.length
    for(let i=0;i<len;i++){
        let str='',reverseStr=''
        for(let j=i;j<len;j++){
            str+=s[j]
            reverseStr=s[j]+reverseStr
            if(reverseStr==str){
                count++
            }
        }
    }
    return count
}

// ？？？？马拉车算法啃不动，先缓一缓。
// 可将时间复杂度降为O(N)，空间复杂度O(N)