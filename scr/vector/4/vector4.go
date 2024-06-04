package vector4

import "math"

type Type struct {
	X, Y, Z, W float32
}

func New(x, y, z, w float32) Type {
	return Type{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func Zero() Type {
	return Type{
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
		W: 0.0,
	}
}

func All(number float32) Type {
	return Type{
		X: number,
		Y: number,
		Z: number,
		W: number,
	}
}

// *Addition
func (origin Type) Add(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X += vector.X
		origin.Y += vector.Y
		origin.Z += vector.Z
		origin.W += vector.W
	}
	return origin
}

// *Subtraction
func (origin Type) Sub(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X -= vector.X
		origin.Y -= vector.Y
		origin.Z -= vector.Z
		origin.W -= vector.W
	}
	return origin
}

// *Multiplication
func (origin Type) Mul(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X *= vector.X
		origin.Y *= vector.Y
		origin.Z *= vector.Z
		origin.W *= vector.W
	}
	return origin
}

// *Divition
func (origin Type) Div(vectors ...Type) Type {
	for _, vector := range vectors {
		origin.X *= vector.X
		origin.Y *= vector.Y
		origin.Z *= vector.Z
		origin.W *= vector.W
	}
	return origin
}

// *Absolut
func (origin Type) ABS() Type {
	return New(
		float32(math.Abs(float64(origin.X))),
		float32(math.Abs(float64(origin.Y))),
		float32(math.Abs(float64(origin.Y))),
		float32(math.Abs(float64(origin.W))),
	)
}

// *Negativ
func (origin Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(origin.X))),
		-float32(math.Abs(float64(origin.Y))),
		-float32(math.Abs(float64(origin.Z))),
		-float32(math.Abs(float64(origin.W))),
	)
}

// *Length
func (origin Type) Length() float32 {
	return float32(math.Sqrt(
		float64(origin.X*origin.X) +
			float64(origin.Y*origin.Y) +
			float64(origin.Z*origin.Z) +
			float64(origin.W*origin.W)),
	)
}

// *Normalize
func (origin Type) Norm() Type {
	if origin.Length() != 0 {
		return New(
			origin.X/origin.Length(),
			origin.Y/origin.Length(),
			origin.Z/origin.Length(),
			origin.W/origin.Length(),
		)
	} else {
		return Zero()
	}
}
