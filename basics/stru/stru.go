package main

import (
	"fmt"
	"time"
)

type custumer struct{
   firstName string 
   lastName string 
   email string 
}
type  order struct {
   id string 
   amount int64 
   status string 
   createAt time.Time
   custumer 

}

func ( o *order) changeString(status string){
	 o.status = status 
}

func neworder ( id string, amount int64 , status string ) *order{
     myorder := order{
	 id : id ,
	 amount : amount ,
	 status : status ,
  }
    return &myorder 
}

func newCustumer( firstName string, lastName string, email string) *custumer{
	mystruct:= custumer{
		firstName : firstName ,
		lastName : lastName ,
		email : email ,
	}	
    return &mystruct	
}


func main (){ 
    order1 := neworder("1" , 100 , "pending")
	custumer1 := newCustumer("Nishchay" , "kumar" , " [EMAIL_ADDRESS]")
   
	  fmt.Println(order1 )
	  fmt.Println(custumer1)
}