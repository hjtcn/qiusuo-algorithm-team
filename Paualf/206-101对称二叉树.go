给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
 
进阶：
你可以运用递归和迭代两种方法解决这个问题吗？

1. Clarificatiion:
递归思路：使用两个指针，一个往左走，一个往右走，然后去比较，思路是有的，在分解子问题上代码没有想清楚。。。。

迭代思路：使用队列，每次向队列中push元素

说实话，上面写的思路很模糊，说明自己还是没有真正把问题理解和想清楚

2. 看题解：

题解：通过同步移动两个指针的方法来遍历这棵树，p指针和q指针一开始都指向这棵树的根，随后p右移时，q左移，p左移时，q右移。每次检查当前p和q节点的值是否相等，如果相等再判断左右子树是否对称。
func isSymmetric(root *TreeNode) bool {
    return check(root, root)
}

func check(p, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left) 
}


题解：
初始化时我们把根节点入队两次，每次提取两个结点并比较它们的值
然后将两个节点的左右子节点按相反的顺序插入队列中
当队列为空时，或者我们检测到树不对称时，该算法结束
队列中每两个连续的节点都应该是相等的，而且它们的子树互为镜像

func isSymmetric(root *TreeNode) bool {
    u, v := root, root
    q := []*TreeNode{}
    q = append(q, u)
    q = append(q, v)
    for len(q) > 0 {
        u, v = q[0], q[1]
        q = q[2:]
        if u == nil && v == nil {
            continue
        }
        if u == nil || v == nil {
            return false
        }
        if u.Val != v.Val {
            return false
        }
        q = append(q, u.Left)
        q = append(q, v.Right)

        q = append(q, u.Right)
        q = append(q, v.Left)
    }
    return true
}

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)

4.总结：
4.1: 有时候自己以为是自己讲不清楚，其实真正本质上是自己不会，所以想不清楚，最后才讲不清楚

4.2: 以前自己很讨厌和不喜欢看文字题解，现在发现是自己没有耐心，真正解决问题的方法其实是在文字中的，把思路给表达出来

4.3: 两个指针，一左一右，向里面放元素的时候对称着放，不是很熟悉


