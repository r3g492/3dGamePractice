package game

type Object interface {
	UnitCube() Cube
	Update(dt float32)
	Id() int32
}
