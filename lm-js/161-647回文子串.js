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
    思路：先写第一种方法。暴力匹配。我还记得当时判断是否是回文串利用js Api反而会超时。因此，这个方式还是比较熟练的
*/
var countSubstrings = function(s) {
    let count=0
    let isPalindrome=(str)=>{
        if(!str.length) return false
        let l=0,r=str.length-1
        while(l<r){
            if(str[l]==str[r]){
                l++;
                r--;
            }
            else{
                return false
            }
        }
        return true
    }
    for(let i=0;i<s.length;i++){
        for(let j=0;j<s.length;j++){
            if(isPalindrome(s.slice(i,j+1))){
                count++;
            }
        }
    }
    return count
}
/*
    时间复杂度：O(N^3)
    空间复杂度：O(1)
*/


//动态规划没有想出来，去看了题解，发现这次dp[]不同的是不记录值了，而是作为标记了。
/*
    和以往不同的点：1.dp[i][j]代表从i到j是否为回文串。
                 2.两层for循环遍历。j在最外层。i作为内层，不断增大，且不能超过j
                 3.对比过程是不断的扩展过程.1）i==j，一个字母，肯定是回文串，count++，dp  [i][j]=1打个标记，从i~j为回文串。
                                        2)j-i==1,两个字母，因此当s[i]==s[j]时，是回文串，count++,dp[i][j]=1打个标记，从i~j为回文串。
                                        3）j-i>1,两个字母以上，dp[i+1][j-1]为true向内缩的字符串为回文串，且s[i]==s[j],因此，count++;
*/

var countSubstrings = function(s) {
    let len=s.length,count=0
    let dp=Array.from(Array(len),()=>Array(len).fill(0))
    for(let j=0;j<len;j++){
        for(let i=0;i<=j;i++){            
            if(i==j){
                dp[i][j]=1
                count++;
            }
            else if(s[i]==s[j]&&j-i==1){
                //两位数
                dp[i][j]=1 //这个是作为标记，而不是存值了
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
//降维优化。
var countSubstrings = function(s) {
    let len=s.length,count=0
    let dp=new Array(len).fill(0)
    for(let j=0;j<len;j++){
        for(let i=0;i<=j;i++){            
            if(i==j){
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
/*
    仔细想了想这个降维，没有总结出它的规律，只能走走case,看看它的合理性
    比如'aaba'
    j==1||j==0,都是回文，dp[]肯定一直为1.
    当j=2时，i的一层dp遍历。i=0时，'aab',j-i>1,s[i]!=s[j],故dp[i]=0   
                         i=1时，'ab',j-i==1,s[i]!=s[j],故dp[i]=0
                         i=2时，'b',j==i,故dp[i]=1

    当j=3时，i的一层dp遍历。i=0时，'aaba',j-i>1,s[i]==s[j],但是dp[i+1]=0,故，dp[i]=0
                         i=1时，'aba',j-i>1,s[i]==s[j],dp[i+1]=1,故dp[i]=1
                         i=2时，'ba',j-i==1,s[i]!=s[j],故dp[i]=0
                         i=3时，'a',j==i,s[i]==s[j].故dp[i]=1
    未降维之前dp[i][j]代表，从i走到j，是否是回文
    降维之后dp[i]：会在上一次循环中已经赋值过dp[i](从0～j),而当次循环查找dp[i+1]对应的末位置为j-1,因此也是和未降维之前是对应的。
*/

/*
    思考：1.今天的题之前做的时候，有种照猫画虎的感觉，回过头来只记得暴力了。回文中心忘完了，动规也无从下手。
        2.今天再做就感觉面对动规更自信了，但是它确实有些巧妙，从存值，变为了存状态。
        3.今天也不去合并if和else的代码了，想让它逻辑性更强，印象更深刻一下
        4.回文中心也先搁一搁，今天就只为动规。
        5.期待这个题再出一遍，再加深一下记忆，这些比网格题难度可高太多了。
*/