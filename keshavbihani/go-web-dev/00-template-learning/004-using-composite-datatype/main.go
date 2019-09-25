package main

import (
	"text/template"
	"log"
	"os"
)

type car struct{
	Name string
	Price float64
}

type person struct{
	Name string
	Age int
}

var tpl *template.Template
var err error
func init(){
	tpl,err = template.ParseFiles("ht.gohtml") // parse blob of files and put a check using MUST
	if err !=nil{
		log.Fatalln(err)
	}
}

func main() {

	c1 := car{
		"Merz",
		122.31,
	}
	c2 := car{
		"Merz1",
		12.543,
	}
	c3 := car{
		"Merz2",
		111.3214,
	}

	p1 := person{
		"Merz11",
		12,
	}
	p2 := person{
		"Merz22",
		13,
	}
	p3 := person{
		"Merz33",
		24,
	}
	c:= []car{c1,c2,c3}
	p:= []person{p1,p2,p3}
	data := struct{
		Cararr [] car
		Perarr [] person
		}{
		c,
		p,
	}
	err := tpl.ExecuteTemplate(os.Stdout,"ht.gohtml",data) //passing 42 as data
	if err !=nil{
		log.Fatalln(err)
	}
}
