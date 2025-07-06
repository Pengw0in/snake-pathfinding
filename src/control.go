package main

func control(instruction int){
	switch instruction {
		case 0:
			snake.X += cellHeight
		case 2:
			snake.X -= cellHeight
		case 1:
			snake.Y += cellHeight
		case 3:
			snake.Y -= cellHeight
	}
	moveTime = 0.0 // Reset to Zero

}