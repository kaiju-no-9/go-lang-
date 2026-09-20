package main

import "fmt"

type OrderStatus int64 

const (
     Resived OrderStatus = iota 
	 Conformed 
	 Shipped 
	 Delivered 
	 Cancelled 
)  

  
func ChaneOrderStatus( status  OrderStatus){
	fmt.Println( " changed order status  :" , status)
} 

func main(){
	ChaneOrderStatus(Cancelled)
}