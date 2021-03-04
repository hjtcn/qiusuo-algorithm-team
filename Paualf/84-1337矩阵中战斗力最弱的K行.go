/*
给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。
如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
示例 1：
输入：mat = 
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]], 
k = 3
输出：[2,0,3]
解释：
每行中的军人数目：
行 0 -> 2 
行 1 -> 4 
行 2 -> 1 
行 3 -> 2 
行 4 -> 5 
从最弱到最强对这些行排序后得到 [2,0,3,1,4]
示例 2：
输入：mat = 
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]], 
k = 2
输出：[0,2]
解释： 
每行中的军人数目：
行 0 -> 1 
行 1 -> 4 
行 2 -> 1 
行 3 -> 1 
从最弱到最强对这些行排序后得到 [0,2,3,1]
 
提示：
m == mat.length
n == mat[i].length
2 <= n, m <= 100
1 <= k <= m
matrix[i][j] 不是 0 就是 1
*/

1. Clearfication:
m*n的矩阵 mat,军人:1, 平民:0
返回矩阵中最弱的k行的索引，按从最弱到最强排序。
如果第i行的军人数量少于第j行，或者两行军人数量相同但i小于j，那么我们认为第i行的战斗力比第j行弱。
军人总是排在 第一行中的靠前位置，也就是说1总是出现在0之前。
找到每行中的军人树木存储到一个map中，map的key是行数，value是军人的个数，也就是1的个数
然后对这个map进行按照 value的排序，然后返回排序后值对应的key的数组，然后根据k的大小返回数组中前k个元素。
伪代码：
row = len(mat)
column = len(mat[0])
m := make(map[int]int, num)
for i : [0,  row)
num := 0
for j : [0, column)
// fail fast
if mat[i][j] == 0 {
break
}
num++
m[i] = num
然后map就构建出来了
怎么根据构建的map对应的值进行排序，然后返回根据排序对应的map的key值，这个是我没有想清楚的

2. Coding
func kWeakestRows(mat [][]int, k int) []int {
    row := len(mat)
    column := len(mat[0])
    map := make(map[int]int, row)
    ret := []int{}
    
    for i := 0;i < row;i++ {
        num := 0
        for j := 0;j < column;j++ {
            if mat[i][j] != 1 {
                break
            }
            num++
        }
        map[i] = num
    }
    
    fmt.Println(map)
    return ret
}
如果用空间换时间呢？加一个数组存储每行1的个数，然后map中存储的是 每行1的个数 对应 key
func kWeakestRows(mat [][]int, k int) []int {
    row := len(mat)
    column := len(mat[0])
    m := make(map[int]int, row)
    ret := []int{}
    values := []int{}
    
    for i := 0;i < row;i++ {
        num := 0
        for j := 0;j < column;j++ {
            if mat[i][j] != 1 {
                break
            }
            num++
        }
        m[num] = i
        values = append(values, num)
    }
    
    sort.Ints(values)
    
    fmt.Println(values)
    for i := 0;i < len(values);i++ {
        mapValue := m[values[i]]
        ret = append(ret, mapValue)
    }
   
    return ret[0:k]
}
哈哈哈，map里面存储的元素可能是相同的，索引可能会被覆盖掉

3. 看题解
定义一个结构体，然后使用 slice sort
type entry struct {
    index int
    soldiers int
}

func kWeakestRows(mat [][]int,k int) []int {
    count := make([]entry, len(mat))
    
    for r,row := range mat {
        soldiers := 0
        
        for _,val := range row {
            soldiers += val
        }
        
        count[r] = entry{r, soldiers}
    }
    
    sort.Slice(count, func(i, j int)bool{
        si,sj := count[i].soldiers,count[j].soldiers
        
        return (si == sj && count[i].index < count[j].index) || si < sj
    })
    
    
    res := []int{}
    
    for i := 0;i < k;i++ {
        res = append(res, count[i].index)
    }
    
    return res
}

func kWeakestRows(mat [][]int,k int) []int {
    m,n := len(mat),len(mat[0])
    tmp := make(map[int][]int)
    temp := []int{}
    res := []int{}
    
    for i := 0;i < m;i++ {
        soldiers := 0
        for j := 0;j < n;j++ {
            if mat[i][j] == 0 {
                break
            }
            soldiers++
        }
        tmp[soldiers] = append(tmp[soldiers], i)
        temp = append(temp, soldiers)
    }
    
    sort.Ints(temp)
    
    for _,v := range temp {
        if len(tmp[v]) == 1 {
            res = append(res, tmp[v][0])
        }else {
            res = append(res, tmp[v]...)
        }
        
        delete(tmp, v)
    }
    
    return res[:k]
}

自己写的时候想到如果是定义map，如果有重复元素会覆盖掉元素，上面的代码定义了map => 切片的结构体，还是挺巧妙的
 tmp := make(map[int][]int)
用到二分的地方：
func kWeakestRows(mat [][]int,k int) []int {
    rawPowers := make([][2]int, len(mat))
    
    for idx,raw := range mat {
        rawPowers[idx] = [2]int{idx, countPower(raw)}
    }
    
    res := make([]int, k)
    
    for i := 0;i < k;i++ {
        for j := len(rawPowers) - 1;j > i;j-- {
            if rawPowers[j][1] < rawPowers[j - 1][1] {
                tmp := rawPowers[j]
                rawPowers[j] = rawPowers[j - 1]
                rawPowers[j-1] = tmp
            }
        }
        res[i] = rawPowers[i][0]
    }
    
    return res
}

func countPower(raw []int) int {
    low,high := 0,len(raw) - 1
    
    for low <= high {
        mid := low + (high - low) / 2
        if raw[mid] == 1 {
            low = mid + 1
        }else {
            high = mid - 1
        }
    }
    
    return low
}

4. 复杂度分析
时间复杂度：正常扫描二维数组 O(m*n), 使用二分查找扫描二维数组 O(m * logn) + 排序冒泡 O(n*n) or 快速排序 O(nlogn)
空间复杂度：开辟map的大小O(n)

5. 总结
5.1: 暴露问题：排序知识欠缺，需要补充对排序知识的了解
5.2：二分查找没有完全掌握
