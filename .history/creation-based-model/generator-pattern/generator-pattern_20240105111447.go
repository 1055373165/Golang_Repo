package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse()
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

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}
