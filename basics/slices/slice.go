package main

import (
	"fmt"
	"math/rand"
	"slices"
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
		   // slice 
		   var num1 = []int{12 , 3 , 4 , 5}
		   var num3 = []int{12 , 3 , 4 , 5}
		    fmt.Println(slices.Compare(num1 , num3))
   
		}