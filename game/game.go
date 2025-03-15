package game

var ObjectMap = make(map[int32]Object)

func Logic(dt float32) {

}

type PlayerShip struct {
	id       int32
	gamePosX int32
	gamePosY int32
}

type Interceptor struct {
	id       int32
	gamePosX int32
	gamePosY int32
}

type Object interface {
	Shape()
	Head()
	Update(dt float32)
	Id() int32
}
