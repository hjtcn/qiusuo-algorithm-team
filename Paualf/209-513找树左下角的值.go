给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。
示例 1:
输入: root = [2,1,3]
输出: 1

示例 2:
输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7
 
提示:
二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1 

1. Clarification:

找树左下角的值

Q:什么是左下角
A: 就是最深的一层：左边的节点

node.Left && node.Left  nil && node.RIght == nil && level 比较深

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Ret struct {
    Level int
    Val int
}

func findBottomLeftValue(root *TreeNode) int { 
    var ret Ret
    ret.Val = root.Val
    level := 0
    helper(root,level,&ret)
    return ret.Val
}

func helper(root *TreeNode,level int,ret *Ret) {
    if root == nil {
        return
    }

    // process current logic
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        if level + 1 > ret.Level {
            ret.Val = root.Left.Val
        }
    }

    helper(root.Left,level+ 1,ret)
    helper(root.Right,level + 1,ret)
}
一开始是这样想的，觉得和昨天的题目挺像的，结果运行了一下发现不行

[0,null,-1] 这个case输出的是 -1，也就是最后一层的节点的最左边的元素

到这里算是理解了 最底层，最左边的元素是什么意思了

方案：最底层，先找出深度，然后遍历的时候将最底层的元素使用前序遍历记录，输出第一个元素

一开始我写的时候是使用了一个元素进行赋值，后来发现会被同一层后面的元素覆盖掉，后来就使用了数组去存储最后一层的值，然后返回这一层的第一个元素
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    dep := depth(root)
    level := 1
    ret := 0
    helper(root,dep,level,&ret)
    return ret
}

func helper(root *TreeNode,dep,level int,ret *int) {
    if root == nil {
        return
    }

    if level == dep {
       // fmt.Println(level,dep,root.Val)
        *ret = root.Val
        return
    }

    helper(root.Left,dep,level + 1,ret)
    helper(root.Right,dep,level + 1,ret)
}

func depth(root *TreeNode)int{
    if root == nil {
        return 0
    }
    left := depth(root.Left)
    right := depth(root.Right)

    return max(left,right) + 1
}

func max(x,y int) int {
    if x > y {
        return x
    }   
    return y
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    dep := depth(root)
    level := 1
    ret := []int{}
    helper(root,dep,level,&ret)
    return ret[0]
}

func helper(root *TreeNode,dep,level int,ret *[]int) {
    if root == nil {
        return
    }

    if level == dep {
       // fmt.Println(level,dep,root.Val)
        *ret = append(*ret,root.Val)
        return
    }

    helper(root.Left,dep,level + 1,ret)
    helper(root.Right,dep,level + 1,ret)
}

func depth(root *TreeNode)int{
    if root == nil {
        return 0
    }
    left := depth(root.Left)
    right := depth(root.Right)

    return max(left,right) + 1
}

func max(x,y int) int {
    if x > y {
        return x
    }   
    return y
}
2. 看题解：
这个为什么不会覆盖值呢，因为它更新值的时候同时更新了最大的深度，同一层的最大深度不会大于它，所以不会被覆盖的哈
var maxDeep int // 全局变量 深度
 var  value  int //全局变量 最终值
func findBottomLeftValue(root *TreeNode) int {
     if root.Left==nil&&root.Right==nil{//需要提前判断一下（不要这个if的话提交结果会出错，但执行代码不会。防止这种情况出现，故先判断是否只有一个节点）
         return root.Val
     }
    findLeftValue (root,maxDeep)
    return value
}
func findLeftValue (root *TreeNode,deep int){
     //最左边的值在左边
     if root.Left==nil&&root.Right==nil{
         if deep>maxDeep{
             value=root.Val
             maxDeep=deep
         }
     }
    //递归
    if root.Left!=nil{
        deep++
        findLeftValue(root.Left,deep)
        deep--//回溯
    }
    if root.Right!=nil{
        deep++
        findLeftValue(root.Right,deep)
        deep--//回溯
    }
}
迭代：
层序遍历更加的简单，每一层的第一个节点覆盖就行，返回最后一个值就可以了
func findBottomLeftValue(root *TreeNode) int {
    queue:=list.New()
    var gradation int
    queue.PushBack(root)
    for queue.Len()>0{
        length:=queue.Len()
        for i:=0;i<length;i++{
            node:=queue.Remove(queue.Front()).(*TreeNode)
            if i==0{gradation=node.Val}
            if node.Left!=nil{
                queue.PushBack(node.Left)
            }
            if node.Right!=nil{
                queue.PushBack(node.Right)
            }
        }
    }
    return gradation
}


//bfs
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[0]
		stack = stack[1:]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return root.Val
}


3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

总结：
4.1: 一开始没有理解清楚题意还是走了一些弯路，所以理解清楚和需求是很重要的哈
