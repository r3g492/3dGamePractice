package graphic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	raylibWindowWidth  int32
	raylibWindowHeight int32
	raylibTargetFrame  int32
)

func RaylibSet(
	windowWidth int,
	windowHeight int,
	targetFrame int,
) {
	raylibWindowWidth = int32(windowWidth)
	raylibWindowHeight = int32(windowHeight)
	raylibTargetFrame = int32(targetFrame)

	rl.InitWindow(raylibWindowWidth, raylibWindowHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(raylibTargetFrame)
}

func RaylibClose() {
	rl.CloseWindow()
}

func RaylibLoop(
	gameLogic func(dt float32),
) {
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		gameLogic(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}
}
