func isValid(s string) bool {
    n := len(s)
    if n % 2 == 1 {
        return false
    }
    pairs := map[rune]rune {
        ')' : '(',
        ']' : '[',
        '}' : '{',
    }
    stack := []rune{}
    for _,char := range s {
        if char == '(' || char == '[' || char == '{' {
            stack = append(stack, char)
        }else if len(stack) > 0 && pairs[char] == stack[len(stack) - 1] {
            stack = stack[:len(stack) - 1]
        }else {
            return false
        }
    }
    
    return len(stack) == 0
}
看题解和总结中看到有些人使用了 range 进行遍历对应类型 rune，使用byte 类型遍历对应 for i := 0;i < len(s);i++ {}
上述代码中stack rune 类型修改为 byte 时候出现了报错：
cannot use char (type rune) as type byte in append (solution.go)
猜测：使用  range 遍历的时候遍历的是 slice,底层存储是数组，可能主动将类型修改为了int32 也就是  rune
使用 for 进行循环的时候，遍历字符串，里面存储的是字节，类型为uint8,写一个小程序验证一下：
func main() {
    s := "([)]"
    for i, char := range s {
        fmt.Printf("index:%d Type:%T value :%v\n", i, char, char)
    }
    for i := 0; i < len(s); i++ {
        fmt.Printf("Type:%T value :%v\n", s[i], s[i])
    }
}
/*
输出结果为：
index:0 Type:int32 value :40
index:1 Type:int32 value :91
index:2 Type:int32 value :41
index:3 Type:int32 value :93
Type:uint8 value :40
Type:uint8 value :91
Type:uint8 value :41
Type:uint8 value :93
*/
总结：
• 题目总结：
一开始知道是使用栈进行做题，题目以前也写过，还是没有写出来，思路也不是很流程，后面看题解以及对比其他人提交的代码才有所收获，如一开始可以判断奇偶性，使用map对左右括号进行映射
题目做过再次做不一定可以做对，知道方向和思路也不一定可以写出来，这些情况说明基础知识以及基本功不够扎实，无论是对编程语言还是数据结构都是有很大欠缺的。
• Go中字符串是不可变的，那么如何遍历字符串呢
• 做题流程
1. 题意  动手画图理解
2. possible solutions
3. compare time/space complex
4. coding & test 
• 关于语言使用方面：
Go中 rune 和 byte是什么呢？有什么区别呢？
byte: uin8 类型，uint8的别名，代表了ASCII码的一个字符
rune：文字符号，为 int32 类型的别名，代表一个UTF-8 字符
UTF-8 和Unicode 有何区别？
Unicode 与 ASCII类似，都是一种字符集。
字符集为每个字符分配一个唯一的ID,我们使用到的所有字符在Unicode字符集中都有一个唯一的ID,a 在 Unicode 与ASCII中的编码都是97。汉字"你"在Unicode 中的编码为20320，在不同国家的字符集中，字符所对应的ID也会不同。无论任何情况下，Unicode中的字符的ID都是不会变化的。
UTF-8是编码规则，将Unicode中字符的ID以某种方式进行编码，UTF-8 是一种变长编码规则，从1到4个字节不等。编码规则如下：
• 0xxxxxx 表示文字符号 0～127，兼容ASCII字符集。
• 从128 到0x10fff 表示其他字符
根据这个规则，拉丁文语系的字符编码一般情况下每个字符占用一个字节，而中文每个字符占用3个字节。
广义的Unicode 指的是一个标准，它定义了字符集及编码规则，即Unicode字符集和UTF-8、UTF-16编码等。
参考：
（Go语言字符类型）[http://c.biancheng.net/view/18.html]
