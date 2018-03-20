package game

// Cell describes one position on the grid
type Cell struct {
	biom *Biom
	dirt *Voxel
}

func newCell(x, y int) *Cell {
	return &Cell{
		biom: newBiom(float32(x), float32(y), -9),
		dirt: newVoxel(float32(x), float32(y), 0, 1, 0, 19, 99, 0.2, true),
	}
}

func (cell *Cell) getZ() float32 { return cell.biom.z }
func (cell *Cell) setZ(newZ float32) {
	cell.biom.z, cell.dirt.height = newZ, newZ
}

func (cell *Cell) update(world *World, x, y, distX, distY float32) {
	cell.dirt.update(world, distX, distY)
	cell.biom.update(world, distX, distY)
}

func (cell *Cell) draw(camera *Camera, x, y float32) {
	if cell.dirt.height > 0 {
		cell.dirt.draw(camera, x, y)
	}
	cell.biom.draw(camera, x, y)
}

func generateTerrain(size, iterations int, smooth bool) [][]*Cell {
	terrain := make([][]*Cell, size)
	for x := 0; x < size; x++ {
		terrain[x] = make([]*Cell, size)
		for y := 0; y < size; y++ {
			terrain[x][y] = newCell(x, y)
		}
	}

	for ; iterations >= 0; iterations-- {
		px, py, r := randRange(0, size), randRange(0, size), randRange(10, 40)
		for x := -r; x <= r; x++ {
			for y := -r; y <= r; y++ {
				// Increase altitude of cell cell with "bell" function factor
				cell := terrain[(px+x+size)%size][(py-y+size)%size]
				cell.setZ(cell.getZ() + 5*exp(-(float32(x)*float32(x)+float32(y)*float32(y))/(float32(r)*2)))
			}
		}
	}

	if !smooth {
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				terrain[x][y].setZ(floor(terrain[x][y].getZ()))
			}
		}
	}

	return terrain
}
