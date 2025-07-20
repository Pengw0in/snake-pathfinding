package main

import (
    algo "github.com/Pengw0in/prc1/src/algorithms"
    rl "github.com/gen2brain/raylib-go/raylib"
)

// ------------------------------------------------------------------------------------
// Main Game Logic
// ------------------------------------------------------------------------------------

func (ui UIText) draw() {
	textWidth := rl.MeasureText(ui.text, ui.fontSize)
	x := screenWidth/2 - textWidth/2
	y := screenHeight/2 - ui.yOffset
	rl.DrawText(ui.text, x, y, ui.fontSize, ui.color)
}

func drawGrid() {
    for row := range rows {      
        for column := range columns {
            rl.DrawRectangleLines(
                int32(column * cellWidth), int32(row * cellHeight),
                int32(cellWidth), int32(cellHeight),
                rl.LightGray,
            )
        }
    }
}

func drawSnake() {
    for i := range snake_1.segments {
        if i == 0 {
            rl.DrawRectangleV(snake_1.segments[i], snake_1.size, rl.DarkBlue)
        } else {
            rl.DrawRectangleV(snake_1.segments[i], snake_1.size, snake_1.color)
        }
    }

}

func reset() {
    snake_1.segments = []rl.Vector2{{X: 0, Y: 0}} // reset segments
    score = 0
    instructionSet =[]int{}
    moveTime = 0.0
    fruit.position = foodSpawn()
    gameOver = false   
}

func collisionCheck() {
    if snake_1.segments[0] == fruit.position {
        fruit.position = foodSpawn()

        score = score + 1
        snake_1.segments = append(snake_1.segments, prevHeadPos) // Add new segment

        instructionSet = []int{}
    }
}

func updateGame() {
    if rl.IsKeyPressed('P') {
        pause = !pause
    }

    if rl.IsKeyPressed('R') && gameOver {
        reset()
    }

    if !gameOver {
        if !pause {
            if len(instructionSet) == 0 {
                // Pixel coordinates to Grid coordinates
                var snakeGrid [][]int
                for _, segments := range snake_1.segments {
                    gridCoord := []int{
                        int(segments.X / cellWidth),
                        int(segments.Y / cellHeight),
                    }
                    snakeGrid = append(snakeGrid, gridCoord)
                }
                fruitGrid := coordinates{
                    int(fruit.position.X / cellWidth),
                    int(fruit.position.Y / cellHeight),
                }

                instructionSet = algo.Bfs(
                    snakeGrid[0][1], snakeGrid[0][0], 
                    fruitGrid.Y, fruitGrid.X, 
                    rows, columns, 
                    snakeGrid,
                )
            }

            deltaTime := rl.GetFrameTime() // Time since last Frame
            moveTime += deltaTime

            prevHeadPos = snake_1.segments[0] // Can this be improved?

            // Process one instruction per one loop
            if len(instructionSet) > 0 && instructionSet[0] != -1 {
                if moveTime >= delayTime {
                    control(instructionSet[0])
                    instructionSet = instructionSet[1:]
                }
            } else {
                gameOver = true
            }

            // Collision logic
            collisionCheck()
        }
    }
}

// ------------------------------------------------------------------------------------
// Game Rendering Logic
// ------------------------------------------------------------------------------------

func drawGame() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.RayWhite)

    // Draw grid
    drawGrid()

    // Draw snake
    drawSnake()

    // Draw fruit
    rl.DrawRectangleV(fruit.position, fruit.size, fruit.color)

    // UI elements
    if pause && !gameOver{
        pauseText := UIText{"GAME PAUSED", 40, 60, rl.Gray}
        pauseText.draw()
        
        // Add instruction text below
        instructionText := UIText{"Press P to Resume", 20, 20, rl.LightGray}
        instructionText.draw()
    }

    if gameOver {
		gameOverTexts := []UIText{
            {"GAME OVER", 50, 80, rl.Red},
            {"Algorithm could not find path", 24, 30, rl.Maroon},
            {"Press R to Restart", 18, -30, rl.Gray},
		}

		for _, text := range gameOverTexts{
			text.draw()
		}
    }

    rl.EndDrawing()
}

// ------------------------------------------------------------------------------------
// Main Update and Draw Function
// ------------------------------------------------------------------------------------

func updateAndDraw() {
    drawGame()
    updateGame()
}