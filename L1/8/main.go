package main

import (
	"fmt"
)

func main() {
	var i int64

	i = 34535

	i = FirstBitChange(i, true)
	BinaryWrite(i)
	i = FirstBitChange(i, false)
	BinaryWrite(i)

}

func FirstBitChange(i int64, b bool) int64 { //т.к первый бит в int64 указывает на то отрицательное число или нет, достаточно умножить число на -1 что бы переключить значение первого бита переменной
	if b {
		if i < 0 {
			return i
		}
		return i * -1
	} else {
		if i < 0 {
			return i * -1
		}
		return i
	}
}

func BinaryWrite(i int64) {
	str := fmt.Sprintf("%064b", i)
	fmt.Println(str, "(len:", len(str), "), (", i, ");")
}
