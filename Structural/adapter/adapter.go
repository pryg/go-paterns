package adapter

import "fmt"

var P = fmt.Println

type VectorGraphics interface {
	DrawLine()
	DrawSquare()
}

type RasterGraphics struct {}

func (rg *RasterGraphics) DrawRasterLine() {
	P("Рисуем линию")
}

func (rg *RasterGraphics) DrawRasterSquare() {
	P("Рисуем квадрат")
}

// адаптер типа наследует то, что нужно адаптипровать
type VectorAdapterFromRaster struct {
	RasterGraphics
}

// для нужного интерфейса просто вызываются методы наследика
func (va *VectorAdapterFromRaster) DrawLine() {
	va.DrawRasterLine()
}

func (va *VectorAdapterFromRaster) DrawSquare() {
	va.DrawRasterSquare()
}

// адаптер содержит объект, который нужно адаприторовать
// так лучше если адаптируемый класс содержит ресурсы, а ен только методы
type VectorAdapterFromRaster2 struct {
	rg *RasterGraphics
}

func NewVectorAdapterFromRaster2(rg *RasterGraphics) VectorGraphics {
	vg := new(VectorAdapterFromRaster2)
	vg.rg = rg
	return vg
}

func (va *VectorAdapterFromRaster2) DrawLine() {
	va.rg.DrawRasterLine()
}

func (va *VectorAdapterFromRaster2) DrawSquare() {
	va.rg.DrawRasterSquare()
}