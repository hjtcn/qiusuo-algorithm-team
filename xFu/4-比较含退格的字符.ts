function backspaceCompare(S: string, T: string): boolean {

    let maxLenList = T.split('');
    let minLenList = S.split('');
    // 使用数据交换，减少split分割数据时间
    if(S.length > T.length){
        let temp = maxLenList;
        maxLenList = minLenList;
        minLenList = temp;
    }

    let stackS: string[] = [];
    let stackT: string[] = [];

    // 只遍历长的栈
    while(maxLenList.length){
        let stackTopMax = maxLenList.shift();
        let stackTopMin = minLenList.shift();
        // S栈操作
        operateStack(stackS, stackTopMax);
        // T栈操作
        operateStack(stackT, stackTopMin);
    }

    // 不使用字符串判断，模拟栈的操作
    // 执行用时: 88 ms
    // 内存消耗: 40.6 MB
    if(stackS.length !== stackT.length){
        return false;
    }else{
        while(stackS.length){
            let outS = stackS.shift();
            let outT = stackT.shift();
            if(outS !== outT){
                return false;
            }
        }
        return true;
    }

    // 字符串写法
    // 执行用时：84 ms, 在所有 typescript 提交中击败了86.89%的用户
    // 内存消耗：40.7 MB, 在所有 typescript 提交中击败了14.29%的用户
    // let compare1 = stackS.join('');
    // let compare2 = stackT.join('');
    // return compare1 === compare2;
};

// 栈的操作
function operateStack (stackList: string[], flag: string | undefined):void {
    if(!flag) return;
    flag === '#' ? stackList.shift() : stackList.unshift(flag)
}