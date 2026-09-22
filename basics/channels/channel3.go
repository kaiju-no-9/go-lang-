package  main 
import(
	"fmt"
	"time"
)
// imp of unbuffered channel in go . 
// we send a limit amount of dats without bloking 

func email(emailChan chan string , done chan bool ){
	 
	for email := range emailChan {
		fmt.Println("sending email", email)
          time.Sleep(time.Second*1)
	}
	done <- true  
}

func main() {
	emailChan := make(chan string , 100)
	done := make(chan bool  , 8)
	go email(emailChan , done) 
	 for i:= 0 ; i<10 ; i++{
		emailChan <- fmt.Sprintf("%dpikachu.com" , i )
	 }
	// bloking the main fction until the reqiuired go routine is achived . 
	fmt.Println("service done ") 
	<-done
}