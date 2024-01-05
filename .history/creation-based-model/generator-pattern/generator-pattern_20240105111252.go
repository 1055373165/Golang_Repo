package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
}

func 