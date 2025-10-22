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

// Add returns a new Vec2 that is the sum of v and other
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// Sub returns a new Vec2 that is the difference of v and other
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

// Scale returns a new Vec2 scaled by factor
func (v Vec2) Scale(factor float64) Vec2 {
	return Vec2{v.X * factor, v.Y * factor}
}

// Length returns the magnitude of the vector
func (v Vec2) Length() float64 {
	return sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a unit vector in the same direction
func (v Vec2) Normalize() Vec2 {
	l := v.Length()
	if l == 0 {
		return Vec2{0, 0}
	}
	return Vec2{v.X / l, v.Y / l}
}

// Entity is the basic object students will interact with.
type Entity struct {
	Name     string
	Shape    Shape
	Pos      Vec2
	Vel      Vec2
	Scale    Vec2
	Color    color.Color
	Gravity  bool
	Visible  bool
	Solid    bool // whether this entity blocks other entities
	Rotation float64

	OnCollide func(other *Entity)
	OnUpdate  func(e *Entity)
}

// AddEntity creates an entity and registers it in the world.
func AddEntity(name string, shape Shape, pos Vec2, gravity bool) *Entity {
	world.mu.Lock()
	defer world.mu.Unlock()
	e := &Entity{
		Name:     name,
		Shape:    shape,
		Pos:      pos,
		Vel:      Vec2{0, 0},
		Scale:    Vec2{16, 16},
		Color:    color.White,
		Gravity:  gravity,
		Visible:  true,
		Solid:    true,
		Rotation: 0,
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

func GetAllEntities() []*Entity {
	world.mu.RLock()
	defer world.mu.RUnlock()
	entities := make([]*Entity, 0, len(world.Entities))
	for _, e := range world.Entities {
		entities = append(entities, e)
	}
	return entities
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

func SetVelocity(e *Entity, vx, vy float64) {
	e.Vel.X = vx
	e.Vel.Y = vy
}

func GetVelocity(e *Entity) Vec2 {
	return e.Vel
}

func SetRotation(e *Entity, angle float64) {
	e.Rotation = angle
}

func Rotate(e *Entity, delta float64) {
	e.Rotation += delta
}

// GetCenter returns the center point of the entity
func GetCenter(e *Entity) Vec2 {
	return Vec2{
		X: e.Pos.X + e.Scale.X/2,
		Y: e.Pos.Y + e.Scale.Y/2,
	}
}

// GetBounds returns the bounding box of the entity
func GetBounds(e *Entity) (x1, y1, x2, y2 float64) {
	return e.Pos.X, e.Pos.Y, e.Pos.X + e.Scale.X, e.Pos.Y + e.Scale.Y
}

// DistanceTo returns the distance between two entities' centers
func DistanceTo(e1, e2 *Entity) float64 {
	c1 := GetCenter(e1)
	c2 := GetCenter(e2)
	return c1.Sub(c2).Length()
}

// IsOverlapping checks if two entities overlap
func IsOverlapping(e1, e2 *Entity) bool {
	return aabbOverlap(e1, e2)
}
