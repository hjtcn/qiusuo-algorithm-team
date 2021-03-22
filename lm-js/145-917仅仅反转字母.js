// 917. 仅仅反转字母
// 给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

 

// 示例 1：

// 输入："ab-cd"
// 输出："dc-ba"
// 示例 2：

// 输入："a-bC-dEf-ghIj"
// 输出："j-Ih-gfE-dCba"
// 示例 3：

// 输入："Test1ng-Leet=code-Q!"
// 输出："Qedo1ct-eeLg=ntse-T!"
 

// 提示：

// S.length <= 100
// 33 <= S[i].ASCIIcode <= 122 
// S 中不包含 \ or "

/*
    思路：双指针。1.如果都是字母，则左右交换
                2.如果一方有非字母，若left为非字母,则left++,若right,则right--

        我是在字母和ASCil值比较用的charCodeAt.
        刚开始不知道，直接用的元素比大小，发现不太可。
        好吧，总结写到这里，想起来正则判断可以写一写，不考虑代码可读性，我可以练习一下正则嘛，嘻嘻。
        /[^A-Za-z]/g.test(a)

*/
/*
 * @lc app=leetcode.cn id=917 lang=javascript
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
/**
 * @param {string} S
 * @return {string}
 */

 var reverseOnlyLetters = function(S) {
    let left=0,right=S.length-1
    let borderLimit=(a)=>{
        // return /[^A-Za-z]/g.test(a)
        return a.charCodeAt()<65||(a.charCodeAt()>90&&a.charCodeAt()<97)||a.charCodeAt()>122
    }
    S=S.split('')
    while(left<right){
        if(borderLimit(S[left])){
            left++;
            continue
        }
        if(borderLimit(S[right])){
            right--;
            continue
        }
            let tmp=S[left]
            S[left]=S[right]
            S[right]=tmp
            left++;
            right--
    }
    return S.join('')
};

/*
    时间复杂度：O(N)
    空间复杂度：O(1)

*/
// @lc code=end

/*
    查看题解，有很多都是把数组筛选字母复制一遍，然后再pop
    再在遍历的时候，将把字母替换为pop出的元素。
    例如：
    var reverseOnlyLetters = function(S) {
        let arr=S.match(/[a-zA-Z]/g)
        //还学到了replace中遍历pop()
        return S.replace(/[a-zA-Z]/g,()=>arr.pop())
    }

*/