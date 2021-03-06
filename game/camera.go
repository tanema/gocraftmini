package game

// Camera describes the views
type Camera struct {
	x, y    int
	size    float32
	half    float32
	visible int
}

func newCamera(screenWidth float32, visible int) *Camera {
	return &Camera{
		size:    screenWidth,
		half:    screenWidth / 2,
		visible: visible,
	}
}

func (camera *Camera) getCellSize() float32 {
	return camera.half / (2*float32(camera.visible) + 1)
}

func (camera *Camera) lookAt(x, y float32) {
	camera.x, camera.y = int(x), int(y)
}

func (camera *Camera) forVisible(world *World, fn func(cell *Cell, x, y, dx, dy float32)) {
	for distX := -camera.visible; distX <= camera.visible; distX++ {
		for distY := -camera.visible; distY <= camera.visible; distY++ {
			x, y := camera.x+distX, camera.y+distY
			fn(world.getCell(x, y), float32(x), float32(y), float32(distX), float32(distY))
		}
	}
}

func (camera *Camera) worldToScreen(x, y, z float32, relative bool) (float32, float32) {
	if relative {
		x, y = x-float32(camera.x), y-float32(camera.y)
	}
	cellSize := camera.getCellSize()
	return camera.half + (x-y)*cellSize, camera.half*3/2 + (x+y-(max(z, 0)*2))*cellSize/2
}
