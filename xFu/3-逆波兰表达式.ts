// 150. 逆波兰表达式求值
// 根据 逆波兰表示法，求表达式的值。

// 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

 

// 说明：

// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况

function evalRPN(tokens: string[]): number {

    let stack:number[] = [];

    while(tokens.length){
        let item = tokens.shift();
        // 若待处理的数据是数字就入栈，否则就执行计算方法
        if(!isNaN(Number(item))){
            stack.push(Number(item))
        }else{
            computeNum(stack, item);
        }
    }

    return stack[0];
};

/**
 * 计算当前运算符下的计算结果
 */
function computeNum (stack: number[], type: string | undefined):void {
    let num2 = stack.pop() || 0;
    let num1 = stack.pop() || 0;
    let total = 0;
    switch(type){
        case '+': total = num1 + num2;break;
        case '-': total = num1 - num2;break;
        case '*': total = num1 * num2;break;
        case '/': total = Math.trunc(num1 / num2);break;
        default: total = 0;break;
    }

    // 将计算结果入栈，等待后续处理
    stack.push(total);
}