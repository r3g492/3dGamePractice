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

type Cube struct {
	GamePosX float32
	GamePosY float32
	Angle    float32
	Width    float32
	Height   float32
	Length   float32
}

type Object interface {
	UnitCube() Cube
	Update(dt float32)
	Id() int32
}
