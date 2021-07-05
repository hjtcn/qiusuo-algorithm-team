1. Clearfication:
      
1.1 朴素点，遍历整个数组

1.2. 二分法：
        排序好的，题眼还是比较重要的
       对排序好的数组处理，一般都可以去想下二分算法可以使用么
       这道题目中不一样的地方在于：
           里面有是  “” 字符串的地方
            如果是空字符串
          1. 记录下该位置ww，然后向后找不是空字符串的地方，
          2. 去判断后面不是空字符串的位置和目标s的大小
          3. 如果不是空字符串的位置 qq，比s大，则更新 right = qq - 1
           如果不是空字符串的位置qq,比s小,则更新 left = ww + 1
比较纠结的地方
       1. 初始化的时候 right 初始化为 len(words) or len(words) - 1
       2. 判断条件判断的时候是 判断 left < right or left <= right 
      3. 对left，right操作的时候是加1，减1还是赋值，这些都是细节。。。
func findString(words []string, s string) int {
    left,right := 0,len(words)
    for left < right {
        mid := left + (right - left) / 2
        if words[mid] == s {
            return mid
        }
        if words[mid] == "" {
            tmp := mid - 1
            for tmp >= 0 {
                if words[tmp] != "" {
                    if words[tmp] == s {
                        return tmp
                    }
                    if words[tmp] > s {
                        right = tmp
                    }
                    if words[tmp] < s {
                        left = mid + 1
                    }
                    break
                }
                tmp--
            }
        }
        if words[mid] > s {
            right = mid
        }
        if words[mid] < s {
            left = mid + 1
        }
    }
    return - 1
}
上面代码多处理了，找到如果words[tmp] != "' 跳出来，然后让外面判断就好了
看题解后想到的

2. 看题解
func findString(words []string, s string) int {
    start, end := 0, len(words)-1
    for start <= end {
        mid := (end + start) / 2
        for ; mid < end; mid++ {
            if words[mid] != "" {
                break
            }
        }
        if words[mid] == s {
            return mid
        }
        if mid == end {
            end = (end+start)/2 - 1
            continue
        }
        if words[mid] < s {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    return -1
}

3. 复杂度分析：
时间复杂度：O(logn)
空间复杂度：O(1)

4. 总结：
4.1: 这道题是二分查找的变形，有空元素的时候怎么处理，这个还是要想一下的
