// 207. 课程表
// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

 

// 示例 1:

// 输入: 2, [[1,0]] 
// 输出: true
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
// 示例 2:

// 输入: 2, [[1,0],[0,1]]
// 输出: false
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 

// 提示：

// 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
// 1 <= numCourses <= 10^5

/*
 * @lc app=leetcode.cn id=207 lang=javascript
 *
 * [207] 课程表
 */

// @lc code=start
/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
var canFinish = function(numCourses, prerequisites) {
    let map=new Map(),queue=[],dependences=new Map()
    for(subject of prerequisites){
          //记录入度
        let count=map.get(subject[0])||0
        map.set(subject[0],++count)
  
        //记录依赖，学完subject[1]，可以学？？？？？[]数组存储
        let dependenceArr=dependences.get(subject[1])||[]
        dependenceArr.push(subject[0])
        dependences.set(subject[1], dependenceArr)
    }
    //如果入度为0，入队
    for(let i=0;i<numCourses;i++){
      if(!map.get(i)){
        queue.push(i)
      }
    }
    while(queue.length){
      let curNode=queue.shift()
      let dependenceArr=dependences.get(curNode)||[]
      //出队过程代表课可以学习了
      numCourses--
      //当前课学完，可以学习的dependenceArr遍历
      for(let i=0;i<dependenceArr.length;i++){
        let count=map.get(dependenceArr[i])
        map.set(dependenceArr[i],--count)
        //入度为0，入队
        if(count==0){
          queue.push(dependenceArr[i])
        }
      }
  
    }
    return numCourses==0
  };
  // @lc code=end
  
/*
    一开始思路都错了，
    1.首先入度为0的课，代表着不需要前提条件就可以学习。直接入队
    2.队首元素。遍历之前统计的依赖课程，入度减一，如果入度为0，则入队。
    3.出队的过程，课程数减一，直到队列长度为0，
    4.判断课程数是否等于0，等于0，则全部都能学习。
    时间复杂度：O(N+M)
    这个不太确定，取决于课程数N*某个课程的依赖课程数M
    空间复杂度：O(N+M)
    队列模拟为O(N)，依赖的邻接表为O(N+M)

*/  