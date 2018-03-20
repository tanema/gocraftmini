package game

import (
	"github.com/tanema/amore/gfx"
)

// Voxel is the main drawn box
type Voxel struct {
	x, y, z       float32
	width, height float32
	h, s, l       float32
	shine         float32
	relative      bool
}

func newVoxel(x, y, z, width, height, h, s, l float32, relative bool) *Voxel {
	return &Voxel{
		x: x, y: y, z: z,
		width: width, height: height,
		h: h, s: s, l: l,
		relative: relative,
		shine:    1,
	}
}

func (voxel *Voxel) update(world *World, x, y float32) {
	// starting luminance based on distance from camera center
	voxel.shine = 1/20 + exp(-(x*x+y*y)/25*2)
	if world.sin > 0 { // if, daytime add sunlight
		voxel.shine += world.sin * (2 - world.sin) * (1 - voxel.shine)
	}
	if voxel.z < 0 { // if water, add shine
		sunx := float32(world.camera.x) + world.sun.x
		suny := float32(world.camera.x) + world.sun.y
		p := (sunx - suny - x + y)
		voxel.shine += 25 * exp(-p*p)
	}
}

func (voxel *Voxel) draw(camera *Camera, px, py float32) {
	cellSize := float32(1)
	if voxel.relative {
		cellSize = camera.getCellSize()
	}

	gfx.SetColor(HSLToRGB(voxel.h/360, voxel.s/100, pow(voxel.l, 1/voxel.shine)))
	x, y := camera.worldToScreen(px, py, voxel.z, voxel.relative)
	width, height := voxel.width*cellSize, voxel.height*cellSize
	coords := []float32{
		x + width, y, x, y + width/2,
		x - width, y,
		x - width, y - height,
		x, y - height - width/2,
		x + width, y - height,
	}
	gfx.Polygon(gfx.FILL, coords)
}
