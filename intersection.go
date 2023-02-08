package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640 //screen width
	sH = 480 //screen height
)

type Game struct {
	//here all the global variables are stored
	width, height int //screen size
	l1, l2        l
}

type v struct { //vector
	x, y int
}

type l struct { //line
	a, b v
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	//Line Draw
	ebitenutil.DrawLine(screen, fl(g.l1.a.x), fl(g.l1.a.y), fl(g.l1.b.x), fl(g.l1.b.y), color.RGBA{255, 100, 100, 255})
	ebitenutil.DrawLine(screen, fl(g.l2.a.x), fl(g.l2.a.y), fl(g.l2.b.x), fl(g.l2.b.y), color.RGBA{100, 255, 100, 255})
}

//-------------------------Functions----------------------------------

func intersectionPoint(l1, l2 l) (p v) {
	a, v, b, u := l1.a, l1.b, l2.a, l2.b

	t1 := (u.x*(a.y-b.y) + u.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	p.x = a.x + v.x*t1
	p.y = a.y + v.y*t1
	return p
}

func fl(v int) float64 { return float64(v) }

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Intersextion Point")
	ebiten.SetWindowResizable(true) //enablening window resizes

	//creating game instance
	g := &Game{width: sW, height: sH,
		l1: l{v{(sW / 2) + 100, (sH / 2) + 50}, v{(sW / 2) - 100, (sH / 2) - 50}},
		l2: l{v{(sW / 2) - 100, (sH / 2) + 100}, v{(sW / 2) + 100, (sH / 2) - 100}}}

	//running game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
