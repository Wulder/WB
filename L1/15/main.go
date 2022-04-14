package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(correctSomeFunc(1024))
	fmt.Println(correctSomeFunc(102))
	fmt.Println(correctSomeFunc(1 << 10))
	//fmt.Println(correctSomeFunc(1 << 37)) //вызвать не получится, потому что мы уточнили что createHugeString() принимает int32
}

var justString string //использование общедоступной переменной - нарушение зоны ответственности, ну нас нет необходимости создавать общие переменные для такой функции

func someFunc() {
	v := createHugeString(1 << 10) //передача целочисленного значения путем битового сдвига, функция createHugeString не сможет принять 1 << 33, тк как принимает int который занимает 32 бита, тем самым выпадети ошибка
	justString = v[:100]           //строка обрезанная до 100го символа, а если строка состоит из 50 символов ? Выпадет ошибка, нужно проверять количество элементов в слайсе, прежде чем его обрезать
}

func createHugeString(i int32) string {

	return strings.Repeat("x", int(i)/10)
}

func correctSomeFunc(symbols int32) string {
	fmt.Println(symbols)
	v := createHugeString(symbols)
	var result string
	if len(v) >= 100 {
		result = v[:100]
	} else {
		fmt.Println("string contains fewer symbols than 100")
		return ""
	}

	return result
}
