package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	tpl,err := template.ParseFiles("../ht.gohtml")// one path down ../  
	// tpl is like container which contains many files
	if err !=nil{
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout,nil) // os.stdout because it is of file type and file is of writer interface type
	// Execute: it execute first file in the tpl container
	// ExecuteTemplate: it executes the file whose name is provided 
	if err !=nil{
		log.Fatalln(err)
	}
}
