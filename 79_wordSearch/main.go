package wordsearch

func exist(board [][]byte, word string) bool {

	// we need to find the first letter
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			// Once we get it we start backstracking
			if backtrack(board, word, r, c, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, row, col, idx int) bool {
	// End condition we are otu of the board
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 || board[row][col] == '0' {
		return false
	}

	// the letter doesn't match
	if board[row][col] != word[idx] {
		return false
	}

	// we are looking at the last letter
	if idx == len(word)-1 {
		return word[idx] == board[row][col]
	}

	// We start a the index idx
	current := board[row][col]

	// Set the current to a dummy
	// board[row][col] = ""

	//search in the 4 direction
	left := backtrack(board, word, row, col-1, idx+1)
	right := backtrack(board, word, row, col+1, idx+1)
	down := backtrack(board, word, row-1, col, idx+1)
	up := backtrack(board, word, row+1, col, idx+1)

	//if we have any matching we return
	if left || right || down || up {
		return true
	}

	board[row][col] = current

	return false
}
