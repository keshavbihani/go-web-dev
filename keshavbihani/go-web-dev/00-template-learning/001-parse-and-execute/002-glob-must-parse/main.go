package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html")) // parse blob of files and put a check using MUST
	// must do error handling by itself
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout,"1.html",nil) // os.stdout because it is of file type and file is of writer interface type
	if err !=nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout,"2.html",nil) // os.stdout because it is of file type and file is of writer interface type
	if err !=nil{
		log.Fatalln(err)
	}
}
