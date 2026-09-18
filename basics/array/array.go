package main

import (
	"fmt"
	"math/rand"
)
 
func main(){
	 var nums [4]int ; 
	
	 //fmt.Println(len(nums))
	  nums[2] = 22
	    var i int 	
		var j int 
	   for  i=0  ; i <len(nums) ;  i++{
            fmt.Println( "index" , i  ,  "value"  , nums[i])
	   }

	 var vals[5]bool 
	   for i= 0 ; i<len(vals) ;  i++{
		if i%2==0{
			vals[i]=true 
		}
	   }

	  fmt.Println(vals)

	  // array. 
	 nums1 := [5][5]int{} 

	 for  i =0 ; i<len(nums1) ; i++{
		for j  =0 ; j<len(nums1)  ;  j++{
			nums1[i][j] = rand.Intn(100)  
		}
	 }
     
         var count int =0 
	 for i= 0 ; i< len(nums1) ; i++{
         for  j = 0 ; j < len(nums1) ; j++{
			if i==j {
				 count += nums1[i][j] ;
			}
	 }
	 }
     fmt.Println(count) ;

	 
} 