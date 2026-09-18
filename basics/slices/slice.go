package  main 

import (
    "fmt"
    "math/rand"
)

// slice -> it is dynainc in nature , most used construct in go .. 
func main (){
	 // parameters in case of nums = make([]int , size  , capacity(max cap ))
	  var  i int 
       var nums = make([]int , 5  , 10)
	    fmt.Println(nums)
		 fmt.Println(len(nums), cap(nums))

		  for  i=0 ; i<10 ; i++{
			nums =  append( nums , rand.Intn(100))
		  }
		
		  fmt.Println(nums)
		  fmt.Println(len(nums), cap(nums))
		
   
		}