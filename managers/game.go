package managers

import (
	"littlejumbo/gong/objects"
	"littlejumbo/gong/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

var ball *objects.Ball

func Game() {
	pallets := setupPallets()
	scores := setupUI()
	setupBall()

	gameEvents()

	InitScore(scores[0], scores[1])
	InitPlayer(pallets[0])
	InitAI(pallets[1])
}

func Reset() {
	events.Emit(values.GAME_RESET)
	ball.Reset()
}

func setupPallets() []*objects.Pallet {
	pw := 40
	ph := 200
	off := 10
	rect := utils.RectSpecs{
		PosX:   0,
		PosY:   0,
		Width:  pw,
		Height: ph,
	}

	pr1 := rect
	pr1.PosX = off
	pr1.PosY = (values.SCREEN_SIZE.Y / 2) - (ph / 2)

	pr2 := rect
	pr2.PosX = values.SCREEN_SIZE.X - pw - off
	pr2.PosY = (values.SCREEN_SIZE.Y / 2) - (ph / 2)

	p1 := objects.NewPallet(values.OBJECT_PALLET_PLAYER, pr1, render.White)
	p2 := objects.NewPallet(values.OBJECT_PALLET_COMPUTER, pr2, render.White)

	return []*objects.Pallet{p1, p2}
}

func setupBall() {
	s := 30
	rect := utils.RectSpecs{
		PosX:   (values.SCREEN_SIZE.X / 2) - (s / 2),
		PosY:   (values.SCREEN_SIZE.Y / 2) - (s / 2),
		Width:  s,
		Height: s,
	}

	ball = objects.NewBall(rect, render.White)
	ball.SetScreenSize(values.SCREEN_SIZE)
}

func setupUI() []*ui.Font {
	font := ui.FontSpecs{
		Name: "CutePixel",
		Path: "assets/font/CutePixel.ttf",
		Size: 48,
	}

	s1 := ui.NewFont(font, values.SCREEN_SIZE)
	s2 := ui.NewFont(font, values.SCREEN_SIZE)

	s1.Init("0", render.White, math.Vector2{})
	s2.Init("0", render.White, math.Vector2{})

	offset := math.Vector2{X: values.SCREEN_SIZE.X / 4, Y: 10}
	s1.AlignText(ui.TopLeft, offset)
	s2.AlignText(ui.TopRight, offset)

	return []*ui.Font{s1, s2}
}

func gameEvents() {
	events.Subscribe(values.SCORE_PLAYER, func(params ...any) error {
		onPlayerScore()
		return nil
	})

	events.Subscribe(values.SCORE_COMPUTER, func(params ...any) error {
		onComputerScore()
		return nil
	})
}

func onPlayerScore() {
	AddScore(Human)
	Reset()
}

func onComputerScore() {
	AddScore(Computer)
	Reset()
}
