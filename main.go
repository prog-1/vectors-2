package main

import (
	"fmt"
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

func sub(a, b vector) vector {
	return vector{a.x - b.x, a.y - b.y}
}

func IntersectionPoint(a, v, b, u vector) (vector, bool) {
	if u.x*v.y == u.y*v.x { // are parallel
		return vector{}, false
	}
	t1 := (u.x*(a.y-b.y) + u.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	t2 := (v.x*(a.y-b.y) + v.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	return vector{
		a.x + v.x*t1,
		a.y + v.y*t1,
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
	halfScreenWidth, halfScreenHeight := float64(screenWidth)/2, float64(screenHeight)/2
	a, a1 := vector{-50 + halfScreenWidth, -20 + halfScreenHeight}, vector{40 + halfScreenWidth, 70 + halfScreenHeight}
	b, b1 := vector{70 + halfScreenWidth, 20 + halfScreenHeight}, vector{-70 + halfScreenHeight, 40 + halfScreenWidth}
	// a, a1 := vector{0, 0}, vector{100, 100}
	// b, b1 := vector{0, 100}, vector{150, 0}
	// a, a1 := vector{0 + halfScreenWidth, 0 + halfScreenHeight}, vector{100 + halfScreenWidth, 100 + halfScreenHeight}
	// b, b1 := vector{0 + halfScreenWidth, 100 + halfScreenHeight}, vector{150 + halfScreenWidth, 0 + halfScreenHeight}
	// a, a1 := vector{20, 20}, vector{screenWidth - 20, screenHeight - 20}
	// b, b1 := vector{20, screenHeight - 20}, vector{screenWidth - 20, 20}
	ebitenutil.DrawLine(g.screenBuffer, a.x, a.y, a1.x, a1.y, color.White)
	ebitenutil.DrawLine(g.screenBuffer, b.x, b.y, b1.x, b1.y, color.White)
	if p, _ := IntersectionPoint(a, sub(a1, a), b, sub(b1, b)); true {
		ebitenutil.DebugPrint(g.screenBuffer, fmt.Sprintf("X: %v, Y: %v", p.x, p.y))
		ebitenutil.DrawCircle(g.screenBuffer, float64(p.x), float64(p.y), 10, color.RGBA{0, 255, 255, 255})
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
