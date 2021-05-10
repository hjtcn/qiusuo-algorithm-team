1. Clearfication:
    列式遍历，如果不是升序则给 sum + 1
    sum := 0
    column := len(strs)
    row := len(strs[0])
    for i := 0;i < row;i++ {
        for j := 0;j < column - 1;j++ {
            if (strs[j][i] > strs[j][i + 1]) {
                sum++
                break
            } 
        }
    }
    return sum

2. Coding:
/*
    列式遍历，如果不是升序则给 sum + 1
    sum := 0
    column := len(strs)
    row := len(strs[0])
    for i := 0;i < row;i++ {
        for j := 0;j < column - 1;j++ {
            if (strs[j][i] > strs[j][i + 1]) {
                sum++
                break
            } 
        }
    }
    return sum
*/
func minDeletionSize(strs []string) int {
    sum := 0
    column := len(strs)
    row := len(strs[0])
    for i := 0;i < row;i++ {
        for j := 0;j < column - 1;j++ {
            if (strs[j][i] > strs[j + 1][i]) {
                sum++
                break
            } 
        }
    }
    return sum
}

3. 看题解：
用了range
func minDeletionSize(strs []string) int {
    if len(strs) <= 1 {
        return 0
    }
    
    var count int
    for i := range strs[0] {
        for j := 0;j < len(strs) - 1;j++ {
            if strs[j][i] > strs[j + 1][i] {
                count++
                break
            }
        }
    }
    
    return count
}

4. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(1)

5. 总结：
5.1: 列式遍历不是很熟悉，一开始想了一会
5.2: Go里面建议是使用range 遍历，具体原因目前还不是特别了解
