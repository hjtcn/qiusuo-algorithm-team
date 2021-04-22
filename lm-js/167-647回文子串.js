/*
647. 回文子串
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

 

示例 1：

输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 

提示：

输入的字符串长度不会超过 1000 。

*/

/*
    思路：先试试动态规划吧
        dp[i][j],代表从i到j是否是回文串,打标记。
        a    a    a
        dp[0][0]   1
        dp[0][1]   1
        dp[0][2]   1
        三种回文串的情况：
        1.i==j,dp[i][j]=1,count+=1
        2.j-i==1&&s[i]==s[j] dp[i][j]=1,count+=1
        3.j-i>1&&s[i]==s[j]&&dp[i+1][j-1] dp[i][j]=1,count+=1
*/

//错误代码：
for(let i=0;i<len;i++)
{
    for(let j=i;j<len;j++){
        if(i==j){
            count+=1
            dp[i][j]=1
        }
        else if(j-i==1&&s[i]==s[j]){
            count+=1
            dp[i][j]=1
        }
        else if(j-i>1&&s[i]==s[j]&&dp[i+1][j-1]){
            count+=1
            dp[i][j]=1
        }
    }
}

/*
错误代码的思路是正确的，问题出在哪呢？
在于我的状态转移的网格模型制造错了。dp[i+1][j-1]代表这i向外扩，j向内缩。这样，我之前的标记是不是就没有任何用了。因为i还没有走到呢。

换乘图形化来说：因为标记的复用，我们网格模型只能从下半块进行标记。
(i,j)    a   a   a
  a      1
  a      1   1
  a      1   1   1
*/

var countSubstrings = function(s) {
    let len=s.length,count=0;
    let dp=Array.from(Array(len),()=>Array(len).fill(0))
      for(let j=0;j<len;j++)
         {
             for(let i=0;i<j;i++){
                 if(i==j){
                     count+=1
                     dp[i][j]=1
                 }
                 else if(j-i==1&&s[i]==s[j]){
                     count+=1
                     dp[i][j]=1
                 }
                 else if(j-i>1&&s[i]==s[j]&&dp[i+1][j-1]){
                     count+=1
                     dp[i][j]=1
                 }
             }
         }
    return count
 }
/*
 时间复杂度：O(N^2)
 空间复杂度：O(N^2)
*/

//想想还是把暴力给写了一遍，熟能生巧嘛
var countSubstrings = function(s) {
    let len=s.length,count=0;
    let isPalindrome=(str)=>{
       let l=0,r=str.length-1
       while(l<r){
           if(str[l]!=str[r]){
               return false
           }
           else{
               l++
               r--
           }
       }
       return true
    }
    for(let i=0;i<len;i++){
        for(let j=i+1;j<=len;j++){
            if(isPalindrome(s.slice(i,j))){
                count+=1
            }
        }   
    }
    return count
 }
 /*

 时间复杂度：O(N^2)
 空间复杂度：O(N) isPalindrome方法应该是用了隐式存储整个字符串的长度
*/

 /*
 思考：前两题做的是思路，而今天把模型拆解的更细致了。
 */
