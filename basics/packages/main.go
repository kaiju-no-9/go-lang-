package packages


import (
	"fmt"

	"github.com/kaiju-no-9/go-lang/auth"
	"github.com/kaiju-no-9/go-lang/user"
)


func main(){
	auth.LoginWithCredentials("Nishchay", "123456")
	fmt.Println(auth.GetSession())

	 user := user.User{
		  Email: "[EMAIL_ADDRESS]",
		 Password: "[PASSWORD]",
	 }

	 fmt.Println(user)
}