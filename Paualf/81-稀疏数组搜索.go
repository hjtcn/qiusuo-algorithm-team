/*
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
示例1:
 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
 输出：-1
 说明: 不存在返回-1。
示例2:
 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
 输出：4
提示:
words的长度在[1, 1000000]之间
*/

1. Clearfication:
排好序的字符串数组，找出给定字符串的位置
对于排好序这种结构，一般来说都可以考虑二分查找，去有序的数据里面查找元素。数组的二分查找比较熟悉，对于字符串有点懵懵的，但是理论上应该是一样的
伪代码：
left = 0,right = len(words) -1 
for left < right {
    mid := left + (right - left) / 2
    if arr[mid] == s {
        return mid
    }else if s > arr[mid] {
        left = mid + 1
    }else if s < arr[mid] {
        right = mid - 1
    }
}
上面的代码有可能出错的点在于判断 arr[mid] 是否等于 s 这个判断条件
如果这个思路不行，那么暴力法，从头到尾遍历字符串判断是否相等
for i,v := range words {
    if v == s {
        return i
    }
}
return -1

2. Coding:
暴力法：
func findString(words []string, s string) int {
    if len(words) == 0 {
        return -1
    }
    for i,v := range words {
        if v == s {
            return i
        }
    }
    return -1
}

二分查找：
下面代码执行不正确，我猜出错的地方是判断字符串是否相等的地方
func findString(words []string, s string) int {
    left,right := 0,len(words) - 1
    for left < right {
        mid := left + (right - left) / 2
        if words[mid] == s {
            return mid
        }else if s > words[mid] {
            left = mid + 1
        }else if s < words[mid] {
            right = mid - 1
        }
    }
    return -1
}

3. 看题解：
看完题解以后，发现自己并没有看清楚题意，自己只是觉得是排序好的字符串，没有看到其中散布着一些空字符串，自己没有去处理这些空字符串的情况（审题不清，昨天开会还刚说过这个问题，自己就犯了，哈哈哈）
思路：判断mid是否为空字符串，如果是空字符串则一直向右寻找，找到不为空的进行比较，如果找到最右边发现最右边还是为空，则将 right 改为 mid -1 继续进行搜索, 将right改为原有 mid -1 这个细节，自己一开始向将原有mid记录一下，觉得还是有点麻烦的，后来看到题解里面写到了如果需要用到这个变量就重新计算一下  mid := (left + right) / 2 - 1, 我们在这个过程中只是移动了 mid 的值，这个时候 left 和 right 都没有变化的，所以可以这样求解的

func findString(words []string, s string) int {
    if len(words) == 0 {
        return -1
    }
    left,right := 0,len(words) - 1
    for left <= right {
        mid := left + (right - left) / 2
        for ;mid < right;mid++ {
            if words[mid] != "" {
                break
            }
        }
        if words[mid] == s {
            return mid
        }
        if mid == right {
            right = (left + (right - left) / 2) - 1
            continue
        }
        if s > words[mid] {
            left = mid + 1
        }else if s < words[mid] {
            right = mid - 1
        }
    }
    return -1
}

更加简洁的代码
func findString(words []string,s string) int {
    l,r := 0,len(words) - 1
    
    for l <= r {
        m := (r + l) >> 1
        for ;words[m] == "" && m < r;m++ {
            
        }
        
        if words[m] == s {
            return m
        }else if m == r {
            r = (r + l) >> 1 - 1
        }else if words[m] < s {
            l = m + 1
        }else {
            r = m - 1
        }
    }
    
    return -1
}

迭代：
func findString(words []string,s string) int {
    var f func(int,int) int
    f = func(l,r int) int {
        if l >= r && words[l] != s {
            return -1
        }
        m := (r + l) >> 1
        for ;words[m] == "" && m < r;m++ {
        
        }
        
        if words[m] == s {
            return m
        }
        
        if m == r || words[m] > s {
            return f(l,m - 1)
        }
        
        if words[m] < s {
            return f(m + 1,r)
        }
        
        return -1
    }
    
    return f(0, len(words) - 1)
}

4. 复杂度分析：

为什么要分析时间复杂度和空间复杂度呢？
时间复杂度可以类比做我们需要CPU计算的时间，我们想要它尽可能的少，这样我们就可以让CPU帮我们做其他复杂的运算
空间复杂度也就是我们需要占用多少内存空间，因为内存是有限的，我们使用越少内存，我们程序的效率会越高

复杂度分析：
时间复杂度：暴力搜索：O(n), 二分查找O(logN) -> 一分为2 -> 2^n = s, n = log(s),n 是以2为底就log的树数值
空间复杂度：O(1) 常量空间

5. 总结：    
5.1: 一定要仔细看题，这道题自己就漏掉了散布着一些空字符串

5.2: 程序真正的核心就是处理这些边界条件和 case
