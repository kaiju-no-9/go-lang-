package main 

 import ( 
	"fmt"
 )
 
 func main(){
       m:=make(map[string]string)
	   m["name"]= "Nishchay"
	   m["age"] = "21"
	   m["city"] = "Delhi"
         
	   for k , v  := range m{
		   fmt.Println("key :", k , "value :", v  )   
	   }
	   n:= make(map[string]int , 5)
	   n["sleep"] = 20 

	   fmt.Println( n["sleep"])
	   fmt.Println(n["rust"])
	   	delete(m , "age")
	   fmt.Println(m) 
	     
	     _, ok := m["city"]
	    if ok {
			fmt.Println("city is present")
		}else{
			fmt.Println("city is not present")
		}
	   
	   
 } 	