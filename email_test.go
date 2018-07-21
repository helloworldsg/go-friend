package friend

import (
	"fmt"
)

func ExampleFindEmails() {
	fmt.Println(FindEmails("hello katie@gmail.com asdf aasdf@asdf.com aasdf@asdf.com"))

	//Output: [aasdf@asdf.com katie@gmail.com]
}

func ExampleValidateEmail() {
	fmt.Println(ValidateEmail("katie@gmail.com"))
	fmt.Println(ValidateEmail("katie+alias@gmail.com"))
	fmt.Println(ValidateEmail("katie+alias"))

	//Output:
	//true
	//true
	//false
}
