// 589. N叉树的前序遍历
// 给定一个 N 叉树，返回其节点值的前序遍历。

// 例如，给定一个 3叉树 :

 

// 返回其前序遍历: [1,3,5,6,2,4]。

 

// 说明: 递归法很简单，你可以使用迭代法完成此题吗?

// 递归

/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder = function(root, mapArr = []) {
    if(root){
        mapArr.push(root.val);
        root.children.forEach(node => {
            preorder(node, mapArr);
        })
    }
    return mapArr;
};

// 迭代

/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder1 = function(root) {
    const stack = [];
    const mapArr = [];

    stack.push(root);

    while(stack.length){
        const node = stack.pop();// 出栈一个元素

        // 结点不一定存在
        if(!node) continue;

        mapArr.push(node.val);

        /**
         * stack.push(...node.children.reverse())
         */
        const child = node.children;
        for(let index = (child.length - 1);index >= 0;index--) {
            stack.push(child[index]);
        }
    }

    return mapArr;
};