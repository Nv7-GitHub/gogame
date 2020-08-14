package gogame

import rl "github.com/gen2brain/raylib-go/raylib"

// Engine is for the main engine
type Engine struct {
	WIDTH      int32
	HEIGHT     int32
	tick       func(delta float32)
	Bg         Color
	Components []Component
}

// NewColor is for creating a color
func NewColor(r, g, b, a uint8) Color {
	return Color{r, g, b, a, rl.NewColor(r, g, b, a)}
}

// NewEngine creates an engine object with parameters
func NewEngine(width, height int32, tick func(delta float32)) Engine {
	return Engine{width, height, tick, NewColor(0, 0, 0, 1), make([]Component, 0)}
}

// AddComponent adds a component to the engine
func (e *Engine) AddComponent(component Component) {
	e.Components = append(e.Components, component)
}

// Start is for creating the window and the gameloop
func (e *Engine) Start(title string, fps int32) {
	rl.InitWindow(e.WIDTH, e.HEIGHT, title)
	rl.SetTargetFPS(fps)

	oldTime := rl.GetTime()

	for _, c := range e.Components {
		c.Load()
	}

	for !rl.WindowShouldClose() {
		e.tick(rl.GetTime() - oldTime)

		rl.BeginDrawing()

		rl.ClearBackground(e.Bg.Color)

		for _, c := range e.Components {
			c.Draw()
		}

		rl.EndDrawing()

		oldTime = rl.GetTime()
	}

	for _, c := range e.Components {
		c.Unload()
	}

	rl.CloseWindow()
}
