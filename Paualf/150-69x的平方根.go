1. Clearfication:
1.1 朴素的方法：
1. i 从1 开始进行循环，找到第一个数的平方大于x，然后终止循环
2. 返回 i - 1
3. 最后不知道返回什么了，返回了  -1 
/*
    1. Clearfication:
        for i : 1 -> -1 找到最大的平方根小于x的数
*/
func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    for i := 1;i <= x;i++ {
        if i * i > x {
            return i  - 1
        }
    }
    
    return -1
}

1.2: 二分查找：
二分法抽象出来，找到第一个大于某个值的地方，然后返回它的后一个下标
1. 如果 mid * mid == x 则返回
2. 如果 > x 则收缩右边界
3. 最后返回的时候判断下边界条件，返回它后一个下标

写的时候还是有点不自信，怕写不出来，硬着写完调试一下，也就过了。。。

func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    start := 1
    end := x
    for start < end {
        mid := start + (end - start) / 2
        if mid * mid == x {
            return mid
        }else if mid * mid > x {
            end = mid 
        }else {
            start = mid + 1
        }
    }
    if end * end == x {
        return x
    }
    return end - 1
}

2. 看题解
看代码是简单，里面的每个细节还是很多的，< or <=,返回值的时候怎么返回，二分感觉还是蛮魔鬼的

3. 复杂度分析：
时间复杂度：遍历：O(n)，二分查找O(logn) 
空间复杂度：O(1)

4. 总结：
不要害怕，想清楚以后，干就完了，实在不会写也没有事情的哈
