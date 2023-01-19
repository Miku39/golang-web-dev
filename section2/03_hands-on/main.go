package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// Southern, Central, Northern
func main() {
	hotels := []hotel{
		hotel{
			Name:    "Southern Hotel",
			Address: "ABC street",
			City:    "ABC",
			Zip:     "123",
			Region:  "Southern",
		},
		hotel{
			Name:    "Central Hotel",
			Address: "DEF street",
			City:    "DEF",
			Zip:     "456",
			Region:  "Central",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
