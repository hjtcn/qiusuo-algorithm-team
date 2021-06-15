累加数是一个字符串，组成它的数字可以形成累加序列。
一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。
说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
示例 å1:
输入: "112358"
输出: true 
解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2:
输入: "199100199"
输出: true 
解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199

1. Clearfication:
   Clearfication:
        怎么去思考这个问题呢？
            怎么去分隔数
            1 99 100 199 怎么从字符串中将这些数据分出来呢？
看了题解以后，就是按照下标分隔，有的时候问题的复杂度本身是存在的，就是按照它本身的难度去做就好了

2. Coding:

3. 看题解：
func isAdditiveNumber(num string)bool {
    n := len(num)
    if n < 3 {
        return false
    }
    for i := 1;i < n;i++ {
        if i > 1 && num[0] == '0' {
            break
        }
        if lst, _ := strconv.Atoi(num[:i]);dfs(num,lst,i) {
            return true
        }
    }
    return false
}
func dfs(num string,lst int,idx int) bool {
    n := len(num)
    isAdditive := false
    
    for nxtStart := idx + 1;nxtSstart < n;nxtStart++ {
        // 排除前导0
        if nxtStart > idx + 1 && num[idx] == '0' {
            break
        }
        
        // 第二个数
        cur,_ := strconv.Atoi(num[idx:nxtStart])
        // 第一个数和第二个数的和
        sum := strconv.Itoa(cur + lst)
        // 第三个数的长度
        dur := len(sum)
        // 第三个数的右区间
        nxtEnd := nxtStart + dur
        // 如果右区间越界，结束循环
        if nxtEnd > n {
            break   
        }
        // 如果和相等
        if sum == num[nxtStart:nxtEnd] {
            // 并且刚好到数组的最后一位，则为累加数
            if nxtEnd == n {
                return true
            }
            // 没到结尾，继续递归
            isAdditive = dfs(num,cur,nxtStart)
        }                      
    }
    return isAdditive
}

4. 复杂度分析：
时间复杂度：O(n^n)
空间复杂度：O(n^n)

5. 总结：
5.1: 先把框架整理出来，然后再去填充细节
5.2: 如果问题本身的复杂度存在，就动手
