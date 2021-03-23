/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:
输入: "race a car"
输出: false
*/
1. Clearfication:
    判断是否是回文串
    只考虑数字和字母，忽略字母的大小写
    首尾进行判断比较
    判断是否是字母和数字，是则才进行比较
     判断是否是字母，转换成小写

2. Coding:
/*
    
    A man, a plan, a canal: Panama
*/
func isPalindrome(s string) bool {
    i,j := 0,len(s) - 1
    for i <= j {
        if !isValid(s[i]) {
            i++
        }
        if !isValid(s[j]) {
            j--
        }
    }
    return true
}
func isValid(s byte) bool {
}
func conventToLower(s byte) byte {
}
上面是自己写的半成品。。。

3. 看题解:
自己写isValid 的时候有点犹豫，还是写的代码太少
func isPalindrome(s string)bool {
    var sgood string
    for i := 0;i < len(s);i++ {
        if isalnum(s[i]) {
            sgood += string(s[i])
        }
    }
    
    n := len(sgood)
    sgood = strings.ToLower(sgood)
    for i := 0;i < n / 2;i++ {
        if sgood[i] != sgood[n - 1 - i] {
            return false
        }
    }
    
    return true
}
func isalnum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
sgood += string(s[i]) 这里一直写成了 sgood += string(s) 找了白天。。。
原字符串上比较：
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left,right := 0,len(s) - 1
    
    for left < right {
        for left < right && !isalnum(s[left]) {
            left++
        }
        
        for left < right && !isalnum(s[right]) {
            right--
        }
        
        if left < right {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
    }
    
    return true
}

func isalnum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
func isPalindrome(s string) bool {
    for i,j := 0,len(s) - 1;i < j;i,j = i + 1,j - 1 {
        for ;i < j && !isChar(s[i]);i++ {}
        
        for ;i < j && !isChar(s[j]);j-- {}
        
        if toLower(s[i]) != toLower(s[j]) {
            return false
        }   
    }
    return true
}
func isChar(c byte) bool {
    return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c - 'A' + 'a'
    }
    
    return c
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1) or O(n)， O(n) 是使用了中间字符串存储，过滤掉不合法字符

5. 总结:
5.1. 对没有写过的代码还是不敢写，有点害怕
5.2. 对字符串的操作不熟悉 
5.3. 代码基本功比较差，要多写代码
