package main

// just a helper file

import (
	rl "github.com/gen2brain/raylib-go/raylib" )

func cellCoordinates(cellX string,cellY string, RowOffset int32, ColOffset int32) {
	rl.DrawText(cellX + "," + cellY, RowOffset, ColOffset, 10, rl.Gray)
}
