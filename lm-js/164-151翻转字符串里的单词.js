// 151. 翻转字符串里的单词
// 给定一个字符串，逐个翻转字符串中的每个单词。

// 说明：

// 无空格字符构成一个 单词 。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 

// 示例 1：

// 输入："the sky is blue"
// 输出："blue is sky the"
// 示例 2：

// 输入："  hello world!  "
// 输出："world! hello"
// 解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 示例 3：

// 输入："a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
// 示例 4：

// 输入：s = "  Bob    Loves  Alice   "
// 输出："Alice Loves Bob"
// 示例 5：

// 输入：s = "Alice does not even like bob"
// 输出："bob like even not does Alice"
 

// 提示：

// 1 <= s.length <= 104
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词


/*
    思路：1.首先我想的是正则和js的api。match(/[\S+]/ig),以空格进行匹配返回数组。
         2.遍历赋值，构建对应的字符串，并用' '进行拼接
         遍历主要还是注重拼接时的细节控制。
         倒序遍历，判断是否为空格：
         1）是：则res+=' '+str(当然res不为空时，用' ',为空，直接res=str赋值即可)
         2）否：则str=s[i]+str，拼接str,记得在上一步及时将str置为空
         细节点：
         1）空格只能保持一个，因此用flag记录是否是单词间第一次出现空格。
         2）返回res时，最后一个单词str会忘了拼接。但是也有可能只有一个单词，即res为空。因此需要兼容控制。
            若res不为空，用res+' '+str拼接
            res为空或str为空，则返回res||str
*/

var reverseWords = function (s) {
    let res='',str='',flag=0
    for(let i=s.length-1;i>=0;i--){
        if(s[i]==' '){
            if(!flag){
                if(res){
                    res+=' '+str
                }
                else{
                    res=str
                }
            }
                str=''
                flag=true
        }
        else{
            str=s[i]+str
            flag=false
        }
    }
    if(res&&str){
        return res+' '+str
    }
    else{
        return res||str
    }
    
};


var reverseWords = function (s) {
    return s.match(/[\S]+/ig).reverse().join(' ')
};

/*
    时间复杂度:O(N).api不确定，但是正则肯定要比遍历复杂一些的。
    空间复杂度:O(N)
*/

/*这次看题解又发现了一个比较简单理解的双端队列，其实主要双端相互赋值吧.
    其实核心难点不在于方法，而在于细节，尤其是空格的控制，而这个方法，把空格匹配清空了，之后的遍历赋值也显得比较容易理解了。
*/
var reverseWords = function (s) {
    s=s.match(/[\S]+/ig)
    let l=0;r=s.length-1
    while(l<=r){
        let tmp=s[l]
        s[l++]=s[r]
        s[r--]=tmp
    }
    return s.join(' ')
};
