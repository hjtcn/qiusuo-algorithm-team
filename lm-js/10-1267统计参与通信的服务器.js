
// 这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。

// 如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。

// 请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。

/**
 * @param {number[][]} grid
 * @return {number}
 */
var countServers = function (grid) {
    let row = grid.length, col = grid[0].length, res = 0, num = 0, w = 0;
    let rowIndexMap = new Map(), colIndexMap = new Map()
    for (let i = 0; i < row; i++) {
        for (let j = 0; j < col; j++) {
            if (grid[i][j] == 1) {
                //记录出现i,j值的次数
                rowIndexMap.set(i, rowIndexMap.has(i) ? rowIndexMap.get(i) + 1 : 1)
                colIndexMap.set(j, colIndexMap.has(j) ? colIndexMap.get(j) + 1 : 1)
            }
        }
    }
    for (let i = 0; i < row; i++) {
        for (let j = 0; j < col; j++) {
            if (grid[i][j] == 1) {
                num++;
                if (rowIndexMap.get(i) == 1 && colIndexMap.get(j) == 1) {
                    w++;
                }
            }
        }
    }
    return num - w

};
/**题解
 暴力解决法。记录出现i，j值的个数，最后再次双层遍历。
 如果该元素grid[i][j]的i，j值都只出现了一次，则该服务器无法通信。记录服务器总个数然后减去无法通信的服务器个数
  复杂度分析：
    平均时间复杂度是:O(N*M)
    双层for循环，分别是行数和列数的遍历
    空间复杂度：O(N+M)
    声明两个map存放i，j的出现次数
 */


/**
* @param {number[][]} grid
* @return {number}
*/
var countServers = function (grid) {
    let sum = 0, row = grid.length, col = grid[0].length;
    for (let i = 0; i < row; i++) {
        for (let j = 0; j < col; j++) {
            if (grid[i][j] == 1) {
                let ret = dfs(grid, i, j);
                if (ret > 1) sum += ret;
            }
        }
    }
    return sum
};

function dfs(grid, x, y) {
    let row = grid.length, col = grid[0].length
    if (x < 0 || x > row - 1 || y < 0 || y > col - 1) return 0;
    if (grid[x][y] != 1) return 0;

    grid[x][y] = -1;
    let sum = 1;

    // 搜索整列
    for (let i = 0; i < row; i++) {
        sum += dfs(grid, i, y);
    }

    // 搜索整行
    for (let i = 0; i < col; i++) {
        sum += dfs(grid, x, i);
    }

    return sum;
}

/**题解
 深搜。整行整列搜索当前包含几个服务器,dfs函数返回服务器的个数值，如果该行该列的服务器个数大于1，则该行该列的服务器都能通信，则res加该个数，遍历过后判断有多少个可通信的服务器
  复杂度分析：
    平均时间复杂度是:O(N*M)
    深搜整行整列搜索，则N*M,事实证明耗时会更长一点
    空间复杂度：O(1）
    声明几个变量
 */