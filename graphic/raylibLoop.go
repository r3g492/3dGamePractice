package graphic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	raylibWindowWidth  int32
	raylibWindowHeight int32
	raylibTargetFrame  int32
	camera             rl.Camera3D
)

func RaylibSet(windowWidth, windowHeight, targetFrame int) {
	raylibWindowWidth = int32(windowWidth)
	raylibWindowHeight = int32(windowHeight)
	raylibTargetFrame = int32(targetFrame)

	rl.InitWindow(raylibWindowWidth, raylibWindowHeight, "raylib [core] example - 3D box")
	rl.SetTargetFPS(raylibTargetFrame)

	camera = rl.Camera3D{
		Position:   rl.NewVector3(4.0, 2.0, 4.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}
}

func RaylibClose() {
	rl.CloseWindow()
}

func RaylibLoop(gameLogic func(dt float32)) {
	shader := rl.LoadShader("resources/shaders/sun.vs", "resources/shaders/sun.fs")
	loc := rl.GetShaderLocation(shader, "lightPos")
	lightPos := rl.NewVector3(0, 0, 50)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		gameLogic(dt)
		rl.UpdateCamera(&camera, rl.CameraOrbital)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		rl.SetShaderValueV(
			shader,
			loc,
			[]float32{lightPos.X, lightPos.Y, lightPos.Z},
			rl.ShaderUniformVec3,
			1,
		)

		rl.BeginShaderMode(shader)
		sunColor := rl.NewColor(255, 223, 0, 255)
		rl.DrawSphere(rl.NewVector3(0, 0, 0), 1, sunColor)
		rl.EndShaderMode()

		rl.DrawGrid(10, 1.0)
		rl.EndMode3D()

		rl.DrawText("Congrats! You created your first 3D box!", 10, 10, 20, rl.DarkGray)
		rl.EndDrawing()
	}
}
