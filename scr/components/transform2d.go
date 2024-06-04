package component

import vector2 "github.com/HaraldWik/go-game/scr/vector/2"

type Transform2D struct {
	Position, Size vector2.Type
	Rotation       float32
}

func NewTransform2D() Transform2D {
	return Transform2D{}
}

func NewTransform2DAdvanced(position, size vector2.Type, rotation float32) Transform2D {
	return Transform2D{
		Position: position,
		Size:     size,
		Rotation: rotation,
	}
}
