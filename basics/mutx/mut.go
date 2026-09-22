package  main 

import (
	"fmt"
	"sync"
)

// is used to deal with race around condition genrally deals with a resouce on the bases of if the resource is 
//   is being used by another process. so the result genrally is not atomic 
// mutx refers to loking of resource  so at a time only one resource can work with a issue  . 
type post struct {
	view int 
	mu sync.Mutex
}

func ( p * post ) Increment (wg * sync.WaitGroup ){
	defer  wg.Done() 
	p.mu.Lock()
     p.view += 1
	p.mu.Unlock()
}

func main (){
     mypost := post{view: 0} 
	 var wg sync.WaitGroup 
	 
	 for i:= 0 ; i<100 ; i++{
		wg.Add(1)
           go  mypost.Increment(&wg)
	 }
	 wg.Wait()
	 fmt.Println("the view Count is : " , mypost.view)
	  
}
