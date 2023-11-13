package valueObjects

import "math/rand"

type Coordinate struct {
	X, Y, Z float64
}

type Temparature struct {
	Value float64
	Scale string
}

type DataOuputRate struct {
	Value  int
	Format string
}

type Ocean3D struct {
	MinX, MaxX float64
	MinY, MaxY float64
	MinZ, MaxZ float64
}

type Transparency int

type Id int

func NewOcean3D(minX, maxX, minY, maxY, minZ, maxZ float64) Ocean3D {
	return Ocean3D{
		MinX: minX, MaxX: maxX,
		MinY: minY, MaxY: maxY,
		MinZ: minZ, MaxZ: maxZ,
	}
}

func (ocean Ocean3D) GetRandomCoordinates3D() Coordinate {
	x := rand.Float64()*(ocean.MaxX-ocean.MinX) + ocean.MinX
	y := rand.Float64()*(ocean.MaxY-ocean.MinY) + ocean.MinY
	z := rand.Float64()*(ocean.MaxZ-ocean.MinZ) + ocean.MinZ
	return Coordinate{X: x, Y: y, Z: z}
}

func NewDataOutputRate(format string) DataOuputRate {
	return DataOuputRate{Value: rand.Intn(59), Format: format}
}
