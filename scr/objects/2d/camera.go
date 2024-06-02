package obj2d

import (
	vec2 "github.com/HaraldWik/go-game/scr/vector/2"
	win "github.com/HaraldWik/go-game/scr/window"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera struct {
	win  win.Window
	pos  vec2.Vector
	rot  float32
	zoom float32
}

func NewCam() Camera {
	return Camera{}
}

func NewCamAdvanced(window win.Window, position vec2.Vector, rotation, zoom float32) Camera {
	return Camera{
		win:  window,
		pos:  position,
		rot:  rotation,
		zoom: zoom,
	}
}

// *Set
func (camera *Camera) SetConnectedWindow(window win.Window) {
	camera.win = window
}

func (camera *Camera) SetPos(position vec2.Vector) {
	camera.pos = position
}

func (camera *Camera) SetRot(rotation float32) {
	camera.rot = rotation
}

func (camera *Camera) SetZoom(zoom float32) {
	camera.zoom = zoom
}

// *Get
func (camera *Camera) GetConnectedWindow() win.Window {
	return camera.win
}

func (camera *Camera) GetPos() vec2.Vector {
	return camera.pos
}

func (camera *Camera) GetRot() float32 {
	return camera.rot
}

func (camera *Camera) GetZoom() float32 {
	return camera.zoom
}

// *Draw
func (camera Camera) Draw() {
	gl.Viewport(0, 0, int32(camera.win.GetSizeWidth()), int32(camera.win.GetSizeHeight()))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// *Calculate aspect ratio
	aspect := float32(camera.win.GetSizeWidth()) / float32(camera.win.GetSizeHeight())
	left := -1*camera.zoom*aspect + camera.pos.X
	right := 1*camera.zoom*aspect + camera.pos.X
	top := 1*camera.zoom + camera.pos.Y
	bottom := -1*camera.zoom + camera.pos.Y
	gl.Ortho(float64(left), float64(right), float64(bottom), float64(top), -3, 3)

	// *Rotation & position
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(float32(camera.pos.X), float32(camera.pos.Y), 0)
	gl.Rotatef(float32(camera.rot), 0, 0, 1)
	gl.Translatef(-float32(camera.pos.X), -float32(camera.pos.Y), 0)
}
