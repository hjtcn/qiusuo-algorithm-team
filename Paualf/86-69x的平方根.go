实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
示例 1:
输入: 4
输出: 2
示例 2:
输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
由于返回类型是整数，小数部分将被舍去。

1. Clearfication:
开根号，4开根号是2，8开根号是2，那么怎么把开根号的过程变成模拟操作呢？
暴力法：
i从1 循环到 x / 2,如果 i * i < x 则x++
如果 i * i > x,则退出
返回i
二分方法, left = 0,right = x /2,[left,right] 里面进行查找,mid = (left + right) / 2
如果 
if left >= right {
return left
}
else if mid * mid > x 则 right = mid - 1
else if mid * mid < x 则 left = mid + 1
mid * mid == x 则返回mid

2. Coding:
有问题的代码，没有正确的考虑边界条件, 特别是返回值的判断
func mySqrt(x int) int {
    for i := 1;i <= x /2;i++ {
        if i * i == x {
            return i
        }else if i * i > x {
            return i - 1
        }else if i * i < x {
            i++
        } 
    }
    
    return -1
}

func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    i := 1
    for ;i <= x /2;i++ {
        if i * i == x {
            return i
        }else if i * i > x {
            return i - 1
        }else if i * i < x {
            
        } 
    }
    
    return i - 1
}

使用二分查找进行实现
left <= right 的时候，退出条件是 left == right 这个时候 left == right 的时候，mid = (left + right) / 2 的还没有判断，所以最后要判断 left * left 是否为 x
然后还有可能left == right 的时候，退出的时候是第一个比它大的元素
func mySqrt(x int)int {
    left,right := 0,x / 2
    
    for left <= right {
        mid := left + (right - left) / 2
        
        if mid * mid == x {
            return mid
        }else if mid * mid > x {
            right = mid - 1
        }else if mid * mid < x {
            left = mid + 1
        }
    }
    
    if left * left == x {
        return left
    }
    return left - 1 
}

上面的分析自己分析错了，最后退出的条件是 left > right 了，所以有可能退出的时候是第一个比它大的元素，所以最后返回条件是 left - 1，我们只要将 x <= 1 进行判断就可以了
func mySqrt(x int)int {
    if x <= 1 {
        return x
    }
    left,right := 1,x / 2
    for left <= right {
        mid := left + (right - left) / 2
        
        if mid * mid == x {
            return mid
        }else if mid * mid > x {
            right = mid - 1
        }else if mid * mid < x {
            left = mid + 1
        }
    }
    return left - 1 
}

如果判断条件是 left < right 的呢，怎么写呢, [left,right), 最后退出的判断条件是 left == right,最后返回条件是什么呢？
这个返回值没有想清楚
func mySqrt(x int)int {
    left,right := 1,x / 2
    
    for left < right {
        mid := left + (right - left) / 2
        if mid * mid == x {
            return x
        }else if mid * mid > x {
            mid = right
        }else if mid * mid < x {
            left = mid + 1
        }
    }
    
    if left * left == x {
        return left
    }
    
    return left - 1
}

上面代码提交了好几次：终于改对了
func mySqrt(x int)int {
    left,right := 1,x / 2
    
    for left < right {
        mid := left + (right - left) / 2
        
        if mid * mid == x {
            return mid
        }else if mid * mid > x {
            right = mid
        }else if mid * mid < x {
            left = mid + 1
        }
    }
    
    if left * left <= x {
        return left
    }
    
    return left - 1
}

比较有意思的是 left * left <= x 的判断条件
3. 看题解：

4. 复杂度分析：
时间复杂度：O(logN）
空间复杂度: O(1)

5. 总结：
5.1： 注意 left 和 right 定义，以及 left <= right 还是 left < right 还有最后是否越界的判断及最后的返回值是否正确
