// 题目及解题思路更新到‘解题思路.md’

function updateBoard(board: string[][], click: number[]): string[][] {
    
    return dfs(click[0], click[1]);
    function dfs(xAxis, yAxis){
        const xBoard = board[xAxis];
        if(xBoard && yAxis >= 0){
            const value = xBoard[yAxis];
        if(value === 'M'){
            board[xAxis][yAxis] = 'X';
        }else if(value === 'E') {
            const num = sumM(board, [xAxis, yAxis]);
            if(num === 0) {
                board[xAxis][yAxis] = 'B';
                dfs(xAxis-1, yAxis-1);
                dfs(xAxis, yAxis-1);
                dfs(xAxis+1, yAxis-1);
                dfs(xAxis-1, yAxis);

                dfs(xAxis, yAxis+1);
                dfs(xAxis-1, yAxis+1);
                dfs(xAxis+1, yAxis);
                dfs(xAxis+1, yAxis+1);
            }else{
                board[xAxis][yAxis] = '' + num;
            }
        }
        
        }
        return board;
    }
    
};

function sumM (board: string[][], click: number[]): number {
    const x = click[0], y = click[1];
    
    return isM(x-1, y-1) + isM(x, y-1)
    + isM(x+1, y-1) + isM(x-1, y)
    + isM(x, y+1) + isM(x-1, y+1)
    + isM(x+1, y) + isM(x+1, y+1)


    function isM(clickX: number, clickY: number) {
        const xBoard = board[clickX];
        return xBoard && xBoard[clickY] === 'M' ? 1 : 0;
    }
}