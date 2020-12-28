// 面试题 16.19. 水域大小
// 你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。

// 示例：

// 输入：
// [
//   [0,2,1,0],
//   [0,1,0,1],
//   [1,1,0,1],
//   [0,1,0,1]
// ]
// 输出： [1,2,4]
// 提示：

// 0 < len(land) <= 1000
// 0 < len(land[i]) <= 1000



/**
 * @param {number[][]} land
 * @return {number[]}
 */
var pondSizes = function(land) {
    let lenX=land.length,lenY=land[0].length,area=null,res=[]
    let dfs=(x,y)=>{
        //边界值
        if(x<0||x>=lenX||y<0||y>=lenY){
            return ;
        }
        if(!land[x][y]){
            //标记走过的路
            land[x][y]=1
            area++;
            //八个方向
            dfs(x+1,y)
            dfs(x,y+1)
            dfs(x+1,y+1)
            dfs(x-1,y)
            dfs(x-1,y-1)
            dfs(x,y-1)
            dfs(x-1,y+1)
            dfs(x+1,y-1)
        }
      
    }
    for(let i=0;i<lenX;i++){
        for(let j=0;j<lenY;j++){
            //统计每一个值为0的区域
            if(!land[i][j]){
                area=0
                dfs(i,j)
                // console.log(area)
                res.push(area)
            }
           
        }
    } 
    return res.sort((a,b)=>a-b)
};

/**
    递归，统计八个方向，核心在于走过的路要标记
    * 时间复杂度:O(N*M)
    * 每个位置只遍历一次
    * 空间复杂度:O(1)
    * 除了结果数组，没声明其余数组变量
 */

let direction=[[1,-1],[1,1],[1,0],[-1,-1],[-1,1],[-1,0],[0,1],[0,-1]]
var pondSizes = function(land) {
    let queue=[],temp_count=0,res=[]
    for(let i=0;i<land.length;i++){
        for(let j=0;j<land[0].length;j++){
            if(land[i][j]==0){
                land[i][j]=1
                queue.push([i,j])//记录下标
                temp_count=1
                while(queue.length){
                    let [x0,y0]=queue.shift()//弹出队列头的下标位置
                    //统计八个方向的值
                    for(let k=0;k<8;k++){
                        let x=x0+direction[k][0]
                        let y=y0+direction[k][1]
                        //在边界内，且值为0
                        if(x>=0&&x<land.length&&y>=0&&y<land[0].length&&land[x][y]==0){
                            //水域大小增加
                            temp_count++;
                            queue.push([x,y])//队列增添新的下标值，多了扩展空间
                            land[x][y]=1//走过的路进行标记
                        }
                    }
                }
                res.push(temp_count)
            }
        }
    }
    return res.sort((a,b)=>a-b)
}

/**
 * 队列模拟迭代，说是广搜，但是已经没有广搜的视觉感了。
 * 踩过一个坑，direction写的双层数组有个位置写错了，没办法，又一个一个打印去找错误的位置。更可怕是数量很多的用例出错，打印的值也没有规律
 * 核心还在标记的过程，并且遍历八个方向，直到没有可拓展的水域为止
 * 时间复杂度:O(N*M)
 * 每个位置只遍历一次
 * 空间复杂度:O(N*M)
 * 队列记录下标值了
 */
