package game

type Enemy struct {
	id    int32
	Speed float32
	Cube  Cube
}

func (e *Enemy) UnitCube() Cube {
	return e.Cube
}

func (e *Enemy) Update(dt float32) {
}

func (e *Enemy) Id() int32 {
	return e.id
}
