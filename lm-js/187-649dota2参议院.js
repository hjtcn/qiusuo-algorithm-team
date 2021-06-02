/*
649. Dota2 参议院
Dota2 的世界里有两个阵营：Radiant(天辉)和 Dire(夜魇)

Dota2 参议院由来自两派的参议员组成。现在参议院希望对一个 Dota2 游戏里的改变作出决定。他们以一个基于轮为过程的投票进行。在每一轮中，每一位参议员都可以行使两项权利中的一项：

禁止一名参议员的权利：

参议员可以让另一位参议员在这一轮和随后的几轮中丧失所有的权利。

宣布胜利：

          如果参议员发现有权利投票的参议员都是同一个阵营的，他可以宣布胜利并决定在游戏中的有关变化。

 

给定一个字符串代表每个参议员的阵营。字母 “R” 和 “D” 分别代表了 Radiant（天辉）和 Dire（夜魇）。然后，如果有 n 个参议员，给定字符串的大小将是 n。

以轮为基础的过程从给定顺序的第一个参议员开始到最后一个参议员结束。这一过程将持续到投票结束。所有失去权利的参议员将在过程中被跳过。

假设每一位参议员都足够聪明，会为自己的政党做出最好的策略，你需要预测哪一方最终会宣布胜利并在 Dota2 游戏中决定改变。输出应该是 Radiant 或 Dire。

 

示例 1：

输入："RD"
输出："Radiant"
解释：第一个参议员来自 Radiant 阵营并且他可以使用第一项权利让第二个参议员失去权力，因此第二个参议员将被跳过因为他没有任何权利。然后在第二轮的时候，第一个参议员可以宣布胜利，因为他是唯一一个有投票权的人
示例 2：

输入："RDD"
输出："Dire"
解释：
第一轮中,第一个来自 Radiant 阵营的参议员可以使用第一项权利禁止第二个参议员的权利
第二个来自 Dire 阵营的参议员会被跳过因为他的权利被禁止
第三个来自 Dire 阵营的参议员可以使用他的第一项权利禁止第一个参议员的权利
因此在第二轮只剩下第三个参议员拥有投票的权利,于是他可以宣布胜利
 

提示：

给定字符串的长度在 [1, 10,000] 之间.
 

*/

/*
    思路：自己遍历的过程有缺失，记录状态和更新状态一直不太对。
         模拟，但要知道如何用栈灵活的模拟。
    题解：双栈模拟：
         先用栈radiant,dire分别统计'R'和'D'所在的位置
         然后循环shift模拟出栈
         dire.shift()
         radiant.shift()
         同时进行，代表一个执行完权利(执行者)，一个被禁掉了。
         执行者权利更新：
         栈顶元素(索引)对比,谁小谁是执行者，执行者相当于它的权利需要转一圈才能使用
         因此所在栈可以push(栈顶元素(索引)+n)

         栈和队列：很厉害，不好理解，需要画画图
         一个用来比较的栈stack,一个放置一次比较胜利者队列winQueue
         stack初始化为空，winQueue初始化为参议员位置数组(这个位置我都有点懵)

         核心逻辑主要做两件事  1.从队列shift中取值对比
                            2.胜利者，从新塞入队列push
         直到队列中没值了，值都在比较队列中，说明对比不出胜利者，剩下的都是一类阵营
         最后stack剩余那一队，就返回那一队
            

         如：DRRDR
         stack     winQueue
第一步      D          R

第二步      R(kill)       
           D(win)     D(新增win)

第三步      R          D  

第四步      D(kill)    R(新增win)
           R(win)     D    

第五步      R          R
                      D

           D(胜利栈增) R(新增win)
第六步      R(win)     R
                      D(kill)

第七步      R(胜利栈增)  R

第八步      R(胜利栈增)
           R(胜利栈增)

*/

var predictPartyVictory = function(senate) {
    let radiant=[],dire=[],len=senate.length
    for(let i=0;i<len;i++){
        if(senate[i]=='R'){
            radiant.push(i)
        }
        else{
            dire.push(i)
        }
    }
    while(radiant.length&&dire.length)
   {         if(radiant[0]<dire[0]){
                radiant.push(radiant[0]+len)
            }
            else{
                dire.push(dire[0]+len)
            }
            dire.shift() 
            radiant.shift()
    }
    return radiant.length?'Radiant':'Dire'
};
// @lc code=end

//栈和队列

var predictPartyVictory = function(senate) {
    let winQueue = senate.split('');
    let stack = [];
    while (winQueue[0]) {
        let data = winQueue.shift();
        if (stack.length === 0 || stack[stack.length-1] === data) {
            stack.push(data);
        } else {
            winQueue.push(stack.shift());
        }
    }
    return stack.pop() === 'R'  ? 'Radiant' : 'Dire';
};


/*
    感觉栈和队列还是不够变通

*/