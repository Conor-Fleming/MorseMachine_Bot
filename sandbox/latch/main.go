//write a program that generates all possible tic tac toe game configurations
//output the number of valid games

package main

import "fmt"

func isWinner(board []byte, player byte) bool {
	winPositions := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	for _, positions := range winPositions {
		if board[positions[0]] == player && board[positions[1]] == player && board[positions[2]] == player {
			return true
		}
	}
	return false
}

func isDraw(board []byte) bool {
	for _, cell := range board {
		if cell == ' ' {
			return false
		}
	}
	return true
}

func generateGames(board []byte, player byte) int {
	if isWinner(board, 'X') || isWinner(board, 'O') || isDraw(board) {
		return 1
	}

	count := 0
	//fmt.Println(board)
	for i := 0; i < 9; i++ {
		if board[i] == ' ' {
			newBoard := make([]byte, len(board))
			copy(newBoard, board)
			newBoard[i] = player
			count += generateGames(newBoard, 'O')
		}
	}

	return count
}

func main() {
	initialBoard := make([]byte, 9)
	for i := range initialBoard {
		initialBoard[i] = ' '
	}

	totalGames := generateGames(initialBoard, 'X')
	fmt.Println("Total valid games of Tic Tac Toe:", totalGames)
}
