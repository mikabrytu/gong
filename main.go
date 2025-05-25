package main

import (
	"littlejumbo/gong/managers"
	"littlejumbo/gong/values"

	gomesengine "github.com/mikabrytu/gomes-engine"
)

func main() {
	gomesengine.Init(
		values.GAME_TITLE,
		int32(values.SCREEN_SIZE.X),
		int32(values.SCREEN_SIZE.Y),
	)

	managers.Game()

	gomesengine.Run()
}
