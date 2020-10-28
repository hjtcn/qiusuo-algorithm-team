func removeKdigits(num string, k int) string {
    stack := []byte{}
    // 遍历字符串入栈
    for i := 0;i < len(num);i++ {
        char := num[i]
        for len(stack) > 0 && char < stack[len(stack) - 1] && k > 0 {
            k--
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, char)
    }
    //fmt.Println(stack)
    // fmt.Println('0')
    for len(stack) > 0 && stack[0] == '0' && len(stack) > k {
        stack = stack[1:len(stack)]
    }
    // // 去除栈中0元素 有bug
    // for i := 0;i < len(stack);i++ {
    //     fmt.Println("i is",i,stack[i])
    //     if stack[i] != '0' {
    //         break
    //     }else {
    //         //  遍历 stack 的时候改变了stack 的值
    //         stack = stack[i + 1:len(stack)]
    //     }
    //     fmt.Println(stack)
    // }
    // 判断 k 是否大于0
    if k > 0 {
        //fmt.Println(k,len(stack))
        stack = stack[:len(stack) - k]
    }
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}

看着代码和注释也是我的一个思考和改进的过程
中间去除栈中头部元素的0的时候出现了一个bug，因为你遍历stack 切片的以后，你后面也改动了切片，所以出现了bug,如果在工程中遇到还是蛮难排查的
    
    // 去除栈中开头0元素
    for i := 0; i < len(stack); i++ {
        fmt.Println("i is", i, stack[i])
        if stack[i] != '0' {
            break
        } else {
            //  遍历 stack 的时候改变了stack 的值
            stack = stack[i+1 : len(stack)]
        }
        fmt.Println(stack)
    }

优化：如果10200 这种情况如果是第一个0可不可以不放进去？
func removeKdigits(num string, k int) string {
    stack := []byte{}
    // 遍历字符串入栈
    for i := 0;i < len(num);i++ {
        char := num[i]
        for len(stack) > 0 && char < stack[len(stack) - 1] && k > 0 {
            k--
            stack = stack[:len(stack) - 1]
        }
        
        // 如果栈为空，并且接下来首个元素是0，则不把这个元素append到栈中
        if len(stack) == 0 && num[i] == '0' {
            continue
        }
        
        stack = append(stack, char)
    }
    
    // 注意这里也是要判断 len(stack) 和 k 到关系，是 >= 哈
    if k > 0 && len(stack) >= k{
        stack = stack[:len(stack) - k]
    }
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}

中间还是很多细节还是很迷人的，比如 10200 k=1这种
比如递增序列：123456789: k=3 这种在工程中也是经常出现的
这种也提示我们要注意边界条件
3个特殊边界值：
1. 10200 k=1
2. 123456789 k = 3
3. 10 k = 2

总结：
1. 学基础知识最好的还是体系化的学习，看一本书或者是从官网中开始学习，然后动手做（Go语言学习）

2. 调试，调试的思考是需要不断思考的，在调试之前要想自己是想要出现什么结果，运行以后出现了什么结果，然后思考改进

3. 明白了为什么大厂或者外企要求基础知识扎实了，因为面对复杂的分布式系统或者大型的系统的时候，基础知识不扎实，你是很难定位问题，解决问题的，是没有办法靠猜去解决的，需要体系化，工程化和科学的方法去解决这些出现的复杂问题
