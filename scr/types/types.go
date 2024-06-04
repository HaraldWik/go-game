package types

import (
	"log"
	"reflect"

	err_ "github.com/HaraldWik/go-game/scr/err"
	vec2 "github.com/HaraldWik/go-game/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game/scr/vector/3"
	vec4 "github.com/HaraldWik/go-game/scr/vector/4"
)

func AreSameType(panicWhenFalse bool, vars ...interface{}) bool {
	if len(vars) < 2 {
		return false
	}

	firstType := reflect.TypeOf(vars[0])

	for _, v := range vars[1:] {
		if reflect.TypeOf(v) != firstType {
			if panicWhenFalse {
				log.Fatalln(err_.WARNING + "Needs to be of same type")
			} else {
				return false
			}
		}
	}
	return true
}

func GetType(variable interface{}) interface{} {
	switch variable.(type) {
	case int8:
		return int8(0)
	case int16:
		return int16(0)
	case int32:
		return int32(0)
	case int64:
		return int64(0)
	case float32:
		return float32(0)
	case float64:
		return float64(0)
	case string:
		return ""
	case vec2.Type:
		return vec2.Type{}
	case vec3.Type:
		return vec3.Type{}
	case vec4.Type:
		return vec4.Type{}
	default:
		return nil
	}
}

func TypeConv(target interface{}, variable interface{}) interface{} {
	targetType := reflect.TypeOf(target).Kind()
	value := reflect.ValueOf(variable)

	switch targetType {
	case reflect.Int8:
		switch value.Kind() {
		case reflect.Int8:
			return int8(value.Interface().(int8))
		case reflect.Int16:
			return int8(value.Interface().(int16))
		case reflect.Int32:
			return int8(value.Interface().(int32))
		case reflect.Int64:
			return int8(value.Interface().(int64))
		case reflect.Float32:
			return int8(value.Interface().(float32))
		case reflect.Float64:
			return int8(value.Interface().(float64))
		default:
			return nil
		}

	case reflect.Int16:
		switch value.Kind() {
		case reflect.Int8:
			return int16(value.Interface().(int8))
		case reflect.Int16:
			return int16(value.Interface().(int16))
		case reflect.Int32:
			return int16(value.Interface().(int32))
		case reflect.Int64:
			return int16(value.Interface().(int64))
		case reflect.Float32:
			return int16(value.Interface().(float32))
		case reflect.Float64:
			return int16(value.Interface().(float64))
		default:
			return nil
		}

	case reflect.Int32:
		switch value.Kind() {
		case reflect.Int8:
			return int32(value.Interface().(int8))
		case reflect.Int16:
			return int32(value.Interface().(int16))
		case reflect.Int32:
			return int32(value.Interface().(int32))
		case reflect.Int64:
			return int32(value.Interface().(int64))
		case reflect.Float32:
			return int32(value.Interface().(float32))
		case reflect.Float64:
			return int32(value.Interface().(float64))
		default:
			return nil
		}

	case reflect.Int64:
		switch value.Kind() {
		case reflect.Int8:
			return int64(value.Interface().(int8))
		case reflect.Int16:
			return int64(value.Interface().(int16))
		case reflect.Int32:
			return int64(value.Interface().(int32))
		case reflect.Int64:
			return int64(value.Interface().(int64))
		case reflect.Float32:
			return int64(value.Interface().(float32))
		case reflect.Float64:
			return int64(value.Interface().(float64))
		default:
			return nil
		}

	case reflect.Float32:
		switch value.Kind() {
		case reflect.Int8:
			return float32(value.Interface().(int8))
		case reflect.Int16:
			return float32(value.Interface().(int16))
		case reflect.Int32:
			return float32(value.Interface().(int32))
		case reflect.Int64:
			return float32(value.Interface().(int64))
		case reflect.Float32:
			return float32(value.Interface().(float32))
		case reflect.Float64:
			return float32(value.Interface().(float64))
		default:
			return nil
		}

	case reflect.Float64:
		switch value.Kind() {
		case reflect.Int8:
			return float64(value.Interface().(int8))
		case reflect.Int16:
			return float64(value.Interface().(int16))
		case reflect.Int32:
			return float64(value.Interface().(int32))
		case reflect.Int64:
			return float64(value.Interface().(int64))
		case reflect.Float32:
			return float64(value.Interface().(float32))
		case reflect.Float64:
			return float64(value.Interface().(float64))
		default:
			return nil
		}

	}

	return nil
}
