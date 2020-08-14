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
	R     uint8
	G     uint8
	B     uint8
	A     uint8
	Color rl.Color
}
