package main

/*
要求:


尝试：
(未完成)

学习：
依次检查行、列、3*3小方格是否无重复

关键点：

*/

func isValidSudoku(board [][]byte) bool { // faster 70.57% less 50.00%
	for row := 0; row < 9; row++ {
		if !isValidSudokuRow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValidSudokuCol(board, col) {
			return false
		}
	}

	for pod := 0; pod < 9; pod++ {
		if !isValidSudokuPod(board, pod) {
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			fmt.Println("Invalid row: ", row)
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			fmt.Println("Invalid col: ", col)
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3

	for drow := 0; drow < 3; drow++ {
		for dcol := 0; dcol < 3; dcol++ {
			n := convertToNumber(board[row+drow][col+dcol])
			if n < 0 {
				continue
			}
			if nums[n] {
				fmt.Println("Invalid pod:", pod)
				fmt.Printf("Found duplicate number %d on row %d, col %d", n, row+drow, col+dcol)
				return false
			}
			nums[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}

func main() {
	board := make([][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........")
	})
	result := isValidSudoku(board)
	fmt.Println(result)
}
