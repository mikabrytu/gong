package managers

import (
	"littlejumbo/gong/objects"
	"littlejumbo/gong/values"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

func Game() {
	setupPallets()
	setupBall()
	setupUI()
}

func setupPallets() {
	pw := 50
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

	objects.NewPallet(pr1, render.White)
	objects.NewPallet(pr2, render.White)
}

func setupBall() {
	rect := utils.RectSpecs{
		PosX:   (values.SCREEN_SIZE.X / 2) - 25,
		PosY:   (values.SCREEN_SIZE.Y / 2) - 25,
		Width:  50,
		Height: 50,
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: func() {
			render.DrawSimpleShapes(rect, render.White)
		},
	})
}

func setupUI() {
	font := ui.FontSpecs{
		Name: "CutePixel",
		Path: "assets/font/CutePixel.ttf",
		Size: 48,
	}

	s1 := ui.NewFont(font, values.SCREEN_SIZE)
	s2 := ui.NewFont(font, values.SCREEN_SIZE)

	s1.RenderText("0", render.White, math.Vector2{})
	s2.RenderText("0", render.White, math.Vector2{})

	offset := math.Vector2{X: 25, Y: 10}
	s1.AlignText(ui.TopLeft, offset)
	s2.AlignText(ui.TopRight, offset)
}
