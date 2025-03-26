package game

type Minion struct {
	id    int32
	Speed float32
	Cube  Cube
}

func (m *Minion) UnitCube() Cube {
	return m.Cube
}

func (m *Minion) Update(dt float32) {
}

func (m *Minion) Id() int32 {
	return m.id
}
