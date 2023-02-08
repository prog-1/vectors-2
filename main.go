package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

type Point struct {
	x, y float64
}

type Parametric struct {
	a Point
	v Point
}

type Line struct {
	a, b Point
}

func (l *Line) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, float64(l.a.x), float64(l.a.y), float64(l.b.x), float64(l.b.y), color.White)
}

func (g *Game) Update() error {
	return nil
}

func ConvertToParametric(line Line) Parametric {
	return Parametric{Point{line.a.x, line.a.y}, Point{line.b.x - line.a.x, line.b.y - line.a.y}}
}

func (a *Parametric) T1(b *Parametric) float64 {
	return ((a.a.y-b.a.y)*b.v.x + (b.a.x-a.a.x)*b.v.y) / (b.v.y*a.v.x - b.v.x*a.v.x)
}

func (a *Parametric) T2(b *Parametric) float64 {
	return ((a.a.y-b.a.y)*a.v.x + (b.a.x-a.a.x)*a.v.y) / (b.v.y*a.v.x - b.v.x*a.v.x)
}

func DrawLineCrossPoint(l1, l2 Line, screen *ebiten.Image) {
	l1.Draw(screen)
	l2.Draw(screen)
	a := ConvertToParametric(l1)
	b := ConvertToParametric(l2)
	T1, T2 := a.T1(&b), a.T2(&b)
	if 0 <= T1 && T1 <= 1 && 0 <= T2 && T2 <= 1 {
		ebitenutil.DrawCircle(screen, (a.a.x + a.v.x*a.T1(&b)), (a.a.y + a.v.y*a.T1(&b)), 3, color.RGBA{255, 0, 0, 255})
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	l1 := Line{Point{0, 0}, Point{100, 100}}
	l2 := Line{Point{0, 100}, Point{150, 0}}
	DrawLineCrossPoint(l1, l2, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
