package game

var (
	allyMap  = make(map[int32]Object)
	enemyMap = make(map[int32]Object)
	player   PlayerShip
)

func GetAllyMap() *map[int32]Object {
	return &allyMap
}

func GetEnemyMap() *map[int32]Object {
	return &enemyMap
}

func GetPlayer() *PlayerShip {
	return &player
}

func Init() {
	player = PlayerShip{
		id: 0,
		Cube: Cube{
			0,
			0,
			0,
			1,
			1,
			2,
		},
	}

	allyMap[0] = &Minion{
		id: 0,
		Cube: Cube{
			GamePosX: 0,
			GamePosY: 10,
			Angle:    30,
			Width:    1,
			Height:   1,
			Length:   1,
		},
	}

	// Example objects
	enemyMap[0] = &Enemy{
		id: 0,
		Cube: Cube{
			GamePosX: 3,
			GamePosY: 2,
			Angle:    45,
			Width:    1,
			Height:   1,
			Length:   1,
		},
	}

	enemyMap[1] = &Enemy{
		id: 1,
		Cube: Cube{
			GamePosX: -4,
			GamePosY: -3,
			Angle:    90,
			Width:    2,
			Height:   1,
			Length:   2,
		},
	}

	enemyMap[2] = &Enemy{
		id: 2,
		Cube: Cube{
			GamePosX: 5,
			GamePosY: -1,
			Angle:    30,
			Width:    1.5,
			Height:   2,
			Length:   1,
		},
	}
}

func Logic(dt float32) {

}
