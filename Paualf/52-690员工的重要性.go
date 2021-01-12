/*
给定一个保存员工信息的数据结构，它包含了员工唯一的id，重要度 和 直系下属的id。
比如，员工1是员工2的领导，员工2是员工3的领导。他们相应的重要度为15, 10, 5。那么员工1的数据结构是[1, 15, [2]]，员工2的数据结构是[2, 10, [3]]，员工3的数据结构是[3, 5, []]。注意虽然员工3也是员工1的一个下属，但是由于并不是直系下属，因此没有体现在员工1的数据结构中。
现在输入一个公司的所有员工信息，以及单个员工id，返回这个员工和他所有下属的重要度之和。
示例 1:
输入: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
输出: 11
解释:
员工1自身的重要度是5，他有两个直系下属2和3，而且2和3的重要度均为3。因此员工1的总重要度是 5 + 3 + 3 = 11。
注意:
一个员工最多有一个直系领导，但是可以有多个直系下属
员工数量不超过2000。
*/

1. Clearfication:
画图分析一下题目给的案例，1 -> 2 or 1 -> 3
1这个节点判断它是否有 importance，有的话则加上，然后判断它的 Subordinates
如果不为空，则继续该过程，继续找，直到 Subordinates 为空

func getImportance(employees []*Employee, id int) int {
    ret := 0
    helper(employees,id,&ret)
    return ret
}
func helper(employees []*Employee,id int,ret *int) {
    for i := 0;i < len(employees);i++ {
        if employees[i].Id == id {
            if employees[i].Importance > 0 {
                *ret = *ret + employees[i].Importance
            }
            if len(employees[i].Subordinates) > 0 {
                for j := 0;j < len(employees[i].Subordinates);j++ {
                    helper(employees,employees[i].Subordinates[j],ret)
                }
            }
        }
    }    
}

一开始自己判断了  if employees[i].Importance > 0，提交以后发现重要度还可以为负数，就改了下
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */
func getImportance(employees []*Employee, id int) int {
    ret := 0
    helper(employees,id,&ret)
    return ret
}
func helper(employees []*Employee,id int,ret *int) {
    for i := 0;i < len(employees);i++ {
        if employees[i].Id == id {
            *ret = *ret + employees[i].Importance
            if len(employees[i].Subordinates) > 0 {
                for j := 0;j < len(employees[i].Subordinates);j++ {
                    helper(employees,employees[i].Subordinates[j],ret)
                }
            }
        }
    }    
}

如果用BFS的话怎么写呢？
1. 如何定义queue
2. 如何遍历queue

伪代码：
将 subordinats 的值放入到queue里面，然后遍历queue,然后每次都遍历 len(emploees) 找到queue 里面的id，将
它的Importance 加一起
queue := []int 

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */
func getImportance(employees []*Employee, id int) int {
    ret := 0
    
    if employees == nil {
        return ret
    }
    
    queue := []int {id}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        for i := 0;i < len(employees);i++{
            if node == employees[i].Id{
                ret += employees[i].Importance
                
                if len(employees[i].Subordinates) > 0 {
                    for j := 0;j < len(employees[i].Subordinates);j++ {
                        queue = append(queue, employees[i].Subordinates[j])
                    }
                }
            }
        }
    }
    
    return ret
}
看了题解还可以用map进行加速查找，这样就不要用遍历整个数组了，时间复杂度也降下来了
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */
func getImportance(employees []*Employee, id int) int {
    ret := 0
    
    if employees == nil {
        return ret
    }
    
    empMap := make(map[int]*Employee)
    for _,v := range employees {
        empMap[v.Id] = v
    }
    
    queue := []int{id}
    
    for len(queue) > 0 {
        curEmployee := empMap[queue[0]]
        queue = queue[1:]
        ret += curEmployee.Importance
        queue = append(queue,curEmployee.Subordinates...)
    }
    return ret
}

复杂度分析：
时间复杂度：O(n * n) ：因为每次要从employees 里面找到对应的id，要找n+n-1 + ...1 次，也就是n(n+1) / 2,所以是O(n*n),使用Map 以后时间复杂度就变成了O(n)
空间复杂度：O(n) 递归使用栈空间 或者使用队列的存储的长度

总结：
1. 人最大的敌人是自己： 本觉得BFS自己写不出来，结果分析清楚以后，自己还是很快的写出来了，还是太懒了，同时还是害怕失败，害怕自己不会写，写不出来，其实没事儿的，努力分析，努力写

2. 对map 还是不敏感，使用的比较少，没有想到使用map去加速查找过程，减少时间复杂度
