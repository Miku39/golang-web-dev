package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, "is speaking.")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("Agent", sa.lname, "is calling.")
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

	fmt.Println(p1.fname)
	p1.pSpeak()

	fmt.Println(sa1.licenseToKill)
	sa1.saSpeak()
	sa1.pSpeak() // sa1.person.pSpeak()
}
