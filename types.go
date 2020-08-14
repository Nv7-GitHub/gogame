package gogame

import rl "github.com/gen2brain/raylib-go/raylib"

// Component is a basic component. To render it, call the draw function
type Component interface {
	Load()
	Draw()
	Unload()
}

// Color is the datatype for colors in gogame
type Color struct {
	r     uint8
	g     uint8
	b     uint8
	a     uint8
	color rl.Color
}
