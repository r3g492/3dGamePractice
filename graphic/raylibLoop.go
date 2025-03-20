package graphic

import (
	"3dGamePractice/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	raylibWindowWidth  int32
	raylibWindowHeight int32
	raylibTargetFrame  int32
	camera             rl.Camera3D

	dragging       bool
	dragStartPoint rl.Vector2
	dragEndPoint   rl.Vector2
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
	var player = game.GetPlayer()
	var objects *map[int32]game.Object = game.GetObjectMap()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// do game logics
		gameLogic(dt)
		player.Update(dt)

		// update camera by player
		camera.Position.X = player.Cube.GamePosX - 10
		camera.Position.Z = player.Cube.GamePosY
		camera.Target.X = player.Cube.GamePosX
		camera.Target.Z = player.Cube.GamePosY

		// Mouse picking logic
		mouseRay := rl.GetScreenToWorldRay(rl.GetMousePosition(), camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)

		// extra
		rl.DrawGrid(10, 1.0)
		rl.DrawText(
			"Congrats! You created your first 3D box!",
			10,
			10,
			20,
			rl.DarkGray,
		)

		// do inputs
		if rl.IsKeyDown(rl.KeyW) {
			game.PlayerSpeedUp(dt)
		}
		if rl.IsKeyDown(rl.KeyS) {
			game.PlayerSpeedDown(dt)
		}
		if rl.IsKeyDown(rl.KeyA) {
			game.PlayerRotateLeft(dt)
		}
		if rl.IsKeyDown(rl.KeyD) {
			game.PlayerRotateRight(dt)
		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			dragging = true
			dragStartPoint = rl.GetMousePosition()
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			dragging = false
		}

		if dragging {
			dragEndPoint = rl.GetMousePosition()
		}

		// draw cubes
		DrawCube(player.Cube, rl.Red)
		for _, obj := range *objects {
			cubeColor := rl.Blue
			if RayHitsCube(mouseRay, obj.UnitCube()) {
				cubeColor = rl.Red
			}
			DrawCube(obj.UnitCube(), cubeColor)

		}

		rl.EndMode3D()

		if dragging {
			DrawDragRectangle(dragStartPoint, dragEndPoint, rl.Green)
		}

		rl.EndDrawing()
	}
}

func DrawCube(
	cube game.Cube,
	color rl.Color,
) {
	rl.PushMatrix()
	rl.Translatef(cube.GamePosX, 0, cube.GamePosY)
	rl.Rotatef(cube.Angle, 0, 1, 0)
	rl.DrawCubeWires(
		rl.Vector3{},
		cube.Width,
		cube.Height,
		cube.Length,
		color,
	)
	rl.PopMatrix()
}

func RayHitsCube(ray rl.Ray, cube game.Cube) bool {
	bounds := rl.NewBoundingBox(
		rl.NewVector3(
			cube.GamePosX-cube.Width/2,
			-cube.Height/2,
			cube.GamePosY-cube.Length/2,
		),
		rl.NewVector3(
			cube.GamePosX+cube.Width/2,
			cube.Height/2,
			cube.GamePosY+cube.Length/2,
		),
	)
	return rl.GetRayCollisionBox(ray, bounds).Hit
}

func DrawDragRectangle(start, end rl.Vector2, color rl.Color) {
	x := start.X
	y := start.Y
	width := end.X - start.X
	height := end.Y - start.Y

	if width < 0 {
		x += width
		width *= -1
	}

	if height < 0 {
		y += height
		height *= -1
	}

	rl.DrawRectangleLines(int32(x), int32(y), int32(width), int32(height), color)
}
