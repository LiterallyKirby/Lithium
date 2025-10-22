package lithium


import (
"image/color"
"log"
"sync"


"github.com/hajimehoshi/ebiten/v2"
)


var (
DefaultWindowW = 640
DefaultWindowH = 480
)


// World is the engine runtime.
type World struct {
Entities map[string]*Entity
Gravity float64


updateCallbacks []func()


mu sync.RWMutex
}


var world = &World{
Entities: make(map[string]*Entity),
Gravity: 0.4,
}


// OnUpdate registers a student callback executed each frame.
func OnUpdate(fn func()) {
world.updateCallbacks = append(world.updateCallbacks, fn)
}


// Run starts the Ebiten loop.
func Run() {
ebiten.SetWindowSize(DefaultWindowW, DefaultWindowH)
ebiten.SetWindowTitle("Lithium")
if err := ebiten.RunGame(&engineGame{}); err != nil {
log.Fatal(err)
}
}


// engineGame implements ebiten.Game
type engineGame struct{}


func (g *engineGame) Update() error {
// input state updated first
updateInputState()


// physics and user callbacks
world.mu.Lock()
defer world.mu.Unlock()


// apply gravity & integrate
for _, e := range world.Entities {
if e.Gravity {
e.Vel.Y += world.Gravity
}
e.Pos.X += e.Vel.X
e.Pos.Y += e.Vel.Y
}


// collision pass
checkCollisions()


// user callbacks
for _, fn := range world.updateCallbacks {
fn()
}


return nil
}


func (g *engineGame) Draw(screen *ebiten.Image) {
world.mu.RLock()
defer world.mu.RUnlock()


}
