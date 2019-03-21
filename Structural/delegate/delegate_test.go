package delegate

import "testing"

func TestDelegate(t *testing.T) {
	painter := new(Painter)

	// Методом "мутатором", делегируются полномочия нужному объекту.
	// Затем нужный метод, который на самом деле выполняется делегатом.
	painter.SetGrafics(new(Triangle))
	painter.Draw()
	painter.SetGrafics(new(Square))
	painter.Draw()
	painter.SetGrafics(new(Circle))
	painter.Draw()
	// Удобно когда нужный метод достаточно большой или их несколько - больших.
}
