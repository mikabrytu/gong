package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Pallet struct {
	Rect      utils.RectSpecs
	color     render.Color
	direction int
	speed     int
	moving    bool
}

func NewPallet(rect utils.RectSpecs, color render.Color) *Pallet {
	pallet := &Pallet{
		Rect:      rect,
		color:     color,
		direction: 0,
		speed:     0,
	}

	lifecycle.Register(lifecycle.GameObject{
		Physics: pallet.physics,
		Render:  pallet.render,
	})

	return pallet
}

func (p *Pallet) Move() {
	p.moving = true
}

func (p *Pallet) Stop() {
	p.moving = false
}

func (p *Pallet) SetDirection(direction int) {
	p.direction = direction
}

func (p *Pallet) SetSpeed(speed int) {
	p.speed = speed
}

func (p *Pallet) physics() {
	if p.moving {
		p.Rect.PosY += p.speed * p.direction
	}
}

func (p *Pallet) render() {
	render.DrawSimpleShapes(p.Rect, p.color)
}
