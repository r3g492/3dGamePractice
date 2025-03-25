package graphic

import (
	"3dGamePractice/game"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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
	const zoomSpeed float32 = 2.0

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

		if dragging {
			DrawDragRectangle(dragStartPoint, dragEndPoint, rl.Green)
		}

		selectedCubes := make(map[int32]bool)

		if dragging {
			for id := range selectedCubes {
				delete(selectedCubes, id)
			}

			for id, obj := range *objects {
				cubeWorldPosition := rl.NewVector3(obj.UnitCube().GamePosX, 0, obj.UnitCube().GamePosY)
				screenPosition := rl.GetWorldToScreen(cubeWorldPosition, camera)

				if RectangleContainsPoint(dragStartPoint, dragEndPoint, screenPosition) {
					selectedCubes[id] = true
				}
			}
		}

		rl.BeginMode3D(camera)

		// extra
		rl.DrawGrid(1000, 1.0)

		// do movement inputs
		if rl.IsKeyDown(rl.KeyW) {
			game.PlayerMoveUp(dt)
		}
		if rl.IsKeyDown(rl.KeyS) {
			game.PlayerMoveDown(dt)
		}
		if rl.IsKeyDown(rl.KeyA) {
			game.PlayerMoveLeft(dt)
		}
		if rl.IsKeyDown(rl.KeyD) {
			game.PlayerMoveRight(dt)
		}
		if !rl.IsKeyDown(rl.KeyW) &&
			!rl.IsKeyDown(rl.KeyS) &&
			!rl.IsKeyDown(rl.KeyA) &&
			!rl.IsKeyDown(rl.KeyD) {
			game.PlayerMoveStop(dt)
		}
		game.PlayerMoveCalc(dt)
		game.PlayerMoveClear(dt)

		wheel := rl.GetMouseWheelMove()
		if wheel != 0 {
			camera.Position = rl.Vector3Add(camera.Position, rl.Vector3Scale(rl.Vector3Normalize(rl.Vector3Subtract(camera.Target, camera.Position)), wheel*zoomSpeed))
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			dragging = false
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

			if selectedCubes[obj.Id()] == true {
				cubeColor = rl.Red
			}
			if RayHitsCube(mouseRay, obj.UnitCube()) {
				cubeColor = rl.Green
			}
			DrawCube(obj.UnitCube(), cubeColor)

		}

		rl.EndMode3D()

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

func RectangleContainsPoint(start, end, point rl.Vector2) bool {
	minX := float32(math.Min(float64(start.X), float64(end.X)))
	maxX := float32(math.Max(float64(start.X), float64(end.X)))
	minY := float32(math.Min(float64(start.Y), float64(end.Y)))
	maxY := float32(math.Max(float64(start.Y), float64(end.Y)))

	return point.X >= minX && point.X <= maxX && point.Y >= minY && point.Y <= maxY
}
