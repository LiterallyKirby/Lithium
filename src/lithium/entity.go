package lithium


import (
"image/color"
)


// Shape is a simple enum for drawable types.
type Shape int


const (
Rect Shape = iota
Circle
)


// Vec2 holds 2D coordinates.
type Vec2 struct{ X, Y float64 }


// Entity is the basic object students will interact with.
type Entity struct {
Name string
Shape Shape
Pos Vec2
Vel Vec2
Scale Vec2
Color color.Color
Gravity bool
Visible bool


OnCollide func(other *Entity)
}


// AddEntity creates an entity and registers it in the world.
func AddEntity(name string, shape Shape, pos Vec2, gravity bool) *Entity {
world.mu.Lock()
defer world.mu.Unlock()
e := &Entity{
Name: name,
Shape: shape,
Pos: pos,
Vel: Vec2{0, 0},
Scale: Vec2{16, 16},
Color: color.White,
Gravity: gravity,
Visible: true,
}
world.Entities[name] = e
return e
}


func GetEntity(name string) *Entity {
world.mu.RLock()
defer world.mu.RUnlock()
return world.Entities[name]
}


func RemoveEntity(name string) {
world.mu.Lock()
defer world.mu.Unlock()
delete(world.Entities, name)
}


// helpers used by students
func Move(e *Entity, dx, dy float64) {
e.Pos.X += dx
e.Pos.Y += dy
}


func Teleport(e *Entity, x, y float64) {
e.Pos.X = x
e.Pos.Y = y
}


func Scale(e *Entity, sx, sy float64) {
e.Scale.X = sx
e.Scale.Y = sy
}


func SetColor(e *Entity, c color.Color) {
e.Color = c
}
