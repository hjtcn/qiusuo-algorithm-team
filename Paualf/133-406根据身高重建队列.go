假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。
示例 1：
输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
解释：
编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
示例 2：
输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
 
提示：
1 <= people.length <= 2000
0 <= hi <= 106
0 <= ki < people.length
题目数据确保队列可以被重建

1. Clearfication:
 关键点在于怎么去思考和解决这种问题：从给的一个二维数组中，根据它对应的数据结构进行重建，构建出新的二维数组

2. Coding:

3. 看题解：
从低到高：第i个人的位置，就是队列中从左到右数第 ki + 1 个空位置
func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(i,j int)bool {
        a,b := people[i],people[j]
        
        return a[0] < b[0] || a[0] == b[0] &&  a[1] > b[1]
    })
    
    ret := make([][]int,len(people))
    
    for _,person := range people {
        spaces := person[1] + 1
        for i := range ret {
            if ret[i] == nil {
                spaces--
                if spaces == 0 {
                    ret[i] = person
                    break
               }
            }   
        }
    }
    
    return ret
}
从高到低进行插入
func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(i,j int)bool {
        a,b := people[i],people[j]
        
        return a[0] > b[0] || a[0] == b[0] &&  a[1] < b[1]
    })
    
    ret := make([][]int,0)
    for _,person := range people {
        idx := person[1]
        ret = append(ret[:idx], append([][]int{person}, ret[idx:]...)...)
    }
    
    return ret
}

4. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O（logn)

5. 总结：
5.1: 排序然后找到合适的位置插入
