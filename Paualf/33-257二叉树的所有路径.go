/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。
示例:
输入:
   1
 /   \
2     3
 \
  5
输出: ["1->2->5", "1->3"]
解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/

Clearfication:
这道题目，思路还是可以通过BFS和DFS进行解决，遍历到根节点，将元素放到数组里面
细节：存储的字符串数组，而且数组里面是通过 '->' 连接起来的   可以通过构建二维数组，如 [[1,2,5],[1,3]] 然后遍历二维数组，将里面的一维数组通过符号 '->' 连接起来

预估复杂度分析：
时间复杂度:O(n)
空间复杂度:O(n) 遍历每个节点树和递归调用栈空间
方法：

DFS
BFS
这是我尝试的代码，函数传参还是会有问题

func binaryTreePaths(root *TreeNode) []string {
    result := [][]string{}
    path := []string{}
    helper(root, &result,path)
    fmt.Println(result)
    return result
}
func helper(root *TreeNode,ret **[]string,path []int) {
    // terminator
    if root == nil {
        // 将path copy 到新的数组中，然后放入到 ret中
        tmp := make([]stirng,len(path))
        copy(path, tmp)
        *ret = append(*ret, tmp)
        return
    }
    // process current logic
    // 将数放入到 path中
    path = append(path, root.Val)
    helper(root.Left,ret,path)
    helper(root.Right,ret,path)
    // dirll down
    // restore current status
}

函数参数经常凉凉
看了官方题解，题解里面定义了全局变量
然后和自己思路不一样的地方在于 先处理根节点，将根节点放到子路径中

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var paths []string
func binaryTreePaths(root *TreeNode) []string {
    paths = []string{}
    constructPaths(root, "")
    return paths    
}
func constructPaths(root *TreeNode,path string) {
    if root != nil {
        pathSB := path
        pathSB += strconv.Itoa(root.Val)
        
        if root.Left == nil && root.Right == nil {
            paths = append(paths, pathSB)
        }else {
            pathSB += "->"
            constructPaths(root.Left, pathSB)
            constructPaths(root.Right, pathSB)
        }
    }
}

题解看着不是很舒服：
看到这个题解里面这样写看着还是蛮舒服的
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    result := []string{}
    helper(root,strconv.Itoa(root.Val),&result)
    return result
}
func helper(root *TreeNode,str string,result *[]string) {
    if root.Left == nil && root.Right == nil {
        *result = append(*result, str)
        return
    }
    if root.Left != nil {
        helper(root.Left, str + "->" + strconv.Itoa(root.Left.Val),result)
    }
    if root.Right != nil {
        helper(root.Right, str + "->" + strconv.Itoa(root.Right.Val), result)
    }
}

BFS:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    paths := []string{}
    if root == nil {
        return paths
    }
    nodeQueue := []*TreeNode{}
    pathQueue := []string{}
    nodeQueue = append(nodeQueue, root)
    pathQueue = append(pathQueue, strconv.Itoa(root.Val))
    for i := 0;i < len(nodeQueue);i++ {
        node,path := nodeQueue[i],pathQueue[i]
        if node.Left == nil && node.Right == nil {
            paths = append(paths, path)
            continue
        }
        if node.Left != nil {
            nodeQueue = append(nodeQueue, node.Left)
            pathQueue = append(pathQueue, path + "->" + strconv.Itoa(node.Left.Val))
        }
        if node.Right != nil {
            nodeQueue = append(nodeQueue, node.Right)
            pathQueue = append(pathQueue, path + "->" + strconv.Itoa(node.Right.Val))
        }
    }
    
    return paths
}

复杂度分析：
时间复杂度；O(n) 遍历每个节点个数
空间复杂度：O(n) 递归开辟栈空间大小 或者 迭代使用节点数组进行存储

总结：
1. 对带路径带参数的题，特别是字符串的题容易懵逼,和上周路径总和思路类似

2. 懵逼的原因在于知识点没有完全掌握，chunk it up，把知识点拆碎，揉烂

3. 多练习
