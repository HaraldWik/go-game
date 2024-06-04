package vec

import (
	"github.com/HaraldWik/go-game/scr/types"
	vector2 "github.com/HaraldWik/go-game/scr/vector/2"
	vector3 "github.com/HaraldWik/go-game/scr/vector/3"
	vector4 "github.com/HaraldWik/go-game/scr/vector/4"
)

func ToVector2(vector interface{}) vector2.Type {
	switch vec := vector.(type) {
	case vector2.Type:
		return vec
	case vector3.Type:
		return vector2.New(vector.(vector3.Type).X, vector.(vector3.Type).Y)
	case vector4.Type:
		return vector2.New(vector.(vector4.Type).X, vector.(vector4.Type).Y)
	default:
		return vector2.Zero()
	}
}

func ToVector3(vector interface{}) vector3.Type {
	switch vec := vector.(type) {
	case vector2.Type:
		return vector3.New(vector.(vector2.Type).X, vector.(vector2.Type).Y, 0.0)
	case vector3.Type:
		return vec
	case vector4.Type:
		return vector3.New(vector.(vector4.Type).X, vector.(vector4.Type).Y, vector.(vector4.Type).Z)
	default:
		return vector3.Zero()
	}
}

func ToVector4(vector interface{}) vector4.Type {
	switch vec := vector.(type) {
	case vector2.Type:
		return vector4.New(vector.(vector2.Type).X, vector.(vector2.Type).Y, 0.0, 0.0)
	case vector3.Type:
		return vector4.New(vector.(vector3.Type).X, vector.(vector3.Type).Y, vector.(vector3.Type).Z, 0.0)
	case vector4.Type:
		return vec
	default:
		return vector4.Zero()
	}
}

func Add(first interface{}, second interface{}) interface{} {
	if types.AreSameType(true, first, second) {
		switch types.GetType(first) {
		case vector2.Type{}:
			return first.(vector2.Type).Add(second.(vector2.Type))
		case vector3.Type{}:
			return first.(vector3.Type).Add(second.(vector3.Type))
		case vector4.Type{}:
			return first.(vector4.Type).Add(second.(vector4.Type))
		default:
			return nil
		}
	} else {
		return nil
	}
}

func VecSub(first interface{}, second interface{}) interface{} {
	if types.AreSameType(true, first, second) {
		switch types.GetType(first) {
		case vector2.Type{}:
			return first.(vector2.Type).Sub(second.(vector2.Type))
		case vector3.Type{}:
			return first.(vector3.Type).Sub(second.(vector3.Type))
		case vector4.Type{}:
			return first.(vector4.Type).Sub(second.(vector4.Type))
		default:
			return nil
		}
	} else {
		return nil
	}
}

func VecMul(first interface{}, second interface{}) interface{} {
	if types.AreSameType(true, first, second) {
		switch types.GetType(first) {
		case vector2.Type{}:
			return first.(vector2.Type).Mul(second.(vector2.Type))
		case vector3.Type{}:
			return first.(vector3.Type).Mul(second.(vector3.Type))
		case vector4.Type{}:
			return first.(vector4.Type).Mul(second.(vector4.Type))
		default:
			return nil
		}
	} else {
		return nil
	}
}

func VecDiv(first interface{}, second interface{}) interface{} {
	if types.AreSameType(true, first, second) {
		switch types.GetType(first) {
		case vector2.Type{}:
			return first.(vector2.Type).Div(second.(vector2.Type))
		case vector3.Type{}:
			return first.(vector3.Type).Div(second.(vector3.Type))
		case vector4.Type{}:
			return first.(vector4.Type).Div(second.(vector4.Type))
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
	case vector2.Type{}:
		return vector.(vector2.Type).ABS()
	case vector3.Type{}:
		return vector.(vector3.Type).ABS()
	case vector4.Type{}:
		return vector.(vector4.Type).ABS()
	default:
		return nil
	}
}

func Neg(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case vector2.Type{}:
		return vector.(vector2.Type).Neg()
	case vector3.Type{}:
		return vector.(vector3.Type).Neg()
	case vector4.Type{}:
		return vector.(vector4.Type).Neg()
	default:
		return nil
	}
}

func Length(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case vector2.Type{}:
		return vector.(vector2.Type).Length()
	case vector3.Type{}:
		return vector.(vector3.Type).Length()
	case vector4.Type{}:
		return vector.(vector4.Type).Length()
	default:
		return nil
	}
}

func Norm(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case vector2.Type{}:
		return vector.(vector2.Type).Norm()
	case vector3.Type{}:
		return vector.(vector3.Type).Norm()
	case vector4.Type{}:
		return vector.(vector4.Type).Norm()
	default:
		return nil
	}
}
