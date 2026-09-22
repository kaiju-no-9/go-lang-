package  main 

import ( 
	 "fmt"
    )
// resiving data from one end point to other . 
// the lowing is a buffer channel in go
func task (done chan bool ){
	defer func () {
		done <- true
	}()   
 fmt.Println(" the Task has started")

}
 
func  main2 (){
     done := make ( chan bool )

	 go task(done)
	 <- done 
	 fmt.Println(" the Task has finished")
}
    