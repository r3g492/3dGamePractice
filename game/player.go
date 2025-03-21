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

func PlayerSpeedUp(dt float32) {
	player.Speed += 5 * dt
}

func PlayerSpeedDown(dt float32) {
	player.Speed -= 90 * dt
	if player.Speed <= 0 {
		player.Speed = 0
	}
}

func PlayerRotateLeft(dt float32) {
	player.Cube.Angle += 90 * dt
}

func PlayerRotateRight(dt float32) {
	player.Cube.Angle -= 90 * dt
}
