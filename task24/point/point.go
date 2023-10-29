package point

import "math"

// Структура точки, содержащая координату x и координату y, который доступны только внутри данного пакета
type Point struct {
	x, y float64
}

// Конструктор новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Геттеры и сеттеры для получения доступа к полям структуры
func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

// Метод возвращающий расстояние между точками
func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
