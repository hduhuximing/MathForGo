package jian_for_offer

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	row, col := len(board), len(board[0])
	//注意：这里初始化，如果不给二维初始化结果不正确
	flag := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		flag[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if hasWord(board, flag, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func hasWord(board [][]byte, flag [][]bool, row int, col int, index int, word string) bool {
	if row < 0 || row >= len(flag) || col < 0 || col >= len(flag[0]) || flag[row][col] || word[index] != board[row][col] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	flag[row][col] = true
	if hasWord(board, flag, row+1, col, index+1, word) ||
		hasWord(board, flag, row, col+1, index+1, word) ||
		hasWord(board, flag, row-1, col, index+1, word) ||
		hasWord(board, flag, row, col-1, index+1, word) {
		return true
	}
	flag[row][col] = false
	return false
}
