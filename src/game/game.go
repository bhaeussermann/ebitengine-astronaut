package game

import (
	_ "image/jpeg"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	backgroundImage *ebiten.Image
	guyImage *ebiten.Image
	width, height float64
	guy guy
}

type guy struct {
	width, height float64
	x, y float64
	rotation float64
	speedX, speedY float64
	rotationSpeed float64
}

func NewGame() *Game {
	backgroundImage, _, error := ebitenutil.NewImageFromFile("../images/background.jpeg")
	if error != nil {
		log.Fatal(error)
	}

	guyImage, _, error := ebitenutil.NewImageFromFile("../images/astronaut.png")
	if error != nil {
		log.Fatal(error)
	}

	guyWidth := float64(guyImage.Bounds().Dx())
	guyHeight := float64(guyImage.Bounds().Dy())
	return &Game{
		backgroundImage: backgroundImage,
		guyImage: guyImage,
		guy: guy {
			width: guyWidth,
			height: guyHeight,
			x: guyWidth/2 + rand.Float64()*float64(800-guyWidth),
			y: guyHeight/2 + rand.Float64()*float64(600-guyHeight),
			speedX: float64(rand.Intn(2) * 2 - 1) * (rand.Float64() / 2 + 0.5) * 0.5,
			speedY: float64(rand.Intn(2) * 2 - 1) * (rand.Float64() / 2 + 0.5) * 0.5,
			rotationSpeed: float64(rand.Intn(2) * 2 - 1) * 0.001,
		},
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.width = float64(outsideWidth)
	g.height = float64(outsideHeight)
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	g.moveGuy()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dragBackground(screen)
	g.drawGuy(screen)
}

func (g *Game) moveGuy() {
	if (g.guy.x < g.guy.width / 2) || (g.guy.x > g.width - g.guy.width / 2) {
		g.guy.speedX *= -1

		left := g.guy.x < g.guy.width / 2
		movingUp := g.guy.speedY < 0
		if left == movingUp {
			// Turn anti-clockwise
			g.guy.rotationSpeed = -math.Abs(g.guy.rotationSpeed)
		} else {
			// Turn clockwise
			g.guy.rotationSpeed = math.Abs(g.guy.rotationSpeed)
		}
	}
	if (g.guy.y < g.guy.height/2) || (g.guy.y > g.height - g.guy.height/2) {
		g.guy.speedY *= -1

		top := g.guy.y < g.guy.height/2
		movingToTheRight := g.guy.speedX > 0
		if top == movingToTheRight {
			// Turn anti-clockwise
			g.guy.rotationSpeed = -math.Abs(g.guy.rotationSpeed)
		} else {
			// Turn clockwise
			g.guy.rotationSpeed = math.Abs(g.guy.rotationSpeed)
		}
	}

	g.guy.x += g.guy.speedX
	g.guy.y += g.guy.speedY
	g.guy.rotation += g.guy.rotationSpeed
}

func (g *Game) dragBackground(screen *ebiten.Image) {
  screenWidth := float64(screen.Bounds().Dx())
  screenHeight := float64(screen.Bounds().Dy())
  imageWidth := float64(g.backgroundImage.Bounds().Dx())
  imageHeight := float64(g.backgroundImage.Bounds().Dy())

  horizontalScale := screenWidth / imageWidth
  verticalScale := screenHeight / imageHeight
  scale := math.Max(horizontalScale, verticalScale)

	imageDrawGeom := ebiten.GeoM{}
  imageDrawGeom.Translate(-imageWidth / 2, -imageHeight / 2)
  imageDrawGeom.Scale(scale, scale)
  imageDrawGeom.Translate(screenWidth / 2, screenHeight / 2)
	screen.DrawImage(g.backgroundImage, &ebiten.DrawImageOptions{GeoM: imageDrawGeom})
}

func (g *Game) drawGuy(screen *ebiten.Image) {
	imageDrawGeom := ebiten.GeoM{}
	imageDrawGeom.Translate(-g.guy.width / 2, -g.guy.height / 2)
	imageDrawGeom.Rotate(g.guy.rotation)
	imageDrawGeom.Translate(g.guy.x, g.guy.y)
	screen.DrawImage(g.guyImage, &ebiten.DrawImageOptions{GeoM: imageDrawGeom})
}
