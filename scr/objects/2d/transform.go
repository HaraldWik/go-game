package obj2d

import vec2 "github.com/HaraldWik/go-game/scr/vector/2"

type Transform struct {
	pos, size vec2.Vector
	rot       float32
}

func NewTransform() Transform {
	return Transform{}
}

func NewTransformAdvanced(position, size vec2.Vector, rotation float32) Transform {
	return Transform{
		pos:  position,
		size: size,
		rot:  rotation,
	}
}

// *Set
func (transform *Transform) SetPos(position vec2.Vector) {
	transform.pos = position
}

func (transform *Transform) SetSize(size vec2.Vector) {
	transform.size = size
}

func (transform *Transform) SetRot(rotation float32) {
	transform.rot = rotation
}

// *Get
func (transform *Transform) GetPos() vec2.Vector {
	return transform.pos
}

func (transform *Transform) GetSize() vec2.Vector {
	return transform.size
}

func (transform *Transform) GetRot() float32 {
	return transform.rot
}
