package managers

import (
	"littlejumbo/gong/objects"

	"github.com/mikabrytu/gomes-engine/events"
)

type Key int

const (
	W Key = iota
	S
)

var playerPallet *objects.Pallet

const SPEED = 5

func InitPlayer(p *objects.Pallet) {
	playerPallet = p
	playerPallet.SetSpeed(SPEED)

	playerEvents()
}

func playerEvents() {
	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_W, func(params ...any) error {
		onKeyPressed(W)
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_S, func(params ...any) error {
		onKeyPressed(S)
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_W, func(params ...any) error {
		onKeyReleased()
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_S, func(params ...any) error {
		onKeyReleased()
		return nil
	})
}

func onKeyPressed(key Key) {
	d := 0

	if key == W {
		d = -1
	}
	if key == S {
		d = 1
	}

	playerPallet.SetDirection(d)
	playerPallet.Move()
}

func onKeyReleased() {
	playerPallet.Stop()
}
