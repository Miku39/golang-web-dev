package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "is speaking.")
}

func (sa secretAgent) speak() {
	fmt.Println("Agent", sa.lname, "is calling.")
}

func communicate(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Miku",
		"Sano",
	}

	sa1 := secretAgent{
		person{
			"Yaki",
			"Tako",
		},
		false,
	}

	communicate(p1)
	communicate(sa1)
}
