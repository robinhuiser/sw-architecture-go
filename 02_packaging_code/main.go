package main

import (
	"fmt"

	"github.com/robinhuiser/sw-architecture-go/02_packaging_code/pkg/zoo"
)

func main() {
	var animals []zoo.Animal

	leo := zoo.Lion{}
	leo.SetName("Leo")
	leo.SetManeLength(10)
	leo.SetLocation(zoo.LatLong{Lat: 10.40, Long: 11.5})
	animals = append(animals, &leo)

	otto := zoo.Pigeon{}
	otto.SetName("Otto")
	otto.SetLocation(zoo.LatLong{Lat: 11.2, Long: 8.6})
	animals = append(animals, &otto)

	makeThemSpeak(animals)
}

func makeThemSpeak(animals []zoo.Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetName() + " says " + animal.Speak())
	}
}
