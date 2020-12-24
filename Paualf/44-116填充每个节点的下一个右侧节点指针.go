/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
进阶：
你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 
示例：
输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
 
提示：
树中节点的数量少于 4096
-1000 <= node.val <= 1000
*/

1. Clearfication:
在纸上画了一下，第一感觉是按照层序遍历，使用level 辅助遍历数组
伪代码：
level = 1
ret :=[]int{}
helper(root *TreeNode,level int,*ret) {
    // terminator
    
    // process current logic
    
    // dirll down
    
    // restore current status
}
自己没想出来，怎么把一层的节点连起来使用
看了题解：
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    
    // 初始化队列同时将第一层节点加入队列中，即根节点
    queue := []*Node{root}
    
    // 循环迭代的是层数
    for len(queue) > 0 {
        tmp := queue
        queue = nil
        
        // 遍历这一层的所有节点
        for i,node := range tmp {
            // 连接
            if i + 1 < len(tmp) {
                node.Next = tmp[i+1]
            }
            
            // 扩展下一层节点
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    // 返回根节点
    return root
}

没有想出来的地方在于 自己把广度优先遍历和层次遍历想的是一样了，其实他们并不一样
层次遍历基于广度优先搜索，它与广度优先搜索的不同之处在于，广度优先搜索只会取出一个节点来扩展，而层次遍历会每次将队列中的所有元素都拿出来拓展，这样能保证每次从队列中拿出来遍历的元素都是属于同一层的。
看到题解中方法二：使用已经建立的next 指针

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    // 每次循环从该层的最左侧节点开始
    for leftmost := root;leftmost.Left != nil;leftmost = leftmost.Left {
        for node := leftmost;node != nil;node = node.Next {
            node.Left.Next = node.Right
            
            if node.Next != nil {
                node.Right.Next = node.Next.Left
            }
        }
    }
    return root
}

看到有使用Map 将每一层的元素连接在一起，还是蛮巧妙的

func connect(root *Node)*Node {
    if root == nil {
        return nil
    }
    m := map[int]*Node{}
    traverse(root,0,m)
    return root
}
func traverse(root *Node,height int,m map[int]*Node) {
    if root == nil {
        return
    }
    traverse(root.Left, height + 1,m)
    traverse(root.Right, height + 1,m)
    if _,ok := m[height];ok {
        m[height].Next = root
        m[height] = m[height].Next
    }else {
        m[height] = root
    }
}

复杂度分析：
时间复杂度：O(n) 遍历每一个节点
空间复杂度：O(n) 开辟队列或者map存储对应关系

总结：
1. 层序遍历和广度优先搜索还是不一样的

2. 要多看题解
