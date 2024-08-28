package main

func main() {

}

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		mcolumn := make(map[byte]struct{})
		mrow := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			_, exists := mrow[board[i][j]]
			if string(board[i][j]) != "." && exists {
				return false
			}
			mrow[board[i][j]] = struct{}{}
			_, exists = mcolumn[board[j][i]]
			if string(board[j][i]) != "." && exists {
				return false
			}
			mcolumn[board[j][i]] = struct{}{}
		}
	}
	for k := 0; k < 9; k += 3 {
		for l := 0; l < 9; l += 3 {
			mblock := make(map[byte]struct{})
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					_, exists := mblock[board[i+k][j+l]]
					if string(board[i+k][j+l]) != "." && exists {
						return false
					}
					mblock[board[i+k][j+l]] = struct{}{}
				}
			}
		}
	}
	return true
}
