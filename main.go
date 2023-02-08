package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Axis struct {
	X int
	Y int
}

type Game struct {
	body      []Axis
	food      Axis
	score     int
	direction int
}

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.direction != dirRight {
			g.direction = dirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.direction != dirLeft {
			g.direction = dirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.direction != dirUp {
			g.direction = dirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.direction != dirDown {
			g.direction = dirUp
		}
	}

	if g.body[0].X == g.food.X && g.body[0].Y == g.food.Y {
		g.body = append(g.body, Axis{X: len(g.body) - 1, Y: len(g.body) - 1})
		g.score++
	}

	for i := len(g.body) - 1; i > 0; i-- {
		g.body[i].X = g.body[i-1].X
		g.body[i].Y = g.body[i-1].Y
	}

	time.Sleep(100)

	switch g.direction {
	case dirLeft:
		g.body[0].X--
	case dirRight:
		g.body[0].X++
	case dirDown:
		g.body[0].Y++
	case dirUp:
		g.body[0].Y--
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x99, 0x00, 0xFF})
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	arcadeFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     24,
		Hinting: font.HintingFull,
	})

	score := fmt.Sprintf("%03d", g.score)
	text.Draw(screen, score, arcadeFont, 280, 15, color.White)
	const grid = 10
	for _, a := range g.body {
		ebitenutil.DrawRect(screen, float64(a.X*grid), float64(a.Y*grid), grid, grid, color.Black)
	}

	ebitenutil.DrawRect(screen, float64(g.food.X*grid), float64(g.food.Y*grid), grid, grid, color.RGBA{0xff, 0, 0, 0xff})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(&Game{body: make([]Axis, 1), food: Axis{X: 50, Y: 100}}); err != nil {
		log.Fatal(err)
	}
}
