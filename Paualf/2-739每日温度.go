暴力法：遍历数组，从每一项后面找比它大的那一项所在的位置

func dailyTemperatures(T []int) []int {
    // 使用make or []int
    ret := make([]int, 0)
    
    for i := 0;i < len(T);i++ {
        temp := 0
        for j := i + 1;j < len(T);j++ {
            if T[j] > T[i] {
                temp = j - i
                break
            }
        }
        ret = append(ret, temp)
    }
    
    return ret
}

单调栈：

func dailyTemperatures(T []int) []int {
    length := len(T)
    stack := make([]int, 0)
    ret := make([]int, length)
    for i := 0;i < length;i++ {
        temperature := T[i]
         for len(stack) > 0 && temperature > T[stack[len(stack) - 1]] {
             prevIndex := stack[len(stack) - 1]
             stack = stack[:len(stack) - 1]
             ret[prevIndex] = i - prevIndex
         }
        stack = append(stack, i)
    }
    return ret
}

总结：
1. 一开始只想到暴力法，没想到使用栈怎么遍历数组

2. 看了官方题解以后，看文字还是很累，看了图解的单调栈以后回来看了文字解释，自己画了下图，才有一些理解，官方的next数组目前没有看懂

3. 使用栈还是很巧妙的，代码很短，里面蕴含的思想和思维还是值得借鉴和学习的
