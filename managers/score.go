package managers

type Player int

const (
	Human = iota
	Computer
)

var playerScore, computerScore int

func InitScore() {
	playerScore = 0
	computerScore = 0
}

func AddScore(player Player) {
	if player == Human {
		playerScore++
	}

	if player == Computer {
		computerScore++
	}
}

func GetHumanScore() int {
	return playerScore
}

func GetComputerScore() int {
	return computerScore
}
