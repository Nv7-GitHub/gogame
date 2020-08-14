package gogame

import rl "github.com/gen2brain/raylib-go/raylib"

// Component is a basic component. To render it, call the draw function
type Component interface {
	Load()
	Draw()
	Unload()
}

// Rect is a rectangle, it draws a rectangle to the screen
type Rect struct {
	x      int32
	y      int32
	width  int32
	height int32
	color  Color
}

// Load function is called on initialization, but is mostly for things that read from disk
func (r Rect) Load() {
	return
}

// Draw is for drawing rect
func (r Rect) Draw() {
	rl.DrawRectangle(r.x, r.y, r.width, r.height, r.color.color)
}

// Unload is for taking things from disk out of RAM, doesnt apply to rect
func (r Rect) Unload() {
	return
}

// NewRect creates a rect with the parameters
func NewRect(x, y, width, height int32, color Color) Rect {
	return Rect{x, y, width, height, color}
}
