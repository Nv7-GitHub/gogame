package gogame

// Component is a basic component. To render it, call the draw function
type Component interface {
	Load()
	Draw()
	Unload()
}
