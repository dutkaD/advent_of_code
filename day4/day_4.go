package day4

import (
	"strconv"
	"strings"
)

func GetLosingBoard(lines []string) int {
	bingoNumbers := sliceToInts(strings.Split(lines[0], ","))

	boards := readBoards(lines)
	losingBoard := -1
	var result int
	bingoCount := -1

	for b := 0; b < len(boards); b++ {
		boardWon := false
		for n := 0; n < len(bingoNumbers); n++ {
			if boardWon {
				break
			}
			for row := 0; row < len(boards[b]); row++ {
				if boardWon {
					break
				}
				for col := 0; col < len(boards[b][row]); col++ {
					if boardWon {
						break
					}
					if bingoNumbers[n] == boards[b][row][col] {
						boards[b][row][col] = -1
						won := checkBoardWon(boards[b], row, col)
						if won {
							boardWon = true
							if n > bingoCount && b != losingBoard {
								losingBoard = b
								bingoCount = n
								result = getResult(boards[b]) * bingoNumbers[n]
							}
							break
						}
					}
				}
			}
		}
	}

	return result
}

func GetBingoWinner(lines []string) int {
	bingoNumbers := sliceToInts(strings.Split(lines[0], ","))
	boards := readBoards(lines)
	for n := 0; n < len(bingoNumbers); n++ {
		for b := 0; b < len(boards); b++ {
			for l := 0; l < len(boards[b]); l++ {
				for p := 0; p < len(boards[b][l]); p++ {
					if bingoNumbers[n] == boards[b][l][p] {
						boards[b][l][p] = -1
						won := checkBoardWon(boards[b], l, p)
						if won {
							return getResult(boards[b]) * bingoNumbers[n]
						}
					}
				}
			}
		}
	}
	return 0
}

func checkBoardWon(board [][]int, row int, column int) bool {
	sumRow := 0
	for c := 0; c < len(board[row]); c++ {
		sumRow += board[row][c]
	}
	sumCol := 0
	for c := 0; c < len(board); c++ {
		sumCol += board[c][column]
	}

	if sumCol == -len(board) || sumRow == -len(board[0]) {
		return true
	}
	return false

}

func readBoards(lines []string) [][][]int {
	var allBoards [][][]int
	var board [][]int
	for a := 2; a < len(lines); a++ {
		if len(lines[a]) == 0 {
			allBoards = append(allBoards, board)
			board = nil
		} else {
			numbersStr := strings.Fields(lines[a])
			numbers := sliceToInts(numbersStr)
			board = append(board, numbers)
		}
	}
	return allBoards
}

func sliceToInts(numbersStr []string) []int {
	var numbers []int
	for c := 0; c < len(numbersStr); c++ {
		num, _ := strconv.Atoi(numbersStr[c])
		numbers = append(numbers, num)
	}
	return numbers
}

func getResult(board [][]int) int {
	sum := 0
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] != -1 {
				sum += board[r][c]
			}
		}
	}
	return sum
}
