package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse()
}

type NormalBuilder struct {
	windowType string
	doorType string
	num
}