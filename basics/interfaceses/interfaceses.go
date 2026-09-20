package  main

import (
	"fmt"
)

type payment interface {
	makepay( amount float64 )
	refund ( amount float64 )
}
	
type paymentmethod struct{
     gateway string 
}

func ( p *paymentmethod ) makepay( amount float64){
	 fmt.Println( "paid by method: " , p.gateway , " amount : " , amount )
} 
func ( p *paymentmethod ) refund( amount float64){
	 fmt.Println( "refunded by method: " , p.gateway , " amount : " , amount )
}

type fake struct {}

func ( f *fake ) makepay( amount float64){
	 fmt.Println( "paid by method: fake gateway" , " amount : " , amount )
}
func ( f *fake ) refund( amount float64){
	 fmt.Println( "refunded by method: fake gateway" , " amount : " , amount )
}

type stripepay struct {}

func ( s *stripepay ) makepay( amount float64){
	 fmt.Println( "paid by method: stripe gateway" , " amount : " , amount )
}	
func ( s *stripepay ) refund( amount float64){
	 fmt.Println( "refunded by method: stripe gateway" , " amount : " , amount )
}

type paypalpay struct {}

func ( p *paypalpay ) makepay(amount float64){
	 fmt.Println( "paid by method: paypal gateway" , " amount : " , amount )
}	
func ( p *paypalpay ) refund(amount float64){
	 fmt.Println( "refunded by method: paypal gateway" , " amount : " , amount )
}
 
func main(){
	var payment payment 
	 payment = &stripepay{}
	 payment.makepay( 100 )

	 payment = &fake{}
	 payment.makepay( 200 )
}
