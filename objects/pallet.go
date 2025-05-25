package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Pallet struct {
	rect  utils.RectSpecs
	color render.Color
}

func NewPallet(rect utils.RectSpecs, color render.Color) *Pallet {
	pallet := &Pallet{
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: pallet.Render,
	})

	return pallet
}

func (p *Pallet) Render() {
	render.DrawSimpleShapes(p.rect, p.color)
}
