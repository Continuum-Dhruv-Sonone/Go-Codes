package main

import (
	"fmt"

	"github.com/ContinuumLLC/GO/Interface/db"
)

func main() {

	md := myDb{}           // Instance of myDb that implement the interface
	fmt.Println(md.Read()) // Accesssing the Read method binded to myDb
	md.Write(1)            // Accessing the Write method binded to myDb
}

// myDb implement the dbAccessor interface
type myDb struct {
}

func (myDb) Read() string {
	return "Reading"
}

func (myDb) Write(a int) {
	fmt.Println("Wrting")
}

// NewDb returns the instance of myDb that implements db.DbAccessor
func NewDb() db.DbAccessor {
	return myDb{}
}
