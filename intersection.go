package main

import (
	"fmt"
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
	width, height int
	l1, l2        line
	c             circle
}

type vector struct {
	x, y float64
}

type line struct {
	a, b vector
	clr  color.RGBA
}

type circle struct {
	r   int
	clr color.RGBA
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	//Line Draw
	ebitenutil.DrawLine(screen, fl(g.width/2)+g.l1.a.x, fl(g.height/2)+g.l1.a.y, fl(g.width/2)+g.l1.b.x, fl(g.height/2)+g.l1.b.y, g.l1.clr)
	ebitenutil.DrawLine(screen, fl(g.width/2)+g.l2.a.x, fl(g.height/2)+g.l2.a.y, fl(g.width/2)+g.l2.b.x, fl(g.height/2)+g.l2.b.y, g.l2.clr)
	p, draw := intersectionPoint(g.l1, g.l2)
	fmt.Println(p)
	if draw {
		ebitenutil.DrawCircle(screen, fl(g.width/2)+p.x, fl(g.height/2)+p.y, fl(g.c.r), g.c.clr)
	}
}

//-------------------------Functions----------------------------------

func intersectionPoint(l1, l2 line) (p vector, draw bool) {
	a, v, b, u := l1.a, l1.b, l2.a, l2.b

	t1 := (u.x*(a.y-b.y) + u.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	t2 := (v.x*(a.y-b.y) + v.y*(b.x-a.x)) / (u.y*v.x - u.x*v.y)
	fmt.Println(t1)
	p.x = a.x + t1*(v.x-a.x)
	p.y = a.y + t1*(v.y-a.y)

	//if 0<=t1<=1 and 0 <=t2<=1 => true, draw the circle
	return p, inRange(t1) && inRange(t2)
}

func inRange(v float64) bool {
	return 0 <= v && v <= 1
}

func fl(v int) float64 { return float64(v) }

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Intersextion Point")
	ebiten.SetWindowResizable(true) //enablening window resizes

	//creating game instance
	g := &Game{width: sW, height: sH,
		l1: line{vector{100, 50}, vector{-100, -50}, color.RGBA{255, 100, 100, 255}},
		l2: line{vector{-100, 100}, vector{100, -100}, color.RGBA{100, 255, 100, 255}},
		//l1: line{vector{20, 40}, vector{40, 10}, color.RGBA{255, 100, 100, 255}},
		//l2: line{vector{20, 10}, vector{40, 20}, color.RGBA{100, 255, 100, 255}},
		c: circle{r: 5, clr: color.RGBA{100, 100, 255, 255}}}

	//running game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
