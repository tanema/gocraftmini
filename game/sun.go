package game

// Sun is the sun!
type Sun struct {
	*Voxel
}

func newSun() *Sun {
	return &Sun{newVoxel(0, 0, 0, 20, 20, 60, 0, 1, false)}
}

func (sun *Sun) update(world *World) {
	visible := -float32(world.camera.visible)
	if world.sin > 0 {
		visible = float32(world.camera.visible)
	}

	sun.y = -1 * visible * cos(world.timeOfDay)
	sun.z = abs(world.sin * visible * 3.2)
}
