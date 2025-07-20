// Package main implements an automated Snake game where pathfinding algorithms 
// control snake movement to find food.
package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

// ------------------------------------------------------------------------------------
// Global Constants Declaration
// ------------------------------------------------------------------------------------

const screenWidth = 800
const screenHeight = 800

const cellWidth = 25
const cellHeight = 25

const columns = screenWidth / cellWidth
const rows = screenHeight / cellHeight

const delayTime = float32(0.01)

// ----------------------------------------------------------------------------------
// Types and Structures Definition
// ----------------------------------------------------------------------------------

type snake struct {
    segments []rl.Vector2
    size     rl.Vector2
    color    rl.Color
}

type food struct {
    position rl.Vector2
    size     rl.Vector2
    color    rl.Color
}
type UIText struct{
	text 	 string
	fontSize int32
	yOffset  int32
	color    rl.Color
}

type coordinates struct {
    X, Y int
}

// ------------------------------------------------------------------------------------
// Global Variables Declaration
// ------------------------------------------------------------------------------------

var score    uint16
var pause    bool
var gameOver bool 

var moveTime       float32
var instructionSet []int

var prevHeadPos rl.Vector2

var snake_1 snake
var fruit food


// ------------------------------------------------------------------------------------
// Entry Point
// ------------------------------------------------------------------------------------

func init() {
	score = 0
	pause = false
	gameOver = false
    moveTime = float32(0.0)

    snake_1 = snake{
        segments: []rl.Vector2{{X: 0, Y: 0}},
        size:     rl.Vector2{X: cellWidth, Y: cellHeight},
        color:    rl.Blue,
    }

    fruit = food{
		position: foodSpawn(),
		size:     rl.Vector2{X: cellWidth, Y: cellHeight},
		color:    rl.Red,
	}
}

func main() {
    rl.InitWindow(screenWidth, screenHeight, "Snake Game")
    defer rl.CloseWindow()
    rl.SetTargetFPS(60)

    // Main Loop
    for !rl.WindowShouldClose() {
        updateAndDraw()
    }
}