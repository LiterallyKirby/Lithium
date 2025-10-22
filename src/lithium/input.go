package lithium

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// Key constants exported for students
var (
	KeyLeft  = ebiten.KeyArrowLeft
	KeyRight = ebiten.KeyArrowRight
	KeyUp    = ebiten.KeyArrowUp
	KeyDown  = ebiten.KeyArrowDown
	KeySpace = ebiten.KeySpace
	KeyA     = ebiten.KeyA
	KeyW     = ebiten.KeyW
	KeyS     = ebiten.KeyS
	KeyD     = ebiten.KeyD
	KeyE     = ebiten.KeyE
	KeyQ     = ebiten.KeyQ
	KeyR     = ebiten.KeyR
	KeyF     = ebiten.KeyF
	KeyEscape = ebiten.KeyEscape
	KeyEnter  = ebiten.KeyEnter
	KeyShift  = ebiten.KeyShift
	KeyCtrl   = ebiten.KeyControl
)

var prevKeys = map[ebiten.Key]bool{}
var prevMouseButtons = map[ebiten.MouseButton]bool{}

func updateInputState() {
	// Update keyboard state
	for k := range prevKeys {
		prevKeys[k] = ebiten.IsKeyPressed(k)
	}

	// Update mouse button state
	for b := range prevMouseButtons {
		prevMouseButtons[b] = ebiten.IsMouseButtonPressed(b)
	}
}

func IsPressed(k ebiten.Key) bool {
	return ebiten.IsKeyPressed(k)
}

func JustPressed(k ebiten.Key) bool {
	cur := ebiten.IsKeyPressed(k)
	prev := prevKeys[k]
	if cur && !prev {
		prevKeys[k] = true
		return true
	}
	if !cur {
		prevKeys[k] = false
	}
	return false
}

func JustReleased(k ebiten.Key) bool {
	cur := ebiten.IsKeyPressed(k)
	prev := prevKeys[k]
	return !cur && prev
}

// Mouse input functions
func GetMousePosition() (int, int) {
	return ebiten.CursorPosition()
}

func GetMousePositionF() Vec2 {
	x, y := ebiten.CursorPosition()
	return Vec2{float64(x), float64(y)}
}

func IsMouseButtonPressed(button ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(button)
}

func JustClickedMouse(button ebiten.MouseButton) bool {
	cur := ebiten.IsMouseButtonPressed(button)
	prev := prevMouseButtons[button]
	if cur && !prev {
		prevMouseButtons[button] = true
		return true
	}
	if !cur {
		prevMouseButtons[button] = false
	}
	return false
}

var (
	MouseButtonLeft   = ebiten.MouseButtonLeft
	MouseButtonRight  = ebiten.MouseButtonRight
	MouseButtonMiddle = ebiten.MouseButtonMiddle
)

// GetMouseWheel returns the scroll wheel offset
func GetMouseWheel() (float64, float64) {
	x, y := ebiten.Wheel()
	return x, y
}
