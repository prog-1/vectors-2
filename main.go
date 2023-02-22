package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type vector struct {
	x, y float64
}

type line struct {
	a, v vector
}

func (l *line) draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, float64(l.a.x)+screenWidth/2, float64(l.a.y)+screenHeight/2, float64(l.v.x)+screenWidth/2, float64(l.v.y)+screenHeight/2, color.White)
}

func IntersectionPoint(l1, l2 *line) (vector, bool) {
	a, v, b, u := l1.a, l1.v, l2.a, l2.v
	if u.x*v.y == u.y*v.x { // are parallel
		return vector{}, false
	}
	t1 := (u.x*(a.y-b.y) + u.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	// t2 := (a.x + v.x*t1 - b.x) / u.x
	t2 := (v.x*(a.y-b.y) + v.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	return vector{
		a.x + (v.x-a.x)*t1,
		a.y + (v.y-a.y)*t1,
	}, (t1 >= 0 && t1 <= 1 && t2 >= 0 && t2 <= 1)
}

type game struct {
	screenBuffer *ebiten.Image
}

func NewGame() *game {
	return &game{
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.screenBuffer, &ebiten.DrawImageOptions{})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()
	l1 := line{vector{-50, -20}, vector{40, 70}}
	l2 := line{vector{70, 20}, vector{-70, 40}}
	l1.draw(g.screenBuffer)
	l2.draw(g.screenBuffer)
	if p, _ := IntersectionPoint(&l1, &l2); true {
		ebitenutil.DrawCircle(g.screenBuffer, float64(p.x)+screenWidth/2, float64(p.y)+screenHeight/2, 10, color.RGBA{0, 255, 255, 255})
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
