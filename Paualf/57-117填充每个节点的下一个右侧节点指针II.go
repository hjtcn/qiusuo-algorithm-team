/*
给定一个二叉树
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
输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
提示：
树中的节点数小于 6000
-100 <= node.val <= 100
*/

1. Clearfication:
将每一层节点放入到一层里面然后连接起来
想到的是层序遍历
层序遍历使用的是 queue :=  []*Node{}
for i := 0;i < len(queue) - 1;i++ {
    queue[i].Next = queue[i+1]
}
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
    
        queue := []*Node{root}
    
    for len(queue) > 0 {
           q := queue
         queue = nil
      
         for i := 0;i < len(q) - 1;i++ {
                 q[i].Next = q[i+1]
             
             if q[i].Left != nil {
                    queue = append(queue,q[i].Left)
             }
             
             if q[i].Right != nil {
                   queue = append(queue, q[i].Right)
             }
             }
    }
    
    return root
}
上面的逻辑是有问题的，因为如果是第一个节点，循环逻辑是进不来的，所以 1的左右节点都加入不进来，如果for 循环的判断条件改为 i < len(q) 的话，访问 q[i+1] 的时候数组会越界，所以我们可以在循环内部对越界再做处理
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
    
        queue := []*Node{root}
    
    for len(queue) > 0 {
          q := queue
              
         queue = nil
         for i := 0;i < len(q) ;i++ {
             if q[i].Left != nil {
                queue = append(queue,q[i].Left)
             }
             
             if q[i].Right != nil {
                queue = append(queue, q[i].Right)
             }
            if i < len(q) - 1 {
                q[i].Next = q[i+1]
                //fmt.Println(q[i].Val,q[i+1].Val)    
            }
             
         }
    }
    
    return root
}

看了题解，感觉变量命名可以改进一下，queue 和 q 还是容易混淆的
遍历的时候也可以使用 range 进行遍历
func connect(root *Node)*Node {
    if root == nil {
        return nil
    }
    q := []*Node{root}
    
    for len(q) > 0 {
        tmp := q
        q = nil
        
        for i,node := range tmp {
            if i + 1 < len(tmp) {
                node.Next = tmp[i+1]
            }
            
            if node.Left != nil {
                q = append(q, node.Left)
            }
            
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
    }
    
    return root
}

方法2:先将当前层节点next链起来，这个时候当前层就会变成了一个链表，然后再去处理它的下一层，将它们的next连接起来
自己没有写出来，看了题解写的
func connect(root *Node) *Node {
    start := root
    
    for start != nil {
        var nextStart,last *Node
        handle := func(cur *Node) {
            if cur == nil {
                return
            }
            
            if nextStart == nil {
                nextStart = cur
            }
            
            if last != nil {
                last.Next = cur
            }
            
            last = cur
        }
        for p := start;p != nil;p = p.Next {
            handle(p.Left)
            handle(p.Right)
        }
        start = nextStart
    }
    
    return root
}

题解写的看的不是特别明白，感觉还是写不出来

看到题解里面还有这样写的，这个看着感觉比较舒服
func connect(root *Node) *Node {
    if root==nil{return nil}
    link(root)
    curr:=root
    for curr!=nil{
        link(curr)
        curr=curr.Next
    }
    connect(root.Left)
    connect(root.Right)
    return root
}
func link(node *Node){
    if node==nil{
        return 
    }
    
    if node.Left!=nil{
        if node.Right!=nil{
            node.Left.Next=node.Right
        }else{
            node.Left.Next=getNext(node)
        }
    }
    if node.Right!=nil{
        node.Right.Next=getNext(node)
    }
}
func getNext(node *Node) *Node{
    next:=node.Next
    for next!=nil{
        if next.Left!=nil{
            return next.Left
        } 
        if next.Right!=nil{
            return next.Right
        }
        next=next.Next
    }
    return nil
}

复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 使用队列存储节点或者递归栈深度使用空间

总结：
1. 层序遍历

2. 将当前层节点连接起来，然后使用当前处理过的节点处理它的写一层节点，和116题个人感觉不同的地方在于，116题的树是完全二叉树，完全二叉树的话，如果有右节点的话，肯定有左节点，但是二叉树就不一样了，可以没有左节点，只有右节点这种情况
