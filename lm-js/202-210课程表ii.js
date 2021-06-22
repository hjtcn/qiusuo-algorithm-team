/*
210. 课程表 II
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

示例 1:

输入: 2, [[1,0]] 
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2:

输入: 4, [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
说明:

输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
提示:

这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。


*/

/*
    思路：首先要知道有向图的出度入度
         自己写这道题的时候，尝试了map去记录依赖关系。
         但是了解懂啊入度出度之后，发现自己依赖关系也记录反了。
         对于这道题来说，应该记录。
         先修完一门课，接下来能修那些课。
         如:[[1,0],[2,0],[3,1],[3,2]]
         遍历记录
         依赖关系：修完0，能修1，     入度：1:1
                 修完0，能修[1,2]        2:1
                 修完1，能修3            3:1
                 修完2，能修3            3:2
        根据依赖关系处理入度：
        队列储存入度为0的课程
        while(queue.length){
            let curSubject=queue.shift() //能修的课程
            //将依赖curSubject的课程入度遍历并减一。
            let nextArr=map.get[curSubject]||[]
            nextArr.forEach(item=>{
                //每一种课程入度减一。
                //当入度为0时再次入队
            })

        }

        结果处理：当queue出队时，target不断push
                如果target.length!=numCourses,则返回[]
                反之返回target

    题解：递归要检测有向图有没有环。不太好理解
        1.记录依赖关系邻接表
        2.遍历递归判断有向图是否存在环
          1）遍历条件：从0～numCourses,同时递归中 会确定是否存在环，一旦存在，停止遍历。
          2）递归过程：
          比如3  [[1,0],[1,2],[0,1]]
          借用visited数组记录是否访问过。
              初始化都为0
              如果访问过dfs(0),visited[0]=1
              遍历，依赖0的数组。
                如果依赖课程也不曾访问过，继续递归(1)
                。。。
             而当dfs(1)时，遍历依赖1的数组：
                嘿：又找到0了，此时visited[0]已经为1，说明出现环了。valid = false;

        3.存在环返回[];
          不存在返回target.

*/


var findOrder = function(numCourses, prerequisites) {
    let len=prerequisites.length,map=new Map(),target=[]
    let inDegree=new Array(numCourses).fill(0)//入度数组
    for(let i=0;i<len;i++){
        inDegree[prerequisites[i][0]]++
        let tSubject=map.get(prerequisites[i][1])
        //记录依赖关系
        if(tSubject){
            map.set(prerequisites[i][1],[...tSubject,prerequisites[i][0]])
        }
        else{
            map.set(prerequisites[i][1],[prerequisites[i][0]])
        }
    }
    let queue=[]
    for(let i=0;i<numCourses;i++){
        if(inDegree[i]==0){
            queue.push(i)
        }
    }
    while(queue.length){
        let curSubject=queue.shift()
        target.push(curSubject)
        let nextArr=map.get(curSubject)||[]
        nextArr.forEach(item=>{
            inDegree[item]--
            if(inDegree[item]==0){
                queue.push(item)
            }
        })
    }
    return target.length==numCourses?target:[]
};



var findOrder = function(numCourses, prerequisites) {
    // 构造临接链表
    var adj = {};
    for(var i = 0;i < numCourses;i++) {
        adj[i] = [];
    }
    for(var [first, second] of prerequisites) {
        adj[second].push(first);
    }
    var valid = true;
    var target = [];
    var visited =new Array(numCourses).fill(0);
    var dfs = function(i) {
        visited[i] = 1;
        //visited记录是否学过i判断环的核心。
        for(var successor of adj[i]) {
            if (visited[successor] === 0) {
                dfs(successor);
            } else if (visited[successor] === 1) {
                //出现环了
                valid = false;
            }
        }
        visited[i] = 2;
        target.unshift(i);
    }
    for(var i = 0;i < numCourses && valid;i++) {
        if (visited[i] === 0) {
            dfs(i);
        }
    }
    if (!valid) {
        return [];
    }
    return target;
};

/*
    时间复杂度：O(N+M)  其中N为课程数，M为先修课程的要求数
    空间复杂度：O(N+M)
*/


/*
    思考：有向图：入度出度
         依赖先后关系更新入度出度：
         后出，先入度减一
         
         dfs拓扑排序：判断是否有环
*/