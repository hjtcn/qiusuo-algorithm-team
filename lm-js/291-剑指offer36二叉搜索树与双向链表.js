/*
剑指 Offer 36. 二叉搜索树与双向链表
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

 

为了让您更好地理解问题，以下面的二叉搜索树为例：

 



 

我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。

 



 

特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

 

注意：本题与主站 426 题相同：https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/

注意：此题对比原题有改动。



*/


/*
    思路：首先中序遍历的结构弄出来，然后记录prev和cur变量。
         prev.right=cur
         cur.left=prev

        最后跳出递归，即遍历完成后，头尾也需要拼接一下。

        迭代中序遍历又忘记了，生气

*/
var treeToDoublyList = function(root) {
    if(!root) return root
    let target=head=new Node()
    let dfs=(root)=>{
        if(!root) return null
        dfs(root.left)
        prev=head
        head=new Node(root.val)
        prev.right=head
        head.left=prev
        dfs(root.right)
    }
    dfs(root)
    head.right=target.right
    target.right.left=head
    return target.right
};




var treeToDoublyList = function(root) {
    if(!root) return root
    let target=head=new Node()
    let stack=[] 
    let p=root

    while(p||stack.length){
        while(p){
            stack.push(p)
            p=p.left
        }
        p=stack.pop()
        let prev=head
        head=new Node(p.val)
        prev.right=head
        head.left=prev
        p=p.right
    }
    head.right=target.right
    target.right.left=head
    return target.right
};


/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：从基础框架，到扩展拼接，多练习就能化繁为简
*/
