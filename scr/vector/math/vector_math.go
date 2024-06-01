package Vec

import (
	"github.com/HaraldWik/go-game/scr/types"
	Vec2 "github.com/HaraldWik/go-game/scr/vector/2"
	Vec3 "github.com/HaraldWik/go-game/scr/vector/3"
	Vec4 "github.com/HaraldWik/go-game/scr/vector/4"
)

func ToVec2(vector interface{}) Vec2.Vector {
	switch vec := vector.(type) {
	case Vec2.Vector:
		return vec
	case Vec3.Vector:
		return Vec2.New(vector.(Vec3.Vector).X, vector.(Vec3.Vector).Y)
	case Vec4.Vector:
		return Vec2.New(vector.(Vec4.Vector).X, vector.(Vec4.Vector).Y)
	default:
		return Vec2.Zero()
	}
}

func ToVec3(vector interface{}) Vec3.Vector {
	switch vec := vector.(type) {
	case Vec2.Vector:
		return Vec3.New(vector.(Vec2.Vector).X, vector.(Vec2.Vector).Y, 0.0)
	case Vec3.Vector:
		return vec
	case Vec4.Vector:
		return Vec3.New(vector.(Vec4.Vector).X, vector.(Vec4.Vector).Y, vector.(Vec4.Vector).Z)
	default:
		return Vec3.Zero()
	}
}

func ToVec4(vector interface{}) Vec4.Vector {
	switch vec := vector.(type) {
	case Vec2.Vector:
		return Vec4.New(vector.(Vec2.Vector).X, vector.(Vec2.Vector).Y, 0.0, 0.0)
	case Vec3.Vector:
		return Vec4.New(vector.(Vec3.Vector).X, vector.(Vec3.Vector).Y, vector.(Vec3.Vector).Z, 0.0)
	case Vec4.Vector:
		return vec
	default:
		return Vec4.Zero()
	}
}

func Add(vector1 interface{}, vector2 interface{}) interface{} {
	if types.AreSameType(true, vector1, vector2) {
		switch types.GetType(vector1) {
		case Vec2.Vector{}:
			return vector1.(Vec2.Vector).Add(vector2.(Vec2.Vector))
		case Vec3.Vector{}:
			return vector1.(Vec3.Vector).Add(vector2.(Vec3.Vector))
		case Vec4.Vector{}:
			return vector1.(Vec4.Vector).Add(vector2.(Vec4.Vector))
		default:
			return nil
		}
	} else {
		return nil
	}
}

func VecSub(vector1 interface{}, vector2 interface{}) interface{} {
	if types.AreSameType(true, vector1, vector2) {
		switch types.GetType(vector1) {
		case Vec2.Vector{}:
			return vector1.(Vec2.Vector).Sub(vector2.(Vec2.Vector))
		case Vec3.Vector{}:
			return vector1.(Vec3.Vector).Sub(vector2.(Vec3.Vector))
		case Vec4.Vector{}:
			return vector1.(Vec4.Vector).Sub(vector2.(Vec4.Vector))
		default:
			return nil
		}
	} else {
		return nil
	}
}

func VecMul(vector1 interface{}, vector2 interface{}) interface{} {
	if types.AreSameType(true, vector1, vector2) {
		switch types.GetType(vector1) {
		case Vec2.Vector{}:
			return vector1.(Vec2.Vector).Mul(vector2.(Vec2.Vector))
		case Vec3.Vector{}:
			return vector1.(Vec3.Vector).Mul(vector2.(Vec3.Vector))
		case Vec4.Vector{}:
			return vector1.(Vec4.Vector).Mul(vector2.(Vec4.Vector))
		default:
			return nil
		}
	} else {
		return nil
	}
}

func VecDiv(vector1 interface{}, vector2 interface{}) interface{} {
	if types.AreSameType(true, vector1, vector2) {
		switch types.GetType(vector1) {
		case Vec2.Vector{}:
			return vector1.(Vec2.Vector).Div(vector2.(Vec2.Vector))
		case Vec3.Vector{}:
			return vector1.(Vec3.Vector).Div(vector2.(Vec3.Vector))
		case Vec4.Vector{}:
			return vector1.(Vec4.Vector).Div(vector2.(Vec4.Vector))
		default:
			return nil
		}
	} else {
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////

func ABS(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case Vec2.Vector{}:
		return vector.(Vec2.Vector).ABS()
	case Vec3.Vector{}:
		return vector.(Vec3.Vector).ABS()
	case Vec4.Vector{}:
		return vector.(Vec4.Vector).ABS()
	default:
		return nil
	}
}

func Neg(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case Vec2.Vector{}:
		return vector.(Vec2.Vector).Neg()
	case Vec3.Vector{}:
		return vector.(Vec3.Vector).Neg()
	case Vec4.Vector{}:
		return vector.(Vec4.Vector).Neg()
	default:
		return nil
	}
}

func Length(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case Vec2.Vector{}:
		return vector.(Vec2.Vector).Length()
	case Vec3.Vector{}:
		return vector.(Vec3.Vector).Length()
	case Vec4.Vector{}:
		return vector.(Vec4.Vector).Length()
	default:
		return nil
	}
}

func Norm(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case Vec2.Vector{}:
		return vector.(Vec2.Vector).Norm()
	case Vec3.Vector{}:
		return vector.(Vec3.Vector).Norm()
	case Vec4.Vector{}:
		return vector.(Vec4.Vector).Norm()
	default:
		return nil
	}
}
