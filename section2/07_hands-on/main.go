package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name string
	Menu []TimeMenu
}

type TimeMenu struct {
	Time  string
	Items []item
}

type item struct {
	Name  string
	Price float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []restaurant{
		restaurant{
			Name: "ABC",
			Menu: []TimeMenu{
				TimeMenu{
					Time: "Breakfast",
					Items: []item{
						item{
							Name:  "egg1",
							Price: 1.25,
						},
						item{
							Name:  "egg2",
							Price: 1.50,
						},
					},
				},
				TimeMenu{
					Time: "Lunch",
					Items: []item{
						item{
							Name:  "potato1",
							Price: 5.25,
						},
						item{
							Name:  "potato2",
							Price: 7.50,
						},
					},
				},
			},
		},
		restaurant{
			Name: "123",
			Menu: []TimeMenu{
				TimeMenu{
					Time: "Lunch",
					Items: []item{
						item{
							Name:  "potato1",
							Price: 5.25,
						},
						item{
							Name:  "potato2",
							Price: 7.50,
						},
					},
				},
				TimeMenu{
					Time: "Dinner",
					Items: []item{
						item{
							Name:  "rice1",
							Price: 8.25,
						},
						item{
							Name:  "rice2",
							Price: 9.75,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
