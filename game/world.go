package game

import (
	//"github.com/tanema/amore/gfx"
	"github.com/tanema/amore/keyboard"
)

// World encapsulates the whole environment
type World struct {
	size      int
	terrain   [][]*Cell
	camera    *Camera
	timeOfDay float32
	sin       float32
	skybox    *Voxel
	sun       *Sun
	player    *Voxel
}

// NewWorld generates a new world to render
func NewWorld(worldSize, visible, iterations int, smooth bool) *World {
	screenSize := float32(1200)
	return &World{
		size:    worldSize,
		terrain: generateTerrain(worldSize, iterations, smooth),
		camera:  newCamera(screenSize, visible),
		skybox:  newVoxel(0, 0, 0, screenSize/2, screenSize/2, 200, 100, 0.9, false),
		sun:     newSun(),
		player:  newVoxel(0, 0, 0, 1, 1, 0, 99, 0.5, true),
	}
}

func (world *World) getCell(x, y int) *Cell {
	i := (x + world.size) % world.size
	j := (y + world.size) % world.size
	return world.terrain[i][j]
}

// Update updates  a step in the world
func (world *World) Update(dt float32) {
	world.updateInput()
	world.timeOfDay += dt / 10
	world.sin = sin(world.timeOfDay)
	world.sun.update(world)
	world.skybox.update(world, float32(world.size), float32(world.size))
	world.camera.forVisible(world, func(cell *Cell, x, y, distX, distY float32) {
		cell.update(world, x, y, distX, distY)
	})
}

// Draw draws one frame
func (world *World) Draw() {
	world.skybox.draw(world.camera, world.skybox.x, world.skybox.y)
	world.sun.draw(world.camera, world.sun.x, world.sun.y)
	world.camera.forVisible(world, func(cell *Cell, x, y, distX, distY float32) {
		cell.draw(world.camera, x, y)
		if x == world.player.x && y == world.player.y {
			world.player.draw(world.camera, x, y)
		}
	})
}

func (world *World) updateInput() {
	if keyboard.IsDown(keyboard.KeyLeft) {
		world.player.x--
	} else if keyboard.IsDown(keyboard.KeyRight) {
		world.player.x++
	}
	if keyboard.IsDown(keyboard.KeyUp) {
		world.player.y--
	} else if keyboard.IsDown(keyboard.KeyDown) {
		world.player.y++
	}

	world.player.x = float32((int(world.player.x) + world.size) % world.size)
	world.player.y = float32((int(world.player.y) + world.size) % world.size)
	world.camera.lookAt(world.player.x, world.player.y)
	cell := world.getCell(int(world.player.x), int(world.player.y))

	// [Space] adds block, [C] removes block/digs soil
	if keyboard.IsDown(keyboard.KeySpace) {
		cell.setZ(cell.getZ() + 1)
	} else if keyboard.IsDown(keyboard.KeyC) {
		cell.setZ(cell.getZ() - 1)
	}

	world.player.z = cell.getZ()

	// [[]/[]] decrease/increase visible size
	if keyboard.IsDown(keyboard.KeyV) {
		world.camera.visible++
	} else if keyboard.IsDown(keyboard.KeyB) {
		world.camera.visible--
	}
}
