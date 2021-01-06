//给定一个保存员工信息的数据结构，它包含了员工唯一的id，重要度 和 直系下属的id。
//
// 比如，员工1是员工2的领导，员工2是员工3的领导。他们相应的重要度为15, 10, 5。那么员工1的数据结构是[1, 15, [2]]，员工2的数据结构是
//[2, 10, [3]]，员工3的数据结构是[3, 5, []]。注意虽然员工3也是员工1的一个下属，但是由于并不是直系下属，因此没有体现在员工1的数据结构中。
//
//
// 现在输入一个公司的所有员工信息，以及单个员工id，返回这个员工和他所有下属的重要度之和。
//
// 示例 1:
//
//
//输入: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
//输出: 11
//解释:
//员工1自身的重要度是5，他有两个直系下属2和3，而且2和3的重要度均为3。因此员工1的总重要度是 5 + 3 + 3 = 11。
//
//
// 注意:
//
//
// 一个员工最多有一个直系领导，但是可以有多个直系下属
// 员工数量不超过2000。
//
// Related Topics 深度优先搜索 广度优先搜索 哈希表
// 👍 123 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */
type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

/**DFS**/

var sum int

func getImportanceOld(employees []*Employee, id int) int {
	sum = 0
	empMap := make(map[int]*Employee)

	for _, v := range employees {
		empMap[v.Id] = v
	}

	nDfs(empMap, id)

	return sum
}

func nDfs(employee map[int]*Employee, search_number int) {
	currEmployee := employee[search_number]
	sum += currEmployee.Importance

	for _, v := range currEmployee.Subordinates {
		nDfs(employee, v)
	}
}

/**
深度优先搜素：
先把给定的切片，构造一个以切片中单个元素中的id为key, 单个元素为value的map, 方便查询
然后根据给定的id,去map中查找,然后将其直系下属的ID依次更新为要查询的id,做递归查找,递归查询时不断累加员工的重要度

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
*/

/**BFS**/
func getImportance(employees []*Employee, id int) int {
	sum := 0
	empMap := make(map[int]*Employee)

	for _, v := range employees {
		empMap[v.Id] = v
	}

	var queue []int
	queue = append(queue, id)
	for len(queue) > 0 {
		currEmployee := empMap[queue[0]]
		sum += currEmployee.Importance

		queue = queue[1:]
		queue = append(queue, currEmployee.Subordinates...)	//追加一个数组，需要加入 "..."
	}

	return sum
}

/**
广度优先搜索：
也是为方便查询，将员工关系的数据放入map中, 将员工ID，从给定的员工ID查询到的其所有下属先后存入队列中
同时不断将员工ID出队列，对重要度进行累加

复杂度分析：
	时间复杂度：O(n)
	空间复杂度：O(n)
*/
//leetcode submit region end(Prohibit modification and deletion)

/**
总的来看，这道题可看作是N叉树先序遍历的变形
*/
