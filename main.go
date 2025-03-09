package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

var (
	windowWidth  int
	windowHeight int
)

func main() {

	fmt.Print("Enter window width: ")
	if _, err := fmt.Scan(&windowWidth); err != nil {
		fmt.Println("Invalid input for width:", err)
		os.Exit(1)
	}

	fmt.Print("Enter window height: ")
	if _, err := fmt.Scan(&windowHeight); err != nil {
		fmt.Println("Invalid input for height:", err)
		os.Exit(1)
	}

	rl.InitWindow(int32(windowWidth), int32(windowHeight), "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		gameLogic(dt)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}
}

func gameLogic(dt float32) {
}
