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

	db := myDb{} // Creating the instance of myDb

	fnRead := db.Read   // Storing Read the method binded to db variable(instance of myDb)
	fnWrite := db.Write // Storing Write the method binded to db variable(instance of myDb)

	fnRead()
	fnWrite(1)

	//Using inside a callback function
	wrapper(fnRead)

}

// wrapper takes a callbck function with definition func() string as argument
func wrapper(callback func() string) {
	fmt.Println(callback())
}
