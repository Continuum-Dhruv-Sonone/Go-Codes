package main

import "fmt"

type myDb struct{}

func main() {

	var a interface{}
	a = 1 // Setting values as int
	switching(a)

	a = "Abc" // Setting values as string
	switching(a)

	a = myDb{} // Setting values as myDb
	switching(a)

	a = nil // Setting values as nil
	switching(a)

	a = 1.0 // Setting values as float for default case
	switching(a)

}

// switching accepts the a interface values and selects a cases according to the values it holds
func switching(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("a holds integer")
	case string:
		fmt.Println("a holds   string")
	case nil:
		fmt.Println("The interface value is nil")
	case myDb:
		fmt.Println("The a hold myDb type value")
	default:
		fmt.Println("Default - No case is satisfied")
	}

}
