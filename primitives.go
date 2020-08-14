package gogame

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Rect is a rectangle, it draws a rectangle to the screen
type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
	Color  Color
}

// Load function is called on initialization, but is mostly for things that read from disk
func (r *Rect) Load() {
	return
}

// Draw is for drawing rect
func (r *Rect) Draw() {
	rl.DrawRectangle(r.X, r.Y, r.Width, r.Height, r.Color.Color)
}

// Unload is for taking things from disk out of RAM, doesnt apply to rect
func (r *Rect) Unload() {
	return
}

// Move is for moving a rectangle to a specific point
func (r *Rect) Move(x, y int32) {
	r.X = x
	r.Y = y
}

// Translate is for translating a rectangle a certain amount
func (r *Rect) Translate(changeX, changeY int32) {
	r.X += changeX
	r.Y += changeY
}

// NewRect creates a rect with the parameters
func NewRect(x, y, width, height int32, color Color) *Rect {
	return &Rect{x, y, width, height, color}
}

// Circle is a circle, it draws a circle to the screen
type Circle struct {
	X      int32
	Y      int32
	Radius float32
	Color  Color
}

// Load function is called on initialization, but is mostly for things that read from disk
func (c *Circle) Load() {
	return
}

// Draw is for drawing rect
func (c *Circle) Draw() {
	rl.DrawCircle(c.X, c.Y, c.Radius, c.Color.Color)
}

// Unload is for taking things from disk out of RAM, doesnt apply to rect
func (c *Circle) Unload() {
	return
}

// Move is for moving a circle to a specific point
func (c *Circle) Move(x, y int32) {
	c.X = x
	c.Y = y
}

// Translate is for translating a circle a certain amount
func (c *Circle) Translate(changeX, changeY int32) {
	c.X += changeX
	c.Y += changeY
}

// NewCircle creates a rect with the parameters
func NewCircle(x, y int32, radius float32, color Color) *Circle {
	return &Circle{x, y, radius, color}
}

// Image is a image, it draws an image to the screen
type Image struct {
	X     int32
	Y     int32
	Path  string
	tex   rl.Texture2D
	color rl.Color
	image *rl.Image
}

// Load function is called on initialization, but is mostly for things that read from disk
func (i *Image) Load() {
	i.image = rl.LoadImage(i.Path)
	i.tex = rl.LoadTextureFromImage(i.image)
}

// Draw is for drawing oval
func (i *Image) Draw() {
	rl.DrawTexture(i.tex, i.X, i.Y, i.color)
}

// Unload is for taking things from disk out of RAM, doesnt apply to rect
func (i *Image) Unload() {
	rl.UnloadImage(i.image)
}

// Resize is for resizing an image
func (i *Image) Resize(width, height int32) {
	rl.ImageResize(i.image, width, height)
	i.tex = rl.LoadTextureFromImage(i.image)
}

// Move is for moving a rectangle to a specific point
func (i *Image) Move(x, y int32) {
	i.X = x
	i.Y = y
}

// Translate is for translating a rectangle a certain amount
func (i *Image) Translate(changeX, changeY int32) {
	i.X += changeX
	i.Y += changeY
}

// NewImage creates a rect with the parameters
func NewImage(x, y, width, height int32, path string) *Image {
	var img *rl.Image
	var tex *rl.Texture2D
	return &Image{x, y, path, *tex, rl.RayWhite, img}
}
