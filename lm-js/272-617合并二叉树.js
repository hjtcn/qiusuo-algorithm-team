/*
617. 合并二叉树
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

示例 1:

输入: 
	Tree 1                     Tree 2                  
          1                         2                             
         / \                       / \                            
        3   2                     1   3                        
       /                           \   \                      
      5                             4   7                  
输出: 
合并后的树:
	     3
	    / \
	   4   5
	  / \   \ 
	 5   4   7
注意: 合并必须从两个树的根节点开始。


*/

/*
    思路：递归。
        主要四种情况
        1.root1,root都不为空。
        值相加，新建二叉树
        root=new TreeNode(root1.val+root2.val)
        root.left=mergeTrees(root1.left,root2.left)
        root.right=mergeTrees(root1.right,root.right)
        2.root1为空，root2不为空
        root=root2
        ...
        3.root1不为空，root2为空
        root=root1
        ...
        4.root1,root2都为空
        root=null

        拆分出各种情况就比较清晰了。
        

        题解：写迭代的时候，反而没有那么清晰了。递归。使用的是一个函数，两个参数。
        但是到迭代的时候。队列每次都将root1,root连着入队。出队也同时出两个。在试图用root2将root1覆盖或更新的情况。然后出队入队有四种情况。
        1.root1.left&&root2.left
        2.root1.right&&root2.right
        3.!root1.left&&root1.left
        4.!root.right&&root2.right

        最后返回root1

*/

var mergeTrees = function(root1, root2) {
    let root=null
    if(root1&&root2){
        root=new TreeNode(root1.val+root2.val)
            root.left=mergeTrees(root1.left,root2.left) 
    root.right=mergeTrees(root1.right,root2.right)
    }
    else if(root1&&!root2){
        root=new TreeNode(root1.val)
        root.left=root1.left
        root.right=root1.right
    }
    else if(!root1&&root2){
        root=new TreeNode(root2.val)
        root.left=root2.left
        root.right=root2.right
    }

    return root
};


var mergeTrees = function(root1, root2) {
    if(!root1||!root2) return root1||root2
    let queue=[root1,root2]
    while(queue.length){
        let node1=queue.shift()
        let node2=queue.shift()
        node1.val+=node2.val
        if(node1.left&&node2.left)
        {
            queue.push(node1.left)
            queue.push(node2.left)
        }
        if(node1.right&&node2.right){
            queue.push(node1.right)
            queue.push(node2.right)
        }
        if(!node1.left&&node2.left){
            node1.left=node2.left
        }
        if(!node1.right&&node2.right){
            node1.right=node2.right
        }
    }
    return root1
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
/*
    思考：写的顺了，还是递归好弄一些。迭代的细节太多，有的时候容易顾不上。
*/