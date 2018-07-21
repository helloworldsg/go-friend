package friend

import "fmt"

func ExampleNewUser() {
	fmt.Println(NewUser("email@gmail.com").Email)

	// Output: email@gmail.com
}
