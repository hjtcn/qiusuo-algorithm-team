// 690. 员工的重要性
// 给定一个保存员工信息的数据结构，它包含了员工唯一的id，重要度 和 直系下属的id。

// 比如，员工1是员工2的领导，员工2是员工3的领导。他们相应的重要度为15, 10, 5。那么员工1的数据结构是[1, 15, [2]]，员工2的数据结构是[2, 10, [3]]，员工3的数据结构是[3, 5, []]。注意虽然员工3也是员工1的一个下属，但是由于并不是直系下属，因此没有体现在员工1的数据结构中。

// 现在输入一个公司的所有员工信息，以及单个员工id，返回这个员工和他所有下属的重要度之和。

// 示例 1:

// 输入: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
// 输出: 11
// 解释:
// 员工1自身的重要度是5，他有两个直系下属2和3，而且2和3的重要度均为3。因此员工1的总重要度是 5 + 3 + 3 = 11。
// 注意:

// 一个员工最多有一个直系领导，但是可以有多个直系下属
// 员工数量不超过2000。

/*
 * @lc app=leetcode.cn id=690 lang=javascript
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * function Employee(id, importance, subordinates) {
 *     this.id = id;
 *     this.importance = importance;
 *     this.subordinates = subordinates;
 * }
 */

/**
 * @param {Employee[]} employees
 * @param {number} id
 * @return {number}
 */
var GetImportance = function(employees, id) {
    let map=new Map(),subordinateArr=[],sum=0
    // let visited=[]
    for(let i=0;i<employees.length;i++){
        //关系用map记录
        map.set(employees[i].id,[employees[i].importance,employees[i].subordinates])
        //目标id
        if(employees[i].id==id){
            sum+=employees[i].importance
            subordinateArr=employees[i].subordinates
        }
    }
     let dfs=(arr)=>{
        for(let i=0;i<arr.length;i++){
            // 这里本来加了一层判断，防止读过的下属，再别人的下属中读取到
            //后来看到注意信息：一个员工最多有一个直系领导
            // if(!visited[arr[i].id]){
                let [importance,subordinates]=map.get(arr[i])
                sum+=importance
                //标记已经读过此员工
                // visited[arr[i]]=1
                dfs(subordinates)
            // }
        }
    }
    dfs(subordinateArr)

    return sum
};

/** 
 递归。刚开始思路错了，本来这周出的题简单，我就以为出的也太水了，只寻找直属下属。
 后来提交一次发现数据不对的时候又从新读了读题。还是要找下属的下属的
 ok，递归要开始了。其实我觉得用的比较爽的还是map和解构赋值。直接记录每一个id的重要性及其下属，然后从目标id查起，递归找下属

时间复杂度：O(N)
一层遍历记录关系，递归遍历目标员工的下属，下属的下属，和深度有关
空间复杂度：O(N)
这个有点不确定，最大的空间消耗在于记录id对应的重要性和其下属组

*/


var GetImportance = function(employees, id) {
    let queue=[id],visited=[],map=new Map(),sum=0
    for(let i=0;i<employees.length;i++){
        map.set(employees[i].id,[employees[i].importance,employees[i].subordinates])
    }
    while(queue.length){
        let target=queue.shift()
        let [importance,subordinates]=map.get(target)
        sum+=importance
        queue=queue.concat(subordinates)
    }
    return sum
};

/** 
 队列迭代。递归做完，迭代就比较顺利了，因为数据构建过程是自己控制的。
 队列不断添加下属的下属去进行迭代

时间复杂度和空间复杂度基本同上
*/