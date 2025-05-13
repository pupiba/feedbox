package main

import (
	"./internals/model"
	"./internals/storage"
	"./internals/transport"
	"fmt"
)

type Rule interface {
	move()
	stop()
}

type Transport struct {
	name  string
	speed int
}

type AirTransport struct {
	name  string
	speed int
}

func (t Transport) move() {
	fmt.Printf("%v едет со скоростью %v км/ч\n", t.name, t.speed)
}

func (t Transport) stop() {
	fmt.Printf("%v остановился\n", t.name)
}

func (t AirTransport) move() {
	fmt.Printf("%v летит со скоростью %v км/ч\n", t.name, t.speed)
}

func (t AirTransport) stop() {
	fmt.Printf("%v приземлился\n", t.name)

}

func main() {
	v1 := Transport{name: "Car1", speed: 60}
	v2 := Transport{name: "Car2", speed: 120}
	var v3 = AirTransport{name: "Plane", speed: 470}

	var arr = [...]Rule{v1, v2, v3}

	for _, val := range arr {
		val.move()
	}

	for _, val := range arr {
		val.stop()
	}
}
