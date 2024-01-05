package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
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

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Igloo Window"
}
func (i *IglooBuilder) setDoorType() {
	i.doorType = "Igloo Door"
}
func (i *IglooBuilder) setNumFloor() {
	i.floor = 10
}
func (i *IglooBuilder) getHouse() House {
	return House{
		window: i.windowType,
		door:   i.doorType,
		floor:  i.floor,
	}
}

type House struct {
	window string
	door   string
	floor  int
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	IglooBuilder := getBuilder("igloo")
}
