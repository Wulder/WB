package points

import "math"

type Point struct {
	x, y float64
}

func CreatePoint(x, y float64) Point { //конструктор
	var point Point
	point.x = x
	point.y = y
	return point
}

func GetDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
