package lithium

import (
	"image/color"
	"log"
	"sync"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var (
	DefaultWindowW = 640
	DefaultWindowH = 480
)

// World is the engine runtime.
type World struct {
	Entities        map[string]*Entity
	Gravity         float64
	BackgroundColor color.Color

	updateCallbacks []func()
	drawCallbacks   []func(screen *ebiten.Image)

	mu sync.RWMutex
}

var world = &World{
	Entities:        make(map[string]*Entity),
	Gravity:         0.4,
	BackgroundColor: color.RGBA{20, 20, 30, 255},
}

// OnUpdate registers a student callback executed each frame.
func OnUpdate(fn func()) {
	world.updateCallbacks = append(world.updateCallbacks, fn)
}

// OnDraw registers a callback for custom drawing
func OnDraw(fn func(screen *ebiten.Image)) {
	world.drawCallbacks = append(world.drawCallbacks, fn)
}

// Run starts the Ebiten loop.
func Run() {
	ebiten.SetWindowSize(DefaultWindowW, DefaultWindowH)
	ebiten.SetWindowTitle("Lithium")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&engineGame{}); err != nil {
		log.Fatal(err)
	}
}

// SetWindowSize changes the window dimensions
func SetWindowSize(w, h int) {
	DefaultWindowW = w
	DefaultWindowH = h
	ebiten.SetWindowSize(w, h)
}

// GetWindowSize returns current window dimensions
func GetWindowSize() (int, int) {
	return DefaultWindowW, DefaultWindowH
}

// engineGame implements ebiten.Game
type engineGame struct {
	frameCount int
}

func (g *engineGame) Update() error {
	// input state updated first
	updateInputState()

	// physics and user callbacks
	world.mu.Lock()

	// apply gravity & integrate
	for _, e := range world.Entities {
		if e.Gravity {
			e.Vel.Y += world.Gravity
		}

		// Clamp velocity to prevent tunneling
		ClampVelocity(e, DefaultPhysics.MaxVelocity)

		e.Pos.X += e.Vel.X
		e.Pos.Y += e.Vel.Y

		// Call per-entity update callback
		if e.OnUpdate != nil {
			e.OnUpdate(e)
		}
	}

	// collision pass
	checkCollisions()

	world.mu.Unlock()

	// user callbacks (outside lock to prevent deadlocks)
	for _, fn := range world.updateCallbacks {
		fn()
	}

	g.frameCount++
	return nil
}

func (g *engineGame) Draw(screen *ebiten.Image) {
	world.mu.RLock()
	bgColor := world.BackgroundColor
	world.mu.RUnlock()

	// Fill background
	screen.Fill(bgColor)

	// Draw all entities (with proper locking inside drawEntities)
	drawEntities(screen)

	// Custom draw callbacks
	world.mu.RLock()
	callbacks := world.drawCallbacks
	world.mu.RUnlock()
	
	for _, fn := range callbacks {
		fn(screen)
	}
}

func (g *engineGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return DefaultWindowW, DefaultWindowH
}

// GetFrameCount returns the current frame number
func GetFrameCount() int {
	return 0 // Would need to expose this from engineGame
}
