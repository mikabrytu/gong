package managers

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/ui"
)

type Player int

const (
	Human = iota
	Computer
)

var textPlayer, textComputer *ui.Font
var countPlayer, countComputer int

func InitScore(tPlayer, tComputer *ui.Font) {
	textPlayer = tPlayer
	textComputer = tComputer
	countPlayer = 0
	countComputer = 0
}

func AddScore(player Player) {
	if player == Human {
		countPlayer++
		textPlayer.UpdateText(fmt.Sprintf("%v", countPlayer))
	}

	if player == Computer {
		countComputer++
	}
}

func GetHumanScore() int {
	return countPlayer
}

func GetComputerScore() int {
	return countComputer
}
