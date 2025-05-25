package objects

import (
	"littlejumbo/gong/values"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Ball struct {
	rect      utils.RectSpecs
	color     render.Color
	direction math.Vector2
	speed     int
}

func NewBall(rect utils.RectSpecs, color render.Color) *Ball {
	ball := &Ball{
		rect:  rect,
		color: color,
		direction: math.Vector2{
			X: -1,
			Y: 0,
		},
		speed: 10,
	}

	lifecycle.Register(lifecycle.GameObject{
		Start:   ball.start,
		Physics: ball.physics,
		Render:  ball.render,
	})

	return ball
}

func (b *Ball) start() {
	physics.RegisterBody(&b.rect, values.OBJECT_BALL)
}

func (b *Ball) physics() {
	body := physics.GetBodyByName(values.OBJECT_BALL)
	if body.Name != "nil" {
		collider := physics.CheckCollision(body)
		if collider.Name != "nil" {
			if collider.Rect.PosX < body.Rect.PosX {
				b.direction.X = 1
			}

			if collider.Rect.PosX > body.Rect.PosX {
				b.direction.X = -1
			}
		}
	}

	pos := math.Vector2{
		X: b.rect.PosX + (b.speed * b.direction.X),
		Y: b.rect.PosY + (b.speed * b.direction.Y),
	}
	b.rect.PosX = pos.X
	b.rect.PosY = pos.Y
}

func (b *Ball) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}
