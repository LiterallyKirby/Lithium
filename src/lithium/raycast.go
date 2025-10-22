package lithium

// Raycast performs a simple stepped raycast from start to end and
// returns the first entity hit (naive but fine for demos).
func Raycast(start, end Vec2, steps int) (bool, *Entity) {
	return RaycastIgnore(start, end, steps, nil)
}

// RaycastIgnore performs raycast but ignores a specific entity
func RaycastIgnore(start, end Vec2, steps int, ignore *Entity) (bool, *Entity) {
	if steps <= 0 {
		steps = 100
	}
	dx := (end.X - start.X) / float64(steps)
	dy := (end.Y - start.Y) / float64(steps)
	x := start.X
	y := start.Y

	world.mu.RLock()
	defer world.mu.RUnlock()

	for i := 0; i <= steps; i++ {
		for _, e := range world.Entities {
			if e == ignore || !e.Visible {
				continue
			}
			if pointInsideEntity(x, y, e) {
				return true, e
			}
		}
		x += dx
		y += dy
	}
	return false, nil
}

// RaycastAll returns all entities hit by the ray
func RaycastAll(start, end Vec2, steps int) []*Entity {
	if steps <= 0 {
		steps = 100
	}
	dx := (end.X - start.X) / float64(steps)
	dy := (end.Y - start.Y) / float64(steps)
	x := start.X
	y := start.Y

	world.mu.RLock()
	defer world.mu.RUnlock()

	hitSet := make(map[*Entity]bool)
	var hits []*Entity

	for i := 0; i <= steps; i++ {
		for _, e := range world.Entities {
			if !e.Visible {
				continue
			}
			if pointInsideEntity(x, y, e) && !hitSet[e] {
				hitSet[e] = true
				hits = append(hits, e)
			}
		}
		x += dx
		y += dy
	}
	return hits
}

func pointInsideEntity(x, y float64, e *Entity) bool {
	return x >= e.Pos.X && x <= e.Pos.X+e.Scale.X && y >= e.Pos.Y && y <= e.Pos.Y+e.Scale.Y
}

// LineIntersectsEntity checks if a line segment intersects an entity
func LineIntersectsEntity(start, end Vec2, e *Entity) bool {
	// Simple AABB line intersection using raycast
	hit, _ := Raycast(start, end, 50)
	return hit
}
