/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：
[
  [3],
  [20,9],
  [15,7]
]
*/

1. Clearfication:
锯齿形层序遍历：
感觉像记录奇偶，然后根据奇偶一个从头开始，一个从尾部开始
正常的层序遍历是一层层的遍历节点，和正常的遍历有啥区别呢，我们可不可以在正常遍历的基础上改一下呢，先按照正常的遍历输出结果以后，然后遍历输出结果，按照奇偶顺序将每一层的节点给反转
正常的层序遍历是使用队列来进行层序遍历，将root节点放入到队列中，然后计算当前层次的节点个数，然后使用for 循环将当前层次节点的个数取出来
size := len(queue)
for i := 0;i < size;i++ {
    
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    ret := [][]int{}
    
    queue := []*TreeNode{root}
    size := len(queue)
    
    for len(queue) > 0 {
      
         tmp := []int{}
         for i := 0;i < size;i++ {
             node := queue[0]
             queue = queue[1:]
             
             tmp = append(tmp, node.Val)
             
             if node.Left != nil {
                 queue = append(queue, node.Left)
             }
             
             if node.Right != nil {
                 queue = append(queue, node.Right)
             }
         }
        fmt.Println(tmp)
        ret = append(ret, tmp)
    }
  
    
    return ret
}

上面的代码输出结果是
[[3],[9],[20],[15],[7]]
和自己想要的输出结果不一样，看到了计算size放的地方没有使用正确

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    ret := [][]int{}
    
    queue := []*TreeNode{root}
   
    for len(queue) > 0 {
         size := len(queue)
         tmp := []int{}
         for i := 0;i < size;i++ {
             node := queue[0]
             queue = queue[1:]
             
             tmp = append(tmp, node.Val)
             
             if node.Left != nil {
                 queue = append(queue, node.Left)
             }
             
             if node.Right != nil {
                 queue = append(queue, node.Right)
             }
         }
        //fmt.Println(tmp)
        ret = append(ret, tmp)
    }
  
    
    return ret
}

修改过size放的位置以后，输出结果就是正常的层序遍历了
[[3],[9,20],[15,7]]
我们接下来就是对层序遍历进行奇偶反转操作
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    ret := [][]int{}
    
    queue := []*TreeNode{root}
   
    for len(queue) > 0 {
         size := len(queue)
         tmp := []int{}
         for i := 0;i < size;i++ {
             node := queue[0]
             queue = queue[1:]
             
             tmp = append(tmp, node.Val)
             
             if node.Left != nil {
                 queue = append(queue, node.Left)
             }
             
             if node.Right != nil {
                 queue = append(queue, node.Right)
             }
         }
        ret = append(ret, tmp)
    }
  
    for i := 0;i < len(ret);i++ {
        if i % 2 == 1 {
            for left,right := 0,len(ret[i]) - 1;left < right;left++ {
                ret[i][left],ret[i][right] = ret[i][right],ret[i][left]
                right--
            }
        }
    }
    
    return ret
}

再加一下边界条件的判断,如果不加判断的话，使用root 变量的时候就会panic掉
所以我们使用每一个变量之前都要判断这个变量是否为空，对它做处理

if root == nil {
    return ret
}
func zigzagLevelOrder(root *TreeNode) [][]int {
    ret := [][]int{}
    
    if root == nil {
        return ret
    }
    queue := []*TreeNode{root}
   
    for len(queue) > 0 {
         size := len(queue)
         tmp := []int{}
         for i := 0;i < size;i++ {
             node := queue[0]
             queue = queue[1:]
             
             tmp = append(tmp, node.Val)
             
             if node.Left != nil {
                 queue = append(queue, node.Left)
             }
             
             if node.Right != nil {
                 queue = append(queue, node.Right)
             }
         }
        ret = append(ret, tmp)
    }
  
    for i := 0;i < len(ret);i++ {
        if i % 2 == 1 {
            for left,right := 0,len(ret[i]) - 1;left < right;left++ {
                ret[i][left],ret[i][right] = ret[i][right],ret[i][left]
                right--
            }
        }
    }
    
    return ret
}

看了其他同学写的题解：
我们最后进行的数组元素的反转，我们可以放入tmp数组的时候就可以对它进行反转，使用level 记录当前是那一层，判断level 的奇偶性，然后判断是否需要反转数组，
反转的时候还看到题解里面这样反转的
for i,n := 0,len(vals);i < n / 2;i++ {
    vals[i],vals[n-1-i] = vals[n-1-i],vals[i]
}
只需要维护了一个变量i,然后尾部元素使用 n-1-i 来进行替换，感觉循环里面还是蕴含着数学的思想在里面的

复杂度分析：
时间复杂度：O(n): 遍历每个节点
空间复杂度：O(n): 使用队列存储每个节点

总结：
1. 使用每个变量的时候我们都要去想它是否为空，如果为空，我们需要对它进行处理么，怎么判断

2. 数学思维，去慢慢接触数学，了解它在编程里面的魅力
