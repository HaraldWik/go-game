package obj2d

import (
	"math"

	"github.com/HaraldWik/go-game/scr/objects/color"
	"github.com/go-gl/gl/v2.1/gl"
)

type Shape struct {
	transform Transform
	color     color.RGB
}

func NewShape2D(transform Transform, rgb color.RGB) Shape {
	return Shape{
		transform: transform,
		color:     rgb,
	}
}

func (shape Shape) DrawRectangle() {
	gl.Begin(gl.QUADS)
	gl.Color3f(shape.color.R, shape.color.G, shape.color.B)

	// *Calculate vertices after rot
	sinR := float32(math.Sin(float64(shape.transform.Rot)))
	cosR := float32(math.Cos(float64(shape.transform.Rot)))

	x0 := -0.5*shape.transform.Size.X*cosR - -0.5*shape.transform.Size.Y*sinR + shape.transform.Pos.X
	y0 := -0.5*shape.transform.Size.X*sinR + -0.5*shape.transform.Size.Y*cosR + shape.transform.Pos.Y

	x1 := 0.5*shape.transform.Size.X*cosR - -0.5*shape.transform.Size.Y*sinR + shape.transform.Pos.X
	y1 := 0.5*shape.transform.Size.X*sinR + -0.5*shape.transform.Size.Y*cosR + shape.transform.Pos.Y

	x2 := 0.5*shape.transform.Size.X*cosR - 0.5*shape.transform.Size.Y*sinR + shape.transform.Pos.X
	y2 := 0.5*shape.transform.Size.X*sinR + 0.5*shape.transform.Size.Y*cosR + shape.transform.Pos.Y

	x3 := -0.5*shape.transform.Size.X*cosR - 0.5*shape.transform.Size.Y*sinR + shape.transform.Pos.X
	y3 := -0.5*shape.transform.Size.X*sinR + 0.5*shape.transform.Size.Y*cosR + shape.transform.Pos.Y

	// *Draw vertices
	gl.Vertex2f(x0, y0) // *Bottom-left
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Top-right
	gl.Vertex2f(x3, y3) // *Top-left

	gl.End()
}

func (shape Shape) DrawTriangle(flip bool) {
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(shape.color.R, shape.color.G, shape.color.B)

	// *Calculate vertices after rotation
	sinR := float32(math.Sin(float64(shape.transform.Rot)))
	cosR := float32(math.Cos(float64(shape.transform.Rot)))

	x0 := shape.transform.Pos.X
	y0 := shape.transform.Pos.Y + shape.transform.Size.Y/2

	var x1, y1, x2, y2 float32

	if flip {
		x1 = shape.transform.Pos.X - shape.transform.Size.X/2*cosR + shape.transform.Size.Y/2*sinR
		y1 = shape.transform.Pos.Y - shape.transform.Size.X/2*sinR - shape.transform.Size.Y/2*cosR

		x2 = shape.transform.Pos.X + shape.transform.Size.X/2*cosR + shape.transform.Size.Y/2*sinR
		y2 = shape.transform.Pos.Y + shape.transform.Size.X/2*sinR - shape.transform.Size.Y/2*cosR
	} else {
		x1 = shape.transform.Pos.X + shape.transform.Size.X/2*cosR - shape.transform.Size.Y/2*sinR
		y1 = shape.transform.Pos.Y + shape.transform.Size.X/2*sinR + shape.transform.Size.Y/2*cosR

		x2 = shape.transform.Pos.X - shape.transform.Size.X/2*cosR - shape.transform.Size.Y/2*sinR
		y2 = shape.transform.Pos.Y - shape.transform.Size.X/2*sinR + shape.transform.Size.Y/2*cosR
	}

	// *Draw vertices
	gl.Vertex2f(x0, y0) // *Top
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Bottom-left

	gl.End()
}

func (shape Shape) DrawCircle(segments int) {
	gl.Begin(gl.LINE_LOOP)
	gl.Color3f(shape.color.R, shape.color.G, shape.color.B)

	theta := 2.0 * math.Pi / float64(segments)

	// *Draw circle points
	for i := 0; i < segments; i++ {
		x := float32(math.Cos(float64(i)*theta)) * shape.transform.Size.X / 2.0
		y := float32(math.Sin(float64(i)*theta)) * shape.transform.Size.Y / 2.0

		// *Apply position & rotation
		cosR := float32(math.Cos(float64(shape.transform.Rot)))
		sinR := float32(math.Sin(float64(shape.transform.Rot)))
		rotatedX := x*cosR - y*sinR + shape.transform.Pos.X
		rotatedY := x*sinR + y*cosR + shape.transform.Pos.Y

		gl.Vertex2f(rotatedX, rotatedY)
	}

	gl.End()
}
