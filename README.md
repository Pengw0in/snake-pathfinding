# Snake Game Operated by Algorithm
A Snake game where the snake automatically finds its path to food using Breadth-First Search algorithm.

## Project Details
- **Created:** July 2025
- **Language:** GoLang
- **Graphics Library:** Raylib-go

## Running the Project
```bash
# Run directly
go mod tidy
go run src/*.go

# Or build and run
go build -o snake-game ./src
./snake-game
```

## Building the Project
```bash
# Build for major platforms
./build.sh
```
## Controls
- `P` - Pause/Resume game
- `R` - Restart game (when game over)

## Dependencies
- [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for raylib