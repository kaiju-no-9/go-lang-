package  main 
 import(
	"fmt"
 )
 // here we are changeing the the value by value nopt efects the other . 
func chagenum( nums *int){
	*nums = 0 
	fmt.Println( "in change nums : ", &nums , " ",nums)
}
func main(){
    nums := 1 
	fmt.Println( "in main nums : ", &nums , " ", nums)
	  chagenum(&nums)
	  fmt.Println( "in main nums : ", &nums , " ", nums)
	
}
   