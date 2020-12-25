/*
在一个小镇里，按从 1 到 N 标记了 N 个人。传言称，这些人中有一个是小镇上的秘密法官。
如果小镇的法官真的存在，那么：
小镇的法官不相信任何人。
每个人（除了小镇法官外）都信任小镇的法官。
只有一个人同时满足属性 1 和属性 2 。
给定数组 trust，该数组由信任对 trust[i] = [a, b] 组成，表示标记为 a 的人信任标记为 b 的人。
如果小镇存在秘密法官并且可以确定他的身份，请返回该法官的标记。否则，返回 -1。
示例 1：
输入：N = 2, trust = [[1,2]]
输出：2
示例 2：
输入：N = 3, trust = [[1,3],[2,3]]
输出：3
示例 3：
输入：N = 3, trust = [[1,3],[2,3],[3,1]]
输出：-1
示例 4：
输入：N = 3, trust = [[1,2],[2,3]]
输出：-1
示例 5：
输入：N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
输出：3
 
提示：
1 <= N <= 1000
trust.length <= 10000
trust[i] 是完全不同的
trust[i][0] != trust[i][1]
1 <= trust[i][0], trust[i][1] <= N
*/

1. Clearfication:    
小镇的法官不相信任何人
每个人都信任小镇的法官
只有1个人同时满足属性1和属性2
思路：

我们可以使用反证法，假设i为法官
1. 小镇的法官不相信任何人 
如果 a == i 则这个人不是法官 continue
2. 如何存储和标记所有人都相信他这个条件

使用map存储，找到信任它的标记为true,最后遍历结果为处理自己不为true,其他都信任它，则返回true
伪代码：
for i := 1;i <= N;i++ {
    map := make([int]bool, 0)
    for j := 0;i < len(trust);j++ {
        if trust[j][0] = i {
            break
        }
        
        a = trust[j][0]
        b = trust[j][1]
        
        if b == i {
            map[a] = true
        }
    }
    
    for z := 1;z <= N;z++ {
        if z == i {
            continue
        }
        
        if !map[z] {
            return -1
        } 
    }
   
    return i
}

代码改到这里感觉怪怪的，就去看题解了
func findJudge(N int, trust [][]int) int {
    for i := 1;i <= N;i++ {
        hash := make(map[int]bool, 0)
        for j := 0;i < len(trust);j++ {
            a := trust[j][0]
            b := trust[j][1]
            if a == i {
                break
            }
            
            if b == i {
                hash[a] = true
            }
        }
        
         for z := 1;z <= N;z++ {
                if z == i {
                    continue
                }
                
                fmt.Println(hash[z])
                if !hash[z] {
                    return -1
                } 
            }
            return i
    }
    return -1
}

看了题解以后，发现自己的思路还是太简单，没有抽象出来，也可能是图做的题少，抽象出来就是入度和出度的问题，找到入读为N-1，出度为0的节点元素

func findJudge(N int, trust [][]int) int {
    if len(trust) == 0 && N == 1 {
        return 1
    }
    // out, in
    edges := make([][2]int, N+1)
    candidates := []int{}
    for _, t := range trust {
        edges[t[0]][0]++
        edges[t[1]][1]++
        if edges[t[1]][1] == N-1 {
            candidates = append(candidates, t[1])
        }
    }
    for _, c := range candidates {
        if edges[c][0] == 0 {
            return c
        }
    }
    return -1
}

看到这些代码还是蛮优秀的
func findJudge(N int,trust [][]int) int {
    if len(trust) == 0 {
        return 1
    }
    people := make(map[int]int)
    judge := make(map[int]int)
    
    for _,v := range trust {
        people[v[0]]++
        judge[v[1]]++
    }
    
    for i,v := range judge {
        if _,ok := people[i];!ok && v == N - 1 {
            return i
        }
    }
    
    return -1
}

复杂度分析：
时间复杂度：O(n) 遍历二维数组
空间复杂度：O(n) 开辟map进行存储

总结：
1. 提升自己对问题的抽象能力

2. 对数据结构如map 用的还是不是很熟悉，要多多锻炼
