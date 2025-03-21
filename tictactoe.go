package main

import (
	"fmt"
	"strconv"
)

type ticactoe struct {
	tArray [3][3]string
	turn   string
	winner string
}

func (t *ticactoe) drawBoard() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("        A           B         C     ")
	fmt.Println("  ==================================")
	fmt.Println("  |          |          |          |")
	fmt.Println("1 |    ", t.tArray[0][0], "   |    ", t.tArray[0][1], "   |   ", t.tArray[0][2], "    |")
	fmt.Println("  |          |          |          |")
	fmt.Println("  ==================================")
	fmt.Println("  |          |          |          |")
	fmt.Println("2 |    ", t.tArray[1][0], "   |    ", t.tArray[1][1], "   |   ", t.tArray[1][2], "    |")
	fmt.Println("  |          |          |          |")
	fmt.Println("  ==================================")
	fmt.Println("  |          |          |          |")
	fmt.Println("3 |    ", t.tArray[2][0], "   |    ", t.tArray[2][1], "   |   ", t.tArray[2][2], "    |")
	fmt.Println("  |          |          |          |")
	fmt.Println("  ==================================")
}

func (t *ticactoe) putSymbol(pos string) bool {
	dict := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
	}
	col, _ := strconv.Atoi(pos[1:])
	col--

	row := dict[pos[0]]

	// make sure position not taken
	if t.tArray[col][row] != " " {
		return false
	}

	// validate pos is under area
	if col < 0 || col > 2 {
		return false
	}
	if row < 0 || row > 2 {
		return false
	}

	//put symbol into array
	if t.turn == "X" {
		t.tArray[col][row] = "X"
	} else {
		t.tArray[col][row] = "O"
	}
	return true
}

func (t *ticactoe) finish() bool {
	// winning conditional
	var winCond [8]string

	r := true
	//run through every posible position
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//diagonal 1
			if i == j {
				winCond[0] += t.tArray[i][j]
			}
			//column
			if i == 0 && j == 0 {
				winCond[1] += t.tArray[i][j]
			}
			if i == 1 && j == 1 {
				winCond[2] += t.tArray[i][j]
			}
			if i == 2 && j == 2 {
				winCond[3] += t.tArray[i][j]
			}
			//row
			if j == 0 && i == 0 {
				winCond[4] += t.tArray[i][j]
			}
			if j == 1 && i == 1 {
				winCond[5] += t.tArray[i][j]
			}
			if j == 2 && i == 2 {
				winCond[6] += t.tArray[i][j]
			}

			//diagonal2
			winCond[7] = t.tArray[2][0] + t.tArray[1][1] + t.tArray[0][2]

			//if there is still space, then not finished yet
			if t.tArray[i][j] == " " {
				r = false
			}
		}
	}

	//check every possible winning condition
	for i := 0; i < 8; i++ {
		if winCond[i] == "XXX" {
			t.winner = "X"
			r = true
		}
		if winCond[i] == "OOO" {
			t.winner = "O"
			r = true
		}
	}
	return r
}

func (t *ticactoe) nextTurn() {
	if t.turn == "X" {
		t.turn = "O"
	} else {
		t.turn = "X"
	}
}

// init object
func newTictactoe() ticactoe {
	return ticactoe{
		tArray: [3][3]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
		turn:   "X",
		winner: " ",
	}
}

func main() {
	fmt.Print("\033[H\033[2J")
	T := newTictactoe()
	T.drawBoard()
	for !T.finish() {
		fmt.Println("winner : ", T.winner)
		for {
			fmt.Println(T.turn, " turn's: (ex: a1, b2 etc..)")
			var pos string
			fmt.Scanln(&pos)
			if T.putSymbol(pos) {
				break
			}
			fmt.Println("fail, make sure in range AND position is not taken!")
		}
		T.drawBoard()

		T.nextTurn()
	}

	fmt.Println("The Winner is ", T.winner)
}
