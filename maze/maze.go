package maze

import "github.com/hajimehoshi/ebiten/v2"

type Maze struct {
	length        int
	empty         map[[4]int]struct{}
	start, finish [4]int
	pos           [4]int
}

const (
	MazeLength  = 20
	innerWidth  = 400
	innerHeight = 200
	outerWidth  = 600
	outerHeight = 400
)

var (
	maze = Maze{}
)

func init() {
	maze.length = MazeLength
	maze.empty = map[[4]int]struct{}{}
	for i := 0; i < maze.length; i++ {
		var new [4]int
		new[0] = i
		maze.empty[new] = struct{}{}
	}
	maze.start = [4]int{0, 0, 0, 0}
	maze.finish = [4]int{
		maze.length - 1,
		0,
		0,
		0,
	}
	maze.pos = maze.start
}

var (
	labels = [4][6]string{
		{"Xp", "Xn", "Yp", "Yn", "Zp", "Zn"},
		{"Yp", "Yn", "Zp", "Zn", "Wp", "Wn"},
		{"Zp", "Zn", "Wp", "Wn", "Xp", "Xn"},
		{"Wp", "Wn", "Xp", "Xn", "Yp", "Yn"},
	}
	labelpos = [6][2]float64{
		{WindowWidth / 2, WindowHeight / 2},
		{WindowWidth / 2, 50},
		{150, WindowHeight / 2},
		{WindowWidth - 150, WindowHeight / 2},
		{WindowWidth / 2, 150},
		{WindowWidth / 2, WindowHeight - 150},
	}
	trans = [4]func([2]float64) [2]float64{
		func(xy [2]float64) [2]float64 {
			x, y := xy[0], xy[1]
			return [2]float64{x / 2, y / 2}
		},
		func(xy [2]float64) [2]float64 {
			x, y := xy[0], xy[1]
			return [2]float64{x/2 + WindowWidth/2, y / 2}
		},
		func(xy [2]float64) [2]float64 {
			x, y := xy[0], xy[1]
			return [2]float64{x / 2, y/2 + WindowHeight/2}
		},
		func(xy [2]float64) [2]float64 {
			x, y := xy[0], xy[1]
			return [2]float64{x/2 + WindowWidth/2, y/2 + WindowHeight/2}
		},
	}
	mazelines = [][2][2]float64{
		{{WindowWidth/2 - innerWidth/2, WindowHeight/2 - innerHeight/2}, {WindowWidth/2 + innerWidth/2, WindowHeight/2 - innerHeight/2}},
		{{WindowWidth/2 + innerWidth/2, WindowHeight/2 - innerHeight/2}, {WindowWidth/2 + innerWidth/2, WindowHeight/2 + innerHeight/2}},
		{{WindowWidth/2 + innerWidth/2, WindowHeight/2 + innerHeight/2}, {WindowWidth/2 - innerWidth/2, WindowHeight/2 + innerHeight/2}},
		{{WindowWidth/2 - innerWidth/2, WindowHeight/2 + innerHeight/2}, {WindowWidth/2 - innerWidth/2, WindowHeight/2 - innerHeight/2}},

		{{WindowWidth/2 - innerWidth/2, WindowHeight/2 - innerHeight/2}, {WindowWidth/2 - outerWidth/2, WindowHeight/2 - outerHeight/2}},
		{{WindowWidth/2 + innerWidth/2, WindowHeight/2 - innerHeight/2}, {WindowWidth/2 + outerWidth/2, WindowHeight/2 - outerHeight/2}},
		{{WindowWidth/2 + innerWidth/2, WindowHeight/2 + innerHeight/2}, {WindowWidth/2 + outerWidth/2, WindowHeight/2 + outerHeight/2}},
		{{WindowWidth/2 - innerWidth/2, WindowHeight/2 + innerHeight/2}, {WindowWidth/2 - outerWidth/2, WindowHeight/2 + outerHeight/2}},

		{{WindowWidth/2 - outerWidth/2, WindowHeight/2 - outerHeight/2}, {WindowWidth/2 + outerWidth/2, WindowHeight/2 - outerHeight/2}},
		{{WindowWidth/2 + outerWidth/2, WindowHeight/2 - outerHeight/2}, {WindowWidth/2 + outerWidth/2, WindowHeight/2 + outerHeight/2}},
		{{WindowWidth/2 + outerWidth/2, WindowHeight/2 + outerHeight/2}, {WindowWidth/2 - outerWidth/2, WindowHeight/2 + outerHeight/2}},
		{{WindowWidth/2 - outerWidth/2, WindowHeight/2 + outerHeight/2}, {WindowWidth/2 - outerWidth/2, WindowHeight/2 - outerHeight/2}},
	}
)

var (
	wallImage *ebiten.Image
)

func init() {
	wallImage = ebiten.NewImage(1, 1)
}
