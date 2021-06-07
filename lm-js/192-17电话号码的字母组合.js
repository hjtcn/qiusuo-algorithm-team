/*17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



 

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
 

提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

*/


/*
    思路:一看见全排列就发懵。先把对应2-9的映射先写出来了。然后就剩下全排列模拟。
    回头去看了看之前写过的46全排列题。
    然后才尝试写的。
    let dfs=(arr)=>{
        1.如果arr的长度和字符串长度一样，说明这是找的目标映射字符串

        2.遍历，拼接映射字符串
        for(let i=0;i<len;i++){
            (1)映射可从中选择数组toArr。为["a","b","c"]或["d","e","f"]等等
            for(let j=0;j<toArr.length;j++){
                //可以从里面选了。
                arr.push(toArr[j])
                dfs(arr)
                arr.pop()再把它拿出来。
            }
        }
    }

    遇到的坑：
    1.若当前选择2对应的其中一个字母，则不能同时选2对应的字母。因此遍历原字符串的i要在dfs的时候++i.当然dfs后面还要把i--减回来。
    因此我将dfs方法的参数变为dfs(arr,++i)
    i初始值也对应beginIndex

    题解：迭代模拟。队列用法还是很巧妙的。
    1.遍历原字符串
      2.遍历当前层的节点个数(当前队列长度)
        逐个让当前层节点curStr出列
        3.遍历映射数组。将curStr+item拼接入队。
    不好理解。我提交完，又打印log还是有点懵。不知其所以然
*/

/**
 * @param {string} digits
 * @return {string[]}
 */
 var letterCombinations = function(digits) {
    let mapArr=[],target=[],len=digits.length
    if(!len) return []
    mapArr[2]=["a","b","c"]
    mapArr[3]=["d","e","f"]
    mapArr[4]=["g","h","i"]
    mapArr[5]=["j","k","l"]
    mapArr[6]=["m","n","o"]
    mapArr[7]=["p","q","r","s"]
    mapArr[8]=["t","u","v"]
    mapArr[9]=["w","x","y","z"]
    let dfs=(arr,beginIndex)=>{
        if(arr.length==len){
            target.push(arr.join(''))
             arr=[]
            return ;
        }
        for(let i=beginIndex;i<len;i++){
            let toArr=mapArr[digits[i]]
            for(let j=0;j<toArr.length;j++){
              arr.push(toArr[j])
                dfs(arr,++i)
                i--
                arr.pop()
            }
        }
    }
    dfs([],0)
    return target
};
/*
    时间复杂度：O(3^m*4^n)
    空间复杂度：O(m+n)

    看题解了解的，本人有点懵。
*/

var letterCombinations = function(digits) {
    let len=digits.length,queue=['']
    if(!len) return []
    let mapArr={
        "2":["a","b","c"],
        "3":["d","e","f"],
        "4":["g","h","i"],
        "5":["j","k","l"],
        "6":["m","n","o"],
        "7":["p","q","r","s"],
        "8":["t","u","v"],
        "9":["w","x","y","z"]
    }
   for(let i=0;i<len;i++){
       let levelSize=queue.length
       for(let j=0;j<levelSize;j++){
           let curStr=queue.shift()
           let toArr=mapArr[digits[i]]
           for(item of toArr){
             queue.push(curStr+item)
           }
       }
   }
   return queue
}
/*
    时间复杂度：O(3^m*4^n)
    空间复杂度：O(3^m*4^n)
*/


/*
    全排列对于递归可以多练习
    迭代，栈和队列有点考灵感了
*/