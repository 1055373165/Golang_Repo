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
		
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

type House struct {
	window string
	door   string
	floor  int
}
