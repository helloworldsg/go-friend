package user_test

import (
	"fmt"
	"github.com/hmoniaga/go-friend/pkg/user"
)

func ExampleNewUser() {
	fmt.Println(user.NewUser("email@gmail.com").Email)

	// Output: email@gmail.com
}
