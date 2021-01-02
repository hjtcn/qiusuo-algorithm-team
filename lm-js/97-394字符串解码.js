// 394. 字符串解码
// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

 

// 示例 1：

// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
// 示例 2：

// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
// 示例 3：

// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
// 示例 4：

// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"


/*
 * @lc app=leetcode.cn id=394 lang=javascript
 *
 * [394] 字符串解码
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var decodeString = function(s) {
    //存储轮的重复次数，存储之前的字符串，记录当前轮的重复次数，当前轮的字符串
    let countArr=[],strArr=[],sum=0,result=''
    for(let i=0;i<s.length;i++){ 
        //数字  
        if(!isNaN(s[i])){
            //这里猜了坑，字符串拼接会多出0，需要将字符串换为数字
          sum=sum*10+Number(s[i])
        }
        else if(s[i]=='['){ 
           strArr.push(result)
           countArr.push(sum)
           sum=0
           result=''
        }
        else if(s[i]==']'){
            //先解决内部，再解决外部，这里踩了坑，用了shift()
          let count=countArr.pop()
          result=strArr.pop()+result.repeat(count)
          //这句拼接是核心
        }
        else{
            result+=s[i]
        }
    }
    return result
};
// @lc code=end
/**
    栈，记录重复次数和之前存储的字符串
    核心即碰到']'时，当前轮结束，开始拼接。拼接过程，countArr进行pop得到k次,strArr进行pop+当前轮的result重复k次，这一轮的结果有了(能拼的都拼了)
    时间复杂度:O(N)
    一层for循环
    空间复杂度:O(N)
    存储轮的重复次数countArr，存储之前的字符串strArr
 */


var decodeString = function(s) {
    let dfs=(s,x)=>{
        let result='',sum=0,index=0,i=x
        while(i<s.length)
        {   
            if(!isNaN(s[i])){
                //记录遍历次数
              sum=sum*10+Number(s[i])
              i++
            }
            else if(s[i]=='['){ 
                //如果是左括号，就可以递归，从下一个元素到遇到右括号截止，并且获得结束的索引和括号中的字符串
               let [nextIndex,str]=dfs(s,i+1)
               result+=str.repeat(sum)
               //初始遍历次数和当前的i
                i=nextIndex
               sum=0
            }
            else if(s[i]==']'){
               // 一直卡在记录index,提前就return了
                return [i+1,result]
            }
            else{
                result+=s[i]
                i++
            }
        }
        return result
    }
    return dfs(s,0)
}


/**
 * 递归，记录索引值和当前结果值，更新result并初始化下一个括号内的遍历次数和当前i
 * 时间复杂度：O(N)
 * 记录每一个元素
 * 空间复杂度：O(N)
 * 看题解了解到极端情况下递归深度会到达线性级别
 */