package managers

import (
	"littlejumbo/gong/objects"
	"littlejumbo/gong/values"
	"math/rand"

	"github.com/mikabrytu/gomes-engine/events"
)

var aiPallet *objects.Pallet

const INITIAL_SPEED int = 5

func InitAI(pallet *objects.Pallet) {
	aiPallet = pallet
	aiEvents()
	initPallet()
}

func aiEvents() {
	events.Subscribe(values.PALLET_MOVE_LIMIT_UP, func(params ...any) error {
		invertPallet(params...)
		return nil
	})

	events.Subscribe(values.PALLET_MOVE_LIMIT_DOWN, func(params ...any) error {
		invertPallet(params...)
		return nil
	})

	events.Subscribe(values.BALL_COLLISION_PALLET, func(params ...any) error {
		setRandomSpeed()
		return nil
	})

	events.Subscribe(values.GAME_RESET, func(params ...any) error {
		initPallet()
		return nil
	})
}

func invertPallet(params ...any) {
	p := params[0].([]any)[0].([]any)[0].(*objects.Pallet)
	if p == aiPallet {
		aiPallet.InvertDirection()
	}
}

func setRandomSpeed() {
	s := []int{5, 10, 15}[rand.Intn(3)]

	aiPallet.SetSpeed(s)
	aiPallet.InvertDirection()
}

func initPallet() {
	d := []int{-1, 1}[rand.Intn(2)]
	aiPallet.SetSpeed(INITIAL_SPEED)
	aiPallet.SetDirection(d)
	aiPallet.Move()
}
