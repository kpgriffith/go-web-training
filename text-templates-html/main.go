package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("one.gohtml")
	if(err != nil){
		log.Fatal("Failed creating template from file")
	}

	err = tpl.Execute(os.Stdout, nil)
	if(err != nil){
		log.Fatal("Error writing template to stdout")
	}

}
