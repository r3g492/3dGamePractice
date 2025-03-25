package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type PlayerShip struct {
	id    int32
	Speed float32
	Cube  Cube
}

func (ps *PlayerShip) UnitCube() Cube {
	return ps.Cube
}

func (ps *PlayerShip) Update(dt float32) {
	rad := float64(player.Cube.Angle) * (math.Pi / 180.0)
	forwardDir := rl.Vector3{
		X: float32(math.Sin(rad)),
		Z: float32(math.Cos(rad)),
	}
	player.Cube.GamePosX += forwardDir.X * player.Speed * dt
	player.Cube.GamePosY += forwardDir.Z * player.Speed * dt
}

func (ps *PlayerShip) Id() int32 {
	return ps.id
}

var (
	targetAngle        float32
	upButtonPressed    bool
	downButtonPressed  bool
	rightButtonPressed bool
	leftButtonPressed  bool
)

func PlayerMoveUp(dt float32) {
	upButtonPressed = true
}

func PlayerMoveDown(dt float32) {
	downButtonPressed = true
}

func PlayerMoveRight(dt float32) {
	rightButtonPressed = true
}

func PlayerMoveLeft(dt float32) {
	leftButtonPressed = true
}

func PlayerMoveCalc(dt float32) {
	x, y := 0, 0

	if upButtonPressed {
		y += 1
	}
	if downButtonPressed {
		y -= 1
	}
	if rightButtonPressed {
		x += 1
	}
	if leftButtonPressed {
		x -= 1
	}

	if x == 0 && y == 0 {
		return
	}

	switch {
	case x == 0 && y == 1:
		targetAngle = 90
	case x == 1 && y == 1:
		targetAngle = 45
	case x == 1 && y == 0:
		targetAngle = 0
	case x == 1 && y == -1:
		targetAngle = 315
	case x == 0 && y == -1:
		targetAngle = 270
	case x == -1 && y == -1:
		targetAngle = 225
	case x == -1 && y == 0:
		targetAngle = 180
	case x == -1 && y == 1:
		targetAngle = 135
	}

	angleDiff := shortestAngleDistance(player.Cube.Angle, targetAngle)

	rotationSpeed := float32(200)

	angleChange := rotationSpeed * dt
	if math.Abs(float64(angleDiff)) < float64(angleChange) {
		player.Cube.Angle = targetAngle
	} else {
		if angleDiff > 0 {
			player.Cube.Angle += angleChange
		} else {
			player.Cube.Angle -= angleChange
		}
	}

	player.Cube.Angle = normalizeAngle(player.Cube.Angle)

	if player.Speed < 10 {
		player.Speed += dt * 5
	}
}

func PlayerMoveClear(dt float32) {
	upButtonPressed = false
	downButtonPressed = false
	rightButtonPressed = false
	leftButtonPressed = false
}

func PlayerMoveStop(dt float32) {
	player.Speed = 0
}

func normalizeAngle(angle float32) float32 {
	for angle < 0 {
		angle += 360
	}
	for angle >= 360 {
		angle -= 360
	}
	return angle
}

func shortestAngleDistance(current, target float32) float32 {
	diff := normalizeAngle(target - current)
	if diff > 180 {
		diff -= 360
	}
	return diff
}
