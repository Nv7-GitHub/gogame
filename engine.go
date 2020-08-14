package GPGame

import rl "github.com/gen2brain/raylib-go/raylib"

// Engine is for the main engine
type Engine struct {
	WIDTH  int32
	HEIGHT int32
	tick   func(delta float64)
}

func createEngine(width, height int32, tick func(delta float64)) Engine {
	return Engine{width, height, tick}
}

func (e *Engine) run(title string, fps int32) {
	rl.InitWindow(e.WIDTH, e.HEIGHT, title)
	rl.SetTargetFPS(fps)
}
