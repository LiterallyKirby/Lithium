package lithium

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

// basic colors available to students
var (
	Red     = color.RGBA{R: 220, G: 50, B: 50, A: 255}
	Green   = color.RGBA{R: 50, G: 200, B: 50, A: 255}
	Blue    = color.RGBA{R: 50, G: 100, B: 220, A: 255}
	White   = color.RGBA{255, 255, 255, 255}
	Black   = color.RGBA{0, 0, 0, 255}
	Yellow  = color.RGBA{240, 220, 70, 255}
	Cyan    = color.RGBA{80, 200, 200, 255}
	Magenta = color.RGBA{200, 50, 200, 255}
	Orange  = color.RGBA{255, 165, 0, 255}
	Purple  = color.RGBA{128, 0, 128, 255}
	Pink    = color.RGBA{255, 192, 203, 255}
	Gray    = color.RGBA{128, 128, 128, 255}
)

func drawEntities(screen *ebiten.Image) {
	world.mu.RLock()
	defer world.mu.RUnlock()
	
	for _, e := range world.Entities {
		if !e.Visible {
			continue
		}

		switch e.Shape {
		case Rect:
			drawRect(screen, e)
		case Circle:
			drawCircle(screen, e)
		}
	}
}

func drawRect(screen *ebiten.Image, e *Entity) {
	if e.Rotation == 0 {
		// Fast path for non-rotated rectangles
		ebitenutil.DrawRect(screen, e.Pos.X, e.Pos.Y, e.Scale.X, e.Scale.Y, e.Color)
	} else {
		// Draw rotated rectangle using vector
		cx := e.Pos.X + e.Scale.X/2
		cy := e.Pos.Y + e.Scale.Y/2

		// Calculate corner positions
		hw := e.Scale.X / 2
		hh := e.Scale.Y / 2

		cos := float32(cosine(float64(e.Rotation)))
		sin := float32(sine(float64(e.Rotation)))

		// Rotate corners around center
		corners := []Vec2{
			{-hw, -hh},
			{hw, -hh},
			{hw, hh},
			{-hw, hh},
		}

		var path vector.Path
		for i, corner := range corners {
			rx := float32(corner.X)*cos - float32(corner.Y)*sin
			ry := float32(corner.X)*sin + float32(corner.Y)*cos

			x := float32(cx) + rx
			y := float32(cy) + ry

			if i == 0 {
				path.MoveTo(x, y)
			} else {
				path.LineTo(x, y)
			}
		}
		path.Close()

		vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)
		for i := range vertices {
			vertices[i].ColorR = 1
			vertices[i].ColorG = 1
			vertices[i].ColorB = 1
			vertices[i].ColorA = 1
		}

		op := &ebiten.DrawTrianglesOptions{}
		op.ColorM.ScaleWithColor(e.Color)
		screen.DrawTriangles(vertices, indices, emptySubImage, op)
	}
}

func drawCircle(screen *ebiten.Image, e *Entity) {
	cx := float32(e.Pos.X + e.Scale.X/2)
	cy := float32(e.Pos.Y + e.Scale.Y/2)
	radius := float32(min(e.Scale.X, e.Scale.Y) / 2)

	vector.DrawFilledCircle(screen, cx, cy, radius, e.Color, false)
}

// DrawText is a small helper for debugging or scoreboard
func DrawText(screen *ebiten.Image, s string, x, y float64) {
	ebitenutil.DebugPrintAt(screen, s, int(x), int(y))
}

// DrawLine draws a line between two points
func DrawLine(screen *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 2, c, false)
}

// DrawCircleOutline draws a circle outline
func DrawCircleOutline(screen *ebiten.Image, x, y, radius float64, c color.Color) {
	vector.StrokeCircle(screen, float32(x), float32(y), float32(radius), 2, c, false)
}

// FillScreen fills the screen with a color
func FillScreen(screen *ebiten.Image, c color.Color) {
	screen.Fill(c)
}

// SetBackgroundColor sets the background color for the next frame
func SetBackgroundColor(c color.Color) {
	world.mu.Lock()
	defer world.mu.Unlock()
	world.BackgroundColor = c
}

// GetFPS returns the current frames per second
func GetFPS() float64 {
	return ebiten.ActualFPS()
}

// DrawDebugInfo draws FPS and entity count
func DrawDebugInfo(screen *ebiten.Image) {
	world.mu.RLock()
	entityCount := len(world.Entities)
	world.mu.RUnlock()

	info := fmt.Sprintf("FPS: %.1f | Entities: %d", GetFPS(), entityCount)
	ebitenutil.DebugPrint(screen, info)
}

var emptySubImage = ebiten.NewImage(3, 3)

func init() {
	emptySubImage.Fill(color.White)
}
