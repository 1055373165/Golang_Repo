package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse()
}

type NormalBuilder struct {
	window
}