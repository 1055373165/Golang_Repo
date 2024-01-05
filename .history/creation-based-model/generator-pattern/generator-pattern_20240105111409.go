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

func ( )
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}
