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

function pondSizes(land: number[][]): number[] {
    const areaList:number[] = [];
    let areaSize = 0;

    land.forEach((landLayout, yAxis) => {
        landLayout.forEach((num, xAxis) => {
            DFS(xAxis, yAxis);
            // 只有判断出了水域大小才保存值
            if(areaSize){
                areaList.push(areaSize);
                areaSize = 0;
            }
        })
    })

    function DFS(xAxis: number, yAxis: number){
        // 判断外层边界
        const landRound = land[yAxis];
        if(!landRound) return;

        const num = landRound[xAxis];

        if(num === 0){
            areaSize ++;
            land[yAxis][xAxis] = -1;
            // 有判断这四个方向的必要，因为再深度遍历的时候可能只有一个方向判断过
            DFS(xAxis - 1, yAxis); // 左
            DFS(xAxis - 1, yAxis - 1); // 左上
            DFS(xAxis, yAxis - 1); // 上
            DFS(xAxis + 1, yAxis - 1); // 右上

            DFS(xAxis + 1, yAxis); // 右
            DFS(xAxis - 1, yAxis + 1); // 左下
            DFS(xAxis, yAxis + 1); // 下
            DFS(xAxis + 1, yAxis + 1); // 右下
            
        }
    }

    return areaList.sort((a, b) => a-b);
};