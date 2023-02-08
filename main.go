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

func DrawCrosspoint(img *ebiten.Image, a, b, c, d Point) {
	ebitenutil.DrawLine(img, a.x, a.y, b.x, b.y, color.RGBA{R: 245, G: 114, B: 227, A: 255})
	ebitenutil.DrawLine(img, c.x, c.y, d.x, d.y, color.RGBA{R: 195, G: 114, B: 245, A: 255})
	t1 := ((a.y-c.y)*(d.x-c.x) + (c.x-a.x)*(d.y-c.y)) / ((d.y-c.y)*(b.x-a.x) - (d.x-c.x)*(b.y-a.y))
	t2 := ((a.y-c.y)*(b.x-a.x) + (c.x-a.x)*(b.y-a.y)) / ((d.y-c.y)*(b.x-a.x) - (d.x-c.x)*(b.y-a.y))
	if 0 <= t1 && t1 <= 1 && 0 <= t2 && t2 <= 1 {
		ebitenutil.DrawCircle(img, a.x+t1*(b.x-a.x), a.y+t1*(b.y-a.y), 1, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawCrosspoint(screen, Point{200, 400}, Point{400, 100}, Point{200, 100}, Point{400, 200})
}
func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
