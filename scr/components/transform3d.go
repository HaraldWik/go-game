package component

import vector3 "github.com/HaraldWik/go-game/scr/vector/3"

type Transform3D struct {
	Position, Size, Rotation vector3.Type
}

func NewTransform3D() Transform3D {
	return Transform3D{}
}

func NewTransform3DAdvanced(position, size, rotation vector3.Type) Transform3D {
	return Transform3D{
		Position: position,
		Size:     size,
		Rotation: rotation,
	}
}
