// 递归

/**
 * Definition for node.
 * class Node {
 *     val: number
 *     children: Node[]
 *     constructor(val?: number) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.children = []
 *     }
 * }
 */

function postorder(root: Node | null): number[] {
        // 左 右 中
        let mapArr = [];

        function mapTree (node: Node | null) {
            if(node) {
                node.children.forEach(item => {
                     mapTree(item);
                });
                mapArr.push(node.val);
            }
        }

        mapTree(root);

        return mapArr;
};

// 迭代

/**
 * Definition for node.
 * class Node {
 *     val: number
 *     children: Node[]
 *     constructor(val?: number) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.children = []
 *     }
 * }
 */

function postorder1(root: Node | null): number[] {
    // 左 右 中
    let stack = [];
    let mapArr = [];

    root && stack.unshift(root);

    while(stack.length) {
        // 出栈一个元素
        const topNode = stack.shift();

        // 直接推入结果数组
        mapArr.push(topNode.val);
        
        // 这里就利用ES6语法简化了代码，按照从右到左的顺序将子节点一个一个从栈数组起始位置插入
        // 这里也可以用contact拼接，那就需要替换原stack数组，有点脱离栈的含义
        if(topNode.children && topNode.children.length > 0) {
            stack.unshift(...topNode.children.reverse());
        }
        
    }

    return mapArr.reverse();
};