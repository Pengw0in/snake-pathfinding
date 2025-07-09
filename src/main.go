/*
Author: Pengw0in
Date  : 07-07-2025
*/

package main

import (
    // "strconv"
    rl "github.com/gen2brain/raylib-go/raylib"
)

// ------------------------------------------------------------------------------------
// Global Constants Declaration
// ------------------------------------------------------------------------------------

const screenWidth = 1600
const screenHeight = 900

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

type coordinates struct {
    X, Y int
}

// ------------------------------------------------------------------------------------
// Global Variables Declaration
// ------------------------------------------------------------------------------------

var moveTime = float32(0.0)
var instructionSet []int
var score uint16

var snake_1 = snake{
    []rl.Vector2{{X: 0, Y: 0}},
    rl.Vector2{X: cellWidth, Y: cellHeight},
    rl.Blue,
}

var prevHeadPos rl.Vector2

var fruit = food{
    foodSpawn(),
    rl.Vector2{X: cellWidth, Y: cellHeight},
    rl.SkyBlue,
}

var pause bool = false
var gameOver bool = false

// ------------------------------------------------------------------------------------
// Entry Point
// ------------------------------------------------------------------------------------

func main() {
    rl.InitWindow(screenWidth, screenHeight, "BFS operated Snake")
    defer rl.CloseWindow()
    rl.SetTargetFPS(60)

    // Main Loop
    for !rl.WindowShouldClose() {
        updateAndDraw()
    }
}