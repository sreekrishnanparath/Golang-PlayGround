package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	address
}

type address struct {
	country string
	zip int
}

func main() {
	//sree := person{firstname:"sreekrishnan"}
	//sree.print()

	//var krish person
	//krish.firstname = "sree"
	//krish.lastname = "krishnan"
	//krish.print()

	hari := person{firstname:"hari", lastname:"krishnan", address:address{"INDIA",680589}}
	//hari = hari.updateName("updateName")
	//hari.updateNameValue("updateName")
	//hari.updateNameValue("updateName")
	//hariPointer := &hari
	hari.updateNameWithPointer("updateName")
	hari.print()
	
	namePointer := &hari
 
 	fmt.Println(&namePointer)
	 printPointer(namePointer)

}

func printPointer(namePointer *person) {
	fmt.Println(&namePointer)
   }

func (p person) updateNameValue(update string){
	p.firstname = update
	fmt.Printf("%+v",p)
}

func (pointerToHari *person) updateNameWithPointer(update string){
	(*pointerToHari).firstname = update
}

func (p person) updateName(update string) person{
	p.firstname = update
	fmt.Printf("%+v",p)
	return p
}
func (p person) print() {
	fmt.Printf("%+v",p)
}