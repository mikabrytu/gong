package managers

import (
	"littlejumbo/gong/objects"
	"math/rand"
)

var aiPallet *objects.Pallet

func InitAI(pallet *objects.Pallet) {
	d := rand.Intn(3) - 1
	println(d)

	aiPallet = pallet
	aiPallet.SetSpeed(5)
	aiPallet.SetDirection(-1)
	aiPallet.Move()
}
