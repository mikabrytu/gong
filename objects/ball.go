package objects

import (
	"littlejumbo/gong/values"

	"github.com/mikabrytu/gomes-engine/events"
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
	screen    math.Vector2
}

const INITIAL_SPEED int = 10

func NewBall(rect utils.RectSpecs, color render.Color) *Ball {
	ball := &Ball{
		rect:  rect,
		color: color,
		direction: math.Vector2{
			X: -1,
			Y: 0,
		},
		speed: INITIAL_SPEED,
	}

	lifecycle.Register(lifecycle.GameObject{
		Start:   ball.start,
		Physics: ball.physics,
		Render:  ball.render,
	})

	return ball
}

func (b *Ball) Reset() {
	b.rect.PosX = (b.screen.X / 2) - (b.rect.Width / 2)
	b.rect.PosY = (b.screen.Y / 2) - (b.rect.Height / 2)
	b.direction.X *= -1
	b.direction.Y = 0
	b.speed = INITIAL_SPEED
}

func (b *Ball) SetSpeed(speed int) {
	b.speed = speed
}

func (b *Ball) SetScreenSize(size math.Vector2) {
	b.screen = size
}

func (b *Ball) start() {
	physics.RegisterBody(&b.rect, values.OBJECT_BALL)
}

func (b *Ball) physics() {
	checkCollision(b)
	checkScreenBoundaries(b)

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

func checkCollision(b *Ball) {
	body := physics.GetBodyByName(values.OBJECT_BALL)
	if body.Name != "nil" {
		collider := physics.CheckCollision(body)
		if collider.Name != "nil" {
			events.Emit(values.BALL_COLLISION_PALLET)

			if collider.Rect.PosX < body.Rect.PosX {
				b.direction.X = 1
			}

			if collider.Rect.PosX > body.Rect.PosX {
				b.direction.X = -1
			}

			bMidY := body.Rect.PosY + (body.Rect.Height / 2)
			cMidY := collider.Rect.PosY + (collider.Rect.Height / 2)

			if cMidY == bMidY {
				b.direction.Y = 0
				b.speed = INITIAL_SPEED * 2
			}

			if cMidY < bMidY {
				b.direction.Y = 1
			}

			if cMidY > bMidY {
				b.direction.Y = -1
			}
		}
	}
}

func checkScreenBoundaries(b *Ball) {
	if b.rect.PosY < 0 {
		b.direction.Y = 1
	}

	if (b.rect.PosY + b.rect.Height) > b.screen.Y {
		b.direction.Y = -1
	}

	if (b.rect.PosX + b.rect.Width) < 0 {
		events.Emit(values.SCORE_COMPUTER)
	}

	if b.rect.PosX > values.SCREEN_SIZE.X {
		events.Emit(values.SCORE_PLAYER)
	}
}
