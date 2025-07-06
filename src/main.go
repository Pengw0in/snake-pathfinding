package main

import (
	"math/rand"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// -----CONSTANTS------
const screenWidth = 1600
const screenHeight = 900

const cellWidth = 50
const cellHeight = 50


const columns = screenWidth / cellWidth
const rows = screenHeight / cellHeight

const delayTime = float32(0.05)

// -----VARIABLES------
var moveTime = float32(0.0)
var instructions []int

var snake = rl.Vector2{X: 0, Y: 0}
var foodPosition = rl.Vector2{
	X: float32(rand.Intn(columns)*50), 
	Y: float32(rand.Intn(rows)*50),
}

var snakeSize = rl.Vector2{X: cellWidth, Y: cellHeight}
var foodSize = rl.Vector2{X: cellWidth, Y: cellWidth}


// Entry Point
func main() {

	rl.InitWindow(screenWidth, screenHeight, "BFS operated Snake")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Main Loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for i := range rows { 			// i ==> --
			for j := range columns {	// j ==> |
				rl.DrawRectangleLines(
					int32(j*cellWidth),
					int32(i*cellHeight),
					int32(cellWidth),
					int32(cellHeight),
					rl.LightGray,
				)

				cellCoordinates(
					strconv.Itoa(j), 
					strconv.Itoa(i),
					int32(j*cellWidth),
					int32(i*cellHeight),
				)
			}
		}

		rl.DrawRectangleV(snake, snakeSize, rl.Blue)
		rl.DrawRectangleV(foodPosition, foodSize, rl.Red)
		
		if len(instructions) == 0 {
		// Pixel coordinates to Grid coordinates
			snakeGridX := int(snake.X / cellWidth)
			snakeGridY := int(snake.Y / cellHeight)
			foodGridX := int(foodPosition.X / cellWidth)
			foodGridY := int(foodPosition.Y / cellHeight)

			instructions = bfs(snakeGridY, snakeGridX, foodGridY, foodGridX)
		}
		
		

		deltaTime := rl.GetFrameTime() // Time since last Frame
		moveTime += deltaTime

		// Process one instruction per one loop
		if len(instructions) > 0 && moveTime >= delayTime {
			control(instructions[0])
			instructions = instructions[1:]
		}


		// Collision logic
		snakeRect := rl.Rectangle{X: snake.X, Y: snake.Y, Width: snakeSize.X, Height: snakeSize.Y}
        foodRect := rl.Rectangle{X: foodPosition.X, Y: foodPosition.Y, Width: foodSize.X, Height: foodSize.Y}
        if rl.CheckCollisionRecs(snakeRect, foodRect){
            foodPosition.X = float32(rand.Intn(columns)) * cellWidth
            foodPosition.Y = float32(rand.Intn(rows)) * cellHeight
            instructions = []int{}  // Clear instructions to recalculate path
        }

		rl.EndDrawing()
	}
}
