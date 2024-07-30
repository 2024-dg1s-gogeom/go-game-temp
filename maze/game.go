package maze

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth, WindowHeight = 800, 600
)

type scene int

const (
	sceneFront scene = iota
	scenePlay
	sceneFinish
)

var (
	draw = map[scene]func(*Game, *ebiten.Image){
		sceneFront:  drawFront,
		scenePlay:   drawPlay,
		sceneFinish: drawFinish,
	}
)

type Game struct {
	scene
}

func (g *Game) Update() error {
	switch g.scene {
	case sceneFront:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.scene = scenePlay
		}
	case scenePlay:
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.scene = sceneFinish
		}
	case sceneFinish:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.scene = sceneFront
		}
	}
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	draw[g.scene](g, scr)
}

func (g *Game) Layout(int, int) (int, int) {
	return WindowWidth, WindowHeight
}
