package Vec2

import "math"

type Vector struct {
	X, Y float32
}

func New(vectorX, vectorY float32) Vector {
	return Vector{
		X: vectorX,
		Y: vectorY,
	}
}

func Zero() Vector {
	return Vector{
		X: 0.0,
		Y: 0.0,
	}
}

func All(number float32) Vector {
	return Vector{
		X: number,
		Y: number,
	}
}

//* Addition
func (origin Vector) Add(vectors ...Vector) Vector {
	for _, vector := range vectors {
		origin.X += vector.X
		origin.Y += vector.Y
	}
	return origin
}

//* Subtraction
func (origin Vector) Sub(vectors ...Vector) Vector {
	for _, vector := range vectors {
		origin.X -= vector.X
		origin.Y -= vector.Y
	}
	return origin
}

//* Multiplication
func (origin Vector) Mul(vectors ...Vector) Vector {
	for _, vector := range vectors {
		origin.X *= vector.X
		origin.Y *= vector.Y
	}
	return origin
}

//* Divition
func (origin Vector) Div(vectors ...Vector) Vector {
	for _, vector := range vectors {
		origin.X *= vector.X
		origin.Y *= vector.Y
	}
	return origin
}

//* Absolut
func (origin Vector) ABS() Vector {
	return New(
		float32(math.Abs(float64(origin.X))),
		float32(math.Abs(float64(origin.Y))),
	)
}

//* Negativ
func (origin Vector) Neg() Vector {
	return New(
		-float32(math.Abs(float64(origin.X))),
		-float32(math.Abs(float64(origin.Y))),
	)
}

//* Length
func (origin Vector) Length() float32 {
	return float32(math.Sqrt(
		float64(origin.X*origin.X) +
			float64(origin.Y*origin.Y)),
	)
}

//* Normalize
func (origin Vector) Norm() Vector {
	if origin.Length() != 0 {
		return New(
			origin.X/origin.Length(),
			origin.Y/origin.Length(),
		)
	} else {
		return Zero()
	}
}
