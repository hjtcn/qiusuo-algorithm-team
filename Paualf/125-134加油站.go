在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
说明: 
如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。
示例 1:
输入: 
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。
示例 2:
输入: 
gas  = [2,3,4]
cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。

1. Clearfication
     枚举开始下标进行寻找
      然后绕圈以后，判断是否可以返回，终止逻辑为 开始索引 == 到达的索引
 如何模拟这个转圈的过程

2. Coding:
下面模拟的代码超时了
// 枚举开始下标进行寻找
//    然后绕圈以后，判断是否可以返回，终止逻辑为 开始索引 == 到达的索引
//    如何模拟这个转圈的过程
func canCompleteCircuit(gas []int, cost []int) int {
    
    for i := 0;i < len(gas);i++ {
        //fmt.Println(i)
        if isCircle(gas,cost,i) {
            return i
        }
    }
    return -1
}

func isCircle(gas []int,cost []int, index int) bool {
    n := len(gas) - 1
    gasNum := gas[index]
    flag := true
    for start := index ;flag || start != index;start++{
        flag = false
        if start > n {
            start = 0
        }
        if gasNum < cost[start] {
            return false
        }
        next := start + 1
        if next > n {
            next = 0
        }
        gasNum = gasNum - cost[start] + gas[next]
    }
    return true
}

本以为是代码逻辑不能模拟，看了超时的case，发现输入结果为
[3,1,1] 
[1,2,2]
想了下，应该是自己代码逻辑有问题，就将这个case输入进行调试发现：
处理边界条件和退出的时候有问题：退出条件为 start != index，start 上面的代码没有为0的时候，所以start++ 判断的逻辑放到了最后
还有一个点，我用了 flag 标记，让它先进循环，然后后面的逻辑都是判断 start != index 的时候再进入循环
// 枚举开始下标进行寻找
//    然后绕圈以后，判断是否可以返回，终止逻辑为 开始索引 == 到达的索引
//    如何模拟这个转圈的过程
func canCompleteCircuit(gas []int, cost []int) int {
    
    for i := 0;i < len(gas);i++ {
        //fmt.Println(i)
        if isCircle(gas,cost,i) {
            return i
        }
    }
    return -1
}
func isCircle(gas []int,cost []int, index int) bool {
    n := len(gas) - 1
    gasNum := gas[index]
    flag := true
    for start := index ;flag || start != index;{
        //fmt.Println(start)
        flag = false
        if gasNum < cost[start] {
            return false
        }
        next := start + 1
        if next > n {
            next = 0
        }
        gasNum = gasNum - cost[start] + gas[next]
        start = start + 1
        if start > n {
            start = 0 
        }
    }
    return true
}

3. 看题解：
官方题解：
模拟遍历，我看了感觉比较巧妙的在 判断 cnt == n 
还有首先判断检查第0个加油站，并试图判断能否环绕一周，如果不能，就从第一个无法到达的加油站开始继续检查，感觉从这里看着像是数学问题
func canCompleteCircuit(gas []int, cost []int) int {
    for i, n := 0, len(gas); i < n; {
        sumOfGas, sumOfCost, cnt := 0, 0, 0
        for cnt < n {
            j := (i + cnt) % n
            sumOfGas += gas[j]
            sumOfCost += cost[j]
            if sumOfCost > sumOfGas {
                break
            }
            cnt++
        }
        if cnt == n {
            return i
        } else {
            i += cnt + 1
        }
    }
    return -1
}

看了下面这个代码，感觉绝了，感觉，它的思考纬度，感觉比较高
func canCompleteCircuit(gas []int, cost []int) int {
    len := len(gas)
    spare := 0 
    minSpare := maths.MaxInt64
    minIndex := 0
    
    for i := 0;i < len;i++ {
        spare += gas[i] - cost[i]
        
        if spare < minSpare {
            minSpare = spare
            minIndex  = i
        }
    }
    
    if spare < 0 {
        return - 1
    }
    
    return  (minIndex + 1) % len
}

4. 复杂度分析：
时间复杂度：模拟时间复杂度：O(n * n）从前向后遍历时间复杂度：O(n)
空间复杂度：O(1)


5. 总结：
5.1：写代码的能力还是有待提高，模拟这个加油过程写的时候有点磕磕绊绊的
