package main

import (
	"math/rand"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO: Make this modular
func foodSpawn() rl.Vector2 {
	for {
		randRow := rand.Intn(rows)
		randCol := rand.Intn(columns)

		foodPos := rl.Vector2{
			X: float32(randCol * cellWidth),
			Y: float32(randRow * cellHeight),
		}

		occupied := false
		for _, segment := range snake_1.segments {
			if segment.X == foodPos.X && segment.Y == foodPos.Y {
				occupied = true
				break
			}
		}

		if !occupied {
			return foodPos
		}
	}
}
