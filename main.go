package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height int
}

func DrawLineSegments(img *ebiten.Image, a, b, c, d Point) {
	if (b.y-a.y)/(b.x-a.x) > 1 {
		for x, y := a.x, a.y; y <= b.y; x, y = x+(b.x-a.x)/(b.y-a.y), y+1 {
			img.Set(int(x), int(y), color.RGBA{0, 255, 0, 255})
		}
	} else {
		for x, y := a.x, a.y; x <= b.x; x, y = x+1, y+(b.y-a.y)/(b.x-a.x) {
			img.Set(int(x), int(y), color.RGBA{0, 255, 0, 255})
		}
	}
	if (d.y-c.y)/(d.x-c.x) > 1 {
		for x, y := c.x, c.y; y <= d.y; x, y = x+(d.x-c.x)/(d.y-c.y), y+1 {
			img.Set(int(x), int(y), color.RGBA{0, 0, 255, 255})
		}
	} else {
		for x, y := c.x, c.y; x <= d.x; x, y = x+1, y+(d.y-c.y)/(d.x-c.x) {
			img.Set(int(x), int(y), color.RGBA{0, 0, 255, 255})
		}
	}
	if (d.y-c.y)*(b.x-a.x)-(d.x-c.x)*(b.y-a.y) != 0 {
		t1 := ((a.y-c.y)*(d.x-c.x) + (c.x-a.x)*(d.y-c.y)) / ((d.y-c.y)*(b.x-a.x) - (d.x-c.x)*(b.y-a.y))
		t2 := ((a.y-c.y)*(b.x-a.x) + (c.x-a.x)*(b.y-a.y)) / ((d.y-c.y)*(b.x-a.x) - (d.x-c.x)*(b.y-a.y))
		if 0 <= t1 && t1 <= 1 && 0 <= t2 && t2 <= 1 {
			ebitenutil.DrawCircle(img, a.x+(b.x-a.x)*t1, a.y+(b.y-a.y)*t1, 2, color.RGBA{255, 0, 0, 255})
		}
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLineSegments(screen, Point{100, 100}, Point{400, 400}, Point{100, 400}, Point{400, 100})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
