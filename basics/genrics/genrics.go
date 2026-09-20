package main

import "fmt"
 // the concept of genrics stops user dublication . 
  
func printSlice[T any ](item []T) {

	for _, item := range item {
		fmt.Println(item)
	}
}

type User[T any] struct{
	name string 
	age T
}

func main(){
	 num := []int{1,2,3,4,5}
	 names := []string{"nishy","kumar","gandhi"}
		printSlice(num)
		printSlice(names)
		user := User[int]{
			name:"nishant",
			age:32, 
		}
		fmt.Println(user)
		
	
}
 