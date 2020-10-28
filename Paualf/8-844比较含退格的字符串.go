/*
    使用两个栈，然后遍历S和T,去除#号前的字符，最后比较去除后的字符串是否相等
*/
func backspaceCompare(S string, T string) bool {
    stack1 := []rune{}
    stack2 := []rune{}
    for _,s := range S {
        if s == '#' {
            if len(stack1) > 0 {
                stack1 = stack1[:len(stack1) - 1]
            }
        }else {
            stack1 = append(stack1, s)
        }
    }
    for _,s := range T {
        if s == '#' {
            if len(stack2) > 0 {
                stack2 = stack2[:len(stack2) - 1]
            }
        }else {
            stack2 = append(stack2, s)
        }
    }
    return string(stack1) == string(stack2)
}

有重复的逻辑，是否可以将重复的逻辑抽象出来作为函数
/*
    使用两个栈，然后遍历S和T,去除#号前的字符，最后比较去除后的字符串是否相等
*/
func backspaceCompare(S string, T string) bool {
    stack1 := backspace(S)
    stack2 := backspace(T)
    return string(stack1) == string(stack2)
}
// 抽象出来函数，输入字符串 返回为 rune 的切片
func backspace(S string)[]rune {
    stack :=  []rune{}
    for _,s := range S {
        if s == '#' {
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }
        }else {
            stack = append(stack, s)
        }
    }
    return stack
}

修改if判断语句，增加可读性和减少代码量
/*
    使用两个栈，然后遍历S和T,去除#号前的字符，最后比较去除后的字符串是否相等
*/
func backspaceCompare(S string, T string) bool {
    stack1 := backspace(S)
    stack2 := backspace(T)
    return string(stack1) == string(stack2)
}
// 抽象出来函数，输入字符串 返回为 rune 的切片
func backspace(S string)[]rune {
    stack :=  []rune{}
    for _,s := range S {
        if s != '#' {
           stack = append(stack, s)
        }else if len(stack) > 0 {
            stack = stack[:len(stack)-1]
        }
    }
    return stack
}
复杂度分析：
时间复杂度：O(M+N)，遍历S和T字符串的长度和
空间复杂度：O(M+N)，开辟栈空间的大小
双指针法：
/*
    双指针法，从后向前遍历，如果是#号就向前遍历，然后比较字符是否相等，直到比较的字符不相等，然后return false
*/
func backspaceCompare(S string, T string) bool {
    skipS,skipT := 0,0
    i,j := len(S)-1,len(T) - 1
    for i >= 0 || j >= 0 {
        for i >= 0 {
            if S[i] == '#' {
                skipS++
                i--
            }else if skipS > 0 {
                skipS--
                i--
            }else {
                break
            }
        }
        for j >= 0 {
            if T[j] == '#' {
                skipT++
                j--
            }else if skipT > 0 {
                skipT
                j--
            }else {
                break
            }
        }
        if i >= 0 && j >= 0 {
            if S[i] != T[j] {
                return false
            }
        }else if  i >= 0 || j >= 0 {
            return false
        }
        i--
        j--
    }
    return true
}

复杂度分析：
时间复杂度：O(M+N): 字符串S和T的长度
空间复杂度：O(1): 定义变量i,j,skipS,skipT

总结：
1. 双指针法没有栈容易想到

2. 程序的优化，最后一定是数学程度上的优化,从题目可以看出空间复杂度从O(M+N) -> O(1):数学的魅力

3. 优化代码和修改自己代码的过程需要不断抽象，思考，改进，是挺费脑力的一个过程,enjoy it
