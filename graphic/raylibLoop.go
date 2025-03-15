package graphic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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
		Position:   rl.NewVector3(10.0, 10.0, 0.0),
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

	var playerAngle float32 = 0
	var playerSpeed float32 = 0
	playerPosition := rl.Vector3{X: 0, Y: 0, Z: 0}
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		gameLogic(dt)
		// rl.UpdateCamera(&camera, rl.CameraOrbital)

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

		if rl.IsKeyDown(rl.KeyW) {
			playerSpeed += 5 * dt
		}

		if rl.IsKeyDown(rl.KeyS) {
			playerSpeed -= 90 * dt
			if playerSpeed <= 0 {
				playerSpeed = 0
			}
		}

		if rl.IsKeyDown(rl.KeyA) {
			playerAngle += 90 * dt
		}

		if rl.IsKeyDown(rl.KeyD) {
			playerAngle -= 90 * dt
		}

		rad := float64(playerAngle) * (math.Pi / 180.0)
		forwardDir := rl.Vector3{
			X: float32(math.Sin(rad)),
			Z: float32(math.Cos(rad)),
		}

		playerPosition.X += forwardDir.X * playerSpeed * dt
		playerPosition.Z += forwardDir.Z * playerSpeed * dt

		rl.DrawGrid(10, 1.0)

		rl.PushMatrix()
		rl.Translatef(playerPosition.X, playerPosition.Y, playerPosition.Z)
		rl.Rotatef(playerAngle, 0, 1, 0)
		rl.DrawCubeWires(
			rl.Vector3{
				0,
				0,
				0,
			},
			5,
			5,
			10,
			rl.Black,
		)
		rl.PopMatrix()
		rl.EndMode3D()

		rl.DrawText("Congrats! You created your first 3D box!", 10, 10, 20, rl.DarkGray)
		rl.EndDrawing()
	}
}
