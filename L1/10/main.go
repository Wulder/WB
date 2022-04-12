package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	tmpSequence := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	var groups = make(map[int][]float64, len(tmpSequence)) //группы с подмножествами показателей температуры

	sort.Float64s(tmpSequence[:]) //сортировка массива с показателями температуры

	min := tmpSequence[0] //минимальная температура в группе (для отмера шага в 10 градусов)
	currentGroupIndex := 0
	for _, i := range tmpSequence {
		if math.Abs(min-i) < 10 {
			groups[currentGroupIndex] = append(groups[currentGroupIndex], i)
		} else {
			min = i
			currentGroupIndex++
			groups[currentGroupIndex] = append(groups[currentGroupIndex], i)
		}
	}

	for i := 0; i < len(groups); i++ {
		fmt.Println("Группа", i, ":", groups[i])
	}

}
