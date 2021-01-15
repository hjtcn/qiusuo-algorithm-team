/*
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
示例 1:
输入: 2, [[1,0]] 
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
提示：
输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
*/

1. Clearfication:
将数据结构抽象成图的话，查找图中是否存在环路，发了一会儿呆，发现自己并不知道怎么写,就去看了题解：
func canFinish(numCourses int,prerequisites [][]int) bool {
        var (
            edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        valid = true
    )
    
    for _,info := range prerequisites {
            edges[info[1]] = append(edges[info[1]],info[0])
    }
    
    for i := 0;i < numCourses && valid;i++ {
            if visited[i] == 0 {
                dfs(i, &visited,&valid)
        }
    }
    
    return valid
}
func dfs(u int,visited *[]int,valid *int) {
      visited[u] = 1
    for _,v := range edges[u] {
            if visited[v] == 0 {
                dfs(v,visited,valid)
            if !*valid{ 
                    return
            }
        }else if visited[v] == 1 {
                *valid = false
            return
        }
    }
    visited[u] = 2
    result = append(result, u)
}

上面是有报错的，visited[u] 报 invalid operation:visited[u] (type *[]int does not support indexing)
后来看到应该是这样写  (*visited)[u] = 1
上面的代码还是有问题的哈，看下面的代码

func canFinish(numCourses int,prerequisites [][]int) bool {
        var (
            edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        valid = true
    )
    
    for _,info := range prerequisites {
            edges[info[1]] = append(edges[info[1]],info[0])
    }
    
    for i := 0;i < numCourses && valid;i++ {
            if visited[i] == 0 {
                dfs(i, visited,&valid,edges,result)
        }
    }
    
    return valid
}
func dfs(u int,visited []int,valid *bool,edges [][]int,result []int) {
    visited[u] = 1
    for _,v := range edges[u] {
            if visited[v] == 0 {
                dfs(v,visited,valid,edges,result)
            if !*valid{ 
                return
            }
        }else if visited[v] == 1 {
                *valid = false
            return
        }
    }
    visited[u] = 2
    result = append(result, u)
}
不明白的一个地方是，go里面传入数组直接用形参就可以了，不用使用指针来传递数组，这是我一直比较疑惑的地方
也明白了为什么要用闭包了，不用传递参数,题解里面使用量闭包，自己把dfs函数给拆了出来
func canFinish(numCourses int, prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        valid = true
        dfs func(u int)
    )
    dfs = func(u int) {
        visited[u] = 1
        for _, v := range edges[u] {
            if visited[v] == 0 {
                dfs(v)
                if !valid {
                    return
                }
            } else if visited[v] == 1 {
                valid = false
                return
            }
        }
        visited[u] = 2
        result = append(result, u)
    }
    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
    }
    for i := 0; i < numCourses && valid; i++ {
        if visited[i] == 0 {
            dfs(i)
        }
    }
    return valid
}

广度优先遍历：
将入度为0的加入到队列中，然后从队列取出元素，放入到结果数组中，然后减少它出度元素对应入度的数量，然后将入读为0的节点放入到队列中，继续该过程，最后判断结果数组中元素是否和课程数量相等
func canFinish(numCourses int,prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        indeg = make([]int, numCourses)
        result []int
    )
    
    for _,info := range prerequisites {
        edges[info[1]] = append(edges[info[1]],info[0])
        indeg[info[0]]++
    }
    
    q := []int{}
    
    for i := 0;i < numCourses;i++ {
        if indeg[i] == 0 {
            q = append(q, i)
        }
    }
    
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        result = append(result, u)
        
        for _,v := range edges[u] {
            indeg[v]--
            if indeg[v] == 0 {
                q = append(q,v)
            }
        }
    }
    
    return len(result) == numCourses
} 

复杂度分析：
时间复杂度：O(n+m) 对图进行遍历
空间复杂度：O(n+m) 转化成邻接表，然后使用栈或者队列存储

总结：
1. 这种题不会做，但是这里面用到的数据结构，如邻接表的转换，使用队列和栈的思想，然后解决一个是否依赖成环的问题，整个过程感觉还是蛮有意思的
