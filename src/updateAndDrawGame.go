package main

// TODO: ADD GAME OVER SCREEN

import (
	algo "github.com/Pengw0in/prc1/src/algorithms"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateGame() {
	if rl.IsKeyPressed('P') {
		pause = !pause
	}

	if !pause {
		if len(instructions) == 0 {
			// Pixel coordinates to Grid coordinates
			snakeGridX := int(snake_1.segments[0].X / cellWidth)
			snakeGridY := int(snake_1.segments[0].Y / cellHeight)
			foodGridX := int(fruit.position.X / cellWidth)
			foodGridY := int(fruit.position.Y / cellHeight)

			instructions = algo.Bfs(snakeGridY, snakeGridX, foodGridY, foodGridX, rows, columns, snake_1.segments, cellHeight)
		}

		deltaTime := rl.GetFrameTime() // Time since last Frame
		moveTime += deltaTime
		prevHeadPos = snake_1.segments[0] // Can this be improved

		// Process one instruction per one loop
		if len(instructions) > 0 && moveTime >= delayTime {
			control(instructions[0])
			instructions = instructions[1:]
		}

		// Collision logic
		if snake_1.segments[0] == fruit.position {
			fruit.position = foodSpawn() 

			score = score + 1
			snake_1.segments = append(snake_1.segments, prevHeadPos) // Add new segment

			instructions = []int{} // Clear instructions to recalculate path
		}
	}
}

func drawGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	for i := range rows { // i ==> --
		for j := range columns { // j ==> |
			rl.DrawRectangleLines(
				int32(j*cellWidth),
				int32(i*cellHeight),
				int32(cellWidth),
				int32(cellHeight),
				rl.LightGray,
			)

			// cellCoordinates(
			// 	strconv.Itoa(j),
			// 	strconv.Itoa(i),
			// 	int32(j*cellWidth),
			// 	int32(i*cellHeight),
			// )
		}
	}
	for i := range snake_1.segments {
		rl.DrawRectangleV(snake_1.segments[i], snake_1.size, snake_1.color)
	}
	rl.DrawRectangleV(fruit.position, fruit.size, fruit.color)

	if pause {
		rl.DrawText("GAME PAUSED", screenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, screenHeight/2-40, 40, rl.Gray)
	}

	rl.EndDrawing()
}

func updateAndDraw() {
	drawGame()
	updateGame()
}
