package main

import "fmt"

// myDb implement the dbAccessor interface
type myDb struct {
}

func (myDb) Read() string {
	return "Reading"
}

func (myDb) Write(a int) {
	fmt.Println("Wrting")
}

func main() {

	fnRead := (myDb).Read   // Storing Read the method that can be called with any instance of myDb
	fnWrite := (myDb).Write // Storing Write the method that can be called with any instance of myDb

	// When using method expressio, the first function argument
	// is explicitely instance of the type for which method expression is derieved
	db1 := myDb{}
	fnRead(db1)
	fnWrite(db1, 1)

	//Using inside a callback function
	wrapper(fnRead)

}

// wrapper takes a callbck function with definition func() string as argument
func wrapper(callback func(myDb) string) {
	db1 := myDb{}
	fmt.Println(callback(db1))
}
