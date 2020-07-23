package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Bark"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow"
}

type Lion struct {
}

func (l Lion) Speak() string {
	return "Woof"
}
func Hello(i interface{}) interface{} {
	return i
}

func main() {
	// a := []Animal{Dog{}, Cat{}, Lion{}}

	c := Cat{}
	fmt.Println((&c).Speak())
	// for _, animal := range a {
	// 	fmt.Println(animal.Speak())
	// }

	// fmt.Println(Hello("fwqw"))
}
