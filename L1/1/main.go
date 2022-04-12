package main

import "fmt"

func main() {

	act := Action{}
	act.Human = Human{name: "Vladimir", age: 24, sex: "male"}

	//использование методов унаследованных от Human в структуре Action (act)
	act.NameInfo()
	act.AgeInfo()
	act.SexInfo()
	act.Walk(5)

}

type Human struct {
	name string
	age  int
	sex  string

	mileage float64 //километраж пройденый Human за всю жизнь
}

//-----Методы типа Human---------
func (h Human) NameInfo() {
	fmt.Println(h.name)
}

func (h Human) AgeInfo() {
	fmt.Println(h.age)
}

func (h Human) SexInfo() {
	fmt.Println(h.sex)
}

//-------------------------------

type Action struct {
	Human //Внедрение структуры Human
}

//-----методы типа Action--------
func (a Action) Walk(distance float64) {
	fmt.Println(a.name, "погулял", distance, "км")
	a.mileage += distance
}

//-------------------------------
