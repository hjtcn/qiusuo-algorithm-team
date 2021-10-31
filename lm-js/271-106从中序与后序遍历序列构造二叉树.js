/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

/*
   思路:中序遍历  左根右
       后序遍历  左右根
        用两个数组构建一个二叉树。递归的方案，那就找出根节点。不断切分左右子节点。
        1.根节点是谁？从后序遍历的数组看。他的val就是后序遍历数组的最后一个元素rootVal。
        let root=nwe TreeNode(rootVal)
        2.左右子节点在哪？知道了根的位置。把中序数组切分。
        let index=inorder.indexOf(rootVal)
        中序数组变为[0,index] [index+1,结尾]
        后序数组变为[0,index] [index,后序数组结尾,不包括最后一个节点] 
        因此递归的过程为
        root.left=buildTree(inorder.slice(0,index),postorder.slice(0,index)
        root.right=buildTree(inorder.slice(index+1),postorder.slice(index,len-1))
   题解：迭代。。不太好理解。
        后序遍历中 一个连续的a,b两个节点。左右根
        会出现两种情况
        1.a是b的右节点
        2.b没有右儿子，a是b的祖先元素上(包括本身)的左儿子。即可能：a是b的左儿子。如果b没有左儿子，就向上找祖先元素。直到找到有左儿子(且a不在它的左儿子中)的节点C,则a是C的左儿子。

        原理：用栈来维护所有还没有考虑过左儿子的祖先节点。指针处理中序遍历。找子节点。
        1.初始状态。根节点(后序遍历最后一个元素)入栈。指针指向指向序遍历最后一个元素。
        2.以此枚举后序遍历中除了第一个节点外的每一个元素。
            (1)如果index恰好指向栈顶节点，则不断的弹出栈顶元素，并向左移动index，将当前节点作为最后一个弹出的栈顶元素(C)的左儿子。
            (2)如果index的指向栈顶元素不同。则当前节点为栈顶节点的右儿子。
        3.无论那种情况，都将当前节点入栈。


*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */
 var buildTree = function(inorder, postorder) {
    let len=postorder.length
    if(len<=0) return null;
    let rootVal=postorder[len-1]
    let root=new TreeNode(rootVal)
    let index=inorder.indexOf(rootVal)
    root.left=buildTree(inorder.slice(0,index),postorder.slice(0,index)
    )
    root.right=buildTree(inorder.slice(index+1),postorder.slice(index,len-1))
    return root
};


var buildTree = function(inorder, postorder) {
    if(postorder.length==0) return null
    let stack=[]
    let len=postorder.length-1
    let root=new TreeNode(postorder[len])
    stack.push(root)
    let index=inorder.length-1
    for(let i=len-1;i>=0;i--){
        let nodeVal=postorder[i]
        let node=stack[stack.length-1]
        if(node.val!=inorder[index]){
            node.right=new TreeNode(nodeVal)
            stack.push(node.right)
        }
        else{
            while(stack.length&&stack[stack.length-1].val==inorder[index]){
                //这个位置踩了坑。stack是存储节点的。对比val时需要取节点值
                node=stack.pop()
                index--;
            }
            node.left=new TreeNode(nodeVal)
            stack.push(node.left)
        }
    }
    return root
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：递归切分数组还是比较好理解的，分而治之。
         我也回过头看了之前写过的105题从前序遍历与中序遍历构建二叉树的题目。
         递归还是可以持续优化的，比如一些用到js专属api的地方。indexOf，可以用map去记录位置去代替。
         slice切换。可以将递归函数的参数仅记录索引位置，不记录切分后的数组。会减少空间复杂度。

         迭代是真的巧妙。往上找祖先元素，有点太难理解了。多写两次，慢慢熟悉。
*/