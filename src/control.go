package main

// make this modular enough
func control(instruction int){
	
	for i := len(snake_1.segments) -  1; i > 0; i-- {
		snake_1.segments[i] = snake_1.segments[i-1]
	}

	switch instruction {
		case 0:
			snake_1.segments[0].X += cellHeight
		case 2:
			snake_1.segments[0].X -= cellHeight
		case 1:
			snake_1.segments[0].Y += cellHeight
		case 3:
			snake_1.segments[0].Y -= cellHeight
	}
	moveTime = 0.0 // Reset to Zero

}