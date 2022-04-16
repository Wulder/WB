package main

import (
	"WB/L1/24/points"
	"fmt"
)

func main() {

	p := points.CreatePoint(123, 43)
	p2 := points.CreatePoint(345.3, 24.556)
	fmt.Println(points.GetDistance(p, p2))

	//fmt.Println(p.x) //не можем обратиться к координатам точек

}
