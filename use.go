package main

import "fmt"

type Builder interface {
	Build()
}

type User struct {
	Name string
	Age  int
}

func (b BrickBuilder) Build() {
	fmt.Println("build with brick")
}

func (w WoodBuilder1) Build() {
	fmt.Println("build with wood")
}

type BrickBuilder struct {
	User
}

type WoodBuilder1 struct {
	User
}

type PaperBuilder struct {
	User
}

func (p PaperBuilder) Build() {
	fmt.Println("build with paper")
}

type Building struct {
	Name string
	Builder
}

func main() {
	woodenBuilding := Building{
		Builder: WoodBuilder1{User{
			Name: "garaaa",
			Age:  23,
		}},
		Name: "didar",
	}
	woodenBuilding.Build()
	brickBuilding := Building{
		Builder: BrickBuilder{User{
			Name: "hinata",
			Age:  23,
		}},
		Name: "adlet",
	}
	paperBuilding := Building{
		Name:    "adlet",
		Builder: PaperBuilder{User{Name: "df", Age: 323}},
	}
	brickBuilding.Build()
	paperBuilding.Build()
}
