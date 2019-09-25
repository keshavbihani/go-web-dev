package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template
var err error
func init(){
	tpl,err = template.ParseFiles("ht.gohtml") // parse blob of files and put a check using MUST
	if err !=nil{
		log.Fatalln(err)
	}
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout,"ht.gohtml",42) //passing 42 as data
	if err !=nil{
		log.Fatalln(err)
	}
}
