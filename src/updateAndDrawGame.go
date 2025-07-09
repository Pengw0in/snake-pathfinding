package main

import (
    algo "github.com/Pengw0in/prc1/src/algorithms"
    rl "github.com/gen2brain/raylib-go/raylib"
)

// ------------------------------------------------------------------------------------
// Game Update Logic
// ------------------------------------------------------------------------------------

func updateGame() {
    if rl.IsKeyPressed('P') {
        pause = false
    }

    if rl.IsKeyPressed('R') && gameOver {
        snake_1.segments = []rl.Vector2{{X: 0, Y: 0}} // reset segments

        score = 0
        instructionSet =[]int{}
        moveTime = 0.0
        fruit.position = foodSpawn()
        gameOver = false
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

                instructionSet = algo.Bfs(snakeGrid[0][1], snakeGrid[0][0], fruitGrid.Y, fruitGrid.X, rows, columns, snakeGrid)
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
            if snake_1.segments[0] == fruit.position {
                fruit.position = foodSpawn()

                score = score + 1
                snake_1.segments = append(snake_1.segments, prevHeadPos) // Add new segment

                instructionSet = []int{} // Clear instructions to recalculate path
            }
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
    for i := range rows {      // i ==> --
        for j := range columns { // j ==> |
            rl.DrawRectangleLines(
                int32(j*cellWidth), int32(i*cellHeight),
                int32(cellWidth), int32(cellHeight),
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

    // Draw snake
    for i := range snake_1.segments {
        rl.DrawRectangleV(snake_1.segments[i], snake_1.size, snake_1.color)
    }

    // Draw fruit
    rl.DrawRectangleV(fruit.position, fruit.size, fruit.color)

    // UI elements
    if pause && !gameOver{
        rl.DrawText("GAME PAUSED", screenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, screenHeight/2-40, 40, rl.Gray)
    }

    if gameOver {
        rl.DrawText("GAME OVER", screenWidth/2-rl.MeasureText("GAME OVER", 40)/2, screenHeight/2-40, 40, rl.Gray)
        rl.DrawText("ALGORITHM EXHAUSTED", screenWidth/2-rl.MeasureText("ALGORITHM EXHAUSTED", 20)/2, screenHeight/2+25, 20, rl.Red)
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