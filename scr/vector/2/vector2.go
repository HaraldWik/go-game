package vector2

import "math"

type Type struct {
	X, Y float32
}

func New(x, y float32) Type {
	return Type{
		X: x,
		Y: y,
	}
}

func Zero() Type {
	return Type{
		X: 0.0,
		Y: 0.0,
	}
}

func All(number float32) Type {
	return Type{
		X: number,
		Y: number,
	}
}

// *Addition
func (origin Type) Add(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X += vector.X
		origin.Y += vector.Y
	}
	return origin
}

// *Subtraction
func (origin Type) Sub(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X -= vector.X
		origin.Y -= vector.Y
	}
	return origin
}

// *Multiplication
func (origin Type) Mul(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X *= vector.X
		origin.Y *= vector.Y
	}
	return origin
}

// *Divition
func (origin Type) Div(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X *= vector.X
		origin.Y *= vector.Y
	}
	return origin
}

// *Absolut
func (origin Type) ABS() Type {
	return New(
		float32(math.Abs(float64(origin.X))),
		float32(math.Abs(float64(origin.Y))),
	)
}

// *Negativ
func (origin Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(origin.X))),
		-float32(math.Abs(float64(origin.Y))),
	)
}

// *Length
func (origin Type) Length() float32 {
	return float32(math.Sqrt(
		float64(origin.X*origin.X) +
			float64(origin.Y*origin.Y)),
	)
}

// *Normalize
func (origin Type) Norm() Type {
	if origin.Length() != 0 {
		return New(
			origin.X/origin.Length(),
			origin.Y/origin.Length(),
		)
	} else {
		return Zero()
	}
}
