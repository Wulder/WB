package main

import "fmt"

func main() {

	bok := Book{Str: "HIHIHI"}
	Test(&bok)
}

type Stringer interface {
	String(i int) string
}

type Book struct {
	Str string
}

func (b *Book) String(i int) string {
	return b.Str
}

func Test(s Stringer) {
	fmt.Println(s.String(1))
}
