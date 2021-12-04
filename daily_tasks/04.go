package main

// Every solution is run along with utils.go
// e.g `go run 04.go utils.go`

import (
	"fmt"
	"strconv"
	"strings"
)

type Field struct {
	value  int
	marked bool
}

type Board struct {
	size   int
	fields [][]Field
}

func (b Board) markNumber(num int) {
	for row := 0; row < b.size; row++ {
		for col := 0; col < b.size; col++ {
			if b.fields[row][col].value == num {
				b.fields[row][col].marked = true
			}
		}
	}
}

func (b Board) hasBingo() bool {
	var rowBingo bool
	var colBingo bool
	for i := 0; i < b.size; i++ {
		rowBingo = true
		colBingo = true
		for j := 0; j < b.size; j++ {
			if rowBingo && !b.fields[i][j].marked {
				rowBingo = false
			}
			if colBingo && !b.fields[j][i].marked {
				colBingo = false
			}
			if !rowBingo && !colBingo {
				break
			}
		}
		if rowBingo || colBingo {
			return true
		}
	}
	return false
}

func (b Board) sumUnmarked(called_num int) int {
	var sum_nums int = 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if !b.fields[i][j].marked {
				sum_nums += b.fields[i][j].value
			}
		}
	}
	return sum_nums * called_num
}

func parse_input(lines []string) ([]int, []Board) {
	var bingo_nums []string = strings.Split(lines[0], ",")
	var nums []int
	for _, num_str := range bingo_nums {
		num, _ := strconv.Atoi(num_str)
		nums = append(nums, num)
	}

	var boards []Board
	var board [][]Field
	for _, row_nums := range lines[2:] {
		if row_nums != "" {
			var board_row_str []string = strings.Fields(row_nums)
			var board_row []Field
			for _, num_str := range board_row_str {
				num, _ := strconv.Atoi(num_str)
				board_row = append(board_row, Field{value: num, marked: false})
			}
			board = append(board, board_row)
		} else {
			boards = append(boards, Board{size: len(board), fields: board})
			board = nil
		}
	}
	return nums, boards
}

func part1(lines []string) int {
	nums_to_mark, boards := parse_input(lines)

	// For every board mark number and check for bingo
	for _, num := range nums_to_mark {
		for _, board := range boards {
			board.markNumber(num)
			if board.hasBingo() {
				return board.sumUnmarked(num)
			}
		}
	}
	return -1 // no bingo on any board
}

func part2(lines []string) int {
	nums_to_mark, boards := parse_input(lines)
	var last_winner_board Board
	var last_drawn_num int

	// For every board mark number and check for bingo
	for _, num := range nums_to_mark {
		var not_finished_boards []Board
		for _, board := range boards {
			board.markNumber(num)
			if board.hasBingo() {
				last_winner_board = board
				last_drawn_num = num
			} else {
				not_finished_boards = append(not_finished_boards, board)
			}
		}
		// Keep iterating only through unfinished bingo boards
		boards = not_finished_boards
	}
	return last_winner_board.sumUnmarked(last_drawn_num)
}

func main() {
	lines := getInputLines("data/04.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
