package maze

import (
	"bytes"
	"image/color"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Title                = "Maze!"
	TitleFontSize        = 48
	TitleMessage         = "Press  Enter  to  start"
	TitleMessageFontSize = 24
	Finish               = "Finish!"
	FinishMessage        = "Press  Enter  to  restart"
)

var (
	InstructionMessage = []string{
		"Q  W  E  R  for  X  Y  Z  W  movement",
		"SHIFT  for  backward",
	}
	TitleFontFace *text.GoTextFaceSource
	//go:embed arcadeclassic.ttf
	ArcadeClassic_ttf []byte
)

func init() {
	TitleFontFace, _ = text.NewGoTextFaceSource(bytes.NewReader(ArcadeClassic_ttf))
}

func drawText(screen *ebiten.Image, msg string, size float64, x, y float64) {
	text.Draw(screen, msg, &text.GoTextFace{
		Source: TitleFontFace,
		Size:   size,
	}, NewTextDrawOption(x, y))
}

func NewTextDrawOption(x, y float64) *text.DrawOptions {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.White)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	return op
}

func drawFront(g *Game, scr *ebiten.Image) {
	drawText(scr, Title, TitleFontSize, WindowWidth/2, WindowHeight/2-150)
	drawText(scr, TitleMessage, TitleMessageFontSize, WindowWidth/2, WindowHeight/2)
	for i, msg := range InstructionMessage {
		drawText(scr, msg, TitleMessageFontSize, WindowWidth/2, WindowHeight/2+100+float64(i)*30)
	}
}

func drawPlay(g *Game, scr *ebiten.Image) {
	// draw dividers
	for i := 0; i < WindowWidth; i++ {
		scr.Set(i, WindowHeight/2, color.White)
	}
	for i := 0; i < WindowHeight; i++ {
		scr.Set(WindowWidth/2, i, color.White)
	}
	for i := 0; i < 4; i++ {
		drawProjection(g, scr, i)
	}
	// draw position box
	vector.DrawFilledRect(scr, WindowWidth/2-100, WindowHeight/2-50, 200, 100, color.Black, false)
	vector.StrokeRect(scr, WindowWidth/2-100, WindowHeight/2-50, 200, 100, 1, color.White, false)
	drawText(scr, "Position  ", 24, WindowWidth/2, WindowHeight/2-24)
}

func drawProjection(g *Game, scr *ebiten.Image, i int) {
	// fill colors
	pos := maze.pos
	pos[i] += 1
	if _, ok := maze.empty[pos]; !ok { // not empty

	}
	tr := trans[i]
	// draw borders
	for _, line := range mazelines {
		p, q := line[0], line[1]
		pp, qq := tr(p), tr(q)
		vector.StrokeLine(scr, float32(pp[0]), float32(pp[1]), float32(qq[0]), float32(qq[1]), 1, color.White, false)
	}
	// draw labels
	for j, label := range labels[i] {
		newpos := tr(labelpos[j])
		drawText(scr, label, 24, newpos[0], newpos[1])
	}
}

func drawFinish(g *Game, scr *ebiten.Image) {
	drawText(scr, Finish, TitleFontSize, WindowWidth/2, WindowHeight/2-150)
	drawText(scr, FinishMessage, TitleMessageFontSize, WindowWidth/2, WindowHeight/2)
}
