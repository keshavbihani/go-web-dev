package main

import "fmt"

type dictionary struct{
	dict map[int]string
}

func (d *dictionary) set(k int,v string){
	if d.dict==nil{
		d.dict = make(map[int]string)
	}
	d.dict[k] = v
}

func (d *dictionary) delete(k int){

	if _,ok:= d.dict[k];ok{
		delete(d.dict,k)
	}
}

func (d *dictionary) has(k int) bool{

	_,ok:= d.dict[k]
	return ok
}

func (d *dictionary) get(k int) string{

	return d.dict[k]
}

func (d *dictionary) clear() {

	d.dict = make(map[int]string)
}

func main(){
	/*
	d:= dictionary{
		map[int]string{},
	}*/
	// or 
	//d:= dictionary{} 
		// or 
	var d dictionary	
	d.set(1,"Keshav")
	d.set(2,"Keshav1")
	d.delete(2)
	d.delete(2)
	fmt.Println(d)	
	fmt.Println(d.has(1))
	fmt.Println(d.get(1))
	//d.clear()
	fmt.Println(len(d.dict))
}