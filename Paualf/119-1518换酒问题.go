1. Clearfication:
Clearfication:
        有一个循环的状态：
            去根据numExchange 的数量去消耗numBottles 
            ret := numBottles
           
            for numBottles > 0 {
                current := numBottles / numExchange
                left = numBottles % numExchange
                ret += current
                numBottles = current + left
            }
            return ret

2. Coding:
func numWaterBottles(numBottles int, numExchange int) int {
             ret := numBottles
            
            for numBottles > 0 && numBottles >= numExchange {
                current := numBottles / numExchange
                left := numBottles % numExchange
                ret += current
                numBottles = current + left
            }
            return ret
}

3. 看题解：
func numWaterBottles(numBottles int, numExchange int) int {
    var result,carry int
    for numBottles > 0 {
        result += numBottles
        numBottles,carry = (numBottles + carry) / numExchange,(numBottles + carry) % numExchange
    }
    
    return result
}
数学解法：

func numWaterBottles(numBottles int, numExchange int) int { 
    return numBottles + (numBottles - 1) / (numExchange - 1)
 }

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

5. 总结：
5.1: 计算机适合解决重复计算，本质上还是 if/else for loop,recuresion 这些操作
5.2: 感觉编程本质上是数学，数学转化为代码
