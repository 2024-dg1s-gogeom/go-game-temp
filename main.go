package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hyosangkang/multi-game/maze/maze"
)

func main() {
	ebiten.SetWindowSize(maze.WindowWidth, maze.WindowHeight)
	ebiten.SetWindowTitle("Maze!")
	if err := ebiten.RunGame(&maze.Game{}); err != nil {
		log.Fatal(err)
	}
}
