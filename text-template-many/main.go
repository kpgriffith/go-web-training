package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init(){
	t = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main(){

	err := t.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if(err != nil){
		log.Fatal("Error writing one template to stdout")
	}

	err = t.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if(err != nil){
		log.Fatal("Error writing two template to stdout")
	}

}
