package obj2d

import vec2 "github.com/HaraldWik/go-game/scr/vector/2"

type Transform struct {
	Pos, Size vec2.Vector
	Rot       float32
}

func NewTransform() Transform {
	return Transform{}
}

func NewTransformAdvanced(position, size vec2.Vector, rotation float32) Transform {
	return Transform{
		Pos:  position,
		Size: size,
		Rot:  rotation,
	}
}

// *Set
func (transform *Transform) SetPos(position vec2.Vector) {
	transform.Pos = position
}

func (transform *Transform) SetSize(size vec2.Vector) {
	transform.Size = size
}

func (transform *Transform) SetRot(rotation float32) {
	transform.Rot = rotation
}

// *Get
func (transform *Transform) GetPos() vec2.Vector {
	return transform.Pos
}

func (transform *Transform) GetSize() vec2.Vector {
	return transform.Size
}

func (transform *Transform) GetRot() float32 {
	return transform.Rot
}
