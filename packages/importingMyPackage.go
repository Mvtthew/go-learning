package main

import (
	"go-learning/packages/myPackage"
	"fmt"
)

func main() {

	caomei := myPackage.NewPerson("Caomei", 19)
	mvtthew := myPackage.NewPerson("Mvtthew", 20)

	earth := myPackage.NewWorld("earth")

	earth.AddToWorld(caomei)
	earth.AddToWorld(mvtthew)

	fmt.Println(earth)

}
