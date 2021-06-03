给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：
输入：digits = ""
输出：[]
示例 3：
输入：digits = "2"
输出：["a","b","c"]
 
提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

1. Clearfication:
   使用  map 数字与字母对应的关系信息
            如 2 =》 s[0]:abc
                遍历 abc
                    遍历 s[1] : def
                        ad ae af bd be bf      
        更像是一种搜索算法  

2. Coding:
map 的声明不正确
var m map[rune][string]
 m map[rune][string] = [
    '2':"abc",
    '3':"def",
    '4':"ghi",
    '5':"jkl",
    '6':"mno",
    '7':'pqrs',
    '8':'tuv',
    '9':'wxyz'
]

func letterCombinations(digits string) []string {
    res := []string
    tmp := ""
    helper(0,digits,tmp,&res)
    return res
}

func helper(index int,digits string,tmp string,res *[]string) {
    if index >= len(digits) {
        res = append(res, tmp)
        return
    }
    for i := index;i < len(digits);i++ {
        if value,ok := m[digits[i]];ok {
            for j := 0;i < len(value);j++ {
                tmp = tmp + value[j]
                helper(index + 1,digits,tmp,res)
            }
        }
    }
}

3. 看题解：
看了下题解，对上面的代码进行了debug
 var m map[byte]string = map[byte]string {
    '2':"abc",
    '3':"def",
    '4':"ghi",
    '5':"jkl",
    '6':"mno",
    '7':"pqrs",
    '8':"tuv",
    '9':"wxyz",
 }

func letterCombinations(digits string) []string {
    res := []string{}
    if len(digits) == 0 {
        return res
    }    
    tmp := ""
    helper(0,digits,tmp,&res)
    return res
}

func helper(index int,digits string,tmp string,res *[]string) {
    if index >= len(digits) {
        *res = append(*res, tmp)
        return
    }
    if value,ok := m[digits[index]];ok {
        for j := 0;j < len(value);j++ {
            helper(index + 1,digits,tmp + string(value[j]) ,res)
        }
    }
  
}

Go里面对map进行遍历尽量使用 for range 
func letterCombinations(digits string) []string {
    ans := []string{}
    if len(digits) == 0 {return ans}
    helper(digits, 0, "", &ans)
    return ans
}

var M = map[byte]string {
    '1': "",
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
    '0': " ",
}

func helper(digits string, pos int, s string, strs *[]string) {
    if pos == len(digits) {
        *strs = append(*strs, s)
        return
    }
    
    for _, c := range M[digits[pos]] {
        helper(digits, pos + 1, s + string(c), strs)
    }
    return
}

4. 复杂度分析：
时间复杂度：O(3^m x 4 ^n),其中m是输入中对应3个字母的数字个数（2，3，4，5，6，8），n是输入中对应4个字母的数字个数（7，9）
空间复杂度：O(m+n)

5. 总结：
5.1: 有递归框架的话感觉没有那么难，就是处理的一些细节，有点搜索的感觉
