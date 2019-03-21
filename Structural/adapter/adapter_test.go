package adapter

import "testing"

// адаптер через наследование
func TestAdapterExtend(t *testing.T) {
	var vectorGraphics VectorGraphics = new(VectorAdapterFromRaster)
	vectorGraphics.DrawLine()
	vectorGraphics.DrawSquare()
}

// адаптер через композицию
func TestAdapterComposite(t *testing.T) {
	var vectorGraphics VectorGraphics = new(VectorAdapterFromRaster2)
	vectorGraphics.DrawLine()
	vectorGraphics.DrawSquare()
}

// чуть более гибко через конструктор
func TestAdapterConstructor(t *testing.T) {
	var vectorGraphics VectorGraphics = NewVectorAdapterFromRaster2(new(RasterGraphics))
	vectorGraphics.DrawLine()
	vectorGraphics.DrawSquare()
}