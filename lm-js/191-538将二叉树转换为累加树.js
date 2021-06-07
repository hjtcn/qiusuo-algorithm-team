/*
538. 把二叉搜索树转换为累加树
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同

 

示例 1：



输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
示例 2：

输入：root = [0,null,1]
输出：[1,null,1]
示例 3：

输入：root = [1,0,2]
输出：[3,3,2]
示例 4：

输入：root = [3,2,4,1]
输出：[7,9,4,10]
 

提示：

树中的节点数介于 0 和 104 之间。
每个节点的值介于 -104 和 104 之间。
树中的所有值 互不相同 。
给定的树为二叉搜索树。

*/

/*
    思路：想到了先遍历右侧，但是左侧怎么处理不知道了。递归过程的顺序和记录深入进去就特别容易乱..
        二叉搜索树都要试图往中序遍历上思考
    题解：1.中序遍历反着来。右->根->左
    举例说明执行顺序
             4(5)
     1(8)             6(3)
0(9)     2(7)     5(4)    7(2)
           3(6)              8(1)

    首先中序遍历反向。先右,后记录和,并赋值根，再右
    let sum=0
    let dfs=(root){
        if(!root) return 
        dfs(root.right)
        sum+=root.val
        root.val=sum
        dfs(root.left)
        return root
    }
    2.迭代：栈和前缀和
    借用出入栈模拟后->根->左的过程
    迭代模拟中序遍历的过程好陌生呀。
    有点懵
*/ 

var convertBST = function(root) {
    let prevNode=new TreeNode()
    prevNode.next=root
    let sum=0
    let dfs=(root)=>{
        if(!root){
            return 0
        }
        dfs(root.right)
        sum+=root.val
        root.val=sum
        dfs(root.left)
        return root   
    }
    dfs(root)
    return prevNode.next
};

var convertBST = function(root) {
    let prevNode=new TreeNode(0)
    prevNode.next=root
    let stack=[]
    let sum=0
    while(stack.length||root!=null){
        if(root){
            stack.push(root)
            root=root.right
        }
        else{
            root=stack.pop()
            root.val+=sum
            sum=root.val
            root=root.left
        }
    }
    return prevNode.next
}


/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    迭代用的不太熟呀，栈,递归,二叉树混为一体的一道题
*/