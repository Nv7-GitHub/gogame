package gogame

import rl "github.com/gen2brain/raylib-go/raylib"

// Engine is for the main engine
type Engine struct {
	WIDTH  int32
	HEIGHT int32
	tick   func(delta float64)
}

func CreateEngine(width, height int32, tick func(delta float64)) Engine {
	return Engine{width, height, tick}
}

func (e *Engine) Start(title string, fps int32) {
	rl.InitWindow(e.WIDTH, e.HEIGHT, title)
	rl.SetTargetFPS(fps)
}
