package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Normal Window"
}
func (b *NormalBuilder) setDoorType() {
	b.doorType = "Normal Door"
}
func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}
func (b *NormalBuilder) getHouse() House {
	return House{
		window: b.windowType,
		door:   b.doorType,
		floor:  b.floor,
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Igloo Window"
}
func (i *IglooBuilder) setDoorType() {
	i.doorType = "Igloo Door"
}


type House struct {
	window string
	door   string
	floor  int
}
