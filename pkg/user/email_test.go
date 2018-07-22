package user_test

import (
	"fmt"
	"github.com/hmoniaga/go-friend/pkg/user"
)

func ExampleFindEmails() {
	fmt.Println(user.FindEmails("hello katie@gmail.com asdf aasdf@asdf.com aasdf@asdf.com"))

	//Output: [aasdf@asdf.com katie@gmail.com]
}

func ExampleValidateEmail() {
	fmt.Println(user.ValidateEmail("katie@gmail.com"))
	fmt.Println(user.ValidateEmail("katie+alias@gmail.com"))
	fmt.Println(user.ValidateEmail("katie+alias"))

	//Output:
	//true
	//true
	//false
}
