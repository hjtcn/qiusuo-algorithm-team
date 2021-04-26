Hercy 想要为购买第一辆车存钱。他 每天 都往力扣银行里存钱。
最开始，他在周一的时候存入 1 块钱。从周二到周日，他每天都比前一天多存入 1 块钱。在接下来每一个周一，他都会比 前一个周一 多存入 1 块钱。
给你 n ，请你返回在第 n 天结束的时候他在力扣银行总共存了多少块钱。
示例 1：
输入：n = 4
输出：10
解释：第 4 天后，总额为 1 + 2 + 3 + 4 = 10 。
示例 2：
输入：n = 10
输出：37
解释：第 10 天后，总额为 (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4) = 37 。注意到第二个星期一，Hercy 存入 2 块钱。
示例 3：
输入：n = 20
输出：96
解释：第 20 天后，总额为 (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4 + 5 + 6 + 7 + 8) + (3 + 4 + 5 + 6 + 7 + 8) = 96 。
 
提示：
1 <= n <= 1000
1. Clearfication:
 怎么算这个问题呢？
        分几个7
       num :=  n / 7  num 标记几个整的7
       part := n % 7 
        
然后根据标记整7的个数去改变起始值的变化

2. Coding:
/*
    1. 怎么算这个问题呢？
        分几个7
 
       num :=  n / 7  num 标记几个整的7
       part := n % 7 
       sum := 0
       i := 1
       for ;i <= num;i++ {
           // i 就是它的 start
           tmp := (i + i + 7) * 7 / 2
           sum += tmp
       }
       // 整份的7处理完了，处理剩余的
       for j := 0;j < part;j++ {
           sum += i
           i++
       }
       return sum
*/
/*
    1. 怎么算这个问题呢？
        分几个7
 
       num :=  n / 7  num 标记几个整的7
       part := n % 7 
       sum := 0
       i := 1
       for ;i <= num;i++ {
           // i 就是它的 start
           tmp := (i + i + 7) * 7 / 2
           sum += tmp
       }
       // 整份的7处理完了，处理剩余的
       for j := 0;j < part;j++ {
           sum += i
           i++
       }
       return sum
*/

func totalMoney(n int) int {
    sum := 0
    start := 1
    circule := n / 7
    part := n % 7
    for ;start <= circule;start++ {
        tmp := (start + start + 7) * 7 / 2
        sum += tmp
    }
    for j := 0;j < part;j++ {
        sum += start
        start++
    }
    return sum
}
start 不是应该加7 要加6

/*
    1. 怎么算这个问题呢？
        分几个7
 
       num :=  n / 7  num 标记几个整的7
       part := n % 7 
       sum := 0
       i := 1
       for ;i <= num;i++ {
           // i 就是它的 start
           tmp := (i + i + 7) * 7 / 2
           sum += tmp
       }
       // 整份的7处理完了，处理剩余的
       for j := 0;j < part;j++ {
           sum += i
           i++
       }
       return sum
*/
func totalMoney(n int) int {
    sum := 0
    start := 1
    circule := n / 7
    part := n % 7
    for ;start <= circule;start++ {
        tmp := (start + start + 6) * 7 / 2
        sum += tmp
    }
    // fmt.Println(sum)
    // fmt.Println(start)
    for j := 0;j < part;j++ {
        sum += start
        start++
    }
    return sum
}

3. 看题解
等差数列
func totalMoney(n int)int {
    week,day := n / 7,n % 7
    firstWeek := 7 * 8 / 2
    
    // n * a1 + d * n * (n - 1) / 2
    return (week * firstWeek + 7 * week * (week - 1) / 2) + week * day + day * (day + 1) / 2
}

4. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)

5. 总结：
5.1:分析，然后动手去把思路去实现出来，不要想特别多，也不是什么都不想
