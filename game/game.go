package game

var (
	objectMap = make(map[int32]Object)
	player    PlayerShip
)

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
}

func Logic(dt float32) {

}
