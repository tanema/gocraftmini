package game

// Biom shows the top part of the voxel
type Biom struct {
	*Voxel
}

func newBiom(x, y, z float32) *Biom {
	return &Biom{newVoxel(x, y, z, 1, 0, 0, 0, 0, true)}
}

func (biom *Biom) update(world *World, x, y float32) {
	biom.Voxel.update(world, x, y)
	if biom.z < 0 { // water
		biom.h = 180 - biom.z*20/3 // display depth of water, deeper is darker
		biom.s, biom.l = 99, 0.5
		biom.height = (1.4 + sin(world.timeOfDay*25+biom.y)) / 6 // waves
	} else if biom.z == 0 { // sand
		biom.h, biom.s, biom.l, biom.height = 60, 99, 0.6, 0
	} else if biom.z > 15 { // snow
		biom.h, biom.s, biom.l, biom.height = 0, 0, 0.9, 0
	} else { // grass
		biom.h, biom.s, biom.l, biom.height = 120, 99, 0.3, 0
	}
}
