给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]

示例 2：
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
 

提示：

树的高度不会超过 1000
树的节点总数在 [0, 10^4] 之间

1. Clarification:
层序遍历：一层一层进行遍历，记录每一层元素的个数，然后去遍历

使用queueNode 进行存储
退出条件为 len(queueNode) > 0
输出节点的值到数组中

使用队列记录
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    ret := [][]int{}
   
    if root == nil {
        return ret
    }

    queueNode := []*Node{root}

    for len(queueNode) > 0 {
        size := len(queueNode)
        tmp := []int{}

        for size > 0 {
            size--

            node := queueNode[0]
            queueNode = queueNode[1:]

            tmp = append(tmp, node.Val)

            for i := 0;i < len(node.Children);i++ {
                queueNode = append(queueNode, node.Children[i])
            }
        }

        ret = append(ret, tmp)
    }

    return ret
}
对比下：有的啥时候会用到 size 的时候是层序遍历
BFS的时候，这个地方有时候自己会有点疑惑，目前感觉哈：应该是需要把一层的节点放入到一个数组中的时候，所以需要用到了 size

dfs 带层数的进行初始化进行遍历

核心代码：每次遍历加上当前层数
放入数组的时候判断当前层数的数组是否存在，如果不存在，则初始化它

初始化的代码没有写出来。。。，少了判断逻辑
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    ret := [][]int{}
    dfs(root,&ret,0)
    return ret
}

func dfs(root *Node,ret *[][]int,level int) {
    if root == nil {
        return
    }

    ret[level] = append(ret[level], root.Val)

    for i := 0;i < len(root.Children);i++ {
        dfs(root,ret,level + 1)
    }
}

2.看题解：

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    ret := [][]int{}
    dfs(root,&ret,0)
    return ret
}

func dfs(root *Node,ret *[][]int,level int) {
    if root == nil {
        return
    }

    if len(*ret) == level {
        *ret = append(*ret,[]int{})
    }

    // *ret 需要加括号？那为什么需要加括号呢，目前还没有想清楚哈
    (*ret)[level] = append((*ret)[level], root.Val)

    for i := 0;i < len(root.Children);i++ {
        dfs(root.Children[i],ret,level + 1)
    }
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4.总结：
4.1: 每个人遇到的问题都不太一样，去思考和动手去解决自己遇到的问题，可能是最快的方式，逃避往往没有用的
