package main

import "fmt"
import "time"
func main() {
    // while loop in go 

	i := 44		

    switch i { 
	case 1 : 
	case 2 : 
		fmt.Println("two")
	case 3 : 
		fmt.Println("three")
	case 4 : 
		fmt.Println("four")
	case 5 : 
		fmt.Println("five")
	default :
		fmt.Println("Not found")
	}
	

	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday :
			fmt.Println("It's the weekend")
		default :
			fmt.Println("It's a weekday")		
	}

	// type switch in go 
	  whomi := func( i interface{}){
		switch t := i.(type){
		case int :
			fmt.Println("i am int", t)
		case string :
			fmt.Println(" i am strig ", t)
	    case bool :
			fmt.Println("i am boolean", t)
	    default :
			fmt.Println("not found")
		}
	  }

	  whomi(45)

	  
	  
}