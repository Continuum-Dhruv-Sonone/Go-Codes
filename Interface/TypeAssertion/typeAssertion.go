package main

import (
	"fmt"

	"github.com/ContinuumLLC/GO/Interface/db"
)

func main() {

	db1 := NewDb1() // Gettig instance of myDb1

	// Type casting db.DbAccessor as myDb1 type
	// Will cause panic if the db1 is not of type myDb1
	// Wrong way of extracting values from Interface
	db := db1.(myDb1)

	//Correcting way of checking type assertion

	if db, ok := db1.(myDb1); !ok {
		fmt.Println("Type Assertion Failed....Expected type was myDb1")
	}

	// Type casting fails here and will cause panic
	// db1 is of type myDb1 and type casting it done for myDb2
	db2 := db1.(myDb2)

	fmt.Println(db)
	fmt.Println(db2)

}

// myDb1 implement the dbAccessor interface
type myDb1 struct {
}

// myDb2 implement the dbAccessor interface
type myDb2 struct {
}

func (myDb1) Read() string {
	return "Reading myDb1"
}

func (myDb1) Write(a int) {
	fmt.Println("Wrting myDb1")
}

func (myDb2) Read() string {
	return "Reading myDb2"
}

func (myDb2) Write(a int) {
	fmt.Println("Wrting myDb2")
}

// NewDb1 returns the instance of myDb1 that implements db.DbAccessor
func NewDb1() db.DbAccessor {
	return myDb1{}
}

// NewDb2 returns the instance of myDb2 that implements db.DbAccessor
func NewDb2() db.DbAccessor {
	return myDb2{}
}
