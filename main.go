package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winWidth  = 1000
	winHeight = 1000
	winTitle  = "intersection"
)

type point struct {
	x, y float64
}

type Game struct {
	width, height int
	a, b, c, d    point
}

func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(x2-x1) <= math.Abs(y2-y1) {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	} else {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(int(x), int(y), c)
		}
	}

}

var c = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      point{100, 100},
		b:      point{900, 900},
		c:      point{100, 900},
		d:      point{900, 100},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, g.a.x, g.a.y, g.b.x, g.b.y, c)
	DrawLineDDA(screen, g.c.x, g.c.y, g.d.x, g.d.y, c)
	if (g.d.y-g.c.y)*(g.b.x-g.a.x)-(g.d.x-g.c.x)*(g.b.y-g.a.y) != 0 {
		t1 := ((g.a.y-g.c.y)*(g.d.x-g.c.x) + (g.c.x-g.a.x)*(g.d.y-g.c.y)) / ((g.d.y-g.c.y)*(g.b.x-g.a.x) - (g.d.x-g.c.x)*(g.b.y-g.a.y))
		t2 := ((g.a.y-g.c.y)*(g.b.x-g.a.x) + (g.c.x-g.a.x)*(g.b.y-g.a.y)) / ((g.d.y-g.c.y)*(g.b.x-g.a.x) - (g.d.x-g.c.x)*(g.b.y-g.a.y))
		if 0 <= t1 && t1 <= 1 && 0 <= t2 && t2 <= 1 {
			ebitenutil.DrawCircle(screen, g.a.x+(g.b.x-g.a.x)*t1, g.a.y+(g.b.y-g.a.y)*t1, 3, color.RGBA{0, 0xFF, 0xFF, 0xFF})
		}
	}
}

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(NewGame(winWidth, winHeight)); err != nil {
		log.Fatal(err)
	}
}
