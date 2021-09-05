序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。

编码的字符串应尽可能紧凑。

 
示例 1：

输入：root = [2,1,3]
输出：[2,1,3]
示例 2：

输入：root = []
输出：[]
 

提示：

树中节点数范围是 [0, 104]
0 <= Node.val <= 104
题目数据 保证 输入的树是一棵二叉搜索树。
 

注意：不要使用类成员/全局/静态变量来存储状态。 你的序列化和反序列化算法应该是无状态的。

1. Clarification:

将二叉搜索树序列化二叉搜索树

放到数组中，中序遍历

怎么从中序遍历的数组中反序列化出来呢？
只有一个中序遍历的能反序列化出来么？

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    Ret []int
}

func Constructor() Codec {
    return Codec{
        Ret:[]int{},
    }
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    helper(root,&this.Ret)
    str := ""

    for _,v := range this.Ret {
        str += string(v)
    }

    return str
}

func helper(root *TreeNode,ret *[]int) {
    helper(root.Left,ret)
    *ret = append(*ret, root.Val)
    helper(root.Right,ret)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    
    return
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
反序列化的我不知道怎么处理

这里又涉及到二叉树的理论知识了？
有中序遍历还原不出来一颗二叉树


2. 2. 看题解：
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		if queue[0] != nil {
			res = append(res, strconv.Itoa(queue[0].Val))
			queue = append(queue, queue[0].Left, queue[0].Right)
		} else {
			res = append(res, "#")
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    fmt.Println(data)
	if len(data) == 0 {
		return nil
	}
	var res = strings.Split(data, ",")
	var root = &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	var queue = []*TreeNode{root}
	res = res[1:]
	for 0 < len(queue) {
		if res[0] != "#" {
			l, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{Val: l}
			queue = append(queue, queue[0].Left)
		}

		if res[1] != "#" {
			r, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNode{Val: r}
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		res = res[2:]
	}
	return root
}

type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.res = strings.Split(data, ",")
	return this.dfsDeserialize()
}

func (this *Codec) dfsDeserialize() *TreeNode {
	node := this.res[0]
	this.res = this.res[1:]
	if node == "#" {
		return nil
	}

	v, _ := strconv.Atoi(node)
	root := &TreeNode{
		Val:   v,
		Left:  this.dfsDeserialize(),
		Right: this.dfsDeserialize(),
	}
	return root
}
看了题解以后发现好像自己还是真的不太理解二叉搜索树诶

先初始化根，然后再写入左，再写入右

3.复杂度分析：
时间复杂度：O(n) 遍历每个节点
空间复杂度：O(n) 递归深度

总结：
4.1: 二叉搜索树的特性：左子树 < 根节点 < 右子树
4.2: 怎么把知识消耗变成自己的知识，需要自己理解，琢磨，练习

