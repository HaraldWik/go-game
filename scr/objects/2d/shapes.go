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
	sinR := float32(math.Sin(float64(shape.transform.rot)))
	cosR := float32(math.Cos(float64(shape.transform.rot)))

	x0 := -0.5*shape.transform.size.X*cosR - -0.5*shape.transform.size.Y*sinR + shape.transform.pos.X
	y0 := -0.5*shape.transform.size.X*sinR + -0.5*shape.transform.size.Y*cosR + shape.transform.pos.Y

	x1 := 0.5*shape.transform.size.X*cosR - -0.5*shape.transform.size.Y*sinR + shape.transform.pos.X
	y1 := 0.5*shape.transform.size.X*sinR + -0.5*shape.transform.size.Y*cosR + shape.transform.pos.Y

	x2 := 0.5*shape.transform.size.X*cosR - 0.5*shape.transform.size.Y*sinR + shape.transform.pos.X
	y2 := 0.5*shape.transform.size.X*sinR + 0.5*shape.transform.size.Y*cosR + shape.transform.pos.Y

	x3 := -0.5*shape.transform.size.X*cosR - 0.5*shape.transform.size.Y*sinR + shape.transform.pos.X
	y3 := -0.5*shape.transform.size.X*sinR + 0.5*shape.transform.size.Y*cosR + shape.transform.pos.Y

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
	sinR := float32(math.Sin(float64(shape.transform.rot)))
	cosR := float32(math.Cos(float64(shape.transform.rot)))

	x0 := shape.transform.pos.X
	y0 := shape.transform.pos.Y + shape.transform.size.Y/2

	var x1, y1, x2, y2 float32

	if flip {
		x1 = shape.transform.pos.X - shape.transform.size.X/2*cosR + shape.transform.size.Y/2*sinR
		y1 = shape.transform.pos.Y - shape.transform.size.X/2*sinR - shape.transform.size.Y/2*cosR

		x2 = shape.transform.pos.X + shape.transform.size.X/2*cosR + shape.transform.size.Y/2*sinR
		y2 = shape.transform.pos.Y + shape.transform.size.X/2*sinR - shape.transform.size.Y/2*cosR
	} else {
		x1 = shape.transform.pos.X + shape.transform.size.X/2*cosR - shape.transform.size.Y/2*sinR
		y1 = shape.transform.pos.Y + shape.transform.size.X/2*sinR + shape.transform.size.Y/2*cosR

		x2 = shape.transform.pos.X - shape.transform.size.X/2*cosR - shape.transform.size.Y/2*sinR
		y2 = shape.transform.pos.Y - shape.transform.size.X/2*sinR + shape.transform.size.Y/2*cosR
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
		x := float32(math.Cos(float64(i)*theta)) * shape.transform.size.X / 2.0
		y := float32(math.Sin(float64(i)*theta)) * shape.transform.size.Y / 2.0

		// *Apply position & rotation
		cosR := float32(math.Cos(float64(shape.transform.rot)))
		sinR := float32(math.Sin(float64(shape.transform.rot)))
		rotatedX := x*cosR - y*sinR + shape.transform.pos.X
		rotatedY := x*sinR + y*cosR + shape.transform.pos.Y

		gl.Vertex2f(rotatedX, rotatedY)
	}

	gl.End()
}
