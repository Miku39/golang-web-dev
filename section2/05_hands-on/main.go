package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
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

// Breakfast, Lunch, and Dinner
func main() {
	menus := []menu{
		menu{
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
		menu{
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
		menu{
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
	}

	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}
