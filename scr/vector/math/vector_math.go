package vec

import (
	"github.com/HaraldWik/go-game/scr/types"
	vec2 "github.com/HaraldWik/go-game/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game/scr/vector/3"
	vec4 "github.com/HaraldWik/go-game/scr/vector/4"
)

func ToVec2(vector interface{}) vec2.Vector {
	switch vec := vector.(type) {
	case vec2.Vector:
		return vec
	case vec3.Vector:
		return vec2.New(vector.(vec3.Vector).X, vector.(vec3.Vector).Y)
	case vec4.Vector:
		return vec2.New(vector.(vec4.Vector).X, vector.(vec4.Vector).Y)
	default:
		return vec2.Zero()
	}
}

func ToVec3(vector interface{}) vec3.Vector {
	switch vec := vector.(type) {
	case vec2.Vector:
		return vec3.New(vector.(vec2.Vector).X, vector.(vec2.Vector).Y, 0.0)
	case vec3.Vector:
		return vec
	case vec4.Vector:
		return vec3.New(vector.(vec4.Vector).X, vector.(vec4.Vector).Y, vector.(vec4.Vector).Z)
	default:
		return vec3.Zero()
	}
}

func ToVec4(vector interface{}) vec4.Vector {
	switch vec := vector.(type) {
	case vec2.Vector:
		return vec4.New(vector.(vec2.Vector).X, vector.(vec2.Vector).Y, 0.0, 0.0)
	case vec3.Vector:
		return vec4.New(vector.(vec3.Vector).X, vector.(vec3.Vector).Y, vector.(vec3.Vector).Z, 0.0)
	case vec4.Vector:
		return vec
	default:
		return vec4.Zero()
	}
}

func Add(vector1 interface{}, vector2 interface{}) interface{} {
	if types.AreSameType(true, vector1, vector2) {
		switch types.GetType(vector1) {
		case vec2.Vector{}:
			return vector1.(vec2.Vector).Add(vector2.(vec2.Vector))
		case vec3.Vector{}:
			return vector1.(vec3.Vector).Add(vector2.(vec3.Vector))
		case vec4.Vector{}:
			return vector1.(vec4.Vector).Add(vector2.(vec4.Vector))
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
		case vec2.Vector{}:
			return vector1.(vec2.Vector).Sub(vector2.(vec2.Vector))
		case vec3.Vector{}:
			return vector1.(vec3.Vector).Sub(vector2.(vec3.Vector))
		case vec4.Vector{}:
			return vector1.(vec4.Vector).Sub(vector2.(vec4.Vector))
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
		case vec2.Vector{}:
			return vector1.(vec2.Vector).Mul(vector2.(vec2.Vector))
		case vec3.Vector{}:
			return vector1.(vec3.Vector).Mul(vector2.(vec3.Vector))
		case vec4.Vector{}:
			return vector1.(vec4.Vector).Mul(vector2.(vec4.Vector))
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
		case vec2.Vector{}:
			return vector1.(vec2.Vector).Div(vector2.(vec2.Vector))
		case vec3.Vector{}:
			return vector1.(vec3.Vector).Div(vector2.(vec3.Vector))
		case vec4.Vector{}:
			return vector1.(vec4.Vector).Div(vector2.(vec4.Vector))
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
	case vec2.Vector{}:
		return vector.(vec2.Vector).ABS()
	case vec3.Vector{}:
		return vector.(vec3.Vector).ABS()
	case vec4.Vector{}:
		return vector.(vec4.Vector).ABS()
	default:
		return nil
	}
}

func Neg(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case vec2.Vector{}:
		return vector.(vec2.Vector).Neg()
	case vec3.Vector{}:
		return vector.(vec3.Vector).Neg()
	case vec4.Vector{}:
		return vector.(vec4.Vector).Neg()
	default:
		return nil
	}
}

func Length(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case vec2.Vector{}:
		return vector.(vec2.Vector).Length()
	case vec3.Vector{}:
		return vector.(vec3.Vector).Length()
	case vec4.Vector{}:
		return vector.(vec4.Vector).Length()
	default:
		return nil
	}
}

func Norm(vector interface{}) interface{} {
	switch types.GetType(vector) {
	case vec2.Vector{}:
		return vector.(vec2.Vector).Norm()
	case vec3.Vector{}:
		return vector.(vec3.Vector).Norm()
	case vec4.Vector{}:
		return vector.(vec4.Vector).Norm()
	default:
		return nil
	}
}
